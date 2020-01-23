/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"github.com/aiziyuer/registryV2/cmd"
	"github.com/aiziyuer/registryV2/impl/util"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func init() {
	_ = godotenv.Load()

	level, err := logrus.ParseLevel(util.GetEnvAnyWithDefault("LOG_LEVEL", "info"))
	if err != nil {
		level = logrus.DebugLevel
	}

	logrus.SetLevel(level)
}

func main() {
	cmd.Execute()
}
