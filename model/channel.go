package model

import "github.com/new-aspect/nexus-api/common"

type Channel struct {
	Id           int    `json:"id"`
	Type         int    `json:"type" gorm:"default:0"`
	Key          string `json:"key" gorm:"not null"`
	Status       int    `json:"status" gorm:"default:1"`
	Name         string `json:"name" gorm:"index"`
	Weight       int    `json:"weight"`
	CreatedTime  int64  `json:"created_time" gorm:"bigint"`
	AccessedTime int64  `json:"accessed_time" gorm:"bigint"`
}

func (c *Channel) Insert() error {
	return DB.Create(c).Error
}

func GetAllChannels() ([]*Channel, error) {
	var channels []*Channel
	err := DB.Find(&channels).Error
	return channels, err
}

func (c *Channel) Update() error {
	return DB.Updates(c).Error
}

func (c *Channel) Delete() error {
	return DB.Delete(c).Error
}

func (c *Channel) HasKey() bool {
	return c.Key != ""
}

func GetChannelById(id int) (*Channel, error) {
	var channel *Channel
	err := DB.Where("id = ?", id).Find(&channel).Error
	return channel, err
}

func GetRandomChannel() (*Channel, error) {
	var channel Channel
	var err error
	if common.UsingSQLite {
		err = DB.Where("status = ?", common.ChannelStatusEnabled).Order("RANDOM()").Limit(1).First(&channel).Error
	} else {
		err = DB.Where("status = ?", common.ChannelStatusEnabled).Order("RAND()").Limit(1).First(&channel).Error
	}

	return &channel, err
}
