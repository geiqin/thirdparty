package oauth

import (
	"encoding/json"
	"errors"
	"github.com/geiqin/thirdparty/result"
	"github.com/geiqin/thirdparty/utils"
	"github.com/xlstudio/wxbizdatacrypt"
)

//微信小程序授权登录（微信小程序）
type AuthWxMini struct {
	BaseRequest
}

func NewAuthWxMini(conf *AuthConfig) *AuthWxMini {
	authRequest := &AuthWxMini{}
	authRequest.Set(utils.RegisterSourceWxMini, conf)

	authRequest.TokenUrl = "https://api.weixin.qq.com/sns/jscode2session"
	authRequest.RefreshUrl = "https://api.weixin.qq.com/sns/jscode2session"
	authRequest.AccessTokenUrl = "https://api.weixin.qq.com/cgi-bin/token"

	return authRequest
}

//获取小程序会话token，一般返回 sessionKey OpenId UnionId
func (a *AuthWxMini) GetSessionKey(code string) (*result.TokenResult, error) {
	url := utils.NewUrlBuilder(a.TokenUrl).
		AddParam("grant_type", "authorization_code").
		AddParam("appid", a.config.ClientId).
		AddParam("secret", a.config.ClientSecret).
		AddParam("js_code", code).
		Build()

	body, err := utils.Post(url)

	if err != nil {
		return nil, err
	}
	//log.Println("body:", body)
	m := utils.JsonToMSS(body)
	if _, ok := m["errcode"]; ok {
		return nil, errors.New(m["errmsg"])
	}

	token := &result.TokenResult{
		SessionKey: m["session_key"],
		OpenId:     m["openid"],
		UnionId:    m["unionid"],
	}
	return token, nil
}

//获取小程序全局唯一后台接口调用凭据（access_token）
func (a *AuthWxMini) GetAccessToken() (*result.TokenResult, error) {
	url := utils.NewUrlBuilder(a.AccessTokenUrl).
		AddParam("grant_type", "client_credential").
		AddParam("appid", a.config.ClientId).
		AddParam("secret", a.config.ClientSecret).
		Build()

	body, err := utils.Post(url)
	if err != nil {
		return nil, err
	}
	m := utils.JsonToMSS(body)
	if _, ok := m["error"]; ok {
		return nil, errors.New(m["error_description"])
	}
	token := &result.TokenResult{
		AccessToken: m["access_token"],
		ExpireIn:    m["expires_in"],
	}
	return token, nil
}

//获取用户信息
func (a *AuthWxMini) GetUserInfo(sessionKey string, encryptedData string, iv string) (*result.UserResult, error) {
	pc := wxbizdatacrypt.WxBizDataCrypt{AppId: a.config.ClientId, SessionKey: sessionKey}
	ret, err := pc.Decrypt(encryptedData, iv, true) //第三个参数解释： 需要返回 JSON 数据类型时 使用 true, 需要返回 map 数据类型时 使用 false

	if err != nil {
		return nil, err
	}

	m := utils.JsonToMSS(ret.(string))
	if _, ok := m["error"]; ok {
		return nil, errors.New(m["error_description"])
	}

	user := &result.UserResult{
		OpenId:    m["openId"],
		UserName:  m["nickName"],
		NickName:  m["nickName"],
		AvatarUrl: m["avatarUrl"],
		City:      m["city"],
		Province:  m["province"],
		Country:   m["country"],
		Source:    a.registerSource,
		Gender:    utils.GetRealGender(m["gender"]).Desc,
	}

	return user, nil
}

//获取手机号码
func (a *AuthWxMini) GetMobileNumber(sessionKey string, encryptedData string, iv string) (*result.WxMobileResult, error) {
	pc := wxbizdatacrypt.WxBizDataCrypt{AppId: a.config.ClientId, SessionKey: sessionKey}
	ret, err := pc.Decrypt(encryptedData, iv, true) //第三个参数解释： 需要返回 JSON 数据类型时 使用 true, 需要返回 map 数据类型时 使用 false

	if err != nil {
		return nil, err
	}

	obj := &result.WxMobileResult{}
	wxErr := json.Unmarshal([]byte(ret.(string)), obj)

	if wxErr != nil {
		return nil, wxErr
	}
	return obj, nil
}
