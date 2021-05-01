package models

import "time"

type Todo struct {
	CommonModelFields
	Title       string    `gorm:"not null" json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"dueDate"`
}
