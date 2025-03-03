// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	strictgin "github.com/oapi-codegen/runtime/strictmiddleware/gin"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /orders)
	GetOrders(c *gin.Context)

	// (POST /orders)
	PostOrders(c *gin.Context)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// GetOrders operation middleware
func (siw *ServerInterfaceWrapper) GetOrders(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetOrders(c)
}

// PostOrders operation middleware
func (siw *ServerInterfaceWrapper) PostOrders(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostOrders(c)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.GET(options.BaseURL+"/orders", wrapper.GetOrders)
	router.POST(options.BaseURL+"/orders", wrapper.PostOrders)
}

type GetOrdersRequestObject struct {
}

type GetOrdersResponseObject interface {
	VisitGetOrdersResponse(w http.ResponseWriter) error
}

type GetOrders200JSONResponse AllOrders

func (response GetOrders200JSONResponse) VisitGetOrdersResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetOrders500Response struct {
}

func (response GetOrders500Response) VisitGetOrdersResponse(w http.ResponseWriter) error {
	w.WriteHeader(500)
	return nil
}

type PostOrdersRequestObject struct {
	Body *PostOrdersJSONRequestBody
}

type PostOrdersResponseObject interface {
	VisitPostOrdersResponse(w http.ResponseWriter) error
}

type PostOrders201Response struct {
}

func (response PostOrders201Response) VisitPostOrdersResponse(w http.ResponseWriter) error {
	w.WriteHeader(201)
	return nil
}

type PostOrders400Response struct {
}

func (response PostOrders400Response) VisitPostOrdersResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type PostOrders500Response struct {
}

func (response PostOrders500Response) VisitPostOrdersResponse(w http.ResponseWriter) error {
	w.WriteHeader(500)
	return nil
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {

	// (GET /orders)
	GetOrders(ctx context.Context, request GetOrdersRequestObject) (GetOrdersResponseObject, error)

	// (POST /orders)
	PostOrders(ctx context.Context, request PostOrdersRequestObject) (PostOrdersResponseObject, error)
}

type StrictHandlerFunc = strictgin.StrictGinHandlerFunc
type StrictMiddlewareFunc = strictgin.StrictGinMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// GetOrders operation middleware
func (sh *strictHandler) GetOrders(ctx *gin.Context) {
	var request GetOrdersRequestObject

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetOrders(ctx, request.(GetOrdersRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetOrders")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(GetOrdersResponseObject); ok {
		if err := validResponse.VisitGetOrdersResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// PostOrders operation middleware
func (sh *strictHandler) PostOrders(ctx *gin.Context) {
	var request PostOrdersRequestObject

	var body PostOrdersJSONRequestBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostOrders(ctx, request.(PostOrdersRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostOrders")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(PostOrdersResponseObject); ok {
		if err := validResponse.VisitPostOrdersResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}
