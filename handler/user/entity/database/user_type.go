package database

import (
	IUser "backend/handler/user"
	"backend/handler/user/models"

	"gorm.io/gorm"
)

type userType struct {
	dbGorm *gorm.DB
}

func NewUserTypeRepo(dbGorm *gorm.DB) IUser.IUserTypeRepo {
	return userType{dbGorm: dbGorm}
}

func (u userType) SelectUserTypeByID(id int) (res models.UserType, err error) {
	err = u.dbGorm.Where("id = ?", id).First(&res).Error
	return
}
