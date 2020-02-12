package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DynamoRepo interface {
	InfoTable(string) (*dynamodb.DescribeTableOutput, error)
	GetTable(string) (*dynamodb.ScanOutput, error)
	ListTable() (*dynamodb.ListTablesOutput, error)
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

func (c Dynamo) GetTable(tableName string) (*dynamodb.ScanOutput, error) {
	return c.Scan(&dynamodb.ScanInput{
		TableName: aws.String(tableName),
	})
}

func (c Dynamo) ListTable() (*dynamodb.ListTablesOutput, error) {
	return c.ListTables(&dynamodb.ListTablesInput{})
}