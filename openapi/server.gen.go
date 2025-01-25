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
	// ルームのメタデータを取得
	// (GET /rooms/{roomId}/metadata)
	GetRoomMetadata(ctx echo.Context, roomId openapi_types.UUID) error
	// ルームのメタデータを更新
	// (PATCH /rooms/{roomId}/metadata)
	UpdateRoomMetadata(ctx echo.Context, roomId openapi_types.UUID) error
	// ルームでの発言権限を変更
	// (PATCH /rooms/{roomId}/participants)
	ChangeParticipantRole(ctx echo.Context, roomId openapi_types.UUID) error
	// サウンドボード用の音声一覧を取得
	// (GET /soundboard)
	GetSoundboardList(ctx echo.Context) error
	// サウンドボード用の短い音声ファイルをアップロード
	// (POST /soundboard)
	PostSoundboard(ctx echo.Context) error
	// アップロード済み音声を LiveKit ルームで再生
	// (POST /soundboard/play)
	PostSoundboardPlay(ctx echo.Context) error
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

// GetRoomMetadata converts echo context to params.
func (w *ServerInterfaceWrapper) GetRoomMetadata(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "roomId" -------------
	var roomId openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "roomId", ctx.Param("roomId"), &roomId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter roomId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetRoomMetadata(ctx, roomId)
	return err
}

// UpdateRoomMetadata converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateRoomMetadata(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "roomId" -------------
	var roomId openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "roomId", ctx.Param("roomId"), &roomId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter roomId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.UpdateRoomMetadata(ctx, roomId)
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

// GetSoundboardList converts echo context to params.
func (w *ServerInterfaceWrapper) GetSoundboardList(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetSoundboardList(ctx)
	return err
}

// PostSoundboard converts echo context to params.
func (w *ServerInterfaceWrapper) PostSoundboard(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostSoundboard(ctx)
	return err
}

// PostSoundboardPlay converts echo context to params.
func (w *ServerInterfaceWrapper) PostSoundboardPlay(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostSoundboardPlay(ctx)
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
	router.GET(baseURL+"/rooms/:roomId/metadata", wrapper.GetRoomMetadata)
	router.PATCH(baseURL+"/rooms/:roomId/metadata", wrapper.UpdateRoomMetadata)
	router.PATCH(baseURL+"/rooms/:roomId/participants", wrapper.ChangeParticipantRole)
	router.GET(baseURL+"/soundboard", wrapper.GetSoundboardList)
	router.POST(baseURL+"/soundboard", wrapper.PostSoundboard)
	router.POST(baseURL+"/soundboard/play", wrapper.PostSoundboardPlay)
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
