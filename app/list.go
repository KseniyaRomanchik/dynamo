package app

import (
	"dynamo/cli/printer"
	"dynamo/db"
)

func ListTable(opts Options) error {
	list, err := db.Client.ListTable()
	if err != nil {
		return printer.Printer.PrintAWSErr(err)
	}

	return printer.Printer.PrintJSON(list)
}


func ListItem(opts Options) error {
	tableName := opts.TableName

	table, err := db.Client.ListItem(tableName)
	if err != nil {
		return printer.Printer.PrintAWSErr(err)
	}

	return printer.Printer.PrintJSON(table)
}
