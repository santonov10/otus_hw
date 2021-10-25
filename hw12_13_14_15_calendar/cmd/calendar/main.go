package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os/signal"
	"syscall"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/santonov10/otus_hw/hw12_13_14_15_calendar/internal/config"
	internalhttp "github.com/santonov10/otus_hw/hw12_13_14_15_calendar/internal/server/http"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "./configs/default.json", "Path to configuration file")
}

func main() {
	flag.Parse()

	if flag.Arg(0) == "version" {
		PrintVersion()
		return
	}

	config, err := config.NewConfig(configFile)
	if err != nil {
		log.Fatalln(err)
	}
	_ = config

	server := internalhttp.NewServer()

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	server.Start(ctx)

	<-ctx.Done()

	fmt.Println("закрываем приложение")
}
