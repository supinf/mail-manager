package dynamodb

import (
	"reflect"
	"strings"

	"github.com/supinf/supinf-mail/app-api/aws/utils"
)

// Tags model からタグの情報を取得します
func Tags(model interface{}) (modelName, hashKey, rangeKey string, localSecondaryIndexes LocalSecondaryIndexes, globalSecondaryIndexes GlobalSecondaryIndexes) {
	typ := reflect.TypeOf(model)
	fieldCount := typ.NumField()

	modelName = strings.ToLower(typ.Name())

	for i := 0; i < fieldCount; i++ {
		tag := typ.Field(i).Tag

		defs := strings.Split(tag.Get("dynamo"), ",")
		localIdx := tag.Get("localIndex")
		globalIdx := tag.Get("index")

		if len(defs) > 1 {
			if defs[1] == "hash" {
				hashKey = defs[0]

			} else if defs[1] == "range" {
				rangeKey = defs[0]
			}
		}
		if len(localIdx) > 0 {
			lsi := strings.Split(localIdx, ",")
			localSecondaryIndexes = append(localSecondaryIndexes, LSI{
				name:     lsi[0],
				rangeKey: defs[0],
			})
		}
		if len(globalIdx) > 0 {
			if globalSecondaryIndexes == nil {
				globalSecondaryIndexes = make(GlobalSecondaryIndexes, 0)
			}

			gsi := strings.Split(globalIdx, ",")

			idx, _ := globalSecondaryIndexes.FindByName(gsi[0])
			if idx < 0 {
				globalSecondaryIndexes = append(globalSecondaryIndexes, GSI{
					name: gsi[0],
				})
				idx = len(globalSecondaryIndexes) - 1
			}

			if gsi[1] == "hash" {
				globalSecondaryIndexes[idx].hashKey = defs[0]
			} else if gsi[1] == "range" {
				globalSecondaryIndexes[idx].rangeKey = defs[0]
			}
		}
	}

	return modelName, hashKey, rangeKey, localSecondaryIndexes, globalSecondaryIndexes
}

// columnValue 特定カラムの値を取得します
func columnValue(modelElm reflect.Value, colName string) interface{} {
	rv := modelElm.FieldByNameFunc(func(name string) bool {
		return name == utils.SnakeToCamelCase(colName, nil) ||
			name == utils.SnakeToCamelCase(colName, map[string]string{"id": "ID"})
	})
	return utils.DereferenceInterface(rv.Interface())
}
