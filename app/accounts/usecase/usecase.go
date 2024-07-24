package usecase

import (
	"context"
	"time"

	"github.com/bufbuild/protovalidate-go"
	"github.com/darwishdev/bzns_pro_api/app/accounts/adapter"
	"github.com/darwishdev/bzns_pro_api/app/accounts/repo"
	"github.com/darwishdev/bzns_pro_api/common/auth"
	db "github.com/darwishdev/bzns_pro_api/common/db/gen"
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/bznspro/v1"
	"github.com/darwishdev/bzns_pro_api/common/redisclient"
)

type AccountsUsecaseInterface interface {
	RoleCreate(ctx context.Context, req *rmsv1.RoleCreateRequest) (*rmsv1.RoleCreateResponse, error)
	RoleFindForUpdate(ctx context.Context, req *rmsv1.RoleFindForUpdateRequest) (*rmsv1.RoleUpdateRequest, error)
	RoleUpdate(ctx context.Context, req *rmsv1.RoleUpdateRequest) (*rmsv1.RoleUpdateResponse, error)
	PermissionsList(ctx context.Context, req *rmsv1.PermissionsListRequest) (*rmsv1.PermissionsListResponse, error)
	RolesList(ctx context.Context, req *rmsv1.RolesListRequest) (*rmsv1.RolesListResponse, error)
	RoleDeleteRestore(ctx context.Context, req *rmsv1.RoleDeleteRestoreRequest) (*rmsv1.RoleDeleteRestoreResponse, error)
	RolesInputList(ctx context.Context, req *rmsv1.RolesInputListRequest) (*rmsv1.RolesInputListResponse, error)

	// user
	UserLogin(ctx context.Context, req *rmsv1.UserLoginRequest) (*rmsv1.UserLoginResponse, error)
	UserAuthorize(ctx context.Context, req *auth.Payload) (*rmsv1.UserLoginResponse, error)
	UserResetPassword(ctx context.Context, req *rmsv1.UserResetPasswordRequest) error
	UsersList(ctx context.Context, req *rmsv1.UsersListRequest) (*rmsv1.UsersListResponse, error)
	UserDeleteRestore(ctx context.Context, req *rmsv1.UserDeleteRestoreRequest) (*rmsv1.UserDeleteRestoreResponse, error)
	UserCreate(ctx context.Context, req *rmsv1.UserCreateRequest) (*rmsv1.UserCreateResponse, error)
	UserFindForUpdate(ctx context.Context, req *rmsv1.UserFindForUpdateRequest) (*rmsv1.UserUpdateRequest, error)
	UserUpdate(ctx context.Context, req *rmsv1.UserUpdateRequest) (*rmsv1.UserUpdateResponse, error)
}

type AccountsUsecase struct {
	repo          repo.AccountsRepoInterface
	validator     *protovalidate.Validator
	tokenMaker    auth.Maker
	tokenDuration time.Duration
	adapter       adapter.AccountsAdapterInterface
	redisClient   redisclient.RedisClientInterface
}

func NewAccountsUsecase(store db.Store, validator *protovalidate.Validator, tokenMaker auth.Maker, tokenDuration time.Duration, redisClient redisclient.RedisClientInterface) AccountsUsecaseInterface {
	repo := repo.NewAccountsRepo(store)
	adapter := adapter.NewAccountsAdapter()
	return &AccountsUsecase{
		repo:          repo,
		tokenMaker:    tokenMaker,
		validator:     validator,
		tokenDuration: tokenDuration,
		adapter:       adapter,
		redisClient:   redisClient,
	}
}
