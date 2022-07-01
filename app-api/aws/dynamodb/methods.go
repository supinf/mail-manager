package dynamodb

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"

	guregu "github.com/guregu/dynamo"
	"github.com/supinf/supinf-mail/app-api/aws/utils"
)

// TODO: 最低限しか実装していないので拡張していく

// Query クエリ
type Query struct {
	db *DynamoDB

	model           interface{}
	tableNamePrefix string

	hashKey                string
	rangeKey               string
	localSecondaryIndexes  LocalSecondaryIndexes
	globalSecondaryIndexes GlobalSecondaryIndexes

	table *guregu.Table
	scan  *guregu.Scan
	query *guregu.Query
	put   *guregu.Put

	lazyFilterFuncList []func(out interface{})

	useGSI         bool
	usedRange      bool
	usedIndexRange bool
}

// LSI Local Secondary Index
type LSI struct {
	name     string
	rangeKey string
}

// GSI Global Secondary Index
type GSI struct {
	name     string
	hashKey  string
	rangeKey string
}

// LocalSecondaryIndexes LSI リスト
type LocalSecondaryIndexes []LSI

// GlobalSecondaryIndexes GSI リスト
type GlobalSecondaryIndexes []GSI

// NewQuery 新しいクエリを生成します
func NewQuery(db *DynamoDB, tableNamePrefix string) *Query {
	return &Query{
		db:              db,
		tableNamePrefix: tableNamePrefix,
	}
}

func (q *Query) Table(model interface{}) *Query {
	mn, hk, rk, lsi, gsi := Tags(model)
	table := q.db.Client().Table(fmt.Sprintf("%s-%s", q.tableNamePrefix, mn))

	q.model = model
	q.hashKey = hk
	q.rangeKey = rk
	q.localSecondaryIndexes = lsi
	q.globalSecondaryIndexes = gsi
	q.table = &table

	return q
}

func (q *Query) Scan() *Query {
	q.scan = q.table.Scan()
	return q
}

func (q *Query) Get(name string, value interface{}) *Query {
	// ベーステーブルの HASH でない場合
	if q.hashKey != name {
		fmt.Printf("%s is not hash key", name)
		return nil
	}
	// 上記以外の場合
	q.query = q.table.Get(name, value)
	return q
}

func (q *Query) IndexGet(index string, name string, value interface{}) *Query {
	// インデックス(GSI)が存在しない、もしくはインデックス(GSI)の HASH と一致しない場合
	if i, found := utils.ContainsStrWithIndex(q.globalSecondaryIndexes.Names(), index); found {
		if q.globalSecondaryIndexes[i].hashKey != name {
			fmt.Printf("%s is not %s's hash key", name, index)
			return nil
		}
	} else {
		fmt.Printf("%s is not gsi", index)
		return nil
	}
	// 上記以外の場合
	q.query = q.table.Get(name, value).Index(index)
	q.useGSI = true
	return q
}

func (q *Query) Range(name string, op string, value ...interface{}) *Query {
	// ベーステーブルの RANGE でない場合
	if q.rangeKey != name {
		fmt.Printf("%s is not range key", name)
		return nil
	}
	// GSI を利用した検索の場合
	if q.useGSI {
		fmt.Printf("need to use IndexRange func")
		return nil
	}
	// Range / IndexRange は合わせて 1回 しか使えないので、IndexRange 利用済みの場合は Filter を利用
	// RANGE はテーブルに1つしかない（=ここが複数回呼ばれるとしたら条件の上書き）ので、usedRange フラグが立っていても Filter は呼ばない
	if q.usedIndexRange {
		return q.Filter(name, op, value...)
	}
	// 上記以外の場合
	q.query = q.query.Range(name, ConvertOperator(op), value...)
	q.usedRange = true
	return q
}

func (q *Query) IndexRange(name string, op string, value ...interface{}) *Query {
	if q.useGSI {
		// GSI を利用した検索の場合
		_, found := utils.ContainsStrWithIndex(q.globalSecondaryIndexes.RangeKeys(), name)
		if !found {
			fmt.Printf("%s is not gsi range key", name)
			return nil
		}
		// GSI による検索の場合は Range 関数は呼べない & RANGE は GSI に1つしかない（=ここが複数回呼ばれるとしたら条件の上書き）ので、
		// usedRange, usedIndexRange フラグが立っていても Filter は呼ばない
		q.query = q.query.Range(name, ConvertOperator(op), value...)

	} else {
		// LSI を利用した検索の場合
		i, found := utils.ContainsStrWithIndex(q.localSecondaryIndexes.RangeKeys(), name)
		if !found {
			fmt.Printf("%s is not lsi range key", name)
			return nil
		}
		// Range / IndexRange は合わせて 1回 しか使えないので、利用済みの場合は Filter を利用
		if q.usedRange || q.usedIndexRange {
			return q.Filter(name, op, value...)
		}
		indexName := q.localSecondaryIndexes[i].name
		q.query = q.query.Index(indexName).Range(name, ConvertOperator(op), value...)
	}

	q.usedIndexRange = true
	return q
}

func (q *Query) Filter(name string, op string, value ...interface{}) *Query {
	// LSI による検索を利用済みの場合（射影される項目がキーのみの場合に Filter を利用できないため、最後に独自フィルタリング）
	if q.usedIndexRange {
		q.lazyFilterFuncList = append(q.lazyFilterFuncList, func(out interface{}) {
			filter(out, name, op, value...)
		})
		return q
	}

	// 上記以外の場合
	expr := ""
	isNotComparisonOperator, _ := regexp.MatchString("^[a-z]+$", op)
	if isNotComparisonOperator {
		// 比較演算子でない場合
		expr = fmt.Sprintf("%s($, ?)", op)
	} else {
		// 比較演算子の場合
		expr = fmt.Sprintf("$ %s ?", op)
	}

	args := append([]interface{}{name}, value...)
	q.query = q.query.Filter(expr, args...)

	return q
}

func (q *Query) FilterByExpr(expr string, args ...interface{}) *Query {
	q.query = q.query.Filter(expr, args...)
	return q
}

func (q *Query) Put(item interface{}) *Query {
	q.put = q.table.Put(item)
	return q
}

func (q *Query) All(out interface{}) error {
	if q.scan != nil {
		return q.scan.All(out)

	} else if q.query != nil {
		if err := q.query.All(out); err != nil {
			return err
		}
		for _, fn := range q.lazyFilterFuncList {
			fn(out)
		}
		return nil

	}
	return errors.New("invalid query")
}

func (q *Query) AllWithAutoFetch(out interface{}) error {
	if q.query != nil {
		if err := q.query.All(out); err != nil {
			return err
		}

		// IndexRange 利用（LSIによる検索）時はプライマリキー（Hash & Range）しか取得しないのでそれをもとに再検索
		if q.usedIndexRange {
			// クエリのレスポンス（テーブルのアイテムリスト）でループ
			elms := reflect.ValueOf(out).Elem()
			for i := 0; i < elms.Len(); i++ {
				// 対象インデックスのアイテムを取得
				elm := elms.Index(i)

				// プライマリキーの値を取得
				hk := columnValue(elm, q.hashKey)
				rk := columnValue(elm, q.rangeKey)

				// プライマリキーで再検索
				model := reflect.New(reflect.TypeOf(q.model)).Interface()
				query := NewQuery(q.db, q.tableNamePrefix)
				if err := query.Table(q.model).Get(q.hashKey, hk).Range(q.rangeKey, "=", rk).One(model); err != nil {
					return err
				}

				// レスポンスの値を上書き
				elm.Set(reflect.ValueOf(model).Elem())
			}
		}

		for _, fn := range q.lazyFilterFuncList {
			fn(out)
		}

		return nil

	}
	return errors.New("invalid query")
}

func (q *Query) One(out interface{}) error {
	return q.query.One(out)
}

func (q *Query) Run() error {
	return q.put.Run()
}

// FindByName 名前から対象の LSI を探索します
func (lsi LocalSecondaryIndexes) FindByName(name string) (int, LSI) {
	for i, idx := range lsi {
		if idx.name == name {
			return i, idx
		}
	}
	return -1, LSI{}
}

// Names LSI の name のスライスを返します
func (lsi LocalSecondaryIndexes) Names() []string {
	names := make([]string, len(lsi))
	for i, idx := range lsi {
		names[i] = idx.name
	}
	return names
}

// RangeKeys LSI の rangeKey のスライスを返します
func (lsi LocalSecondaryIndexes) RangeKeys() []string {
	keys := make([]string, len(lsi))
	for i, idx := range lsi {
		keys[i] = idx.rangeKey
	}
	return keys
}

// FindByName 名前から対象の GSI を探索します
func (gsi GlobalSecondaryIndexes) FindByName(name string) (int, GSI) {
	for i, idx := range gsi {
		if idx.name == name {
			return i, idx
		}
	}
	return -1, GSI{}
}

// Names GSI の name のスライスを返します
func (gsi GlobalSecondaryIndexes) Names() []string {
	names := make([]string, len(gsi))
	for i, idx := range gsi {
		names[i] = idx.name
	}
	return names
}

// HashKeys GSI の hashKey のスライスを返します
func (gsi GlobalSecondaryIndexes) HashKeys() []string {
	keys := make([]string, len(gsi))
	for i, idx := range gsi {
		keys[i] = idx.hashKey
	}
	return keys
}

// RangeKeys GSI の rangeKey のスライスを返します
func (gsi GlobalSecondaryIndexes) RangeKeys() []string {
	keys := make([]string, len(gsi))
	for i, idx := range gsi {
		keys[i] = idx.rangeKey
	}
	return keys
}
