package flags

const (
	TableName = "table-name"
	Key = "key"
	ItemAttributes = "item-attributes"
	TableAttributes = "table-attributes"
	Prefix = "prefix"

	Output = "output"

	ItemsList = "items_list"
	AwsAccount = "aws_account"
	RawItemKeyName = "raw_item_key_name"
	RawItemKeyValue = "raw_item_key_value"
	EnvType = "env_type"
	SsmName = "ssm_name"
	Content = "content"

	versionTemplate = `Version: %s
Commit: %s
Image: %s
Timestamp: %s
`
)

var (
	Image  string
	Commit string
	Time   string
)