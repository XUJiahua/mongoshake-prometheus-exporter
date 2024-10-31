package main

import (
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"github.com/xujiahua/mongoshake-prometheus-exporter/pkg/config"
	"github.com/xujiahua/mongoshake-prometheus-exporter/pkg/mongoshake"
	"github.com/xujiahua/mongoshake-prometheus-exporter/pkg/probe"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		logrus.Fatalf("usage: %s <config.toml>", args[0])
	}
	cfg, err := config.LoadConfig(args[1])
	if err != nil {
		logrus.Fatalf("load config error: %v", err)
	}
	if err := cfg.Validate(); err != nil {
		logrus.Fatalf("validate config error: %v", err)
	}

	if cfg.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	for _, probeCfg := range cfg.Probes {
		options := []mongoshake.Option{}
		if probeCfg.Alias != "" {
			options = append(options, mongoshake.WithAlias(probeCfg.Alias))
		}
		if probeCfg.Timeout > 0 {
			options = append(options, mongoshake.WithTimeout(time.Duration(probeCfg.Timeout)*time.Second))
		}
		client := mongoshake.NewClient(probeCfg.BaseURL, options...)
		prober := probe.NewProber(client, probeCfg.Interval)
		go prober.Probe()
	}

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
