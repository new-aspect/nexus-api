package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"

	"os"
)

type RequestBody struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

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
	v1 := V1{apiKey: apiKey, forwardHost: forwardHost}

	router := gin.Default()
	router.POST("/v1/chat/completions", v1.apiV1ChatCompletions)

	if err := router.Run(":3000"); err != nil {
		fmt.Println(err)
	}
}

type V1 struct {
	apiKey      string
	forwardHost string
}

func (v *V1) apiV1ChatCompletions(c *gin.Context) {
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

	request, err := buildForwardRequest(bodyBytes, http.MethodPost, v.forwardHost+"/v1/chat/completions", v.apiKey)
	if err != nil {
		c.JSON(500, gin.H{
			"status": 500,
			"err":    err.Error(),
		})
		return
	}

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
func buildForwardRequest(bodyBytes []byte, method string, targetUrl string, apiKey string) (*http.Request, error) {
	request, err := http.NewRequest(method, targetUrl, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	request.Header.Set("Content-Type", "application/json")
	return request, nil
}
