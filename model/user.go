package model

import (
	"errors"
	"hitszedu-go/database"
	"hitszedu-go/entity"
)

// 使用openid获取user
func GetUserByOpenid(openid string, identity string) *entity.User {
	user := new(entity.User)
	err := database.First(user, &entity.User{Openid: openid, Identity: identity})
	if err != nil {
		return nil
	} else {
		return user
	}
}

// 使用tel获取user
func GetUserByTel(tel string, identity string) *entity.User {
	user := new(entity.User)
	err := database.First(user, &entity.User{Tel: tel, Identity: identity})
	if err != nil {
		return nil
	} else {
		return user
	}
}

// 使用userid获取user
func GetUser(userID string) *entity.User {
	user := new(entity.User)
	err := database.First(user, "user_id = ?", userID)
	if err != nil {
		return nil
	} else {
		return user
	}
}

// 添加一个user
func AddUser(user *entity.User) error {
	return database.Create(user)
}

// 更新一个user
func UpdateUser(user *entity.User) error {
	return database.Save(user)
}

// 添加一个学生
func AddStudent(user *entity.User, hitszID string) error {
	tx := database.Begin()
	student := entity.NewStudent(hitszID, user.UserID)
	//TODO 添加默认简历
	err := tx.Create(user).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Create(student).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// 添加一个管理员
func AddAdmin(user *entity.User, hitszID string) error {
	admin := entity.NewAdmin("", "")
	err := database.First(admin, "admin_id = ?", hitszID)
	if err != nil {
		return err
	}
	if admin.UserID != "" {
		return errors.New("userExist")
	}
	admin.UserID = user.UserID
	tx := database.Begin()
	err = tx.Create(user).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Save(admin).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
