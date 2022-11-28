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

func (m middleware) ValidateAccess(apiID int) func(c *gin.Context) {
	return func(c *gin.Context) {

		ctx := goutil.ParseContext(c)
		userInfo := goutil.GetContext(ctx)

		_, errs := m.useCaseUser.GetUserByID(ctx, userInfo.UserID)
		if errs.Error != nil {
			goutil.ResponseError(c, http.StatusForbidden, errors.New(http.StatusText(http.StatusForbidden)), nil)
			c.Abort()
			return
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
