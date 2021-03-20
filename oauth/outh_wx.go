package oauth

import (
	"encoding/json"
	"errors" //	"fmt"
	"regexp"
	"strconv"
)

//获取登录地址
func (e *AuthWx) GetRedirectUrl(state string) string {
	return "https://open.weixin.qq.com/connect/qrconnect?appid=" + e.Conf.ClientId + "&redirect_uri=" + e.Conf.RedirectUrl + "&response_type=code&scope=snsapi_login&state=" + state
}

//获取token
func (e *AuthWx) GetToken(code string) (*AuthWxSuccRes, error) {

	str, err := HttpPost("https://api.weixin.qq.com/sns/oauth2/access_token?appid=" + e.Conf.ClientId + "&secret=" + e.Conf.ClientSecret + "&code=" + code + "&grant_type=authorization_code")
	if err != nil {
		return nil, err
	}

	ismatch, _ := regexp.MatchString("errcode", str)
	if ismatch {

		p := &AuthWxErrRes{}
		err := json.Unmarshal([]byte(str), p)
		if err != nil {
			return nil, err
		}
		return nil, errors.New("Error:" + strconv.Itoa(p.Error) + " Error_description:" + p.ErrorDescription)

	} else {

		p := &AuthWxSuccRes{}
		err := json.Unmarshal([]byte(str), p)
		if err != nil {
			return nil, err
		}
		return p, nil
	}

}

//获取第三方用户信息
func (e *AuthWx) GetUserInfo(accessToken string, openid string) (string, error) {

	str, err := HttpGet("https://api.weixin.qq.com/sns/userinfo?access_token=" + accessToken + "&openid=" + openid)
	if err != nil {
		return "", err
	}
	return string(str), nil
}

//构造微信授权登录
func NewAuthWx(config *AuthConf) *AuthWx {
	return &AuthWx{
		Conf: config,
	}
}
