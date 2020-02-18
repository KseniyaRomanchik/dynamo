package app

import (
	"dynamo/cli/printer"
	"dynamo/db"
)

func DeleteTable(opts Options) error {
	table, err := db.Client.DelTable(opts.TableName)
	if err != nil {
		return printer.Printer.PrintAWSErr(err)
	}

	return printer.Printer.PrintJSON(table.TableDescription)
}


func DeleteItem(opts Options) error {
	item, err := db.Client.DelItem(opts.TableName, opts.Key)
	if err != nil {
		return printer.Printer.PrintAWSErr(err)
	}

	return printer.Printer.PrintJSON(item)
}
