package oauth

import (
	"github.com/geiqin/thirdparty/utils"
)

//基本配置
type AuthConfig struct {
	ClientId     string
	ClientSecret string
	RedirectUrl  string
}

type BaseRequest struct {
	authorizeUrl   string      //授权登录URL
	TokenUrl       string      //获得令牌URL
	AccessTokenUrl string      //获得访问令牌URL
	RefreshUrl     string      //刷新令牌URL
	userInfoUrl    string      //获取用户信息URL
	config         *AuthConfig //配置信息
	registerSource int32       //注册来源
}

func (b *BaseRequest) Set(sourceName utils.RegisterSource, cfg *AuthConfig) {
	b.config = cfg
	b.registerSource = int32(sourceName)
}

func (*BaseRequest) GetState(state string) string {
	if state == "" {
		return utils.GetUUID()
	}
	return state
}
