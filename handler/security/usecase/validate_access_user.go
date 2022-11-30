package usecase

import (
	"backend/logs"
	"backend/util"
	"context"

	goutil "github.com/muhammadrivaldy/go-util"
	"gorm.io/gorm"
)

func (u useCase) ValidateAccessUser(ctx context.Context, apiID int64) (res bool, errs util.Error) {

	userInfo := goutil.GetContext(ctx)

	// prepare a filter
	filter := util.FilterQuery{}
	filter.Conditions = append(filter.Conditions, util.Condition{Field: "api_id", Operation: "=", Value: apiID})
	filter.Conditions = append(filter.Conditions, util.Condition{Operation: "and"})
	filter.Conditions = append(filter.Conditions, util.Condition{Field: "user_type_id", Operation: "=", Value: userInfo.GroupID})

	// check access of user
	_, err := u.securityEntity.AccessRepo.SelectAccessByFilter(filter)
	if err == gorm.ErrRecordNotFound {
		logs.Logging.Error(ctx, err)
		return false, util.ErrorMapping(err)
	} else if err != nil {
		logs.Logging.Error(ctx, err)
		return false, util.ErrorMapping(err)
	}

	return true, util.Error{}
}
