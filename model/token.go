package model

import "github.com/google/uuid"

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
