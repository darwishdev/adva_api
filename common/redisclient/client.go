package redisclient

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

type RedisClientInterface interface {
	AuthSessionCreate(ctx context.Context, req *AuthSession) error
	AuthSessionFind(ctx context.Context, username string) (*AuthSession, error)
}

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(host string, port string, password string, db int) RedisClientInterface {
	log.Debug().Interface("host", host).Interface("port", port).Interface("db", db).Msg("from redis")
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password, // no password set
		DB:       db,       // use default DB
	})

	log.Debug().Interface("client", client).Msg("from redis client")

	return &RedisClient{
		client: client,
	}
}
