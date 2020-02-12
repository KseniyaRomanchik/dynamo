package commands

import (
	"dynamo/cli/flags"
	"dynamo/cli/printer"
	"dynamo/db"
	"github.com/urfave/cli/v2"
)

func get(command Command) func (*cli.Context) error {
	return func(ctx *cli.Context) error {
		return fns["get"+"_"+command](ctx)
	}
}

func getTable(ctx *cli.Context) error {
	tableName := ctx.String(string(flags.TableName))

	table, err := db.Client.GetTable(tableName)
	if err != nil {
		return err
	}

	return printer.Printer.PrintJSON(table)
}