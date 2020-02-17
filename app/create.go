package app

import (
	"dynamo/cli/printer"
	"dynamo/db"
)

func CreateTable(opts Options) error {
	tableName := opts.TableName

	table, err := db.Client.CrTable(tableName)
	if err != nil {
		return printer.Printer.PrintAWSErr(err)
	}

	return printer.Printer.PrintText(table)
}

func CreateItem(opts Options) error {
	item, err := db.Client.CrItem(opts.TableName, opts.HashKey)
	if err != nil {
		return printer.Printer.PrintAWSErr(err)
	}

	return printer.Printer.PrintText(item)
}
