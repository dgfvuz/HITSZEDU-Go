package model

import (
	"hitszedu-go/database"
	"hitszedu-go/entity"
)

//添加一个advice
func AddAdvice(advice *entity.Advice) error {
	return database.Create(advice)
}

//使用userID获取identity
func GetIdentity(userID string) string {
	user := new(entity.User)
	err := database.First(user, "user_id = ?", userID)
	if err != nil {
		return ""
	} else {
		return user.Identity
	}
}
