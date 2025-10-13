package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/new-aspect/nexus-api/model"
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
	channel, err := model.GetAllChannel()
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to get all channels " + err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "success", "data": channel})
}
