package user

import (
	"hitszedu-go/entity"
	"hitszedu-go/middleware"
	"hitszedu-go/model"
	"hitszedu-go/service"

	"github.com/gin-gonic/gin"
)

// 微信小程序登录
func Wxlogin(c *gin.Context) {
	code := c.Query("code")
	identity := c.Query("identity")
	openid, ok := service.GetOpenid(identity, code)
	if !ok {
		c.JSON(200, gin.H{
			"state": "getOpenidFail",
			"data":  "",
		})
	}
	user := model.GetUserByOpenid(openid, identity)
	if user == nil {
		c.JSON(200, gin.H{
			"state": "notExist",
			"data": gin.H{
				"openid": openid,
			},
		})
	} else {
		token, _ := middleware.GenerateToken(user.UserID)
		user.Filter()
		c.JSON(200, gin.H{
			"state": "success",
			"data": gin.H{
				"openid":        openid,
				"authorization": token,
				"user":          user,
			},
		})
	}
}

// 微信小程序注册
func WxRegister(c *gin.Context) {
	// 验证请求合法性
	openid := c.Query("openid")
	code := c.Query("code")
	identity := c.Query("identity")
	hitszID := c.Query("hitszID")
	hitszPassword := c.Query("hitszPassword")
	if identity != "public" && identity != "student" && identity != "admin" {
		c.JSON(200, gin.H{
			"state": "identityError",
			"data":  "",
		})
	}
	//获取手机号
	tel, ok := service.GetPhoneNumber(identity, code)
	if !ok {
		c.JSON(200, gin.H{
			"state": "getPhoneNumberFail",
			"data":  "",
		})
	}
	user := model.GetUserByTel(tel, identity)
	if user == nil {
		//注册
		user = entity.NewUser(tel, "", openid, identity, "normal")
		var err error
		if identity == "public" {
			err = model.AddUser(user)
		} else {
			if !service.HitszLogin(hitszID, hitszPassword) {
				c.JSON(200, gin.H{
					"state": "hitszError",
					"data":  "",
				})
				return
			}
			if identity == "student" {
				err = model.AddStudent(user, hitszID)
			} else {
				err = model.AddAdmin(user, hitszID)
			}
		}
		if err != nil {
			c.JSON(200, gin.H{
				"state": "dbError",
				"data":  "",
			})
			return
		}
		token, _ := middleware.GenerateToken(user.UserID)
		user.Filter()
		c.JSON(200, gin.H{
			"state": "notExist",
			"data": gin.H{
				"authorization": token,
				"user":          user,
			},
		})
	} else {
		//登录
		user.Openid = openid
		err := model.UpdateUser(user)
		if err != nil {
			c.JSON(200, gin.H{
				"state": "dbError",
				"data":  "",
			})
		}
		token, _ := middleware.GenerateToken(user.UserID)
		c.JSON(200, gin.H{
			"state": "success",
			"data": gin.H{
				"authorization": token,
				"user":          user,
			},
		})
	}
}
