package app

import (
	"dynamo/cli/printer"
	"dynamo/db"
)

func UpdateTable(opts Options) error {
	tableName := opts.TableName

	table, err := db.Client.UpdTable(tableName)
	if err != nil {
		return err
	}

	return printer.Printer.PrintJSON(table)
}
