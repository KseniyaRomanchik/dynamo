package app

import (
	"dynamo/cli/printer"
	"dynamo/db"
)

func ListTable(opts Options) error {
	list, err := db.Client.ListTable()
	if err != nil {
		return err
	}

	return printer.Printer.PrintJSON(list)
}
