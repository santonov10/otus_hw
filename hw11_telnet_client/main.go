package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	timeout := flag.Duration("timeout", 10*time.Second, "")
	flag.Parse()

	if len(flag.Args()) != 2 {
		log.Fatalln("должно быть 2 аргумента при запуске: \"host\" \"port\"")
	}

	address := flag.Arg(0)
	port := flag.Arg(1)

	telNetClient := NewTelnetClient(net.JoinHostPort(address, port), *timeout, os.Stdin, os.Stdout)
	defer telNetClient.Close()

	err := telNetClient.Connect()
	if err != nil {
		log.Panic(err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGTERM)
	defer cancel()

	go func() {
		telNetClient.Send()
		cancel()
	}()

	go func() {
		telNetClient.Receive()
		fmt.Println("...Connection was closed by peer")
		cancel()
	}()

	<-ctx.Done()
}
