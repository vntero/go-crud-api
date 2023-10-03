package models

import "time"

type Post struct {
	ID	uint `json:"id" gorm:"primaryKey"`
	Title	uint `json:"title"`
	Content	uint `json:"content"`
	CreatedAt	time.Time `json:"created_at"`
}