package main

import (
	"fmt"
	"hatsumora.com/MiniDNS/utils"
	"os"
	"os/signal"
	"syscall"
)

const port int = 53
const ip string = "127.0.0.1"

func main() {
	fmt.Printf("Starting MiniDNS at %s\n", utils.FormatAddress(ip, port))
	close := make(chan bool)
	setupCloseHandler(close)
	handleTraffic(close)
}

func setupCloseHandler(closeChan chan bool) {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		fmt.Println("Print Ctrl+C to close")
		<-c
		closeChan <- false
		fmt.Println("MiniDNS has stopped")
		os.Exit(0)
	}()
}
