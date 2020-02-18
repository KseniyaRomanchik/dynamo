package app

import (
	"dynamo/cli/printer"
	"dynamo/db"
)

func UpdateTable(opts Options) error {
	table, err := db.Client.UpdTable(opts.TableName)
	if err != nil {
		return printer.Printer.PrintAWSErr(err)
	}

	return printer.Printer.PrintJSON(table)
}


func UpdateItem(opts Options) error {
	item, err := db.Client.UpdItem(opts.TableName, opts.Key, opts.ItemAttributes)
	if err != nil {
		return printer.Printer.PrintAWSErr(err)
	}

	return printer.Printer.PrintText(item)
}
