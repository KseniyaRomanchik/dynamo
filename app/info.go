package app

import (
	"dynamo/cli/printer"
	"dynamo/db"
)

func InfoTable(opts Options) error {
	tableName := opts.TableName

	info, err := db.Client.InfoTable(tableName)
	if err != nil {
		return printer.Printer.PrintAWSErr(err)
	}

	return printer.Printer.PrintText(info)
}
