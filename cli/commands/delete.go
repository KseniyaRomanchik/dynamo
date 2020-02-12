package commands

import "github.com/urfave/cli/v2"

func deleteAction(command Command) func (*cli.Context) error {
	return func(ctx *cli.Context) error {
		return nil
	}
}
