// controller/chat_completions.go
package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/new-aspect/nexus-api/common"
	"io"
	"net/http"
)

type V1 struct {
	ApiKey      string
	ForwardHost string
}

type RequestBody struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// 我这边写思路，这里的reply只做转发，我记得是用io.Copy进行转发，转发前用http生成一个请求，复制请求头
// 这个思路不对，应该改成这样的思路
// 拿到请求的参数->按请求参数转发->拿到转发的响应结果->将结果拼装回本次请求

func ChatCompletions(c *gin.Context) {
	chanType := c.GetInt("channel")
	baseUrl := common.ChannelBaseURLs[chanType]

	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", baseUrl, c.Request.URL.String()), c.Request.Body)
	if err != nil {
		c.JSON(500, gin.H{"error": "无法构造请求" + err.Error()})
		return
	}
	request.Header = c.Request.Header

	client := http.Client{}
	response, err := client.Do(request)

	_, err = io.Copy(c.Writer, response.Body)
	if err != nil {
		c.JSON(500, gin.H{"error": "转发报错" + err.Error()})
		return
	}
}
