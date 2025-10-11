package model

type Channel struct {
	ID int

	Name string
	Key  string
}

func (c *Channel) Insert() error {
	return DB.Create(c).Error
}
