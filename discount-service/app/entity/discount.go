package entity

import (
	"gorm.io/gorm"
)

// Discount entity known as Discount model and create discounts table
type Discount struct {
	ID         uint64 `gorm:"column:id;primaryKey;autoIncrement"`
	Code       string `gorm:"<-;column:code;type:varchar(255);not null;unique;comment:credit code;example:xyzABC12"`
	CodeCredit uint64 `gorm:"<-;column:code_credit;type:bigint;not null;comment:credit amount of the code;example:1000000"`
	Count      uint   `gorm:"<-;column:count;type:int;not null;comment:count of discount code;example:1000"`
	Used       uint   `gorm:"<-;column:used;type:int;not null;default:0;comment:dount of used discount codes;example:400"`
	gorm.Model
}
