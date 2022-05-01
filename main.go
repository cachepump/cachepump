package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/cachepump/cachepump/pump"
	"github.com/cachepump/cachepump/server"

	logger "github.com/AntonYurchenko/log-go"
)

// Application flags.
var endpoint = flag.String("e", ":8080", "http endpoint")
var logLevel = flag.String("l", "INFO", "level for log messages")
var sourceFile = flag.String("s", "./config.yml", "path to yaml file with description of all sources")

func init() {
	flag.Parse()
	logger.SetLevelStr(*logLevel)
}

func main() {

	exit := make(chan os.Signal, 1)
	start()

	// ^C handler.
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	<-exit
	stop()
}

// start is a function for starting of all processes inside this application.
func start() {
	pump.Start(*sourceFile)
	server.Start(*endpoint)
}

// start is a function for stoping of all processes inside this application.
func stop() {
	pump.Stop()
	server.Stop()
}
