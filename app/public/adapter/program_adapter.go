package adapter

import (
	"github.com/darwishdev/bzns_pro_api/common/convertor"
	db "github.com/darwishdev/bzns_pro_api/common/db/gen"
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/bznspro/v1"
)

func (a *PublicAdapter) programsListRowGrpcFromSql(resp *db.ProgramsListRow) *rmsv1.ProgramsListRow {

	return &rmsv1.ProgramsListRow{
		ProgramId:    resp.ProgramID,
		ProgramName:  resp.ProgramName,
		ProgramImage: resp.ProgramImage,
		Breif:        resp.Breif.String,
		CreatedAt:    resp.CreatedAt.Time.Format(a.dateFormat),
		UpdatedAt:    resp.UpdatedAt.Time.Format(a.dateFormat),
		DeletedAt:    resp.DeletedAt.Time.Format(a.dateFormat),
	}
}
func (a *PublicAdapter) ProgramsListGrpcFromSql(resp *[]db.ProgramsListRow) *rmsv1.ProgramsListResponse {
	records := make([]*rmsv1.ProgramsListRow, 0)
	deletedRecords := make([]*rmsv1.ProgramsListRow, 0)
	for _, v := range *resp {
		record := a.programsListRowGrpcFromSql(&v)
		if v.DeletedAt.Valid {
			deletedRecords = append(deletedRecords, record)
		} else {
			records = append(records, record)
		}
	}
	return &rmsv1.ProgramsListResponse{
		Records:        records,
		DeletedRecords: deletedRecords,
	}
}

func (a *PublicAdapter) ProgramCreateSqlFromGrpc(req *rmsv1.ProgramCreateRequest) *db.ProgramCreateParams {

	return &db.ProgramCreateParams{
		ProgramName:  req.ProgramName,
		Breif:        convertor.ToPgType(req.Breif),
		ProgramImage: req.ProgramImage,
	}
}
func (a *PublicAdapter) ProgramEntityGrpcFromSql(resp *db.Program) *rmsv1.Program {

	return &rmsv1.Program{
		ProgramId:    int32(resp.ProgramID),
		ProgramName:  resp.ProgramName,
		ProgramImage: resp.ProgramImage,

		CreatedAt: resp.CreatedAt.Time.Format(a.dateFormat),
		DeletedAt: resp.DeletedAt.Time.Format(a.dateFormat),
		UpdatedAt: resp.UpdatedAt.Time.Format(a.dateFormat),
	}

}
func (a *PublicAdapter) ProgramCreateGrpcFromSql(resp *db.Program) *rmsv1.ProgramCreateResponse {
	return &rmsv1.ProgramCreateResponse{
		Program: a.ProgramEntityGrpcFromSql(resp),
	}
}

func (a *PublicAdapter) ProgramUpdateSqlFromGrpc(req *rmsv1.ProgramUpdateRequest) *db.ProgramUpdateParams {

	return &db.ProgramUpdateParams{
		ProgramID:    req.ProgramId,
		ProgramName:  req.ProgramName,
		ProgramImage: req.ProgramImage,
	}
}
func (a *PublicAdapter) ProgramUpdateGrpcFromSql(resp *db.Program) *rmsv1.ProgramUpdateResponse {
	return &rmsv1.ProgramUpdateResponse{
		Program: a.ProgramEntityGrpcFromSql(resp),
	}
}
func (a *PublicAdapter) ProgramFindForUpdateGrpcFromSql(resp *db.ProgramFindForUpdateRow) *rmsv1.ProgramUpdateRequest {
	return &rmsv1.ProgramUpdateRequest{
		ProgramId:    resp.ProgramID,
		ProgramName:  resp.ProgramName,
		ProgramImage: resp.ProgramImage,
		Breif:        resp.Breif.String,
	}

}

//  ProgramFind(ctx context.Context, req *int32) (*db.ProductsSchemaMenuView, error)
