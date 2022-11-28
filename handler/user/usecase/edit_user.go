package usecase

import (
	"backend/handler/user/models"
	"backend/handler/user/payload"
	"backend/logs"
	"backend/util"
	"context"
	"errors"
	"time"

	goutil "github.com/muhammadrivaldy/go-util"
	"gorm.io/gorm"
)

func (u *useCase) EditUser(ctx context.Context, req payload.RequestEditUser) (errs util.Error) {

	userInfo := goutil.GetContext(ctx)

	modelUser, err := u.userEntity.UserRepo.SelectUserByID(req.ID)
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
		ID:         req.ID,
		Name:       req.Name,
		Phone:      req.Phone,
		Email:      req.Email,
		UserTypeID: req.UserTypeID,
		UpdatedBy:  userInfo.UserID,
		UpdatedAt:  time.Now(),
	}

	modelUser, err = u.userEntity.UserRepo.UpdateUser(modelUser)
	if err != nil {
		logs.Logging.Error(ctx, err)
		return util.ErrorMapping(err)
	}

	return
}
