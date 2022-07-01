package dynamodb

import (
	"strings"

	guregu "github.com/guregu/dynamo"
)

// ConvertOperator 文字列を Operator に変換します
func ConvertOperator(op string) guregu.Operator {
	switch strings.ToLower(strings.TrimSpace(op)) {
	case "=":
		return guregu.Equal
	case "!=":
		return guregu.NotEqual
	case "<":
		return guregu.Less
	case "<=":
		return guregu.LessOrEqual
	case ">":
		return guregu.Greater
	case ">=":
		return guregu.GreaterOrEqual
	case "between":
		return guregu.Between
	case "begins_with":
		return guregu.BeginsWith
	}
	return ""
}
