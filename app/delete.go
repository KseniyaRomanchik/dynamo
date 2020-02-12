package app

import (
	"dynamo/cli/printer"
	"dynamo/db"
)

func DeleteTable(opts Options) error {
	tableName := opts.TableName

	table, err := db.Client.DelTable(tableName)
	if err != nil {
		return err
	}

	return printer.Printer.PrintJSON(table)
}
