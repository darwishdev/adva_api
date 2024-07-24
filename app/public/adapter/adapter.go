package adapter

import (
	db "github.com/darwishdev/bzns_pro_api/common/db/gen"
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/bznspro/v1"
)

type PublicAdapterInterface interface {
	SettingsUpdateSqlFromGrpc(req *rmsv1.SettingsUpdateRequest) *db.SettingsUpdateParams
	SettingsEntityGrpcFromSql(resp []db.Setting) []*rmsv1.Setting
	SettingsFindForUpdateGrpcFromSql(resp *[]db.SettingsFindForUpdateRow) *rmsv1.SettingsFindForUpdateResponse
	categoriesListRowGrpcFromSql(resp *db.CategoriesListRow) *rmsv1.CategoriesListRow
	CategoriesListGrpcFromSql(resp *[]db.CategoriesListRow) *rmsv1.CategoriesListResponse
	CategoryCreateSqlFromGrpc(req *rmsv1.CategoryCreateRequest) *db.CategoryCreateParams
	CategoryEntityGrpcFromSql(resp *db.Category) *rmsv1.Category
	CategoryCreateGrpcFromSql(resp *db.Category) *rmsv1.CategoryCreateResponse
	CategoryUpdateSqlFromGrpc(req *rmsv1.CategoryUpdateRequest) *db.CategoryUpdateParams
	CategoryUpdateGrpcFromSql(resp *db.Category) *rmsv1.CategoryUpdateResponse
	CategoriesInputListGrpcFromSql(resp *[]db.CategoriesInputListRow) *rmsv1.CategoriesInputListResponse
	CategoryFindForUpdateGrpcFromSql(resp *db.CategoryFindForUpdateRow) *rmsv1.CategoryUpdateRequest

	//services
	servicesListRowGrpcFromSql(resp *db.ServicesListRow) *rmsv1.ServicesListRow
	ServicesListGrpcFromSql(resp *[]db.ServicesListRow) *rmsv1.ServicesListResponse
	ServiceCreateSqlFromGrpc(req *rmsv1.ServiceCreateRequest) *db.ServiceCreateParams
	ServiceEntityGrpcFromSql(resp *db.Service) *rmsv1.Service
	ServiceCreateGrpcFromSql(resp *db.Service) *rmsv1.ServiceCreateResponse
	ServiceUpdateSqlFromGrpc(req *rmsv1.ServiceUpdateRequest) *db.ServiceUpdateParams
	ServiceUpdateGrpcFromSql(resp *db.Service) *rmsv1.ServiceUpdateResponse
	ServiceFindForUpdateGrpcFromSql(resp *db.ServiceFindForUpdateRow) *rmsv1.ServiceUpdateRequest

	// testemonials
	testemonialsListRowGrpcFromSql(resp *db.TestemonialsListRow) *rmsv1.TestemonialsListRow
	TestemonialsListGrpcFromSql(resp *[]db.TestemonialsListRow) *rmsv1.TestemonialsListResponse
	TestemonialCreateSqlFromGrpc(req *rmsv1.TestemonialCreateRequest) *db.TestemonialCreateParams
	TestemonialEntityGrpcFromSql(resp *db.Testemonial) *rmsv1.Testemonial
	TestemonialCreateGrpcFromSql(resp *db.Testemonial) *rmsv1.TestemonialCreateResponse
	TestemonialUpdateSqlFromGrpc(req *rmsv1.TestemonialUpdateRequest) *db.TestemonialUpdateParams
	TestemonialUpdateGrpcFromSql(resp *db.Testemonial) *rmsv1.TestemonialUpdateResponse
	TestemonialFindForUpdateGrpcFromSql(resp *db.TestemonialFindForUpdateRow) *rmsv1.TestemonialUpdateRequest

	//team
	teamMembersListRowGrpcFromSql(resp *db.TeamMembersListRow) *rmsv1.TeamMembersListRow
	TeamMembersListGrpcFromSql(resp *[]db.TeamMembersListRow) *rmsv1.TeamMembersListResponse
	TeamMemberCreateSqlFromGrpc(req *rmsv1.TeamMemberCreateRequest) *db.TeamMemberCreateParams
	TeamMemberEntityGrpcFromSql(resp *db.TeamMember) *rmsv1.TeamMember
	TeamMemberCreateGrpcFromSql(resp *db.TeamMember) *rmsv1.TeamMemberCreateResponse
	TeamMemberUpdateSqlFromGrpc(req *rmsv1.TeamMemberUpdateRequest) *db.TeamMemberUpdateParams
	TeamMemberUpdateGrpcFromSql(resp *db.TeamMember) *rmsv1.TeamMemberUpdateResponse
	TeamMemberFindForUpdateGrpcFromSql(resp *db.TeamMemberFindForUpdateRow) *rmsv1.TeamMemberUpdateRequest

	//projects
	projectsListRowGrpcFromSql(resp *db.ProjectsListRow) *rmsv1.ProjectsListRow
	ProjectsListGrpcFromSql(resp *[]db.ProjectsListRow) *rmsv1.ProjectsListResponse
	ProjectCreateSqlFromGrpc(req *rmsv1.ProjectCreateRequest) *db.ProjectCreateParams
	ProjectEntityGrpcFromSql(resp *db.Project) *rmsv1.Project
	ProjectCreateGrpcFromSql(resp *db.Project) *rmsv1.ProjectCreateResponse
	ProjectUpdateSqlFromGrpc(req *rmsv1.ProjectUpdateRequest) *db.ProjectUpdateParams
	ProjectUpdateGrpcFromSql(resp *db.Project) *rmsv1.ProjectUpdateResponse
	ProjectFindForUpdateGrpcFromSql(resp *db.ProjectFindForUpdateRow) *rmsv1.ProjectUpdateRequest

	//program
	programsListRowGrpcFromSql(resp *db.ProgramsListRow) *rmsv1.ProgramsListRow
	ProgramsListGrpcFromSql(resp *[]db.ProgramsListRow) *rmsv1.ProgramsListResponse
	ProgramCreateSqlFromGrpc(req *rmsv1.ProgramCreateRequest) *db.ProgramCreateParams
	ProgramEntityGrpcFromSql(resp *db.Program) *rmsv1.Program
	ProgramCreateGrpcFromSql(resp *db.Program) *rmsv1.ProgramCreateResponse
	ProgramUpdateSqlFromGrpc(req *rmsv1.ProgramUpdateRequest) *db.ProgramUpdateParams
	ProgramUpdateGrpcFromSql(resp *db.Program) *rmsv1.ProgramUpdateResponse
	ProgramFindForUpdateGrpcFromSql(resp *db.ProgramFindForUpdateRow) *rmsv1.ProgramUpdateRequest
}

type PublicAdapter struct {
	dateFormat string
}

func NewPublicAdapter() *PublicAdapter {
	return &PublicAdapter{
		dateFormat: "2006-01-02 15:04:05",
	}
}
