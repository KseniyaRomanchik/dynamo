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

	return printer.Printer.PrintJSON(list.TableNames)
}


func ListItem(opts Options) error {
	item, err := db.Client.ListItem(opts.TableName)
	if err != nil {
		return printer.Printer.PrintAWSErr(err)
	}

	return printer.Printer.PrintJSON(item)
}
