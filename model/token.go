package model

import (
	"errors"
	"github.com/google/uuid"
	"github.com/new-aspect/nexus-api/common"
	"strings"
)

type Token struct {
	Id           int    `json:"id"`
	UserId       int    `json:"user_id"`
	Key          string `json:"key" gorm:"uniqueIndex"`
	Status       int    `json:"status" gorm:"default:1"`
	Name         string `json:"name" gorm:"index" `
	CreatedTime  int64  `json:"created_time" gorm:"bigint"`
	AccessedTime int64  `json:"accessed_time" gorm:"bigint"`
}

func (t *Token) InitKeyIfNotExits() {
	if t.Key != "" {
		return
	}
	t.Key = genUUID()
}

func genUUID() string {
	return uuid.New().String()
}

// 创建Token
func (t *Token) Insert() error {
	return DB.Create(t).Error
}

func GetAllTokens() ([]*Token, error) {
	var tokens []*Token
	err := DB.Find(&tokens).Error
	return tokens, err
}

func (t *Token) Update() error {
	return DB.Updates(t).Error
}

func (t *Token) Delete() error {
	return DB.Delete(t).Error
}

func ValidateUseToken(key string) (*Token, error) {
	if key == "" {
		return nil, errors.New("未提供 token")
	}
	key = strings.TrimPrefix(key, "Bearer ")
	token := &Token{}
	err := DB.Where("key = ?", key).Find(token).Error
	if err != nil {
		return nil, err
	}

	if token.Status != common.TokenStatusEnable {
		return nil, errors.New("该 token 已被禁用")
	}

	go func() {
		token.AccessedTime = common.GetTimestamp()
		if err = token.Update(); err != nil {
			common.SysError("更新 token 访问时间失败：" + err.Error())
		}
	}()

	return token, nil
}
