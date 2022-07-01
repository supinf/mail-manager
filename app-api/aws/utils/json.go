package utils

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
)

// ReplaceAllFieldValues 与えられた JSON 文字列の値部分をすべて置換します
func ReplaceAllFieldValues(jsonStr, new string) (*string, error) {
	if jsonStr == "" {
		return nil, nil
	}

	var mapData map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &mapData); err != nil {
		return nil, err
	}

	newMapData := make(map[string]interface{})
	for key := range mapData {
		newMapData[key] = new
	}

	newJSONStr, err := json.Marshal(newMapData)
	if err != nil {
		return nil, err
	}

	return aws.String(string(newJSONStr)), nil
}

// MapToJSONString map を JSON 文字列に変換します
func MapToJSONString(m map[string]interface{}) (*string, error) {
	if m == nil {
		return nil, nil
	}
	jsonByte, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return aws.String(string(jsonByte)), nil
}

// JSONStringToMap JSON 文字列を map に変換します
func JSONStringToMap(str *string) (map[string]interface{}, error) {
	if str == nil {
		return nil, nil
	}
	var mapData map[string]interface{}
	if err := json.Unmarshal([]byte(aws.StringValue(str)), &mapData); err != nil {
		return nil, err
	}
	return mapData, nil
}
