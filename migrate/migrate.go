package migrate

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/randijulio13/gogin/database"
	"github.com/randijulio13/gogin/model"
)

func Run() {
	user()
	permission()
}

func Seed() {
	permissions := []model.Permission{{ID: uuid.NewString(), Name: "user_management"}, {ID: uuid.NewString(), Name: "search_by_nik"}, {ID: uuid.NewString(), Name: "search_combination"}}
	database.DB.Save(&permissions)
	fmt.Println("permission seed successful")

	user := model.User{
		ID:       uuid.NewString(),
		Name:     "Admin",
		Nip:      "1234",
		Position: "Super Admin",
	}

	user.SetPassword("1234")
	user.SetPermissions([]string{"user_management", "search_by_nik", "search_combination"})
	user.CreateUser()
	fmt.Println("user seed successful")
}

func permission() {
	database.DB.AutoMigrate(&model.Permission{})
}

func user() {
	database.DB.AutoMigrate(&model.User{})
}
