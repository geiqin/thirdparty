# thirdparty

 Go语言实现的第三方授权登录，整合QQ、微信、微信小程序、微博、抖音、支付宝等第三方平台的授权登录

### 微信授权登录
```go
package main

import (
	"fmt"
	"github.com/geiqin/thirdparty/oauth"
)

func main()  {
	wxConf := &oauth.AuthConf{
        ClientId: "your app_id", 
        ClientSecret: "your app_secret", 
        RedirectUrl: "http://www.geiqin.com"}

	wxAuth := oauth.NewAuthWxWechat(wxConf)
	fmt.Print(wxAuth.GetRedirectUrl("sate")) //获取第三方登录地址
	wxRes, err := wxAuth.GetToken("code")
	userInfo, _ := wxAuth.GetUserInfo(wxRes.AccessToken, wxRes.OpenId)
}
``` 