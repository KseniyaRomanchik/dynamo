package commands

import (
	"dynamo/app"
)

const (
	Get = "get"
	List = "list"
	Update = "update"
	Delete = "delete"
	Create = "create"
	Info = "info"

	Table = "table"
	Item = "item"

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

var Fns = map[string]func(app.Options) error {
	"info_table": app.InfoTable,
	"get_table": app.GetTable,
	"list_table": app.ListTable,
	"create_table": app.CreateTable,
	"delete_table": app.DeleteTable,
	"update_table": app.UpdateTable,

	"get_item": app.GetItem,
	"list_item": app.ListItem,
	"create_item": app.CreateItem,
	"delete_item": app.DeleteItem,
	"update_item": app.UpdateItem,
}

