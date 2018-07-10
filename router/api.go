package router

import (
	. "github.com/HAL-RO-Developer/caseTeamA_server/controller"
	"github.com/HAL-RO-Developer/caseTeamA_server/middleware"
	"github.com/gin-gonic/gin"
)

func userRouter(user *gin.RouterGroup) {
	// ユーザー登録、サインアップ、削除
	user.POST("/signup", User.Create)
	user.POST("/signin", middleware.Login)
	// 子ども情報の登録、取得、削除
	user.POST("/child", User.Child)
	user.GET("/child", User.GetChildren)
	user.DELETE("/child/:child_id", User.DeleteChild)
	// デバイスID発行、取得、削除
	user.POST("/device", Device.CreateNewDevice)
	user.GET("/device", Device.ListDevice)
	user.DELETE("/device/:device_id", Device.DeleteDevice)

	// BOCCOAPI
	user.POST("/bocco", Bocco.RegistBocco)
	user.GET("/bocco", Bocco.GetBoccoInfo)
	user.DELETE("/bocco", Bocco.DeleteBoccoInfo)
}

func workRouter(work *gin.RouterGroup) {
	// 回答記録取得
	work.GET("/graph/:child_id", Record.WorkRecordForGraph)
	work.GET("/detail/:child_id", Record.WorkRecordForDetail)
	work.POST("/message", WorkMessage.EditMessage)
	work.GET("/message/:child_id", WorkMessage.GetMessage)
	work.DELETE("/message/:message_id", WorkMessage.DeleteMessage)
}

func thingRouter(thing *gin.RouterGroup) {
	// デバイスID紐付け
	thing.POST("/registration", Device.DeviceRegistration)
	// ICリーダー
	thing.POST("/reader", Reader.SendTag)
}

func questionRouter(question *gin.RouterGroup) {
	question.POST("/create", Question.CreateQuestion)
	question.POST("/genre", Question.CreateGenre)
	question.GET("/genre", Question.GetGenre)
}
