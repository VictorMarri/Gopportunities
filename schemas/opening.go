package schemas

import (
	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model //Unwrapping gorm model to this struct, indicating that this struct is interesting to gorm use
	Role       string
	Company    string
	Location   string
	Remote     bool
	Link       string
	Salary     float64
}
