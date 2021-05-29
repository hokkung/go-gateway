package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/hokkung/go-gateway/server"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	FlagEnv = kingpin.Flag("env", "Enviroment").Default("local").String()
	FlagLog = kingpin.Flag("log", "Log level").Default("info").String()
)

func initKingpin() {
	kingpin.CommandLine.DefaultEnvars()
	kingpin.Parse()
}

func initServer() {
	s := server.New()
	s.Run()
}

func afterInit() {
	// handle signals
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// wait for signal
	<-c
	fmt.Println("gracefully stop")
}

func main() {
	initKingpin()
	initServer()
	afterInit()
}
