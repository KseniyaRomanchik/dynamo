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

func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)

	err := flags.LoadFlags()
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
