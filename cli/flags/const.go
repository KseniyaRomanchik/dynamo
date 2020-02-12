package flags

type Flag string

const (
	TableName Flag = "table-name"
	HashKey Flag = "hash-key"
	Output Flag = "output"

	ItemsList Flag = "items_list"
	AwsAccount Flag = "aws_account"
	Prefix Flag = "prefix"
	RawItemKeyName Flag = "raw_item_key_name"
	RawItemKeyValue Flag = "raw_item_key_value"
	EnvType Flag = "env_type"
	SsmName Flag = "ssm_name"
	Content Flag = "content"
)