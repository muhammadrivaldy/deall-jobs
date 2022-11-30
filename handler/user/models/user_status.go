package models

import "time"

type UserStatusID int

const (
	UserStatusIDNonActive UserStatusID = -1
	UserStatusIDActive    UserStatusID = 1
)

type UserStatus struct {
	ID        int       `gorm:"column:id"`
	Key       string    `gorm:"column:key"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (u *UserStatus) TableName() string {
	return "mst_user_status"
}
