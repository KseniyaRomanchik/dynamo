package flags

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
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

	Flags = []cli.Flag{
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:  string(HashKey),
			Usage: "hash key",
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:  string(Output),
			Usage: "output",
		}),
	}

	return nil
}

var RequiredFlags = []cli.Flag{
	&cli.StringFlag{
		Name:  string(TableName),
		Usage: "table name",
		Required: true,
	},
}

var TerraformFlags = []cli.Flag{
	&cli.StringFlag{
		Name:  string(TableName),
		Usage: "table name",
		Required: true,
	},
}
