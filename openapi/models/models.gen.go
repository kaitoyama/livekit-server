// Package models provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package models

import (
	"time"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

// RoomWithParticipants defines model for RoomWithParticipants.
type RoomWithParticipants struct {
	// Participants ルーム内の参加者一覧
	Participants []struct {
		// Identity ユーザーIDや名前など
		Identity *string `json:"identity,omitempty"`

		// JoinedAt 参加した時刻
		JoinedAt *time.Time `json:"joinedAt,omitempty"`
	} `json:"participants"`

	// RoomId ルームのID
	RoomId openapi_types.UUID `json:"roomId"`
}

// GetLiveKitTokenParams defines parameters for GetLiveKitToken.
type GetLiveKitTokenParams struct {
	// Room 参加するルームのUUID
	Room openapi_types.UUID `form:"room" json:"room"`
}
