package http

import (
	"backend/handler/user/payload"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	goutil "github.com/muhammadrivaldy/go-util"
)

func (e *endpoint) EditUser(c *gin.Context) {

	userID, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		goutil.ResponseError(c, http.StatusInternalServerError, err, nil)
		return
	}

	payload := payload.RequestEditUser{ID: userID}

	if err := c.BindJSON(&payload); err != nil {
		goutil.ResponseError(c, http.StatusBadRequest, errors.New(http.StatusText(http.StatusBadRequest)), nil)
		return
	}

	if err := e.validation.ValidationStruct(payload); err != nil {
		goutil.ResponseError(c, http.StatusBadRequest, errors.New(http.StatusText(http.StatusBadRequest)), nil)
		return
	}

	errs := e.useCaseUser.EditUser(goutil.ParseContext(c), payload)
	if errs.Error != nil {
		goutil.ResponseError(c, errs.Code, errs.Error, nil)
		return
	}

	goutil.ResponseOK(c, http.StatusOK, nil)
}
