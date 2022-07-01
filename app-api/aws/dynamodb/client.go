package dynamodb

import (
	"github.com/guregu/dynamo"
	"github.com/supinf/supinf-mail/app-api/aws"
)

// DynamoDB Amazon DynamoDB
type DynamoDB struct {
	client *dynamo.DB
}

// New 新しい DynamoDB Client を生成します
func New() (*DynamoDB, error) {
	client, err := dynamoDBClient()
	if err != nil {
		return nil, err
	}
	return &DynamoDB{
		client: client,
	}, nil
}

// Client Client を返します
func (db *DynamoDB) Client() *dynamo.DB {
	return db.client
}

func dynamoDBClient() (*dynamo.DB, error) {
	sess, cfg, err := aws.Configure(nil)
	if err != nil {
		return nil, err
	}
	return dynamo.New(sess, cfg), nil
}
