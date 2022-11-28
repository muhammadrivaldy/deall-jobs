package http

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	goutil "github.com/muhammadrivaldy/go-util"
)

func (e *endpoint) Login(c *gin.Context) {

	email, password, _ := c.Request.BasicAuth()

	if err := e.validation.ValidationVariable(email, "email", ""); err != nil {
		goutil.ResponseError(c, http.StatusBadRequest, errors.New(http.StatusText(http.StatusBadRequest)), nil)
	}

	if err := e.validation.ValidationVariable(email, "required,min=8", ""); err != nil {
		goutil.ResponseError(c, http.StatusBadRequest, errors.New(http.StatusText(http.StatusBadRequest)), nil)
	}

	token, refreshToken, errs := e.useCaseSecurity.Login(context.Background(), email, password)
	if errs.Error != nil {
		goutil.ResponseError(c, errs.Code, errs.Error, nil)
		return
	}

	goutil.ResponseOK(c, http.StatusOK, map[string]string{"token": token, "refresh_token": refreshToken})
}
