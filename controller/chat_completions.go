// controller/chat_completions.go
package controller

import (
	"bytes"
	"encoding/json"
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

func ChatCompletions(c *gin.Context) {

	channelType := c.GetInt("channel")
	baseUrl := common.ChannelBaseURLs[channelType]

	// 1. 先把 Body 读到内存
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to read request body " + err.Error()})
		return
	}
	// 别忘了在高并发场景下要关闭 Body
	defer c.Request.Body.Close()

	// 校验
	var req RequestBody
	err = json.Unmarshal(bodyBytes, &req)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to unmarshal body " + err.Error()})
		return
	}

	// 校验看
	if len(req.Messages) == 0 {
		c.JSON(400, gin.H{"error": "message body can't be empty "})
		return
	}

	request, err := buildForwardRequest(bodyBytes, http.MethodPost, fmt.Sprintf("%s%s", baseUrl, c.Request.URL.String()))
	if err != nil {
		c.JSON(500, gin.H{
			"status": 500,
			"err":    err.Error(),
		})
		return
	}
	request.Header = c.Request.Header.Clone()

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		c.JSON(500, gin.H{
			"status": 500,
			"err":    err.Error(),
		})
		return
	}

	c.DataFromReader(resp.StatusCode, resp.ContentLength, "", resp.Body, nil)

}

// 接收原始请求体、目标URL和API Key
// 返回一个构建好的 http.Request 对象，或者一个错误
func buildForwardRequest(bodyBytes []byte, method string, targetUrl string) (*http.Request, error) {
	request, err := http.NewRequest(method, targetUrl, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")

	return request, nil
}
