package commands

import (
	"dynamo/app"
	"fmt"
	"github.com/urfave/cli/v2"
	"dynamo/cli/flags"
	log "github.com/sirupsen/logrus"
)

var (
	Commands []*cli.Command
	Subcommands []*cli.Command
)

func getSubcommands(c string) []*cli.Command {
	fls := append(flags.Flags, flags.RequiredFlags...)

	return []*cli.Command{
		{
			Name:      Info,
			UsageText: fmt.Sprintf("*** info %s", c),
			Usage:     "info",
			Flags:     fls,
			Action:    commandAction(command(c, "info")),
		},
		{
			Name:      Get,
			UsageText: fmt.Sprintf("*** get %s", c),
			Usage:     "get",
			Flags:     fls,
			Action:    commandAction(command(c, "get")),
		},
		{
			Name:      Put,
			UsageText: fmt.Sprintf("*** put %s", c),
			Usage:     "put",
			Flags:     fls,
			Action:    commandAction(command(c, "update")),
		},
		{
			Name:      Delete,
			UsageText: fmt.Sprintf("*** delete %s", c),
			Usage:     "delete",
			Flags:     fls,
			Action:    commandAction(command(c, "delete")),
		},
		{
			Name:      Post,
			UsageText: fmt.Sprintf("*** post %s", c),
			Usage:     "post",
			Flags:     fls,
			Action:    commandAction(command(c, "create")),
		},
		{
			Name:      List,
			UsageText: fmt.Sprintf("*** list %s", c),
			Usage:     "list",
			Flags:     flags.Flags,
			Action:    commandAction(command(c, "list")),
		},

	}
}

func LoadCommands() {
	Commands = []*cli.Command{
		{
			Name:      Table,
			UsageText: "*** table ",
			Usage:     "table",
			Flags:     flags.Flags,
			Subcommands: getSubcommands(Table),
		},
		{
			Name:      Item,
			UsageText: "*** item ",
			Usage:     "item",
			Flags:     flags.Flags,
			Subcommands: getSubcommands(Item),
		},
	}
}

func commandAction(actionFns ...func(*cli.Context) error) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		log.Debugf("\n %s \n", c.Command.UsageText)
		for _, f := range c.Command.Flags {
			log.Debugf("\t* %s: %+v\n", f.Names()[0], c.String(f.Names()[0]))
		}

		for _, fn := range actionFns {
			if err := fn(c); err != nil {
				return err
			}
		}

		return nil
	}
}

func command(c, subc string) func (*cli.Context) error {
	return func(ctx *cli.Context) error {
		return Fns[subc+"_"+c](app.Options{
			TableName: ctx.String(flags.TableName),
		})
	}
}
