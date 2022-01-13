package entity

import (
	"crypto/rand"
	"fmt"
	"hitszedu-go/config"
	"math/big"

	uuid "github.com/satori/go.uuid"
)

type User struct {
	UserID     string `json:"userID" gorm:"primary_key"`
	Nickname   string `json:"nickname"`
	AvatarUrl  string `json:"avatarUrl"`
	Motto      string `json:"motto"`
	Tel        string `json:"tel"`
	Password   string `json:"password"`
	Openid     string `json:"openid"`
	Identity   string `json:"identity"`
	Permission string `json:"permission"`
}

func NewUser(tel string, password string, openid string, identity string, permission string) *User {
	_randInt, _ := rand.Int(rand.Reader, big.NewInt(999999))
	randInt := _randInt.String()
	return &User{
		UserID:     uuid.NewV4().String(),
		Nickname:   fmt.Sprintf("用户%s", randInt),
		AvatarUrl:  config.GetString("entity.user.avatar"),
		Motto:      config.GetString("entity.user.motto"),
		Tel:        tel,
		Password:   password,
		Openid:     openid,
		Identity:   identity,
		Permission: permission,
	}
}

// 过滤user的私密信息
func (user *User) Filter() {
	user.UserID = ""
	user.Openid = ""
	user.Password = ""
	user.Permission = ""
	user.Tel = ""
}
