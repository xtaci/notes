package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	//	"strings"
)

func main() {
	cmd := exec.Command("Rscript", "./test.R")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out.String())
}
