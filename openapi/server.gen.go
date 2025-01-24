// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package openapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	. "github.com/pikachu0310/livekit-server/openapi/models"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// サーバーの生存確認
	// (GET /ping)
	PingServer(ctx echo.Context) error
	// ルームと参加者の一覧を取得
	// (GET /rooms)
	GetRooms(ctx echo.Context) error
	// テスト用
	// (GET /test)
	Test(ctx echo.Context) error
	// LiveKitトークンを取得
	// (GET /token)
	GetLiveKitToken(ctx echo.Context, params GetLiveKitTokenParams) error
	// LiveKit Webhook受信
	// (POST /webhook)
	LiveKitWebhook(ctx echo.Context) error
	// WebSocketエンドポイント
	// (GET /ws)
	GetWs(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// PingServer converts echo context to params.
func (w *ServerInterfaceWrapper) PingServer(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PingServer(ctx)
	return err
}

// GetRooms converts echo context to params.
func (w *ServerInterfaceWrapper) GetRooms(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetRooms(ctx)
	return err
}

// Test converts echo context to params.
func (w *ServerInterfaceWrapper) Test(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Test(ctx)
	return err
}

// GetLiveKitToken converts echo context to params.
func (w *ServerInterfaceWrapper) GetLiveKitToken(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetLiveKitTokenParams
	// ------------- Required query parameter "room" -------------

	err = runtime.BindQueryParameter("form", true, true, "room", ctx.QueryParams(), &params.Room)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter room: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetLiveKitToken(ctx, params)
	return err
}

// LiveKitWebhook converts echo context to params.
func (w *ServerInterfaceWrapper) LiveKitWebhook(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.LiveKitWebhook(ctx)
	return err
}

// GetWs converts echo context to params.
func (w *ServerInterfaceWrapper) GetWs(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetWs(ctx)
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

	router.GET(baseURL+"/ping", wrapper.PingServer)
	router.GET(baseURL+"/rooms", wrapper.GetRooms)
	router.GET(baseURL+"/test", wrapper.Test)
	router.GET(baseURL+"/token", wrapper.GetLiveKitToken)
	router.POST(baseURL+"/webhook", wrapper.LiveKitWebhook)
	router.GET(baseURL+"/ws", wrapper.GetWs)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
