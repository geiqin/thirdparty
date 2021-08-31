package oauth

import (
	"errors"
	"github.com/geiqin/thirdparty/result"
	"github.com/geiqin/thirdparty/utils"
)

//微信授权登录（第三方应用）
type AuthWxWechat struct {
	BaseRequest
}

func NewAuthWxWechat(conf *AuthConfig) *AuthWxWechat {
	authRequest := &AuthWxWechat{}
	authRequest.Set(utils.RegisterSourceWechat, conf)

	authRequest.authorizeUrl = "https://open.weixin.qq.com/connect/qrconnect"
	authRequest.TokenUrl = "https://api.weixin.qq.com/sns/oauth2/access_token"
	authRequest.userInfoUrl = "https://api.weixin.qq.com/sns/userinfo"

	return authRequest
}

//获取登录地址
func (a *AuthWxWechat) GetRedirectUrl(state string) (*result.CodeResult, error) {
	url := utils.NewUrlBuilder(a.authorizeUrl).
		AddParam("response_type", "code").
		AddParam("appid", a.config.ClientId).
		AddParam("redirect_uri", a.config.RedirectUrl).
		AddParam("scope", "snsapi_login").
		AddParam("state", a.GetState(state)).
		Build()

	_, err := utils.Post(url)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

//获取token
func (a *AuthWxWechat) GetToken(code string) (*result.TokenResult, error) {
	url := utils.NewUrlBuilder(a.TokenUrl).
		AddParam("grant_type", "authorization_code").
		AddParam("code", code).
		AddParam("appid", a.config.ClientId).
		AddParam("secret", a.config.ClientSecret).
		AddParam("redirect_uri", a.config.RedirectUrl).
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
		AccessToken:  m["access_token"],
		RefreshToken: m["refresh_token"],
		ExpireIn:     m["expires_in"],
		OpenId:       m["openid"],
		UnionId:      m["unionid"],
		Scope:        m["scope"],
		TokenType:    m["token_type"],
	}
	if token.AccessToken == "" {
		return nil, errors.New("获取AccessToken数据为空！")
	}
	return token, nil
}

//获取第三方用户信息
func (a *AuthWxWechat) GetUserInfo(accessToken string, openId string) (*result.UserResult, error) {
	url := utils.NewUrlBuilder(a.userInfoUrl).
		AddParam("openid", openId).
		AddParam("access_token", accessToken).
		Build()

	body, err := utils.Get(url)
	if err != nil {
		return nil, err
	}
	m := utils.JsonToMSS(body)
	if _, ok := m["error"]; ok {
		return nil, errors.New(m["error_description"])
	}
	user := &result.UserResult{
		OpenId:    m["openid"],
		UnionId:   m["unionid"],
		UserName:  m["nickname"],
		NickName:  m["nickname"],
		AvatarUrl: m["headimgurl"],
		City:      m["city"],
		Province:  m["province"],
		Country:   m["country"],
		Language:  m["language"],
		Source:    a.registerSource,
		Gender:    utils.GetRealGender("sex").Desc,
	}
	if user.OpenId == "" {
		return nil, errors.New("获取用户信息为空！")
	}
	return user, nil
}
