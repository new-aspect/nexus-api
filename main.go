// main.go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/new-aspect/nexus-api/controller"
	"github.com/new-aspect/nexus-api/middleware"
	"github.com/new-aspect/nexus-api/model"
)

func main() {

	// read env
	//apiKey := os.Getenv("API_KEY")
	//forwardHost := os.Getenv("FORWARD_HOST")
	//if apiKey == "" {
	//	panic("环境变量 API_KEY 不能为空")
	//}
	//if forwardHost == "" {
	//	panic("环境变量 FORWARD_HOST 不能为空")
	//}

	if err := model.InitDB(); err != nil {
		panic("初始化数据库失败 " + err.Error())
	}

	//controllerV1 := controller.V1{ApiKey: apiKey, ForwardHost: forwardHost}

	router := gin.Default()
	router.POST("/v1/chat/completions", middleware.TokenAuth(), middleware.Distribute(), controller.ChatCompletions)
	router.POST("/v1/api/channel", controller.AddChannel)
	router.GET("/v1/api/channel", controller.GetAllChannel)
	router.PUT("/v1/api/channel/:id", controller.UpdateChannel)
	router.DELETE("/v1/api/channel/:id", controller.DeleteChannel)

	router.POST("/v1/api/token", controller.AddToken)
	router.GET("/v1/api/token", controller.GetAllToken)

	if err := router.Run(":3000"); err != nil {
		fmt.Println(err)
	}
}
