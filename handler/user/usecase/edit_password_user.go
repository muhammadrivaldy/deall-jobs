package usecase

import (
	"backend/handler/user/models"
	"backend/logs"
	"backend/util"
	"context"
	"errors"
	"time"

	goutil "github.com/muhammadrivaldy/go-util"
	"gorm.io/gorm"
)

func (u *useCase) EditPasswordUser(ctx context.Context, ID int64, password string) (errs util.Error) {

	userInfo := goutil.GetContext(ctx)

	modelUser, err := u.userEntity.UserRepo.SelectUserByID(ID)
	if err == gorm.ErrRecordNotFound {
		logs.Logging.Error(ctx, err)
		return util.ErrorMapping(util.ErrorDataNotFound)
	} else if err != nil {
		logs.Logging.Error(ctx, err)
		return util.ErrorMapping(err)
	}

	if !modelUser.IsActive() {
		logs.Logging.Error(ctx, errors.New("user is not active"))
		return util.ErrorMapping(util.ErrorDataNotFound)
	}

	modelUser = models.User{
		ID:        ID,
		Password:  password,
		UpdatedBy: userInfo.UserID,
		UpdatedAt: time.Now(),
	}

	modelUser.EncryptPassword()

	modelUser, err = u.userEntity.UserRepo.UpdateUser(modelUser)
	if err != nil {
		logs.Logging.Error(ctx, err)
		return util.ErrorMapping(err)
	}

	return
}
