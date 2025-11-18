package helper

import "encoding/json"

func ScanJson(data string, v interface{}) error {
	return json.Unmarshal([]byte(data), v)
}