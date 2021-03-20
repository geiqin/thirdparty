package oauth

import (
	//	"encoding/json"
	"errors" //	"fmt"
	//	"net/url"
	"regexp"
	//	"strconv"
)

//获取登录地址
func (e *AuthGithub) GetRedirectUrl(state string) string {
	return "https://github.com/login/oauth/authorize?client_id=" + e.Conf.ClientId + "&redirect_uri=" + e.Conf.RedirectUrl + "&state=" + state
}

//获取token
func (e *AuthGithub) GetToken(code string) (string, error) {
	str, err := HttpGet("https://github.com/login/oauth/access_token?client_id=" + e.Conf.ClientId + "&client_secret=" + e.Conf.ClientSecret + "&code=" + code + "&redirect_uri=" + e.Conf.RedirectUrl)
	if err != nil {
		return "", err
	}

	ismatch, _ := regexp.MatchString("error", str)
	if ismatch {

		return "", errors.New(str)

	} else {
		re, _ := regexp.Compile("access_token=(.*)&scope")
		newres := re.FindStringSubmatch(str)
		if len(newres) >= 2 {
			return newres[1], nil
		}
		return "", nil
	}

}

//获取第三方用户信息
func (e *AuthGithub) GetUserInfo(accessToken string) (string, error) {

	str, err := HttpGet("https://api.github.com/user?access_token=" + accessToken)
	if err != nil {
		return "", err
	}
	return string(str), nil

}

//构造Github授权登录
func NewAuthGithub(config *AuthConf) *AuthGithub {
	return &AuthGithub{
		Conf: config,
	}
}
