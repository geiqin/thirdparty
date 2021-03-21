package main

import (
	"fmt"
	"github.com/geiqin/thirdparty/config"
	"github.com/geiqin/thirdparty/oauth"
	"log"
)

func main() {
	wxConf := &config.AuthConfig{ClientId: "xxx", ClientSecret: "xxx", RedirectUrl: "http://www.geiqin.com"}

	wxAuth := oauth.NewAuthWxWechat(wxConf)

	fmt.Print(wxAuth.GetRedirectUrl("sate")) //获取第三方登录地址

	wxRes, err := wxAuth.GetToken("code")

	userInfo, _ := wxAuth.GetUserInfo(wxRes.AccessToken, wxRes.OpenId)

	log.Println("ssss:", err, userInfo)
}
