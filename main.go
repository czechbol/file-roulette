/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"os"

	"github.com/czechbol/file-roulette/cmd"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})

	if err := (&cmd.CLIOptions{}).CobraCommand().Execute(); err != nil {
		log.Error(err)
	}
}
