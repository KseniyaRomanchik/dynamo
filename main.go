package main

import (
	"dynamo/cli/printer"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"os"
	"dynamo/cli/commands"
	"dynamo/cli/flags"
	"dynamo/db"
)

const (
	LOG_LEVEL = "LOG_LEVEL"
)

func init() {
	logLevel, err := log.ParseLevel(os.Getenv(LOG_LEVEL))
	if err != nil {
		logLevel = log.WarnLevel
	}
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(logLevel)

	err = flags.LoadFlags()
	if err != nil {
		log.Fatal(err)
	}

	commands.LoadCommands()

	db.Init()
	printer.Init()
}

func main() {
	app := &cli.App{
		Name:     "dynamo",
		Usage:    "dynamo {command} {subcommand} --table-name {name} --arg1 value1  --arg2 value2",
		Version:  "0.0.1",
		Flags:    flags.Flags,
		Commands: commands.Commands,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
