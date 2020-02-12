package commands

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"dynamo/cli/flags"
	log "github.com/sirupsen/logrus"
)

var (
	Commands []*cli.Command
	Subcommands []*cli.Command
)

func getSubcommands(command Command) []*cli.Command {
	fls := append(flags.Flags, flags.RequiredFlags...)

	return []*cli.Command{
		{
			Name:      string(Info),
			UsageText: fmt.Sprintf("*** info %s", command),
			Usage:     "info",
			Flags:     fls,
			Action:    commandAction(info(command)),
		},
		{
			Name:      string(Get),
			UsageText: fmt.Sprintf("*** get %s", command),
			Usage:     "get",
			Flags:     fls,
			Action:    commandAction(get(command)),
		},
		{
			Name:      string(Put),
			UsageText: fmt.Sprintf("*** put %s", command),
			Usage:     "put",
			Flags:     fls,
			Action:    commandAction(put(command)),
		},
		{
			Name:      string(Delete),
			UsageText: fmt.Sprintf("*** delete %s", command),
			Usage:     "delete",
			Flags:     fls,
			Action:    commandAction(deleteAction(command)),
		},
		{
			Name:      string(Post),
			UsageText: fmt.Sprintf("*** post %s", command),
			Usage:     "post",
			Flags:     fls,
			Action:    commandAction(post(command)),
		},
		{
			Name:      string(List),
			UsageText: fmt.Sprintf("*** list %s", command),
			Usage:     "list",
			Flags:     flags.Flags,
			Action:    commandAction(list(command)),
		},

	}
}

func LoadCommands() {
	Commands = []*cli.Command{
		{
			Name:      string(Table),
			UsageText: "*** table ",
			Usage:     "table",
			Flags:     flags.Flags,
			Subcommands: getSubcommands(Table),
		},
		{
			Name:      string(Item),
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
