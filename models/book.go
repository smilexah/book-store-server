package models

import "time"

type Book struct {
	ID          uint      `gorm:"primaryKey"`
	Title       string    `json:"title"`
	ISBN        string    `json:"isbn"`
	Description string    `json:"description"`
	Author      string    `json:"author"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
