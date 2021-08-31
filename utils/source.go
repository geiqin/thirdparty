package utils

//注册来源
type RegisterSource int32

const (
	RegisterSourceHand   RegisterSource = 9999 //手动添加
	RegisterSourceMobile RegisterSource = 1001 //手机一键登录
	RegisterSourceSms    RegisterSource = 1002 //手机短信
	RegisterSourceWxMini RegisterSource = 2101 //微信小程序
	RegisterSourceWechat RegisterSource = 2102 //微信登录（APP通过微信登录）
	RegisterSourceQQ     RegisterSource = 2103 //QQ登录
	RegisterSourceAlipay RegisterSource = 2201 //支付宝登录
	RegisterSourceDouYin RegisterSource = 2301 //抖音登录
	RegisterSourceWeibo  RegisterSource = 2401 //微博登录
)
