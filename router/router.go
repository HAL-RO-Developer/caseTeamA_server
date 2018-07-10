package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors)

	r.GET("/log", log)

	// ユーザー登録
	user := r.Group("/")
	userRouter(user)

	// teamA
	work := r.Group("/work")
	workRouter(work)

	// デバイス側処理
	thing := r.Group("/thing")
	thingRouter(thing)

	// 問題登録
	question := r.Group("/question")
	questionRouter(question)
	return r

}

func cors(c *gin.Context) {
	headers := c.Request.Header.Get("Access-Control-Request-Headers")

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,PUT,PATCH,POST,DELETE")
	c.Writer.Header().Set("Access-Control-Allow-Headers", headers)
	if c.Request.Method == "OPTIONS" {
		c.Status(200)
		c.Abort()
	}
}

func log(c *gin.Context) {
	c2 := *c
	s, _ := c2.GetRawData()
	fmt.Println(string(s))
}
