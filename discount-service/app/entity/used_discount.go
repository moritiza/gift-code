package entity

import (
	"gorm.io/gorm"
)

// UsedDiscount entity known as UsedDiscount model and create used_discounts table
type UsedDiscount struct {
	ID     uint64 `gorm:"column:id;primaryKey;autoIncrement"`
	Code   string `gorm:"<-;column:code;type:varchar(255);not null;comment:credit code;example:xyzABC12"`
	Mobile string `gorm:"<-;column:mobile;type:varchar(11);not null;comment:mobile of the user;example:09122222222"`
	gorm.Model
}
