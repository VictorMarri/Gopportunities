package schemas

import (
	"gorm.io/gorm"
	"time"
)

type Opening struct {
	gorm.Model //Unwrapping gorm model to this struct, indicating that this struct is interesting to gorm use
	Role       string
	Company    string
	Location   string
	Remote     bool
	Link       string
	Salary     int64
}

// Data structure as response from database query
type OpeningResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	DeletedAt time.Time `json:"deletedAt,omitempty"` //If the field came empty, don`t show him
	Role      string    `json:"role"`
	Company   string    `json:"company"`
	Location  string    `json:"location"`
	Remote    bool      `json:"remote"`
	Link      string    `json:"link"`
	Salary    int64     `json:"salary"`
}
