package gapi

import (
	"github.com/bufbuild/protovalidate-go"
	accountsUsecase "github.com/darwishdev/bzns_pro_api/app/accounts/usecase"
	blogUsecase "github.com/darwishdev/bzns_pro_api/app/blog/usecase"
	eventsUsecase "github.com/darwishdev/bzns_pro_api/app/events/usecase"
	placesUsecase "github.com/darwishdev/bzns_pro_api/app/places/usecase"

	publicUsecase "github.com/darwishdev/bzns_pro_api/app/public/usecase"
	"github.com/darwishdev/bzns_pro_api/common/auth"
	db "github.com/darwishdev/bzns_pro_api/common/db/gen"
	"github.com/darwishdev/bzns_pro_api/common/pb/bznspro/v1/bznsprov1connect"
	"github.com/darwishdev/bzns_pro_api/common/redisclient"

	"github.com/darwishdev/bzns_pro_api/config"
)

// Server serves gRPC requests for our banking usecase.
type Api struct {
	bznsprov1connect.UnimplementedBznsProServiceHandler
	config          config.Config
	tokenMaker      auth.Maker
	accountsUsecase accountsUsecase.AccountsUsecaseInterface
	eventsUsecase   eventsUsecase.EventsUsecaseInterface
	placesUsecase   placesUsecase.PlacesUsecaseInterface
	blogUsecase     blogUsecase.BlogsUsecaseInterface
	redisClient     redisclient.RedisClientInterface

	publicUsecase publicUsecase.PublicUsecaseInterface
	store         db.Store
}

// NewServer creates a new gRPC server.
func NewApi(config config.Config, store db.Store, redisClient redisclient.RedisClientInterface) bznsprov1connect.BznsProServiceHandler {
	validator, err := protovalidate.New()

	if err != nil {
		panic("cann't create validator in gapi/api.go")
	}
	tokenMaker, err := auth.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		panic("cann't create paset maker in gapi/api.go")
	}
	accountsUsecase := accountsUsecase.NewAccountsUsecase(store, validator, tokenMaker, config.AccessTokenDuration, redisClient)
	eventsUsecase := eventsUsecase.NewEventsUsecase(store, validator)
	placesUsecase := placesUsecase.NewPlacesUsecase(store, validator)
	blogUsecase := blogUsecase.NewBlogsUsecase(store, validator)

	publicUsecase := publicUsecase.NewPublicUsecase(store, validator)
	return &Api{
		config:          config,
		tokenMaker:      tokenMaker,
		store:           store,
		accountsUsecase: accountsUsecase,
		redisClient:     redisClient,
		eventsUsecase:   eventsUsecase,
		placesUsecase:   placesUsecase,
		blogUsecase:     blogUsecase,
		publicUsecase:   publicUsecase,
	}
}
