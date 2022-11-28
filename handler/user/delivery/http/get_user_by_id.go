package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	goutil "github.com/muhammadrivaldy/go-util"
)

func (e *endpoint) GetUserByID(c *gin.Context) {

	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		goutil.ResponseError(c, http.StatusInternalServerError, err, nil)
		return
	}

	user, errs := e.useCaseUser.GetUserByID(goutil.ParseContext(c), userID)
	if errs.Error != nil {
		goutil.ResponseError(c, errs.Code, errs.Error, nil)
		return
	}

	goutil.ResponseOK(c, http.StatusOK, user)
}
