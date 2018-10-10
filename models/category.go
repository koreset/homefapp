package models

type Category struct {
	ID    uint `gorm:"primary_key"`
	Name  string
	Posts []Post
}
