package model

import (
	"github.com/supinf/supinf-mail/app-api/aws/dynamodb"
	"github.com/supinf/supinf-mail/app-api/config"
)

var db interface{}

func Initialize() error {
	dynamoDB, err := dynamodb.New()
	if err != nil {
		return err
	}
	db = dynamoDB
	return nil
}

type query interface {
	Table(model interface{}) query
	Scan() query
	Get(name string, value interface{}) query
	IndexGet(index string, name string, value interface{}) query
	Range(name string, op string, value ...interface{}) query
	IndexRange(name string, op string, value ...interface{}) query
	Filter(name string, op string, value ...interface{}) query
	FilterByExpr(expr string, args ...interface{}) query
	Put(item interface{}) query
	One(out interface{}) error
	All(out interface{}) error
	AllWithAutoFetch(out interface{}) error
	Run() error
}

func newQuery() query {
	return &dynamoDBQuery{
		impl: dynamodb.NewQuery(db.(*dynamodb.DynamoDB), config.AppStage),
	}
}
