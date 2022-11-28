package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	goutil "github.com/muhammadrivaldy/go-util"
)

func (e *endpoint) RefreshJWT(c *gin.Context) {

	token, refreshToken, errs := e.useCaseSecurity.RefreshJWT(goutil.ParseContext(c))
	if errs.Error != nil {
		goutil.ResponseError(c, errs.Code, errs.Error, nil)
		return
	}

	goutil.ResponseOK(c, http.StatusOK, map[string]string{"token": token, "refresh_token": refreshToken})
}
