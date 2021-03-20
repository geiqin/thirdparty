# thirdparty
 Go语言实现的第三方登录开源库，整合QQ、微信、微博、Github等第三方平台的授权登录
package main

import (
	"fmt"
	"github.com/geiqin/thirdparty/oauth"
)

func main()  {
	wxConf := &oauth.AuthConf{ClientId: "xxx", ClientSecret: "xxx", RedirectUrl: "http://www.change.tm"}

	wxAuth := oauth.NewAuthWx(wxConf)

	fmt.Print(wxAuth.GetRedirectUrl("sate")) //获取第三方登录地址

	wxRes, err := wxAuth.GetToken("code")

	userInfo, _ := wxAuth.GetUserInfo(wxRes.AccessToken, wxRes.OpenId)
}