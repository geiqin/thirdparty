package main

import (
	"fmt"
	"thirdparty/oauth"
)

func main()  {
	wxConf := &oauth.AuthConf{ClientId: "xxx", ClientSecret: "xxx", RedirectUrl: "http://www.change.tm"}

	wxAuth := oauth.NewAuthWx(wxConf)

	fmt.Print(wxAuth.GetRedirectUrl("sate")) //获取第三方登录地址

	wxRes, err := wxAuth.GetToken("code")

	userInfo, _ := wxAuth.GetUserInfo(wxRes.AccessToken, wxRes.Openid)
}