/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"github.com/aiziyuer/registryV2/cmd"
	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load()
}

func main() {
	cmd.Execute()
}
