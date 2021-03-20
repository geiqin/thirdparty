package oauth


//基本配置
type AuthConf struct {
	ClientId     string
	ClientSecret string
	RedirectUrl  string
}

//@ qq 结构 ------------------------------------------------- start

type AuthQq struct {
	Conf *AuthConf
}
type AuthQqErrRes struct {
	Error            int    `json:"error"`
	ErrorDescription string `json:"error_description"`
}
type AuthQqMe struct {
	ClientId string `json:"client_id"`
	OpenId   string `json:"openid"`
}

//@ qq 结构 ------------------------------------------------- end

//@ weibo 结构 ------------------------------------------------- start

type AuthWb struct {
	Conf *AuthConf
}

type AuthWbErrRes struct {
	Error            int    `json:"error_code"`
	ErrorDescription string `json:"error"`
}

type AuthWbSuccRes struct {
	AccessToken string `json:"access_token"`
	OpenId      string `json:"uid"`
}

//@ weibo 结构 ------------------------------------------------- end

//@ weixin 结构 ------------------------------------------------- start

type AuthWx struct {
	Conf *AuthConf
}

type AuthWxErrRes struct {
	Error            int    `json:"errcode"`
	ErrorDescription string `json:"errmsg"`
}

type AuthWxSuccRes struct {
	AccessToken string `json:"access_token"`
	OpenId      string `json:"openid"`
}

//@ weixin 结构 ------------------------------------------------- end

//@ github 结构 ------------------------------------------------- start

type AuthGithub struct {
	Conf *AuthConf
}

type AuthGithubErrRes struct {
	Error            int    `json:"errcode"`
	ErrorDescription string `json:"errmsg"`
}

type AuthGithubSuccRes struct {
	AccessToken string `json:"access_token"`
	OpenId      string `json:"openid"`
}

//@ github 结构 ------------------------------------------------- end

//@ gitee 结构 ------------------------------------------------- start

type AuthGitee struct {
	Conf *AuthConf
}

type AuthGiteeErrRes struct {
	Error            int    `json:"errcode"`
	ErrorDescription string `json:"errmsg"`
}

type AuthGiteeSuccRes struct {
	AccessToken string `json:"access_token"`
}

//@ gitee 结构 ------------------------------------------------- end
