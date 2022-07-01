package dynamodb

import (
	"reflect"
	"strings"
	"time"

	guregu "github.com/guregu/dynamo"
	"github.com/supinf/supinf-mail/app-api/aws/utils"
)

// TODO: テスト実装

// filter リストを任意の条件でフィルタリングします
func filter(out interface{}, name string, op string, value ...interface{}) {
	elms := reflect.ValueOf(out).Elem()
	for i := elms.Len() - 1; i >= 0; i-- {
		elm := elms.Index(i)
		// 対象項目の取得
		val := columnValue(elm, name)
		// 演算してマッチしなければリストから削除
		if !expr(op, val, value...) {
			elms.Set(reflect.AppendSlice(elms.Slice(0, i), elms.Slice(i+1, elms.Len())))
		}
	}
}

func expr(op string, target interface{}, values ...interface{}) bool {
	if len(values) == 0 {
		return false
	}

	switch ConvertOperator(op) {
	case guregu.Equal:
		return equalsInterface(target, values[0])
	case guregu.NotEqual:
		return notEqualsInterface(target, values[0])
	case guregu.Less:
		return lessInterface(target, values[0])
	case guregu.LessOrEqual:
		return lessOrEqualInterface(target, values[0])
	case guregu.Greater:
		return greaterInterface(target, values[0])
	case guregu.GreaterOrEqual:
		return greaterOrEqualInterface(target, values[0])
	case guregu.Between:
		if len(values) <= 1 {
			return false
		}
		return betweenInterface(target, values[0], values[1])
	case guregu.BeginsWith:
		return beginsWith(target, values[0])
	}
	return false
}

func equalsInterface(val1, val2 interface{}) bool {
	v1 := utils.DereferenceInterface(val1)
	v2 := utils.DereferenceInterface(val2)

	vt1 := reflect.TypeOf(v1)
	vt2 := reflect.TypeOf(v2)

	// 型が異なる場合
	if vt1.Kind() != vt2.Kind() {
		return false
	}

	// 比較可能（時間除く）の場合
	if vt1.Comparable() && vt2.Comparable() {
		return v1 == v2
	}

	// 比較可能（時間）の場合
	switch v := v1.(type) { // nolint:gocritic
	case time.Time:
		return v.Equal(v2.(time.Time))
	}

	// 比較不能の場合
	if vt1.Kind() == reflect.Slice {
		v1Slice := reflect.ValueOf(v1).Elem()
		v2Slice := reflect.ValueOf(v2).Elem()
		if v1Slice.Len() != v2Slice.Len() {
			return false
		}
		for i := 0; i < v1Slice.Len(); i++ {
			e1 := utils.DereferenceInterface(v1Slice.Index(i).Interface())
			e2 := utils.DereferenceInterface(v2Slice.Index(i).Interface())

			et1 := reflect.TypeOf(e1)
			et2 := reflect.TypeOf(e2)

			if (!et1.Comparable() && !et2.Comparable()) || (e1 != e2) {
				return false
			}
		}
		return true

	} else if vt1.Kind() == reflect.Map {
		v1Map := reflect.ValueOf(v1).Elem()
		v2Map := reflect.ValueOf(v2).Elem()
		if len(v1Map.MapKeys()) != len(v2Map.MapKeys()) {
			return false
		}
		for _, k := range v1Map.MapKeys() {
			e1 := utils.DereferenceInterface(v1Map.MapIndex(k).Interface())
			e2 := utils.DereferenceInterface(v2Map.MapIndex(k).Interface())

			et1 := reflect.TypeOf(e1)
			et2 := reflect.TypeOf(e2)

			if (!et1.Comparable() && !et2.Comparable()) || (e1 != e2) {
				return false
			}
		}
		return true

	}

	return false
}

func notEqualsInterface(val1, val2 interface{}) bool {
	return !equalsInterface(val1, val2)
}

func lessInterface(val1, val2 interface{}) bool {
	v1 := utils.DereferenceInterface(val1)
	v2 := utils.DereferenceInterface(val2)

	vt1 := reflect.TypeOf(v1)
	vt2 := reflect.TypeOf(v2)

	// 型が異なる場合
	if vt1.Kind() != vt2.Kind() {
		return false
	}

	// キャストして比較
	switch v := v1.(type) {
	case int:
		return v < v2.(int)
	case int8:
		return v < v2.(int8)
	case int16:
		return v < v2.(int16)
	case int32:
		return v < v2.(int32)
	case int64:
		return v < v2.(int64)
	case float32:
		return v < v2.(float32)
	case float64:
		return v < v2.(float64)
	case string:
		return v < v2.(string)
	case time.Time:
		return v.Before(v2.(time.Time))
	}

	return false
}

func lessOrEqualInterface(val1, val2 interface{}) bool {
	return lessInterface(val1, val2) || equalsInterface(val1, val2)
}

func greaterInterface(val1, val2 interface{}) bool {
	v1 := utils.DereferenceInterface(val1)
	v2 := utils.DereferenceInterface(val2)

	vt1 := reflect.TypeOf(v1)
	vt2 := reflect.TypeOf(v2)

	// 型が異なる場合
	if vt1.Kind() != vt2.Kind() {
		return false
	}

	// キャストして比較
	switch v := v1.(type) {
	case int:
		return v > v2.(int)
	case int8:
		return v > v2.(int8)
	case int16:
		return v > v2.(int16)
	case int32:
		return v > v2.(int32)
	case int64:
		return v > v2.(int64)
	case float32:
		return v > v2.(float32)
	case float64:
		return v > v2.(float64)
	case string:
		return v > v2.(string)
	case time.Time:
		return v.After(v2.(time.Time))
	}

	return false
}

func greaterOrEqualInterface(val1, val2 interface{}) bool {
	return greaterInterface(val1, val2) || equalsInterface(val1, val2)
}

func betweenInterface(val, from, to interface{}) bool {
	return greaterOrEqualInterface(val, from) && lessOrEqualInterface(val, to)
}

func beginsWith(base, prefix interface{}) bool {
	baseVal := utils.DereferenceInterface(base)
	prefixVal := utils.DereferenceInterface(prefix)

	baseType := reflect.TypeOf(baseVal)
	prefixType := reflect.TypeOf(prefixVal)

	// 型が異なる場合
	if baseType.Kind() != prefixType.Kind() {
		return false
	}

	// キャストして比較
	switch v := baseVal.(type) {
	case string:
		return strings.HasPrefix(v, prefixVal.(string))
	default:
		return false
	}
}
