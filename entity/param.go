package entity

//请求参数
type RequestParam struct {
	Code  string `json:"code"`
	State string `json:"state"`

	// alipay
	AuthCode string `json:"auth_code"`

	//wechat-mini
	Signature     string `json:"signature"`
	RawData       string `json:"raw_data"`
	EncryptedData string `json:"encrypted_data"`
	Iv            string `json:"iv"`
	IsPull        bool   `json:"is_pull"`
	PullType      string `json:"pull_type"`

	// huawei
	AuthorizationCode string `json:"authorization_code"`

	// twitter
	OauthToken    string `json:"oauthToken"`
	OauthVerifier string `json:"oauthVerifier"`
}
