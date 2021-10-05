package entity

import (
	"gorm.io/gorm"
)

// Credit entity known as Credit model and create credits table
type Credit struct {
	ID     uint64 `gorm:"column:id;primaryKey;autoIncrement"`
	Mobile string `gorm:"<-;column:mobile;type:varchar(11);not null;unique;comment:mobile of the user;example:09122222222"`
	Credit uint64 `gorm:"<-;column:credit;type:bigint;not null;default:0;comment:credit amount of the user;example:250000"`
	gorm.Model
}
