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
	openapi_types "github.com/oapi-codegen/runtime/types"
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
	// ルームでの発言権限を変更
	// (PATCH /rooms/{roomId}/participants)
	ChangeParticipantRole(ctx echo.Context, roomId openapi_types.UUID) error
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

// ChangeParticipantRole converts echo context to params.
func (w *ServerInterfaceWrapper) ChangeParticipantRole(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "roomId" -------------
	var roomId openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "roomId", ctx.Param("roomId"), &roomId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter roomId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.ChangeParticipantRole(ctx, roomId)
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

	// ------------- Optional query parameter "isWebinar" -------------

	err = runtime.BindQueryParameter("form", true, false, "isWebinar", ctx.QueryParams(), &params.IsWebinar)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter isWebinar: %s", err))
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
	router.PATCH(baseURL+"/rooms/:roomId/participants", wrapper.ChangeParticipantRole)
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
