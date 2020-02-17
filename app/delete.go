package app

import (
	"dynamo/cli/printer"
	"dynamo/db"
)

func DeleteTable(opts Options) error {
	tableName := opts.TableName

	table, err := db.Client.DelTable(tableName)
	if err != nil {
		return printer.Printer.PrintAWSErr(err)
	}

	return printer.Printer.PrintText(table)
}


func DeleteItem(opts Options) error {
	item, err := db.Client.DelItem(opts.TableName, opts.HashKey)
	if err != nil {
		return printer.Printer.PrintAWSErr(err)
	}

	return printer.Printer.PrintText(item)
}
