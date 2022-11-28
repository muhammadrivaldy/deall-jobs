package http

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	goutil "github.com/muhammadrivaldy/go-util"
)

func (e *endpoint) EditPasswordUser(c *gin.Context) {

	password := c.Param("password")

	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		goutil.ResponseError(c, http.StatusInternalServerError, err, nil)
		return
	}

	if err := e.validation.ValidationVariable(userID, "required,min=1", ""); err != nil {
		goutil.ResponseError(c, http.StatusBadRequest, errors.New(http.StatusText(http.StatusBadRequest)), nil)
		return
	}

	if err := e.validation.ValidationVariable(password, "required,min=8", ""); err != nil {
		goutil.ResponseError(c, http.StatusBadRequest, errors.New(http.StatusText(http.StatusBadRequest)), nil)
		return
	}

	errs := e.useCaseUser.EditPasswordUser(goutil.ParseContext(c), userID, password)
	if errs.Error != nil {
		goutil.ResponseError(c, errs.Code, errs.Error, nil)
		return
	}

	goutil.ResponseOK(c, http.StatusOK, nil)
}
