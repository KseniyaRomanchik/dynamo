package app

import (
	"dynamo/cli/printer"
	"dynamo/db"
)

func GetTable(opts Options) error {
	table, err := db.Client.GetTable(opts.TableName, opts.Prefix)
	if err != nil {
		return printer.Printer.PrintAWSErr(err)
	}

	return printer.Printer.PrintJSON(table.Items)
}


func GetItem(opts Options) error {
	item, err := db.Client.GItem(opts.TableName, opts.Key)
	if err != nil {
		return printer.Printer.PrintAWSErr(err)
	}

	return printer.Printer.PrintJSON(item.Item)
}