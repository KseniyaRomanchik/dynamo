`make build`

`dynamo table|item create|get|update|delete|list|info`

*--table-name=share_resource*

*--key='{"resource_id":{"S":"work.test1.res1234"},{"creation_delay":{"N":60}}'*

*--item-attributes='{"smth":345,"lolo":"pets"}'*

*--table-attributes='[{"AttributeName":"resource_id","AttributeType":"S","KeyType":"HASH"},{"AttributeName":"test2","AttributeType":"N","KeyType":"RANGE"}]'*