package repo

import (
	"context"

	db "github.com/darwishdev/bzns_pro_api/common/db/gen"
)

type PublicRepoInterface interface {
	SettingsUpdate(ctx context.Context, req *db.SettingsUpdateParams) error
	SettingsFindForUpdate(ctx context.Context) (*[]db.SettingsFindForUpdateRow, error)
	//cats
	CategoryCreate(ctx context.Context, req *db.CategoryCreateParams) (*db.Category, error)
	CategoryUpdate(ctx context.Context, req *db.CategoryUpdateParams) (*db.Category, error)
	CategoriesList(ctx context.Context, req int32) (*[]db.CategoriesListRow, error)
	CategoriesInputList(ctx context.Context, req int32) (*[]db.CategoriesInputListRow, error)
	CategoryDeleteRestore(ctx context.Context, req []int32) error
	CategoryFindForUpdate(ctx context.Context, req *int32) (*db.CategoryFindForUpdateRow, error)
	//services
	ServiceCreate(ctx context.Context, req *db.ServiceCreateParams) (*db.Service, error)
	ServiceUpdate(ctx context.Context, req *db.ServiceUpdateParams) (*db.Service, error)
	ServicesList(ctx context.Context) (*[]db.ServicesListRow, error)
	ServiceDeleteRestore(ctx context.Context, req []int32) error
	ServiceFindForUpdate(ctx context.Context, req *int32) (*db.ServiceFindForUpdateRow, error)

	//testemonials
	TestemonialCreate(ctx context.Context, req *db.TestemonialCreateParams) (*db.Testemonial, error)
	TestemonialUpdate(ctx context.Context, req *db.TestemonialUpdateParams) (*db.Testemonial, error)
	TestemonialsList(ctx context.Context) (*[]db.TestemonialsListRow, error)
	TestemonialDeleteRestore(ctx context.Context, req []int32) error
	TestemonialFindForUpdate(ctx context.Context, req *int32) (*db.TestemonialFindForUpdateRow, error)

	//team
	TeamMemberCreate(ctx context.Context, req *db.TeamMemberCreateParams) (*db.TeamMember, error)
	TeamMemberUpdate(ctx context.Context, req *db.TeamMemberUpdateParams) (*db.TeamMember, error)
	TeamMembersList(ctx context.Context) (*[]db.TeamMembersListRow, error)
	TeamMemberDeleteRestore(ctx context.Context, req []int32) error
	TeamMemberFindForUpdate(ctx context.Context, req *int32) (*db.TeamMemberFindForUpdateRow, error)

	// projects
	ProjectCreate(ctx context.Context, req *db.ProjectCreateParams) (*db.Project, error)
	ProjectUpdate(ctx context.Context, req *db.ProjectUpdateParams) (*db.Project, error)
	ProjectsList(ctx context.Context) (*[]db.ProjectsListRow, error)
	ProjectDeleteRestore(ctx context.Context, req []int32) error
	ProjectFindForUpdate(ctx context.Context, req *int32) (*db.ProjectFindForUpdateRow, error)

	//programs
	ProgramCreate(ctx context.Context, req *db.ProgramCreateParams) (*db.Program, error)
	ProgramUpdate(ctx context.Context, req *db.ProgramUpdateParams) (*db.Program, error)
	ProgramsList(ctx context.Context) (*[]db.ProgramsListRow, error)
	ProgramDeleteRestore(ctx context.Context, req []int32) error
	ProgramFindForUpdate(ctx context.Context, req *int32) (*db.ProgramFindForUpdateRow, error)
}

type PublicRepo struct {
	store        db.Store
	errorHandler map[string]string
}

func NewPublicRepo(store db.Store) *PublicRepo {
	errorHandler := map[string]string{}
	return &PublicRepo{
		store:        store,
		errorHandler: errorHandler,
	}
}
