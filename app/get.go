package app

import (
	"dynamo/cli/printer"
	"dynamo/db"
)

func GetTable(opts Options) error {
	tableName := opts.TableName

	table, err := db.Client.GetTable(tableName)
	if err != nil {
		return err
	}

	return printer.Printer.PrintJSON(table)
}