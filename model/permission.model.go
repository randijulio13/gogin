package model

import (
	"github.com/google/uuid"
	"github.com/randijulio13/gogin/database"
	"gorm.io/gorm"
)

type Permission struct {
	ID   string `json:"ID"`
	Name string `json:"name" gorm:"unique;not null"`
	// Users []*User `json:"users,omitempty" gorm:"many2many:user_permissions"`
	gorm.Model
}

func (p *Permission) SavePermission() (*Permission, error) {
	p.ID = uuid.NewString()
	err := database.DB.Save(&p).Error
	if err != nil {
		return &Permission{}, err
	}
	return p, nil
}

func GetPermissions(p *[]Permission, pString []string) *gorm.DB {
	return database.DB.Where("name IN ?", pString).Find(&p)
}
