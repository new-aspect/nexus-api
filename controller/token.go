package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/new-aspect/nexus-api/model"
)

func AddToken(c *gin.Context) {
	token := model.Token{}
	err := c.ShouldBindJSON(&token)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to unmarshal body " + err.Error()})
		return
	}

	token.InitKeyIfNotExits()

	err = token.Insert()
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to insert channel " + err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "success"})
}

func GetAllToken(c *gin.Context) {
	tokens, err := model.GetAllTokens()
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to get all tokens " + err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "success", "data": tokens})
}
