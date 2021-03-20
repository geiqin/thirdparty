package utils

import (
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

func StrToMSS(s string) map[string]string {
	if s == "" {
		return nil
	}
	fields := strings.Split(s, "&")
	mss := make(map[string]string)
	for _, field := range fields {
		if strings.Contains(field, "=") {
			keyValue := strings.Split(field, "=")
			key, _ := url.PathUnescape(keyValue[0])
			value := ""
			if len(keyValue) == 2 {
				value, _ = url.PathUnescape(keyValue[1])
				value = strings.ReplaceAll(value, "+", " ")
			}
			mss[key] = value
		}
	}
	return mss
}

func JsonToMSS(s string) map[string]string {
	if s == "" {
		return nil
	}
	msi := make(map[string]interface{})
	err := json.Unmarshal([]byte(s), &msi)
	if err != nil {
		return nil
	}
	mss := make(map[string]string)
	for k, v := range msi {
		mss[k] = convertAnyToStr(v)
	}
	return mss
}

// 将任意类型转string
func convertAnyToStr(v interface{}) string {
	if v == nil {
		return ""
	}
	switch d := v.(type) {
	case string:
		return d
	case int, int8, int16, int32, int64:
		return strconv.FormatInt(reflect.ValueOf(v).Int(), 10)
	case uint, uint8, uint16, uint32, uint64:
		return strconv.FormatUint(reflect.ValueOf(v).Uint(), 10)
	case []byte:
		return string(d)
	case float32, float64:
		return strconv.FormatFloat(reflect.ValueOf(v).Float(), 'f', -1, 64)
	case bool:
		return strconv.FormatBool(d)
	default:
		return fmt.Sprint(v)
	}
}
