package main

import (
	"flag"
	"os"

	"go.uber.org/zap"

	"go-clean-arch/config"
	"go-clean-arch/server"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "config-file", "", "Specify config file path")
	flag.Parse()

	defer func() {
		if err := recover(); err != nil {
			zap.S().Errorf("Recover when start project err:%s", err)
			os.Exit(0)
		}
	}()

	//load config
	cfg, err := config.Load(configFile)
	if err != nil {
		zap.S().Errorf("load config fail with err: %v", err)
		panic(err)
	}

	//start new server
	s, err := server.NewServer(cfg)
	if err != nil {
		zap.S().Errorf("create server fail with err: %v", err)
		panic(err)
	}
	s.Init()

	if err := s.ListenHTTP(); err != nil {
		zap.S().Errorf("start server fail with err: %v", err)
		panic(err)
	}
}
