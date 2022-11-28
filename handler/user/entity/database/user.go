package database

import (
	IUser "backend/handler/user"
	"backend/handler/user/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type user struct {
	dbGorm *gorm.DB
}

func NewUserRepo(dbGorm *gorm.DB) IUser.IUserRepo {
	return user{dbGorm: dbGorm}
}

func (u user) InsertUser(req models.User) (res models.User, err error) {
	err = u.dbGorm.Clauses(clause.OnConflict{DoNothing: true}).Create(&req).Error
	return req, err
}

func (u user) SelectUserByID(id int64) (res models.User, err error) {
	err = u.dbGorm.Where("id = ?", id).First(&res).Error
	return
}

func (u user) SelectUserByEmail(email string) (res models.User, err error) {
	err = u.dbGorm.Where("email = ?", email).First(&res).Error
	return
}

func (u user) UpdateUser(req models.User) (res models.User, err error) {
	err = u.dbGorm.Model(&models.User{}).Where("id = ?", req.ID).Updates(req).First(&res).Error
	return
}
