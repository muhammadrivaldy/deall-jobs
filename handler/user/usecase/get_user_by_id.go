package usecase

import (
	"backend/handler/user/models"
	"backend/logs"
	"backend/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (u *useCase) GetUserByID(ctx context.Context, ID int64) (res models.User, errs util.Error) {

	res, err := u.userEntity.UserRepo.SelectUserByID(ID)
	if err == gorm.ErrRecordNotFound {
		logs.Logging.Error(ctx, err)
		return res, util.ErrorMapping(util.ErrorDataNotFound)
	} else if err != nil {
		logs.Logging.Error(ctx, err)
		return res, util.ErrorMapping(err)
	}

	if !res.IsActive() {
		logs.Logging.Error(ctx, errors.New("user is not active"))
		return res, util.ErrorMapping(util.ErrorDataNotFound)
	}

	return
}
