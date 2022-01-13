package service

import (
	"fmt"
	"hitszedu-go/config"
	"hitszedu-go/util"
	"sync"

	"github.com/robfig/cron"
)

var lock *sync.RWMutex = new(sync.RWMutex)
var accessToken [3]string

// 开启微信Token服务
func WxStart() {
	ReflashAccessToken("public")
	ReflashAccessToken("student")
	ReflashAccessToken("admin")
	c := cron.New()
	spec := "@hourly"
	c.AddFunc(spec, func() {
		ReflashAccessToken("public")
		ReflashAccessToken("student")
		ReflashAccessToken("admin")
	})
	c.Start()
}

type AccessTokenRes struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
}

// 获取identity的序号
func getIdentityNum(identity string) int {
	switch identity {
	case "public":
		return 0
	case "student":
		return 1
	case "admin":
		return 2
	default:
		return -1
	}
}

//刷新某端的AccessToken
func ReflashAccessToken(identity string) (success bool) {
	lock.Lock()
	defer lock.Unlock()
	identityNum := getIdentityNum(identity)
	if identityNum != -1 {
		return false
	}
	url := config.GetString("wx.getAccessToken")
	appid := config.GetString(fmt.Sprintf("wx.%sAppID", identity))
	secret := config.GetString(fmt.Sprintf("wx.%sAppID", identity))
	res := new(AccessTokenRes)
	err := util.GetObj(fmt.Sprintf("%s?grant_type=client_credential&appid=%s&secret=%s", url, appid, secret), res)
	if err != nil {
		return false
	}
	if res.ErrCode == 0 {
		accessToken[identityNum] = res.AccessToken
	} else {
		fmt.Println(res.ErrMsg)
		return false
	}
	return true
}

type PhoneNumberReq struct {
	Code string `json:"code"`
}

type PhoneNumberRes struct {
	ErrCode   int `json:"errcode"`
	PhoneInfo struct {
		PurePhoneNumber string `json:"purePhoneNumber"`
	} `json:"phone_info"`
}

// 调用getPhoneNumber接口,获取phoneNumber
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/phonenumber/phonenumber.getPhoneNumber.html
func GetPhoneNumber(identity string, code string) (number string, success bool) {
	lock.RLock()
	defer lock.RUnlock()
	identityNum := getIdentityNum(identity)
	req := &PhoneNumberReq{Code: code}
	_url := config.GetString("wx.getPhoneNumber")
	url := fmt.Sprintf("%s?access_token=%s", _url, accessToken[identityNum])
	obj := new(PhoneNumberRes)
	err := util.PostObj(url, req, "application/json", obj)
	if err != nil || obj.ErrCode != 0 {
		return "", false
	}
	return obj.PhoneInfo.PurePhoneNumber, true
}

type OpenidRes struct {
	ErrCode int    `json:"errcode"`
	Openid  string `json:"openid"`
}

// 调用code2session接口,获取openid
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/login/auth.code2Session.html
func GetOpenid(identity string, code string) (number string, success bool) {
	_url := config.GetString("wx.code2Session")
	appid := config.GetString(fmt.Sprintf("wx.%sAppID", identity))
	secret := config.GetString(fmt.Sprintf("wx.%sAppID", identity))
	url := fmt.Sprintf("%s?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", _url, appid, secret, code)
	obj := new(OpenidRes)
	err := util.GetObj(url, obj)
	if err != nil || obj.ErrCode != 0 {
		return "", false
	}
	return obj.Openid, true
}
