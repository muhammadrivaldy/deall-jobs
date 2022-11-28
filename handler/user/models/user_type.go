package models

import "time"

type UserTypeID int

const (
	UserTypeIDAdmin   UserTypeID = 1
	UserTypeIDRegular UserTypeID = 2
)

type UserType struct {
	ID        int       `gorm:"column:id"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (u *UserType) TableName() string {
	return "mst_user_type"
}
