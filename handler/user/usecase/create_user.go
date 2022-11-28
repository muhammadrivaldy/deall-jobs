package usecase

import (
	"backend/handler/user/models"
	"backend/handler/user/payload"
	"backend/logs"
	"backend/util"
	"context"
	"time"

	goutil "github.com/muhammadrivaldy/go-util"
)

func (u *useCase) CreateUser(ctx context.Context, req payload.RequestCreateUser) (errs util.Error) {

	userInfo := goutil.GetContext(ctx)

	modelUser := models.User{
		Name:       req.Name,
		Phone:      req.Phone,
		Email:      req.Email,
		Status:     int(models.UserStatusIDActive),
		Password:   req.Password,
		UserTypeID: req.UserTypeID,
		CreatedBy:  userInfo.UserID,
		CreatedAt:  time.Now(),
		UpdatedBy:  userInfo.UserID,
		UpdatedAt:  time.Now(),
	}

	modelUser.EncryptPassword()

	modelUser, err := u.userEntity.UserRepo.InsertUser(modelUser)
	if err != nil {
		logs.Logging.Error(ctx, err)
		return util.ErrorMapping(err)
	}

	return
}
