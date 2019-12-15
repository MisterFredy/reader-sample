package seeder

import (
	"github.com/jinzhu/gorm"
	"github.com/reader/models"
)

var users = []models.User{
	models.User{
		Nickname: "Steven victor",
		Email:    "steven@gmail.com",
		Password: "password",
	},
	models.User{
		Nickname: "Martin Luther",
		Email:    "luther@gmail.com",
		Password: "password",
	},
}

func upUser(db *gorm.DB) {

}
