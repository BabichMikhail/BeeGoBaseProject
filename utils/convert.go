package utils

import "encoding/json"

func ToJson(object interface{}) string {
	bytes, err := json.Marshal(object)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func ToStruct(jsonStr string, object interface{}) {
	err :=  json.Unmarshal([]byte(jsonStr), object)
	if err != nil {
		panic(err)
	}
}