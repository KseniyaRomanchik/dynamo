package commands

import (
	"dynamo/app"
	"dynamo/cli/flags"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var (
	Commands []*cli.Command
	Subcommands []*cli.Command
)

func getSubcommands(c string, fls ActionFlags) []*cli.Command {
	//fls := append(flags.Flags, flags.RequiredFlags...)

	return []*cli.Command{
		{
			Name:      Info,
			UsageText: fmt.Sprintf("*** info %s", c),
			Usage:     "info",
			Flags:     fls.Info,
			Action:    commandAction(command(c, "info")),
		},
		{
			Name:      Get,
			UsageText: fmt.Sprintf("*** get %s", c),
			Usage:     "get",
			Flags:     fls.Get,
			Action:    commandAction(command(c, "get")),
		},
		{
			Name:      Update,
			UsageText: fmt.Sprintf("*** update %s", c),
			Usage:     "update",
			Flags:     fls.Update,
			Action:    commandAction(command(c, "update")),
		},
		{
			Name:      Delete,
			UsageText: fmt.Sprintf("*** delete %s", c),
			Usage:     "delete",
			Flags:     fls.Delete,
			Action:    commandAction(command(c, "delete")),
		},
		{
			Name:      Create,
			UsageText: fmt.Sprintf("*** create %s", c),
			Usage:     "create",
			Flags:     fls.Create,
			Action:    commandAction(command(c, "create")),
		},
		{
			Name:      List,
			UsageText: fmt.Sprintf("*** list %s", c),
			Usage:     "list",
			Flags:     fls.List,
			Action:    commandAction(command(c, "list")),
		},

	}
}

type ActionFlags struct {
	Create []cli.Flag
	Get []cli.Flag
	Update []cli.Flag
	Delete []cli.Flag
	List []cli.Flag
	Info []cli.Flag
}

func LoadCommands() {
	tableFls := ActionFlags{
		Create: append(flags.Flags, flags.TName, flags.TableAttr),
		Get: append(flags.Flags, flags.TName),
		Update: append(flags.Flags, flags.TName, flags.TableAttr),
		Delete: append(flags.Flags, flags.TName),
		List: flags.Flags,
		Info: append(flags.Flags, flags.TName),
	}

	itemFls := ActionFlags{
		Create: append(flags.Flags, flags.TName, flags.ItemKey, flags.ItemAttr),
		Get: append(flags.Flags, flags.TName, flags.ItemKey),
		Update: append(flags.Flags, flags.TName, flags.ItemKey, flags.ItemAttr),
		Delete: append(flags.Flags, flags.TName, flags.ItemKey),
		List: append(flags.Flags, flags.TName),
		Info: append(flags.Flags, flags.TName, flags.ItemKey),
	}

	Commands = []*cli.Command{
		{
			Name:      Table,
			UsageText: "*** table ",
			Usage:     "table",
			Subcommands: getSubcommands(Table, tableFls),
		},
		{
			Name:      Item,
			UsageText: "*** item ",
			Usage:     "item",
			Subcommands: getSubcommands(Item, itemFls),
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
			Key: ctx.String(flags.Key),
			TableAttributes: ctx.String(flags.TableAttributes),
			ItemAttributes: ctx.String(flags.ItemAttributes),
		})
	}
}

//func parseFlagsWithType(ctx *cli.Context) {
//	if ctx.IsSet(flags.HashKey) {
//		key := make(map[string]interface{}, 1)
//		hk := strings.Split(ctx.String(flags.HashKey), "=")
//		key, ok := hk[0]
//		t, ok := hk[1]
//		v, ok := hk[2]
//		if len(hk) == 1 {
//			key[hk[0]] = map[string]interface{}{}
//		}
//		if len
//		key[hk[0]] = map[string]interface{}{
//				hk[1]: hk[2],
//			},
//		}
//		ctx.Set(flags.HashKey, key)
//	}
//
//
//}