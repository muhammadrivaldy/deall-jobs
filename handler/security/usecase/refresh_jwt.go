package usecase

import (
	"backend/logs"
	"backend/util"
	"context"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	goutil "github.com/muhammadrivaldy/go-util"
	"gorm.io/gorm"
)

func (u *useCase) RefreshJWT(ctx context.Context) (token, refreshToken string, errs util.Error) {

	userInfo := goutil.GetContext(ctx)

	modelUser, err := u.userEntity.UserRepo.SelectUserByID(userInfo.UserID)
	if err == gorm.ErrRecordNotFound {
		logs.Logging.Error(ctx, err)
		return "", "", util.ErrorMapping(util.ErrorDataNotFound)
	} else if err != nil {
		logs.Logging.Error(ctx, err)
		return
	}

	if !modelUser.IsActive() {
		logs.Logging.Error(ctx, errors.New("user is not active"))
		return "", "", util.ErrorMapping(util.ErrorDataNotFound)
	}

	token, refreshToken, err = goutil.CreateJWT(goutil.JWT{
		UserID:     modelUser.ID,
		Name:       modelUser.Name,
		Phone:      modelUser.Phone,
		Email:      modelUser.Email,
		GroupID:    modelUser.UserTypeID,
		ExpToken:   time.Now().Add(15 * time.Minute),
		ExpRefresh: time.Now().AddDate(0, 0, 15),
		Jti:        userInfo.Jti,
	}, jwt.SigningMethodHS256, u.config.JWTKey)
	if err != nil {
		logs.Logging.Error(ctx, err)
		return
	}

	return
}
