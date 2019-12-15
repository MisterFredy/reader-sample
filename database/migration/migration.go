package migration

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/reader/models"
)

func Migrate(db *gorm.DB) {
	//maybe disini ada yg bisah buat bagian lebih dinamis?untuk migration
	//base migration dari model jadi ketika kita membuat models sekaligus membuat skemanya
	//lebih mudah ketimbang laravel

	err := db.Debug().DropTableIfExists(&models.Post{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Post{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	/*err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "null", "null").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}*/
}
