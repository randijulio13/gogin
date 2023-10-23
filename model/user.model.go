package model

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/randijulio13/gogin/database"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID          string       `json:"ID" gorm:"not null;primary_key"`
	Password    string       `json:"-" gorm:"not null"`
	Email       string       `json:"email,omitempty" gorm:"default:null"`
	Nip         string       `json:"nip" gorm:"unique;not null"`
	Name        string       `json:"name" gorm:"not null"`
	Position    string       `json:"position" gorm:"not null"`
	Permissions []Permission `json:"permissions" gorm:"many2many:user_permissions"`
	gorm.Model
}

func (u *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (u *User) SetPassword(password string) {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		panic("failed set password.")
	}
	u.Password = hashedPassword
}

func (u *User) SetPermissions(permissions []string) {
	GetPermissions(&u.Permissions, permissions)
}

func (u *User) GenerateToken() (string, error) {
	var permissions []string
	for _, permission := range u.Permissions {
		permissions = append(permissions, permission.Name)
	}
	claims := jwt.MapClaims{
		"nip":         u.Nip,
		"name":        u.Name,
		"permissions": permissions,
		"exp":         time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := os.Getenv("JWT_SECRET_KEY")
	tokenString, err := token.SignedString([]byte(secretKey))

	return tokenString, err
}

func (u *User) CreateUser() (*User, error) {
	u.ID = uuid.NewString()
	err := database.DB.Create(&u).Error
	if err != nil {
		return u, err
	}
	return u, nil
}

func (u *User) UpdateUser() (*User, error) {
	if len(u.Permissions) != 0 {
		database.DB.Model(&u).Association("Permissions").Replace(u.Permissions)
	}
	err := database.DB.Save(&u).Error
	if err != nil {
		return u, err
	}
	return u, nil
}

func (u *User) GetById(id string) *gorm.DB {
	return database.DB.Preload("Permissions").First(&u, "id = ?", id)
}

func (u *User) GetByNip(nip string) *gorm.DB {
	return database.DB.Preload("Permissions").Where("nip = ?", nip).Find(&u)
}

func (u *User) GetByEmail(email string) *gorm.DB {
	return database.DB.Where("email = ?", email).Find(&u)
}

func (u *User) GetByEmailOrNip(email string, nip string) *gorm.DB {
	return database.DB.Where("email = ?", email).Or("nip = ?", nip).Find(&u)
}

func (u *User) SavePermissions(permissions *[]Permission) *User {
	database.DB.Model(&u).Association("Permissions").Append(permissions)
	return u
}

func (u *User) Can(p string) bool {
	hasPermission := false
	for _, permission := range u.Permissions {
		if permission.Name == p {
			hasPermission = true
			break
		}
	}

	return hasPermission
}
