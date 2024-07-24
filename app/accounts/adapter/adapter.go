package adapter

import (
	db "github.com/darwishdev/bzns_pro_api/common/db/gen"
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/bznspro/v1"
	"github.com/darwishdev/bzns_pro_api/common/redisclient"
)

type AccountsAdapterInterface interface {
	PermissionsListGrpcFromSql(resp *[][]byte) (*rmsv1.PermissionsListResponse, error)
	rolesListRowGrpcFromSql(resp *db.RolesListRow) *rmsv1.RolesListRow
	RolesListGrpcFromSql(resp *[]db.RolesListRow) *rmsv1.RolesListResponse
	RoleCreateSqlFromGrpc(req *rmsv1.RoleCreateRequest) *db.RoleCreateTXParams
	RoleUpdateSqlFromGrpc(req *rmsv1.RoleUpdateRequest) *db.RoleUpdateTXParams
	RoleUpdateGrpcFromSql(resp *db.AccountsSchemaRole) *rmsv1.RoleUpdateResponse
	RoleEntityGrpcFromSql(resp *db.AccountsSchemaRole) *rmsv1.Role
	RoleFindForUpdateGrpcFromSql(resp *[]byte) (*rmsv1.RoleUpdateRequest, error)
	RoleCreateGrpcFromSql(resp *db.AccountsSchemaRole) *rmsv1.RoleCreateResponse
	RolesInputListGrpcFromSql(resp *[]db.RolesInputListRow) *rmsv1.RolesInputListResponse

	//userUserFindByEmailOrCode(ctx context.Context, req string) (*db.UserFindByEmailOrCodeRow, error)
	UserLoginGrpcFromSql(resp *db.UserFindByEmailOrCodeRow) (*rmsv1.UserLoginResponse, error)
	UserLoginGrpcFromCache(authSession *redisclient.AuthSession) (*rmsv1.UserLoginResponse, error)
	UserResetPasswordSqlFromGrpc(req *rmsv1.UserResetPasswordRequest) *db.UserResetPasswordParams
	UsersListGrpcFromSql(resp [][]byte) (*rmsv1.UsersListResponse, error)
	UserCreateSqlFromGrpc(req *rmsv1.UserCreateRequest) *db.UserCreateTXParams
	UserEntityGrpcFromSql(resp *db.AccountsSchemaUser) *rmsv1.User
	UserCreateGrpcFromSql(resp *db.AccountsSchemaUser) *rmsv1.UserCreateResponse
	UserUpdateSqlFromGrpc(req *rmsv1.UserUpdateRequest) *db.UserUpdateTXParams
	UserUpdateGrpcFromSql(resp *db.AccountsSchemaUser) *rmsv1.UserUpdateResponse
	UserFindForUpdateGrpcFromSql(resp *[]byte) (*rmsv1.UserUpdateRequest, error)
	UsersPermissionsMapFromSql(resp []byte) (*map[string]map[string]bool, error)
	// UsersPermissionsMapFromSql(resp []db.UserPermissionsListRow) (*map[string]map[string]bool, error)
}

type AccountsAdapter struct {
	dateFormat string
}

func NewAccountsAdapter() *AccountsAdapter {
	return &AccountsAdapter{
		dateFormat: "2006-01-02 15:04:05",
	}
}
