package model

import "time"

type User struct {
	ID        uint `gorm:"primarykey"`
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
