package http

import (
	"backend/config"
	"backend/handler/security"
	"backend/handler/security/payload"
	"backend/handler/user"
	"backend/middleware"
	"context"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	goutil "github.com/muhammadrivaldy/go-util"
)

type endpoint struct {
	useCaseSecurity security.ISecurityUseCase
	useCaseUser     user.IUserUseCase
	validation      goutil.Validation
}

// NewEndpoint is a function for override a useCase method
func NewEndpoint(
	config config.Configuration,
	engine *gin.Engine,
	useCaseSecurity security.ISecurityUseCase,
	useCaseUser user.IUserUseCase,
	validation goutil.Validation) error {

	// declare variable
	var edp = endpoint{
		useCaseSecurity: useCaseSecurity,
		useCaseUser:     useCaseUser,
		validation:      validation}

	// register the service
	serviceID, err := useCaseSecurity.RegisterService(context.TODO(), "User")
	if err.Error != nil {
		return err.Error
	}

	// declare the endpoint
	const rootEndpoint = "/api/v1/user"
	var createUser = payload.RegisterApiRequest{Name: "Create User", Endpoint: rootEndpoint, Method: http.MethodPost, ServiceID: serviceID}
	var getUserByID = payload.RegisterApiRequest{Name: "Get User By ID", Endpoint: rootEndpoint + "/:user_id", Method: http.MethodGet, ServiceID: serviceID}
	var editUser = payload.RegisterApiRequest{Name: "Edit User", Endpoint: rootEndpoint + "/:user_id", Method: http.MethodPut, ServiceID: serviceID}
	var editPasswordUser = payload.RegisterApiRequest{Name: "Edit Password User", Endpoint: rootEndpoint + "/:user_id/password", Method: http.MethodPut, ServiceID: serviceID}
	var removeUser = payload.RegisterApiRequest{Name: "Remove User", Endpoint: rootEndpoint + "/:user_id", Method: http.MethodDelete, ServiceID: serviceID}

	// append data apis
	var Apis []*payload.RegisterApiRequest
	Apis = append(Apis, &createUser, &getUserByID, &editUser, &editPasswordUser, &removeUser)

	// register the apis
	for _, i := range Apis {
		useCaseSecurity.RegisterApi(context.TODO(), i)
	}

	// middleware validate access
	middleware := middleware.NewMiddleware(useCaseUser, useCaseSecurity)

	// route the endpoint
	engine.Handle(createUser.Method, createUser.Endpoint, goutil.ParseJWT(config.JWTKey, jwt.SigningMethodHS256), middleware.ValidateAccess(int(createUser.ID)), edp.CreateUser)
	engine.Handle(getUserByID.Method, getUserByID.Endpoint, goutil.ParseJWT(config.JWTKey, jwt.SigningMethodHS256), middleware.ValidateAccess(int(getUserByID.ID)), edp.GetUserByID)
	engine.Handle(editUser.Method, editUser.Endpoint, goutil.ParseJWT(config.JWTKey, jwt.SigningMethodHS256), middleware.ValidateAccess(int(editUser.ID)), edp.EditUser)
	engine.Handle(editPasswordUser.Method, editPasswordUser.Endpoint, goutil.ParseJWT(config.JWTKey, jwt.SigningMethodHS256), middleware.ValidateAccess(int(editPasswordUser.ID)), edp.EditPasswordUser)
	engine.Handle(removeUser.Method, removeUser.Endpoint, goutil.ParseJWT(config.JWTKey, jwt.SigningMethodHS256), middleware.ValidateAccess(int(removeUser.ID)), edp.RemoveUser)

	// send result
	return nil
}
