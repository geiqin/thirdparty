package utils

import (
	"strconv"
	"strings"
)

const (
	UNKNOWN = iota - 1
	FEMALE
	MALE
)

var GenderMap = map[int]string{
	UNKNOWN: "未知",
	FEMALE:  "女",
	MALE:    "男",
}

type AuthUserGender struct {
	Code int    `json:"code"`
	Desc string `json:"desc"`
}

var males = []string{"m", "男", "1", "male"}

func GetRealGender(originalGender string) *AuthUserGender {
	originalGender = strings.ToLower(originalGender)
	if originalGender == "" || originalGender == strconv.Itoa(UNKNOWN) {
		return &AuthUserGender{UNKNOWN, GenderMap[UNKNOWN]}
	}
	for _, v := range males {
		if v == originalGender {
			return &AuthUserGender{MALE, GenderMap[MALE]}
		}
	}
	return &AuthUserGender{FEMALE, GenderMap[FEMALE]}
}

// for wechat real gender
func GetWechatRealGender(originalGender string) *AuthUserGender {
	if originalGender == "" || originalGender == "0" {
		return &AuthUserGender{UNKNOWN, GenderMap[UNKNOWN]}
	}
	return GetRealGender(originalGender)
}
