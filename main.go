package main

import (
	"runtime"

	"github.com/Pxrzival/bitcoind-exporter/config"
	"github.com/Pxrzival/bitcoind-exporter/fetcher"
	"github.com/Pxrzival/bitcoind-exporter/prometheus"
	"github.com/Pxrzival/bitcoind-exporter/zmq"
	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

func init() {
	log.SetFormatter(&prefixed.TextFormatter{
		TimestampFormat:  "2006/01/02 - 15:04:05",
		FullTimestamp:    true,
		QuoteEmptyFields: true,
		SpacePadding:     45,
	})

	log.SetReportCaller(true)

	level, err := log.ParseLevel(config.C.LogLevel)
	if err != nil {
		log.WithError(err).Fatal("Invalid log level")
	}

	log.SetLevel(level)
}

func main() {
	log.WithFields(log.Fields{
		"commit":  commit,
		"runtime": runtime.Version(),
		"arch":    runtime.GOARCH,
	}).Infof("Bitcoind Exporter ₿ %s", version)

	go prometheus.Start()
	go zmq.Start()

	fetcher.Start()
}
