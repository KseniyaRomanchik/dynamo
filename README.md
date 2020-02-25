`make build`

`dynamo table|item create|get|update|delete|list|info`

*--table-name=share_resource*

*--key='{"resource_id":{"S":"work.test1.res1234"},{"creation_delay":{"N":60}}'*

*--item-attributes='{"smth":345,"lolo":"pets"}'*

*--table-attributes='[{"AttributeName":"resource_id","AttributeType":"S","KeyType":"HASH"},{"AttributeName":"test2","AttributeType":"N","KeyType":"RANGE"}]'*

> dynamo item get --table-name=trdeploy_share_resources --key='{"resource_id":{"S":"work.test1.res1256"}}'

> dynamo item update --table-name=trdeploy_share_resources --key='{"resource_id":{"S":"work.test1.res1256"}}' --item-attributes='{"smth":345,"lolo":"pets"}'

> dynamo item create --table-name=trdeploy_share_resources --key='{"resource_id":{"S":"work.test1.res1256"}}'

> dynamo item delete --table-name=trdeploy_share_resources --key='{"resource_id":{"S":"work.test1.res1256"}}'

> dynamo item list --table-name=trdeploy_share_resources

> dynamo table get --table-name=trdeploy_share_resources

> dynamo table update ?????

> dynamo table create --table-name=trdeploy_share_resources123 --table-attributes='[{"AttributeName":"resource_id","AttributeType":"S","KeyType":"HASH"},{"AttributeName":"test2","AttributeType":"N","KeyType":"RANGE"}]'

> dynamo table info --table-name=trdeploy_share_resources

> dynamo table list
