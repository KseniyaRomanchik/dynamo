package commands

import (
	"dynamo/cli/flags"
	"dynamo/cli/printer"
	"dynamo/db"
	"github.com/urfave/cli/v2"
)

func info(command Command) func (*cli.Context) error {
	return func(ctx *cli.Context) error {
		return fns["info"+"_"+command](ctx)
	}
}

func infoTable(ctx *cli.Context) error {
	tableName := ctx.String(string(flags.TableName))

	info, err := db.Client.InfoTable(tableName)
	if err != nil {
		return err
	}

	return printer.Printer.PrintJSON(info)
}

//func item(ctx *cli.Context) error {
//	return nil
//}
