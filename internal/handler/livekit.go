package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/pikachu0310/livekit-server/internal/pkg/util"

	"github.com/labstack/echo/v4"
	"github.com/livekit/protocol/auth"
	"github.com/livekit/protocol/livekit"
	lksdk "github.com/livekit/server-sdk-go/v2"
	"github.com/pikachu0310/livekit-server/openapi/models"
)

// GetLiveKitToken GET /token?room=UUID
// Bearerトークン(ES256)で認証後、LiveKit接続用JWTを生成して返す。
func (h *Handler) GetLiveKitToken(c echo.Context, _ models.GetLiveKitTokenParams) error {
	// 1) roomクエリパラメータ取得 (必須)
	room := c.QueryParam("room")
	if room == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "room query parameter is required",
		})
	}

	if !h.repo.CheckChannelExistence(room) {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Channel not found: " + room,
		})
	}

	isWebinar := c.QueryParam("isWebinar") == "true"

	userID, echoError := util.GetTraqUserID(c)
	if echoError != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": echoError.Error(),
		})
	}

	// 6) LiveKit用APIキー/シークレット
	apiKey := h.repo.ApiKey
	apiSecret := h.repo.ApiSecret
	if apiKey == "" || apiSecret == "" {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "API key and secret must be set in environment variables",
		})
	}

	// 6-2) ルームが存在するか確認
	isExistingRoom := false
	for _, roomState := range h.repo.RoomState {
		if roomState.RoomId.String() == room {
			isExistingRoom = true
			break
		}
	}

	if isExistingRoom {
		// ルームが存在して、webinar=true の場合はCanPublish=false
		for _, roomState := range h.repo.RoomState {
			if roomState.RoomId.String() == room {
				if roomState.IsWebinar != nil && *roomState.IsWebinar {
					isWebinar = true
				}
				break
			}
		}
	}

	// 7) VideoGrant にルーム名、CanPublishData=true を設定
	// ルームが存在しない場合はCanPublish=true
	// ルームが存在して、webinar=true の場合はCanPublish=false
	at := auth.NewAccessToken(apiKey, apiSecret)
	grant := &auth.VideoGrant{
		RoomJoin:             true,
		Room:                 room,
		CanPublish:           util.BoolPtr(!(isWebinar && isExistingRoom)),
		CanPublishData:       util.BoolPtr(true),
		CanUpdateOwnMetadata: util.BoolPtr(true),
	}
	randomUUID := uuid.New().String()
	userIdentity := fmt.Sprintf("%s_%s", userID, randomUUID)
	at.SetVideoGrant(grant).
		SetIdentity(userIdentity).
		SetName(userID).
		SetValidFor(24 * time.Hour)

	livekitToken, err := at.ToJWT()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to generate livekit token",
		})
	}

	// 8) ルーム状態を更新
	if !isExistingRoom {
		metadata := util.Metadata{
			Status:    "",
			IsWebinar: isWebinar,
		}
		metadataStr, err := json.Marshal(metadata)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to marshal metadata",
			})
		}
		lkclient := lksdk.NewRoomServiceClient(h.repo.LiveKitHost, h.repo.ApiKey, h.repo.ApiSecret)
		_, err = lkclient.CreateRoom(c.Request().Context(), &livekit.CreateRoomRequest{
			Name:     room,
			Metadata: string(metadataStr),
		})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to create room",
			})
		}
		// ルームが存在しない場合は新規作成
		emptyMetadata := ""
		roomWithParticipants := models.RoomWithParticipants{
			IsWebinar:    &isWebinar,
			Metadata:     &emptyMetadata,
			RoomId:       uuid.MustParse(room),
			Participants: []models.Participant{},
		}
		h.repo.AddRoomState(roomWithParticipants)
		h.repo.SendStartRoomMessageToTraQ(room)
	}

	// 9) 最終的にトークンをJSONで返す
	return c.JSON(http.StatusOK, models.TokenResponse{
		Token: livekitToken,
	})
}
