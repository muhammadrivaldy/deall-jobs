package user

import (
	"backend/handler/user/models"
	"backend/handler/user/payload"
	"backend/util"
	"context"
)

// IUserUseCase is a interface for layer business
type IUserUseCase interface {
	CreateUser(ctx context.Context, req payload.RequestCreateUser) (errs util.Error)
	GetUserByID(ctx context.Context, ID int64) (res models.User, errs util.Error)
	EditUser(ctx context.Context, req payload.RequestEditUser) (errs util.Error)
	EditPasswordUser(ctx context.Context, ID int64, password string) (errs util.Error)
	RemoveUser(ctx context.Context, ID int64) (errs util.Error)
}

type IUserRepo interface {
	InsertUser(req models.User) (res models.User, err error)
	SelectUserByID(id int64) (res models.User, err error)
	SelectUserByEmail(email string) (res models.User, err error)
	UpdateUser(req models.User) (res models.User, err error)
}

type IUserTypeRepo interface {
	SelectUserTypeByID(id int) (res models.UserType, err error)
}
