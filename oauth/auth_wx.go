package oauth

import (
	"errors"
	"github.com/geiqin/thirdparty/config"
	"github.com/geiqin/thirdparty/result"
	"github.com/geiqin/thirdparty/utils"
)

//微信授权登录
type AuthWx struct {
	BaseRequest
}

func NewAuthWx(conf *config.AuthConfig) *AuthWx {
	authRequest := &AuthWx{}
	authRequest.Set("weixin", conf)

	authRequest.authorizeUrl = "https://open.weixin.qq.com/connect/qrconnect"
	authRequest.TokenUrl = "https://api.weixin.qq.com/sns/oauth2/access_token"
	authRequest.userInfoUrl = "https://api.weixin.qq.com/sns/userinfo"

	return authRequest
}

//获取登录地址
func (a *AuthWx) GetRedirectUrl(state string) (*result.CodeResult, error) {
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
func (a *AuthWx) GetToken(code string) (*result.TokenResult, error) {
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
		Scope:        m["scope"],
		TokenType:    m["token_type"],
	}
	return token, nil
}

//获取第三方用户信息
func (a *AuthWx) GetUserInfo(openId string, accessToken string) (*result.UserResult, error) {
	url := utils.NewUrlBuilder(a.TokenUrl).
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
		UUID:      m["id"],
		UserName:  m["login"],
		NickName:  m["name"],
		Avatar:    m["avatar_url"],
		Company:   m["company"],
		Blog:      m["blog"],
		Location:  m["location"],
		Email:     m["email"],
		Remark:    m["bio"],
		Url:       m["html_url"],
		CreatedAt: m["created_at"],
		UpdatedAt: m["updated_at"],
		Source:    a.sourceName,
		Gender:    utils.GetRealGender("").Desc,
	}
	return user, nil
}
