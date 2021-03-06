package db

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	log "github.com/sirupsen/logrus"
)

type DynamoRepo interface {
	InfoTable(string) (*dynamodb.DescribeTableOutput, error)
	GetTable(string, string) (*dynamodb.ScanOutput, error)
	ListTable() (*dynamodb.ListTablesOutput, error)
	DelTable(string) (*dynamodb.DeleteTableOutput, error)
	CrTable(string, string) (*dynamodb.CreateTableOutput, error)
	UpdTable(string, string) (*dynamodb.UpdateTableOutput, error)

	//InfoItem(string) (*dynamodb.DescribeItemOutput, error)
	GItem(string, string) (*dynamodb.GetItemOutput, error)
	ListItem(string) ([]*string, error)
	DelItem(string, string) (*dynamodb.DeleteItemOutput, error)
	CrItem(string, string) (*dynamodb.PutItemOutput, error)
	UpdItem(string, string, string) (*dynamodb.UpdateItemOutput, error)
}

type Dynamo struct {
	*dynamodb.DynamoDB
}

var Client DynamoRepo

func Init() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	Client = Dynamo{dynamodb.New(sess)}
}

func (c Dynamo) InfoTable(tableName string) (*dynamodb.DescribeTableOutput, error) {
	return c.DescribeTable(&dynamodb.DescribeTableInput{
		TableName: aws.String(tableName),
	})
}

func (c Dynamo) GetTable(tableName string, prefix string) (*dynamodb.ScanOutput, error) {

	return c.Scan(&dynamodb.ScanInput{
		TableName: aws.String(tableName),
	})
}

func (c Dynamo) ListTable() (*dynamodb.ListTablesOutput, error) {
	return c.ListTables(&dynamodb.ListTablesInput{})
}

func (c Dynamo) DelTable(tableName string) (*dynamodb.DeleteTableOutput, error) {
	return c.DeleteTable(&dynamodb.DeleteTableInput{
		TableName: aws.String(tableName),
	})
}

func (c Dynamo) CrTable(tableName string, attrCr string) (*dynamodb.CreateTableOutput, error) {
	var attrDef []*dynamodb.AttributeDefinition
	var key []*dynamodb.KeySchemaElement

	if err := json.Unmarshal([]byte(attrCr), &attrDef); err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(attrCr), &key); err != nil {
		return nil, err
	}

	input := dynamodb.CreateTableInput{
		AttributeDefinitions: attrDef,
		KeySchema: key,
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(DefaultProvisionedThroughput),
			WriteCapacityUnits: aws.Int64(DefaultProvisionedThroughput),
		},
		TableName: aws.String(tableName),
	}

	log.Debug(input)

	return c.CreateTable(&input)
}

func (c Dynamo) UpdTable(tableName string, attrUpd string) (*dynamodb.UpdateTableOutput, error) {
	var attrDef []*dynamodb.AttributeDefinition
	var key []*dynamodb.KeySchemaElement

	if err := json.Unmarshal([]byte(attrUpd), &attrDef); err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(attrUpd), &key); err != nil {
		return nil, err
	}

	input := dynamodb.UpdateTableInput{
		AttributeDefinitions: attrDef,
		TableName: aws.String(tableName),
	}

	log.Debug(input)

	return c.UpdateTable(&input)
}

func (c Dynamo) GItem(tableName string, keys string) (*dynamodb.GetItemOutput, error) {
	var filter map[string]*dynamodb.AttributeValue

	err := json.Unmarshal([]byte(keys), &filter)
	if err != nil {
		return nil, err
	}

	log.Debugf("%+v", dynamodb.GetItemInput{
		Key: filter,
		TableName: aws.String(tableName),
	})

	return c.GetItem(&dynamodb.GetItemInput{
		Key: filter,
		TableName: aws.String(tableName),
	})
}

func (c Dynamo) ListItem(tableName string) ([]*string, error) {
	output, err := c.DescribeTable(&dynamodb.DescribeTableInput{
		TableName: aws.String(tableName),
	})
	if err != nil {
		return nil, err
	}

	items := make([]*string, len(output.Table.AttributeDefinitions))
	for i, attr := range output.Table.AttributeDefinitions {
		items[i] = attr.AttributeName
	}

	return items, nil
}

func (c Dynamo) DelItem(tableName string, keys string) (*dynamodb.DeleteItemOutput, error) {
	var filter map[string]*dynamodb.AttributeValue

	err := json.Unmarshal([]byte(keys), &filter)
	if err != nil {
		return nil, err
	}

	cond := expression.Name(Locked).Equal(expression.Value(false))

	builder, err := expression.
		NewBuilder().
		WithCondition(cond).
		Build()
	if err != nil {
		return nil, fmt.Errorf("cannot build update expression: %+v", err)
	}

	input := dynamodb.DeleteItemInput{
		Key: filter,
		ExpressionAttributeNames: builder.Names(),
		ExpressionAttributeValues: builder.Values(),
		ConditionExpression: builder.Condition(),
		TableName: aws.String(tableName),
	}

	log.Debug(input)

	return c.DeleteItem(&input)
}

func (c Dynamo) CrItem(tableName string, keys string) (*dynamodb.PutItemOutput, error) {
	var filter map[string]*dynamodb.AttributeValue

	if err := json.Unmarshal([]byte(keys), &filter); err != nil {
		return nil, err
	}

	filter[Locked] = &dynamodb.AttributeValue{
		BOOL: aws.Bool(false),
	}

	input := dynamodb.PutItemInput{
		Item: filter,
		TableName: aws.String(tableName),
	}

	log.Debug(input)

	return c.PutItem(&input)
}

func (c Dynamo) UpdItem(tableName, keys, attrUpd string) (*dynamodb.UpdateItemOutput, error) {
	var filter map[string]*dynamodb.AttributeValue
	var updValues map[string]interface{}

	if err := json.Unmarshal([]byte(keys), &filter); err != nil {
		return nil, fmt.Errorf("filter parse error: %+v", err)
	}
	if err := json.Unmarshal([]byte(attrUpd), &updValues); err != nil {
		return nil, fmt.Errorf("updates parse error: %+v", err)
	}

	cond := expression.Name(Locked).Equal(expression.Value(false))

	var update expression.UpdateBuilder
	for k, v := range updValues {
		update = update.Set(expression.Name(k), expression.Value(v))
	}

	builder, err := expression.
		NewBuilder().
		WithUpdate(update).
		WithCondition(cond).
		Build()
	if err != nil {
		return nil, fmt.Errorf("cannot build update expression: %+v", err)
	}

	input := dynamodb.UpdateItemInput{
		Key: filter,
		ConditionExpression: builder.Condition(),
		ExpressionAttributeNames: builder.Names(),
		ExpressionAttributeValues: builder.Values(),
		UpdateExpression: builder.Update(),
		TableName: aws.String(tableName),
	}

	log.Debug(input)

	return c.UpdateItem(&input)
}