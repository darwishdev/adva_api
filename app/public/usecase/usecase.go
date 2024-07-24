package usecase

import (
	"context"

	"github.com/bufbuild/protovalidate-go"
	"github.com/darwishdev/bzns_pro_api/app/public/adapter"
	"github.com/darwishdev/bzns_pro_api/app/public/repo"
	db "github.com/darwishdev/bzns_pro_api/common/db/gen"
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/bznspro/v1"
)

type PublicUsecaseInterface interface {
	SettingsUpdate(ctx context.Context, req *rmsv1.SettingsUpdateRequest) error
	SettingsFindForUpdate(ctx context.Context, req *rmsv1.SettingsFindForUpdateRequest) (*rmsv1.SettingsFindForUpdateResponse, error)
	CategoryCreate(ctx context.Context, req *rmsv1.CategoryCreateRequest) (*rmsv1.CategoryCreateResponse, error)
	CategoryFindForUpdate(ctx context.Context, req *rmsv1.CategoryFindForUpdateRequest) (*rmsv1.CategoryUpdateRequest, error)
	CategoryUpdate(ctx context.Context, req *rmsv1.CategoryUpdateRequest) (*rmsv1.CategoryUpdateResponse, error)
	CategoriesList(ctx context.Context, req *rmsv1.CategoriesListRequest) (*rmsv1.CategoriesListResponse, error)
	CategoryDeleteRestore(ctx context.Context, req *rmsv1.CategoryDeleteRestoreRequest) (*rmsv1.CategoryDeleteRestoreResponse, error)
	CategoriesInputList(ctx context.Context, req *rmsv1.CategoriesInputListRequest) (*rmsv1.CategoriesInputListResponse, error)
	// service
	ServiceCreate(ctx context.Context, req *rmsv1.ServiceCreateRequest) (*rmsv1.ServiceCreateResponse, error)
	ServiceFindForUpdate(ctx context.Context, req *rmsv1.ServiceFindForUpdateRequest) (*rmsv1.ServiceUpdateRequest, error)
	ServiceUpdate(ctx context.Context, req *rmsv1.ServiceUpdateRequest) (*rmsv1.ServiceUpdateResponse, error)
	ServicesList(ctx context.Context, req *rmsv1.ServicesListRequest) (*rmsv1.ServicesListResponse, error)
	ServiceDeleteRestore(ctx context.Context, req *rmsv1.ServiceDeleteRestoreRequest) (*rmsv1.ServiceDeleteRestoreResponse, error)
	//project
	ProjectCreate(ctx context.Context, req *rmsv1.ProjectCreateRequest) (*rmsv1.ProjectCreateResponse, error)
	ProjectFindForUpdate(ctx context.Context, req *rmsv1.ProjectFindForUpdateRequest) (*rmsv1.ProjectUpdateRequest, error)
	ProjectUpdate(ctx context.Context, req *rmsv1.ProjectUpdateRequest) (*rmsv1.ProjectUpdateResponse, error)
	ProjectsList(ctx context.Context, req *rmsv1.ProjectsListRequest) (*rmsv1.ProjectsListResponse, error)
	ProjectDeleteRestore(ctx context.Context, req *rmsv1.ProjectDeleteRestoreRequest) (*rmsv1.ProjectDeleteRestoreResponse, error)

	//testemonial
	TestemonialCreate(ctx context.Context, req *rmsv1.TestemonialCreateRequest) (*rmsv1.TestemonialCreateResponse, error)
	TestemonialFindForUpdate(ctx context.Context, req *rmsv1.TestemonialFindForUpdateRequest) (*rmsv1.TestemonialUpdateRequest, error)
	TestemonialUpdate(ctx context.Context, req *rmsv1.TestemonialUpdateRequest) (*rmsv1.TestemonialUpdateResponse, error)
	TestemonialsList(ctx context.Context, req *rmsv1.TestemonialsListRequest) (*rmsv1.TestemonialsListResponse, error)
	TestemonialDeleteRestore(ctx context.Context, req *rmsv1.TestemonialDeleteRestoreRequest) (*rmsv1.TestemonialDeleteRestoreResponse, error)
	//team
	TeamMemberCreate(ctx context.Context, req *rmsv1.TeamMemberCreateRequest) (*rmsv1.TeamMemberCreateResponse, error)
	TeamMemberFindForUpdate(ctx context.Context, req *rmsv1.TeamMemberFindForUpdateRequest) (*rmsv1.TeamMemberUpdateRequest, error)
	TeamMemberUpdate(ctx context.Context, req *rmsv1.TeamMemberUpdateRequest) (*rmsv1.TeamMemberUpdateResponse, error)
	TeamMembersList(ctx context.Context, req *rmsv1.TeamMembersListRequest) (*rmsv1.TeamMembersListResponse, error)
	TeamMemberDeleteRestore(ctx context.Context, req *rmsv1.TeamMemberDeleteRestoreRequest) (*rmsv1.TeamMemberDeleteRestoreResponse, error)

	ClientInitialize(ctx context.Context, req *rmsv1.ClientInitializeRequest) (*rmsv1.ClientInitializeResponse, error)

	ProgramCreate(ctx context.Context, req *rmsv1.ProgramCreateRequest) (*rmsv1.ProgramCreateResponse, error)
	ProgramFindForUpdate(ctx context.Context, req *rmsv1.ProgramFindForUpdateRequest) (*rmsv1.ProgramUpdateRequest, error)
	ProgramUpdate(ctx context.Context, req *rmsv1.ProgramUpdateRequest) (*rmsv1.ProgramUpdateResponse, error)
	ProgramsList(ctx context.Context, req *rmsv1.ProgramsListRequest) (*rmsv1.ProgramsListResponse, error)
	ProgramDeleteRestore(ctx context.Context, req *rmsv1.ProgramDeleteRestoreRequest) (*rmsv1.ProgramDeleteRestoreResponse, error)
}

type PublicUsecase struct {
	repo      repo.PublicRepoInterface
	validator *protovalidate.Validator
	adapter   adapter.PublicAdapterInterface
}

func NewPublicUsecase(store db.Store, validator *protovalidate.Validator) *PublicUsecase {
	repo := repo.NewPublicRepo(store)
	adapter := adapter.NewPublicAdapter()

	return &PublicUsecase{
		repo:      repo,
		validator: validator,
		adapter:   adapter,
	}
}
