package schema

import (
	"time"

	"gorm.io/gorm"
)

type Base struct {
	Id        int       `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	CreatedBy string
	UpdatedBy string
	DeletedAt gorm.DeletedAt
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		User{},
	)
}
