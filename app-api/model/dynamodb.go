package model

import (
	"github.com/supinf/supinf-mail/app-api/aws/dynamodb"
)

type dynamoDBQuery struct {
	impl *dynamodb.Query
}

func (q *dynamoDBQuery) Table(model interface{}) query {
	q.impl = q.impl.Table(model)
	return q
}

func (q *dynamoDBQuery) Scan() query {
	q.impl = q.impl.Scan()
	return q
}

func (q *dynamoDBQuery) Get(name string, value interface{}) query {
	q.impl = q.impl.Get(name, value)
	return q
}

func (q *dynamoDBQuery) IndexGet(index string, name string, value interface{}) query {
	q.impl = q.impl.IndexGet(index, name, value)
	return q
}

func (q *dynamoDBQuery) Range(name string, op string, value ...interface{}) query {
	q.impl = q.impl.Range(name, op, value...)
	return q
}

func (q *dynamoDBQuery) IndexRange(name string, op string, value ...interface{}) query {
	q.impl = q.impl.IndexRange(name, op, value...)
	return q
}

func (q *dynamoDBQuery) Filter(name string, op string, value ...interface{}) query {
	q.impl = q.impl.Filter(name, op, value...)
	return q
}

func (q *dynamoDBQuery) FilterByExpr(expr string, args ...interface{}) query {
	q.impl = q.impl.FilterByExpr(expr, args...)
	return q
}

func (q *dynamoDBQuery) Put(item interface{}) query {
	q.impl = q.impl.Put(item)
	return q
}

func (q *dynamoDBQuery) One(out interface{}) error {
	return q.impl.One(out)
}

func (q *dynamoDBQuery) All(out interface{}) error {
	return q.impl.All(out)
}

func (q *dynamoDBQuery) AllWithAutoFetch(out interface{}) error {
	return q.impl.AllWithAutoFetch(out)
}

func (q *dynamoDBQuery) Run() error {
	return q.impl.Run()
}
