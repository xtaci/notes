package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	go server()
	<-time.After(2 * time.Second)
	// A public key may be used to authenticate against the remote
	// server by using an unencrypted PEM-encoded private key file.
	//
	// If you have an encrypted private key, the crypto/x509 package
	// can be used to decrypt it.
	key, err := ioutil.ReadFile("/Users/xtaci/.ssh/id_rsa")
	if err != nil {
		log.Fatalf("unable to read private key: %v", err)
	}

	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("unable to parse private key: %v", err)
	}

	config := &ssh.ClientConfig{
		User: "xtaci",
		Auth: []ssh.AuthMethod{
			// Use the PublicKeys method for remote authentication.
			ssh.PublicKeys(signer),
		},
	}

	// Connect to the remote server and perform the SSH handshake.
	client, err := ssh.Dial("tcp", "127.0.0.1:2022", config)
	if err != nil {
		log.Fatalf("unable to connect: %v", err)
	}
	session, _ := client.NewSession()
	stdin, err := session.StdinPipe()
	go io.Copy(stdin, os.Stdin)

	stdout, err := session.StdoutPipe()
	go io.Copy(os.Stdout, stdout)

	stderr, err := session.StderrPipe()
	go io.Copy(os.Stderr, stderr)
	_ = session.Shell()
	defer client.Close()
	select {}
}

func server() {
	f, err := os.Open("/Users/xtaci/.ssh/authorized_keys")
	if err != nil {
		panic(err)
	}
	bts, _ := ioutil.ReadAll(f)
	out, _, _, _, err := ssh.ParseAuthorizedKey(bts)

	fmt.Println(out)
	config := &ssh.ServerConfig{
		PublicKeyCallback: func(conn ssh.ConnMetadata, key ssh.PublicKey) (*ssh.Permissions, error) {
			cc := ssh.CertChecker{}
			cc.IsAuthority = func(auth ssh.PublicKey) bool {
				return false
			}
			log.Println(key.Marshal())
			log.Println(out.Marshal())
			if bytes.Equal(key.Marshal(), out.Marshal()) {
				return nil, nil
			}
			return nil, errors.New("permission")
		},
	}
	privateBytes, err := ioutil.ReadFile("/Users/xtaci/.ssh/id_rsa")
	if err != nil {
		panic("Failed to load private key")
	}

	private, err := ssh.ParsePrivateKey(privateBytes)
	if err != nil {
		panic("Failed to parse private key")
	}

	config.AddHostKey(private)

	// Once a ServerConfig has been configured, connections can be
	// accepted.
	listener, err := net.Listen("tcp", "0.0.0.0:2022")
	if err != nil {
		panic("failed to listen for connection")
	}
	nConn, err := listener.Accept()
	if err != nil {
		panic("failed to accept incoming connection")
	}

	// Before use, a handshake must be performed on the incoming
	// net.Conn.
	_, chans, reqs, err := ssh.NewServerConn(nConn, config)
	if err != nil {
		panic(err)
	}
	// The incoming Request channel must be serviced.
	go ssh.DiscardRequests(reqs)

	// Service the incoming Channel channel.
	for newChannel := range chans {
		// Channels have a type, depending on the application level
		// protocol intended. In the case of a shell, the type is
		// "session" and ServerShell may be used to present a simple
		// terminal interface.
		if newChannel.ChannelType() != "session" {
			newChannel.Reject(ssh.UnknownChannelType, "unknown channel type")
			continue
		}
		channel, requests, err := newChannel.Accept()
		if err != nil {
			panic("could not accept channel.")
		}

		// Sessions have out-of-band requests such as "shell",
		// "pty-req" and "env".  Here we handle only the
		// "shell" request.
		go func(in <-chan *ssh.Request) {
			for req := range in {
				ok := false
				switch req.Type {
				case "shell":
					ok = true
					if len(req.Payload) > 0 {
						// We don't accept any
						// commands, only the
						// default shell.
						ok = false
					}
				}
				req.Reply(ok, nil)
			}
		}(requests)

		term := terminal.NewTerminal(channel, "> ")

		go func() {
			defer channel.Close()
			for {
				line, err := term.ReadLine()
				if err != nil {
					break
				}
				fmt.Println(line)
			}
		}()
	}
}
