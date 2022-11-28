package http

import (
	"backend/handler/user/payload"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	goutil "github.com/muhammadrivaldy/go-util"
)

func (e *endpoint) CreateUser(c *gin.Context) {

	userTypeID, _ := strconv.ParseInt(c.PostForm("user_type_id"), 10, 8)

	var payload = payload.RequestCreateUser{
		Name:       c.PostForm("name"),
		Phone:      c.PostForm("phone"),
		Email:      c.PostForm("email"),
		Password:   c.PostForm("password"),
		UserTypeID: int(userTypeID),
	}

	if err := e.validation.ValidationStruct(payload); err != nil {
		goutil.ResponseError(c, http.StatusBadRequest, err, nil)
		return
	}

	errs := e.useCaseUser.CreateUser(goutil.ParseContext(c), payload)
	if errs.Error != nil {
		goutil.ResponseError(c, errs.Code, errs.Error, nil)
		return
	}

	goutil.ResponseOK(c, http.StatusCreated, nil)
}
