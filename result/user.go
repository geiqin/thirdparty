package result

//用户信息
type UserResult struct {
	UUID      string `json:"uuid"`
	OpenId    string `json:"open_id"`
	UnionId   string `json:"union_id"`
	UserName  string `json:"user_name"`
	NickName  string `json:"nick_name"`
	AvatarUrl string `json:"avatar"`
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
	Source    int32  `json:"source"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`

	Token *TokenResult `json:"token"`
}
