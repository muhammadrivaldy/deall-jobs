package middleware

import (
	"backend/handler/security"
	"backend/handler/user"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	goutil "github.com/muhammadrivaldy/go-util"
)

type middleware struct {
	useCaseUser     user.IUserUseCase
	useCaseSecurity security.ISecurityUseCase
}

func NewMiddleware(useCaseUser user.IUserUseCase, useCaseSecurity security.ISecurityUseCase) middleware {
	return middleware{
		useCaseUser:     useCaseUser,
		useCaseSecurity: useCaseSecurity,
	}
}

func (m middleware) ValidateAccess(apiID int64) func(c *gin.Context) {
	return func(c *gin.Context) {

		ctx := goutil.ParseContext(c)

		userInfo := goutil.GetContext(ctx)

		modelUser, errs := m.useCaseUser.GetUserByID(ctx, userInfo.UserID)
		if errs.Error != nil {
			goutil.ResponseError(c, http.StatusForbidden, errors.New(http.StatusText(http.StatusForbidden)), nil)
			c.Abort()
			return
		}

		// the meaning of 8 is api for refresh jwt
		if apiID != 8 {
			if modelUser.Email != userInfo.Email || modelUser.Phone != userInfo.Phone || modelUser.UserTypeID != userInfo.GroupID {
				goutil.ResponseError(c, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)), nil)
				c.Abort()
				return
			}
		}

		_, errs = m.useCaseSecurity.ValidateAccessUser(ctx, apiID)
		if errs.Error != nil {
			goutil.ResponseError(c, http.StatusForbidden, errors.New(http.StatusText(http.StatusForbidden)), nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
