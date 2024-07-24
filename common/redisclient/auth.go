package redisclient

import (
	"context"
	"encoding/json"

	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/bznspro/v1"
	"github.com/rs/zerolog/log"
)

type AuthSession struct {
	UserName    string
	SideBar     []*rmsv1.SideBarItem
	Permissions map[string]map[string]bool
	UserID      int32
}

func (r *RedisClient) AuthSessionCreate(ctx context.Context, req *AuthSession) error {
	jsonBytes, err := json.Marshal(req)
	if err != nil {
		return err
	}
	log.Debug().Interface("username", req.UserName).Msg("AuthSessionFind")

	err = r.client.Set(ctx, req.UserName, jsonBytes, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisClient) AuthSessionFind(ctx context.Context, username string) (*AuthSession, error) {
	var parsedStruct AuthSession
	jsonBytes, err := r.client.Get(ctx, username).Bytes()
	if err != nil {
		return nil, err
	}

	log.Debug().Interface("username", username).Msg("AuthSessionFind")
	err = json.Unmarshal(jsonBytes, &parsedStruct)
	if err != nil {
		return nil, err
	}

	return &parsedStruct, nil
}
