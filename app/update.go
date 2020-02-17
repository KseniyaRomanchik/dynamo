package app

import (
	"dynamo/cli/printer"
	"dynamo/db"
)

func UpdateTable(opts Options) error {
	tableName := opts.TableName

	table, err := db.Client.UpdTable(tableName)
	if err != nil {
		return printer.Printer.PrintAWSErr(err)
	}

	return printer.Printer.PrintJSON(table)
}


func UpdateItem(opts Options) error {
	item, err := db.Client.UpdItem(opts.TableName, opts.HashKey, opts.AttributeUpdates)
	if err != nil {
		return printer.Printer.PrintAWSErr(err)
	}

	return printer.Printer.PrintText(item)
}
