package app

import (
	"dynamo/cli/printer"
	"dynamo/db"
)

func GetTable(opts Options) error {
	tableName := opts.TableName

	table, err := db.Client.GetTable(tableName)
	if err != nil {
		return printer.Printer.PrintAWSErr(err)
	}

	return printer.Printer.PrintText(table)
}


func GetItem(opts Options) error {
	item, err := db.Client.GItem(opts.TableName, opts.HashKey)
	if err != nil {
		return printer.Printer.PrintAWSErr(err)
	}

	return printer.Printer.PrintText(item)
}