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

	Flags = []cli.Flag{
		//altsrc.NewStringFlag(&cli.StringFlag{
		//	Name:  HashKey,
		//	Usage: "hash key",
		//}),
		//altsrc.NewStringFlag(&cli.StringFlag{
		//	Name:  Output,
		//	Usage: "output",
		//}),
	}

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

//var TerraformFlags = []cli.Flag{
//	&cli.StringFlag{
//		Name:  TableName,
//		Usage: "table name",
//		Required: true,
//	},
//}

var AttrUpdates = &cli.StringFlag{
		Name:  AttributeUpdates,
		Usage: "attribute updates",
		Required: true,
	}

var AttrCreates = &cli.StringFlag{
		Name:  AttributeCreates,
		Usage: "attribute creates",
		Required: true,
	}
