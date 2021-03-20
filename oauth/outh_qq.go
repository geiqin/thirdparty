package oauth

import (
	"encoding/json"
	"errors"
	"regexp"
	"strconv"
)

//获取登录地址
func (e *AuthQq) GetRedirectUrl(state string) string {
	return "https://graph.qq.com/oauth2.0/authorize?response_type=code&client_id=" + e.Conf.ClientId + "&redirect_uri=" + e.Conf.RedirectUrl + "&state=" + state
}

//获取token
func (e *AuthQq) GetToken(code string) (string, error) {

	str, err := HttpGet("https://graph.qq.com/oauth2.0/token?grant_type=authorization_code&client_id=" + e.Conf.ClientId + "&client_secret=" + e.Conf.ClientSecret + "&code=" + code + "&redirect_uri=" + e.Conf.RedirectUrl)
	if err != nil {
		return "", err
	}

	ismatch, _ := regexp.MatchString("error", str)
	if ismatch {
		re, _ := regexp.Compile("({.*})")
		newres := re.FindStringSubmatch(str)
		errstr := newres[0]
		p := &AuthQqErrRes{}
		err := json.Unmarshal([]byte(errstr), p)
		if err != nil {
			return "", err
		}
		return "", errors.New("Error:" + strconv.Itoa(p.Error) + " Error_description:" + p.ErrorDescription)

	} else {
		re, _ := regexp.Compile("access_token=(.*)&expires_in")
		newres := re.FindStringSubmatch(str)
		if len(newres) >= 2 {
			return newres[1], nil
		}
		return "", nil
	}

}

//获取第三方id
func (e *AuthQq) GetMe(accessToken string) (*AuthQqMe, error) {

	str, err := HttpGet("https://graph.qq.com/oauth2.0/me?access_token=" + accessToken)
	if err != nil {
		return nil, err
	}
	ismatch, _ := regexp.MatchString("error", str)
	if ismatch {
		re, _ := regexp.Compile("({.*})")
		newres := re.FindStringSubmatch(str)
		errstr := newres[0]
		p := &AuthQqErrRes{}
		err := json.Unmarshal([]byte(errstr), p)
		if err != nil {
			return nil, err
		}
		return nil, errors.New("Error:" + strconv.Itoa(p.Error) + " Error_description:" + p.ErrorDescription)

	} else {
		re, _ := regexp.Compile("({.*})")
		newres := re.FindStringSubmatch(str)
		errstr := newres[0]
		p := &AuthQqMe{}
		err := json.Unmarshal([]byte(errstr), p)
		if err != nil {
			return nil, err
		}

		return p, nil
	}

}

//获取第三方用户信息
func (e *AuthQq) GetUserInfo(accessToken string, openid string) (string, error) {

	str, err := HttpGet("https://graph.qq.com/user/get_user_info?access_token=" + accessToken + "&oauth_consumer_key=" + e.Conf.ClientId + "&openid=" + openid)
	if err != nil {
		return "", err
	}
	return string(str), nil

}

//构造QQ授权登录
func NewAuthQq(config *AuthConf) *AuthQq {
	return &AuthQq{
		Conf: config,
	}
}
