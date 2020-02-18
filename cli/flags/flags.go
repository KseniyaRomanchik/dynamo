package flags

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

var (
	Flags []cli.Flag
)

func LoadFlags() error {
	cli.VersionPrinter = printVersion

	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	configPath := fmt.Sprintf("%s/.%s.yaml", home, "trdeploy")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		configPath = fmt.Sprintf("/etc/%s.yaml", "trdeploy")
	}

	Flags = []cli.Flag{}

	return nil
}

var TName = &cli.StringFlag{
	Name:  TableName,
	Usage: "table name",
	Required: true,
}

var ItemKey = &cli.StringFlag{
	Name:  Key,
	Usage: "key",
	Required: true,
}

var ItemAttr = &cli.StringFlag{
	Name:  ItemAttributes,
	Usage: "item attributes",
	Required: true,
}

var TableAttr = &cli.StringFlag{
	Name:  TableAttributes,
	Usage: "table attributes",
	Required: true,
}

var Pref = &cli.StringFlag{
	Name:  Prefix,
	Usage: "prefix",
}
