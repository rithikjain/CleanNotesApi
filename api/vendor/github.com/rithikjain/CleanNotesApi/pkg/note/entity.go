package note

import "github.com/jinzhu/gorm"

type Note struct {
	gorm.Model
	UserID float64 `json:"user_id"`
	Title  string  `json:"title"`
	Body   string  `json:"body"`
}
