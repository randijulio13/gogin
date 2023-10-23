package model

type LogActivity struct {
	ID   string `json:"ID"`
	Name string `json:"name" gorm:"not null"`
}
