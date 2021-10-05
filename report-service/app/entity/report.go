package entity

import (
	"gorm.io/gorm"
)

// Report entity known as Report model and create reports table
type Report struct {
	ID         uint64 `gorm:"column:id;primaryKey;autoIncrement"`
	Mobile     string `gorm:"<-;column:mobile;type:varchar(11);not null;uniqueIndex:compositeindex;comment:mobile of the user;example:09122222222"`
	Code       string `gorm:"<-;column:code;type:varchar(255);not null;uniqueIndex:compositeindex;comment:credit code;example:xyzABC12"`
	CodeCredit uint64 `gorm:"<-;column:code_credit;type:bigint;not null;comment:credit amount of the code;example:1000000"`
	gorm.Model
}
