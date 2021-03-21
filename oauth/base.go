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
	authorizeUrl string      //授权登录URL
	TokenUrl     string      //获得令牌URL
	RefreshUrl   string      //刷新令牌URL
	userInfoUrl  string      //获取用户信息URL
	config       *AuthConfig //配置信息
	sourceName   string      //来源名称
}

func (b *BaseRequest) Set(sourceName string, cfg *AuthConfig) {
	b.config = cfg
	b.sourceName = sourceName
}

func (*BaseRequest) GetState(state string) string {
	if state == "" {
		return utils.GetUUID()
	}
	return state
}
