package result

//用户信息
type UserResult struct {
	UUID      string `json:"uuid"`
	OpenId    string `json:"openId"`
	UserName  string `json:"user_name"`
	NickName  string `json:"nick_name"`
	Avatar    string `json:"avatar"`
	Company   string `json:"company"`
	Language  string `json:"language"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
	Blog      string `json:"blog"`
	Location  string `json:"location"`
	Mobile    string `json:"mobile"`
	Email     string `json:"email"`
	Remark    string `json:"remark"`
	Url       string `json:"url"`
	Gender    string `json:"gender"`
	Source    string `json:"source"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`

	Token *TokenResult `json:"token"`
}
