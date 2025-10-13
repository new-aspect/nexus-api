package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/new-aspect/nexus-api/model"
	"strconv"
)

func (v *V1) ApiAddChannel(c *gin.Context) {
	channel := model.Channel{}
	err := c.ShouldBindJSON(&channel)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to unmarshal body " + err.Error()})
		return
	}

	err = channel.Insert()
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to insert channel " + err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "success"})
}

func GetAllChannel(c *gin.Context) {
	channels, err := model.GetAllChannels()
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to get all channels " + err.Error()})
		return
	}

	// 每隔字段保密
	for i, _ := range channels {
		channels[i].Key = secretKey(channels[i].Key)
	}

	c.JSON(200, gin.H{"message": "success", "data": channels})
}

// key的第5为到倒数5位加密
func secretKey(key string) string {
	// 如果key长度小于等于10，则不处理
	if len(key) <= 10 {
		return key
	}

	// 提取前4位
	prefix := key[:4]

	middle := key[4 : len(key)-4]

	suffix := key[len(key)-4:]

	// 对中间部分进行加密（这里使用简单的星号替换，实际应使用真实加密算法）
	encryptedMiddle := ""
	for i := 0; i < len(middle); i++ {
		encryptedMiddle += "*"
	}

	// 组合结果
	return prefix + encryptedMiddle + suffix
}

type UpdateChannelReq struct {
	Name string
}

func UpdateChannel(c *gin.Context) {
	req := UpdateChannelReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to unmarshal body " + err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to unmarshal body " + err.Error()})
		return
	}

	channel := model.Channel{
		ID:   id,
		Name: req.Name,
	}

	err = channel.Update()
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to update channel " + err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "success"})
}
