package utils

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
)

// Equals ２つの文字列が同じかどうかを判定します
func Equals(str1, str2 string) bool {
	return str1 == str2
}

// EqualsValue ２つのポインタ文字列が値として同じかどうかを判定します
func EqualsValue(str1, str2 *string) bool {
	return Equals(aws.StringValue(str1), aws.StringValue(str2))
}

// ContainsStrWithIndex リストに文字列が含まれるかどうかを判定し、見つかった場合はそのインデックスを返します
func ContainsStrWithIndex(list []string, str string) (int, bool) {
	for i, s := range list {
		if s == str {
			return i, true
		}
	}
	return -1, false
}

// SnakeToCamelCase SnakeCase を CamelCase に変換します。 特殊な変換をしたい場合には wordMap で指定します（ex. id → ID）。
func SnakeToCamelCase(str string, wordMap map[string]string) string {
	result := ""

	list := strings.Split(str, "_")
	for _, s := range list {
		if wordMap != nil && wordMap[s] != "" {
			// マップで指定されている単語はそれに従って変換
			result += wordMap[s]
		} else {
			// それ以外は先頭大文字に変換
			result += strings.Title(s)
		}
	}

	return result
}
