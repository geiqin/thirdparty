package oauth

import (

	//	"encoding/json"
	"encoding/json"
	"errors"

	//	"net/url"
	"regexp"
	//	"strconv"
)

//获取登录地址
func (e *AuthGitee) GetRedirectUrl(state string) string {
	return "https://gitee.com/oauth/authorize?client_id=" + e.Conf.ClientId + "&redirect_uri=" + e.Conf.RedirectUrl + "&response_type=code"
}

//获取token
func (e *AuthGitee) GetToken(code string) (*AuthGiteeSuccRes, error) {

	str, err := HttpPost("https://gitee.com/oauth/token?grant_type=authorization_code&code=" + code + "&client_id=" + e.Conf.ClientId + "&redirect_uri=" + e.Conf.RedirectUrl + "&client_secret=" + e.Conf.ClientSecret)
	if err != nil {
		return nil, err
	}

	ismatch, _ := regexp.MatchString("error", str)
	if ismatch {

		return nil, errors.New(str)

	} else {
		p := &AuthGiteeSuccRes{}
		err := json.Unmarshal([]byte(str), p)
		if err != nil {
			return nil, err
		}
		return p, nil
	}

}

//获取第三方用户信息
func (e *AuthGitee) GetUserInfo(accessToken string) (string, error) {

	str, err := HttpGet("https://gitee.com/api/v5/user?access_token=" + accessToken)
	if err != nil {
		return "", err
	}
	return string(str), nil

}

//构造Gitee授权登录
func NewAuthGitee(config *AuthConf) *AuthGitee {
	return &AuthGitee{
		Conf: config,
	}
}
