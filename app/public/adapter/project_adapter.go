package adapter

import (
	db "github.com/darwishdev/bzns_pro_api/common/db/gen"
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/bznspro/v1"
)

func (a *PublicAdapter) projectsListRowGrpcFromSql(resp *db.ProjectsListRow) *rmsv1.ProjectsListRow {

	return &rmsv1.ProjectsListRow{
		ProjectId:    resp.ProjectID,
		ProjectName:  resp.ProjectName,
		ProjectImage: resp.ProjectImage,
		CategoryId:   resp.CategoryID,
		CategoryName: resp.CategoryName,
		CreatedAt:    resp.CreatedAt.Time.Format(a.dateFormat),
		UpdatedAt:    resp.UpdatedAt.Time.Format(a.dateFormat),
		DeletedAt:    resp.DeletedAt.Time.Format(a.dateFormat),
	}
}
func (a *PublicAdapter) ProjectsListGrpcFromSql(resp *[]db.ProjectsListRow) *rmsv1.ProjectsListResponse {
	records := make([]*rmsv1.ProjectsListRow, 0)
	deletedRecords := make([]*rmsv1.ProjectsListRow, 0)
	for _, v := range *resp {
		record := a.projectsListRowGrpcFromSql(&v)
		if v.DeletedAt.Valid {
			deletedRecords = append(deletedRecords, record)
		} else {
			records = append(records, record)
		}
	}
	return &rmsv1.ProjectsListResponse{
		Records:        records,
		DeletedRecords: deletedRecords,
	}
}

func (a *PublicAdapter) ProjectCreateSqlFromGrpc(req *rmsv1.ProjectCreateRequest) *db.ProjectCreateParams {

	return &db.ProjectCreateParams{
		ProjectName:  req.ProjectName,
		ProjectImage: req.ProjectImage,
		CategoryID:   req.CategoryId,
	}
}
func (a *PublicAdapter) ProjectEntityGrpcFromSql(resp *db.Project) *rmsv1.Project {

	return &rmsv1.Project{
		ProjectId:    int32(resp.ProjectID),
		ProjectName:  resp.ProjectName,
		ProjectImage: resp.ProjectImage,
		CategoryId:   resp.CategoryID,
		CreatedAt:    resp.CreatedAt.Time.Format(a.dateFormat),
		DeletedAt:    resp.DeletedAt.Time.Format(a.dateFormat),
		UpdatedAt:    resp.UpdatedAt.Time.Format(a.dateFormat),
	}

}
func (a *PublicAdapter) ProjectCreateGrpcFromSql(resp *db.Project) *rmsv1.ProjectCreateResponse {
	return &rmsv1.ProjectCreateResponse{
		Project: a.ProjectEntityGrpcFromSql(resp),
	}
}

func (a *PublicAdapter) ProjectUpdateSqlFromGrpc(req *rmsv1.ProjectUpdateRequest) *db.ProjectUpdateParams {

	return &db.ProjectUpdateParams{
		ProjectID:    req.ProjectId,
		ProjectName:  req.ProjectName,
		ProjectImage: req.ProjectImage,
		CategoryID:   req.CategoryId,
	}
}
func (a *PublicAdapter) ProjectUpdateGrpcFromSql(resp *db.Project) *rmsv1.ProjectUpdateResponse {
	return &rmsv1.ProjectUpdateResponse{
		Project: a.ProjectEntityGrpcFromSql(resp),
	}
}
func (a *PublicAdapter) ProjectFindForUpdateGrpcFromSql(resp *db.ProjectFindForUpdateRow) *rmsv1.ProjectUpdateRequest {
	return &rmsv1.ProjectUpdateRequest{
		ProjectId:    resp.ProjectID,
		ProjectName:  resp.ProjectName,
		ProjectImage: resp.ProjectImage,
		CategoryId:   resp.CategoryID,
	}

}

//  ProjectFind(ctx context.Context, req *int32) (*db.ProductsSchemaMenuView, error)
