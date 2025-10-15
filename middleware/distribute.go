package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/new-aspect/nexus-api/common"
	"github.com/new-aspect/nexus-api/model"
	"net/http"
	"strconv"
)

// 思路
// 在请求头分配渠道的auth token，目的是在渠道api的获得调用权限
// 随机或制定拿到渠道key，带入请求头
func Distribute() gin.HandlerFunc {
	return func(c *gin.Context) {
		channel := &model.Channel{}
		channelId, exists := c.Get("channelId")
		if exists {
			id, err := strconv.Atoi(channelId.(string))
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": gin.H{
						"message": "无效的渠道 ID",
						"type":    "one_api_error",
					},
				})
				c.Abort()
				return
			}

			channel, err = model.GetChannelById(id)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": gin.H{
						"message": "无效的渠道 ID",
						"type":    "one_api_error",
					},
				})
				c.Abort()
				return
			}

			if channel.Status != common.ChannelStatusEnabled {
				c.JSON(200, gin.H{
					"error": gin.H{
						"message": "该渠道已被禁用",
						"type":    "one_api_error",
					},
				})
				c.Abort()
				return
			}
		} else {
			// Select a channel for the user
			var err error
			channel, err = model.GetRandomChannel()
			if err != nil {
				c.JSON(200, gin.H{
					"error": gin.H{
						"message": "无可用渠道",
						"type":    "one_api_error",
					},
				})
				c.Abort()
				return
			}
		}
		c.Set("channel", channel.Type)
		c.Request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", channel.Key))
		c.Next()
	}
}
