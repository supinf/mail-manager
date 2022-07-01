package utils

import "reflect"

// DereferenceInterface interface がポインタ型の場合に値型に変換します
func DereferenceInterface(value interface{}) interface{} {
	return reflect.Indirect(reflect.ValueOf(value)).Interface()
}
