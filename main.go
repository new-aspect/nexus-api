// main.go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/new-aspect/nexus-api/controller/chat"
	"os"
)

func main() {

	// read env
	apiKey := os.Getenv("API_KEY")
	forwardHost := os.Getenv("FORWARD_HOST")
	if apiKey == "" {
		panic("环境变量 API_KEY 不能为空")
	}
	if forwardHost == "" {
		panic("环境变量 FORWARD_HOST 不能为空")
	}
	chatController := chat.V1{ApiKey: apiKey, ForwardHost: forwardHost}

	router := gin.Default()
	router.POST("/v1/chat/completions", chatController.ApiV1ChatCompletions)

	if err := router.Run(":3000"); err != nil {
		fmt.Println(err)
	}
}
