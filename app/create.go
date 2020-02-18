package app

import (
	"dynamo/cli/printer"
	"dynamo/db"
)

func CreateTable(opts Options) error {
	table, err := db.Client.CrTable(opts.TableName, opts.AttributeCreates)
	if err != nil {
		return printer.Printer.PrintAWSErr(err)
	}

	return printer.Printer.PrintJSON(table)
}

func CreateItem(opts Options) error {
	item, err := db.Client.CrItem(opts.TableName, opts.Key)
	if err != nil {
		return printer.Printer.PrintAWSErr(err)
	}

	return printer.Printer.PrintText(item)
}
