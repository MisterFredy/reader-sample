package seeder

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/reader/models"
)

var posts = []models.Post{
	models.Post{
		Title:   "Title 1",
		Content: "Hello world 1",
	},
	models.Post{
		Title:   "Title 2",
		Content: "Hello world 2",
	},
}

func SeedPost(db *gorm.DB) {
	var err error
	for i, _ := range users {
		posts[i].AuthorID = users[i].ID

		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}
}
