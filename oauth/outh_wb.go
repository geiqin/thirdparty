package oauth

import (
	"encoding/json"
	"errors"
	"regexp"
	"strconv"
)

//获取登录地址
func (e *AuthWb) GetRedirectUrl(state string) string {
	return "https://api.weibo.com/oauth2/authorize?client_id=" + e.Conf.ClientId + "&response_type=code&display=page&redirect_uri=" + e.Conf.RedirectUrl + "&state=" + state
}

//获取token
func (e *AuthWb) GetToken(code string) (*AuthWbSuccRes, error) {

	str, err := HttpPost("https://api.weibo.com/oauth2/access_token?client_id=" + e.Conf.ClientId + "&client_secret=" + e.Conf.ClientSecret + "&code=" + code + "&grant_type=authorization_code&redirect_uri=" + e.Conf.RedirectUrl)
	if err != nil {
		return nil, err
	}

	ismatch, _ := regexp.MatchString("error", str)
	if ismatch {

		p := &AuthWbErrRes{}
		err := json.Unmarshal([]byte(str), p)
		if err != nil {
			return nil, err
		}
		return nil, errors.New("Error:" + strconv.Itoa(p.Error) + " Error_description:" + p.ErrorDescription)

	} else {

		p := &AuthWbSuccRes{}
		err := json.Unmarshal([]byte(str), p)
		if err != nil {
			return nil, err
		}
		return p, nil
	}

}

//获取第三方用户信息
func (e *AuthWb) GetUserInfo(accessToken string, openid string) (string, error) {

	str, err := HttpGet("https://api.weibo.com/2/users/show.json?access_token=" + accessToken + "&uid=" + openid)
	if err != nil {
		return "", err
	}
	return string(str), nil
}

//构造Weibo授权登录
func NewAuthWb(config *AuthConf) *AuthWb {
	return &AuthWb{
		Conf: config,
	}
}
