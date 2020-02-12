package commands

import (
	"dynamo/cli/printer"
	"dynamo/db"
	"github.com/urfave/cli/v2"
)

func list(command Command) func (*cli.Context) error {
	return func(ctx *cli.Context) error {
		return fns["list"+"_"+command](ctx)
	}
}

func listTable(ctx *cli.Context) error {
	list, err := db.Client.ListTable()
	if err != nil {
		return err
	}

	return printer.Printer.PrintJSON(list)
}
