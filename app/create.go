package app

import (
	"dynamo/cli/printer"
	"dynamo/db"
)

func CreateTable(opts Options) error {
	tableName := opts.TableName

	table, err := db.Client.CrTable(tableName)
	if err != nil {
		return err
	}

	return printer.Printer.PrintJSON(table)
}
