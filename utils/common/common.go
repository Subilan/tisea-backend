package common

import (
	"encoding/json"
	"time"
)

func GetTimestampSeconds() int64 {
	return time.Now().Unix()
}

func GetTimestampMilli() int64 {
	return time.Now().UnixMilli()
}

// 将给定的数据利用 json.Marshal 编码为 JSON，然后转换为字符串输出。如果有错误，输出字符串为空字符串 ""，并会返回一个 error
func Jsonify(data interface{}) (string, error) {
	result, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

// 忽略错误版的 Jsonify。建议用在确定 data 可以被正确编码为 JSON 时。
func Jsonify_(data interface{}) string {
	result, err := Jsonify(data)
	if err != nil {
		return ""
	}
	return result
}