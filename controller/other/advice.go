package other

import (
	"fmt"
	"hitszedu-go/entity"
	"hitszedu-go/middleware"
	"hitszedu-go/model"
	"hitszedu-go/util"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// 添加一个建议
func AddAdvice(c *gin.Context) {

	//先接收从前端来的内容

	device := c.GetHeader("device")
	adviceDetail := c.GetHeader("adviceDetail")
	time := util.TimeStamp()
	token := c.GetHeader("token")
	userID, _ := middleware.ParseToken(token)
	identity := model.GetIdentity(userID)
	//获得所有需要的字段
	//建立对象advice
	advice := &entity.Advice{
		AdviceID:     uuid.NewV4().String(),
		UserID:       userID,
		Device:       device,
		AdviceDetail: adviceDetail,
		Time:         fmt.Sprintf("%d", time),
		Identity:     identity,
	}

	//再将advice交给modles处理
	model.AddAdvice(advice)
}
