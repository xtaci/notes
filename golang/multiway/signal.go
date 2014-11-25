package main

import (
	"os"
	"os/signal"
	"syscall"
)

import (
	"utils"
)

//----------------------------------------------- handle unix signals
func sig_handler() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP)
	for {
		msg := <-ch
		switch msg {
		case syscall.SIGHUP:
			utils.INFO("Recevied signal:", msg)
		}
	}
}
