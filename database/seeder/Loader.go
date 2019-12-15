package seeder

import (
	"github.com/jinzhu/gorm"
)

func Seedload(db *gorm.DB) {
	SeedPost(db)
	SeedUser(db)
}
