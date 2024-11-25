// Package users provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.3 DO NOT EDIT.
package users

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	strictecho "github.com/oapi-codegen/runtime/strictmiddleware/echo"
)

// User defines model for User.
type User struct {
	Created *time.Time `json:"created,omitempty"`
	Email   *string    `json:"email,omitempty"`
	Id      *uint      `json:"id,omitempty"`
	Updated *time.Time `json:"updated,omitempty"`
}

// UserAdd defines model for UserAdd.
type UserAdd struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserUpdate defines model for UserUpdate.
type UserUpdate struct {
	Email    *string `json:"email,omitempty"`
	Password *string `json:"password,omitempty"`
}

// PostUsersJSONRequestBody defines body for PostUsers for application/json ContentType.
type PostUsersJSONRequestBody = UserAdd

// PatchUsersIdJSONRequestBody defines body for PatchUsersId for application/json ContentType.
type PatchUsersIdJSONRequestBody = UserUpdate

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get all users
	// (GET /users)
	GetUsers(ctx echo.Context) error
	// Add new user
	// (POST /users)
	PostUsers(ctx echo.Context) error
	// Delete user by id
	// (DELETE /users/{id})
	DeleteUsersId(ctx echo.Context, id uint) error
	// Get user by id
	// (GET /users/{id})
	GetUsersId(ctx echo.Context, id uint) error
	// Update user by id
	// (PATCH /users/{id})
	PatchUsersId(ctx echo.Context, id uint) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetUsers converts echo context to params.
func (w *ServerInterfaceWrapper) GetUsers(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetUsers(ctx)
	return err
}

// PostUsers converts echo context to params.
func (w *ServerInterfaceWrapper) PostUsers(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostUsers(ctx)
	return err
}

// DeleteUsersId converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteUsersId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id uint

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteUsersId(ctx, id)
	return err
}

// GetUsersId converts echo context to params.
func (w *ServerInterfaceWrapper) GetUsersId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id uint

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetUsersId(ctx, id)
	return err
}

// PatchUsersId converts echo context to params.
func (w *ServerInterfaceWrapper) PatchUsersId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id uint

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PatchUsersId(ctx, id)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/users", wrapper.GetUsers)
	router.POST(baseURL+"/users", wrapper.PostUsers)
	router.DELETE(baseURL+"/users/:id", wrapper.DeleteUsersId)
	router.GET(baseURL+"/users/:id", wrapper.GetUsersId)
	router.PATCH(baseURL+"/users/:id", wrapper.PatchUsersId)

}

type GetUsersRequestObject struct {
}

type GetUsersResponseObject interface {
	VisitGetUsersResponse(w http.ResponseWriter) error
}

type GetUsers200JSONResponse []User

func (response GetUsers200JSONResponse) VisitGetUsersResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostUsersRequestObject struct {
	Body *PostUsersJSONRequestBody
}

type PostUsersResponseObject interface {
	VisitPostUsersResponse(w http.ResponseWriter) error
}

type PostUsers201JSONResponse User

func (response PostUsers201JSONResponse) VisitPostUsersResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type DeleteUsersIdRequestObject struct {
	Id uint `json:"id"`
}

type DeleteUsersIdResponseObject interface {
	VisitDeleteUsersIdResponse(w http.ResponseWriter) error
}

type DeleteUsersId200Response struct {
}

func (response DeleteUsersId200Response) VisitDeleteUsersIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type DeleteUsersId404Response struct {
}

func (response DeleteUsersId404Response) VisitDeleteUsersIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type GetUsersIdRequestObject struct {
	Id uint `json:"id"`
}

type GetUsersIdResponseObject interface {
	VisitGetUsersIdResponse(w http.ResponseWriter) error
}

type GetUsersId200JSONResponse User

func (response GetUsersId200JSONResponse) VisitGetUsersIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetUsersId404Response struct {
}

func (response GetUsersId404Response) VisitGetUsersIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type PatchUsersIdRequestObject struct {
	Id   uint `json:"id"`
	Body *PatchUsersIdJSONRequestBody
}

type PatchUsersIdResponseObject interface {
	VisitPatchUsersIdResponse(w http.ResponseWriter) error
}

type PatchUsersId200Response struct {
}

func (response PatchUsersId200Response) VisitPatchUsersIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type PatchUsersId404Response struct {
}

func (response PatchUsersId404Response) VisitPatchUsersIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Get all users
	// (GET /users)
	GetUsers(ctx context.Context, request GetUsersRequestObject) (GetUsersResponseObject, error)
	// Add new user
	// (POST /users)
	PostUsers(ctx context.Context, request PostUsersRequestObject) (PostUsersResponseObject, error)
	// Delete user by id
	// (DELETE /users/{id})
	DeleteUsersId(ctx context.Context, request DeleteUsersIdRequestObject) (DeleteUsersIdResponseObject, error)
	// Get user by id
	// (GET /users/{id})
	GetUsersId(ctx context.Context, request GetUsersIdRequestObject) (GetUsersIdResponseObject, error)
	// Update user by id
	// (PATCH /users/{id})
	PatchUsersId(ctx context.Context, request PatchUsersIdRequestObject) (PatchUsersIdResponseObject, error)
}

type StrictHandlerFunc = strictecho.StrictEchoHandlerFunc
type StrictMiddlewareFunc = strictecho.StrictEchoMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// GetUsers operation middleware
func (sh *strictHandler) GetUsers(ctx echo.Context) error {
	var request GetUsersRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetUsers(ctx.Request().Context(), request.(GetUsersRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetUsers")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetUsersResponseObject); ok {
		return validResponse.VisitGetUsersResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PostUsers operation middleware
func (sh *strictHandler) PostUsers(ctx echo.Context) error {
	var request PostUsersRequestObject

	var body PostUsersJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostUsers(ctx.Request().Context(), request.(PostUsersRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostUsers")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostUsersResponseObject); ok {
		return validResponse.VisitPostUsersResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// DeleteUsersId operation middleware
func (sh *strictHandler) DeleteUsersId(ctx echo.Context, id uint) error {
	var request DeleteUsersIdRequestObject

	request.Id = id

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteUsersId(ctx.Request().Context(), request.(DeleteUsersIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteUsersId")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(DeleteUsersIdResponseObject); ok {
		return validResponse.VisitDeleteUsersIdResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetUsersId operation middleware
func (sh *strictHandler) GetUsersId(ctx echo.Context, id uint) error {
	var request GetUsersIdRequestObject

	request.Id = id

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetUsersId(ctx.Request().Context(), request.(GetUsersIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetUsersId")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetUsersIdResponseObject); ok {
		return validResponse.VisitGetUsersIdResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PatchUsersId operation middleware
func (sh *strictHandler) PatchUsersId(ctx echo.Context, id uint) error {
	var request PatchUsersIdRequestObject

	request.Id = id

	var body PatchUsersIdJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PatchUsersId(ctx.Request().Context(), request.(PatchUsersIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PatchUsersId")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PatchUsersIdResponseObject); ok {
		return validResponse.VisitPatchUsersIdResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}