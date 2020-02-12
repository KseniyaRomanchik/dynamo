package commands

import "github.com/urfave/cli/v2"

type Command string
type Subcommand string

const (
	Get Subcommand = "get"
	List Subcommand = "list"
	Put Subcommand = "put"
	Delete Subcommand = "delete"
	Post Subcommand = "post"
	Info Subcommand = "info"

	Table Command = "table"
	Item Command = "item"

	show_table = "show_table"
	list_items = "list_items"
	get_items = "get_items"
	get_aws = "get_aws"
	get_prefix = "get_prefix"
	update_raw_item_key = "update_raw_item_key"
	update_item_keys = "update_item_keys"
	delete_item_keys = "delete_item_keys"
	item_keys = "item_keys"
	get_vpc = "get_vpc"
	get_db = "get_db"
	get_eks_config = "get_eks_config"
	get_app_var = "get_app_var"
	load_json = "load_json"
	get_ssm = "get_ssm"
	load_ssm = "load_ssm"
)

var fns = map[Command]func(ctx *cli.Context) error {
	"info_table": infoTable,
	"get_table": getTable,
	"list_table": listTable,
}

