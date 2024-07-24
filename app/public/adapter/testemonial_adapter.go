package adapter

import (
	"github.com/darwishdev/bzns_pro_api/common/convertor"
	db "github.com/darwishdev/bzns_pro_api/common/db/gen"
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/bznspro/v1"
	"github.com/rs/zerolog/log"
)

func (a *PublicAdapter) testemonialsListRowGrpcFromSql(resp *db.TestemonialsListRow) *rmsv1.TestemonialsListRow {

	return &rmsv1.TestemonialsListRow{
		TestemonialId:    resp.TestemonialID,
		TestemonialName:  resp.TestemonialName,
		TestemonialTitle: resp.TestemonialTitle,
		TestemonialImage: resp.TestemonialImage.String,
		Breif:            resp.Breif,
		CreatedAt:        resp.CreatedAt.Time.Format(a.dateFormat),
		UpdatedAt:        resp.UpdatedAt.Time.Format(a.dateFormat),
		DeletedAt:        resp.DeletedAt.Time.Format(a.dateFormat),
	}
}
func (a *PublicAdapter) TestemonialsListGrpcFromSql(resp *[]db.TestemonialsListRow) *rmsv1.TestemonialsListResponse {
	records := make([]*rmsv1.TestemonialsListRow, 0)
	deletedRecords := make([]*rmsv1.TestemonialsListRow, 0)
	for _, v := range *resp {
		record := a.testemonialsListRowGrpcFromSql(&v)
		if v.DeletedAt.Valid {
			deletedRecords = append(deletedRecords, record)
		} else {
			records = append(records, record)
		}
	}
	return &rmsv1.TestemonialsListResponse{
		Records:        records,
		DeletedRecords: deletedRecords,
	}
}

func (a *PublicAdapter) TestemonialCreateSqlFromGrpc(req *rmsv1.TestemonialCreateRequest) *db.TestemonialCreateParams {
	log.Debug().Interface("aasdasdasd", req).Msg("hjoolals")
	return &db.TestemonialCreateParams{
		TestemonialName:  req.TestemonialName,
		TestemonialTitle: req.TestemonialTitle,
		Breif:            req.Breif,
		TestemonialImage: convertor.ToPgType(req.TestemonialImage),
	}
}
func (a *PublicAdapter) TestemonialEntityGrpcFromSql(resp *db.Testemonial) *rmsv1.Testemonial {

	return &rmsv1.Testemonial{
		TestemonialId:    int32(resp.TestemonialID),
		TestemonialName:  resp.TestemonialName,
		TestemonialImage: resp.TestemonialImage.String,
		TestemonialTitle: resp.TestemonialTitle,
		CreatedAt:        resp.CreatedAt.Time.Format(a.dateFormat),
		DeletedAt:        resp.DeletedAt.Time.Format(a.dateFormat),
		UpdatedAt:        resp.UpdatedAt.Time.Format(a.dateFormat),
	}

}
func (a *PublicAdapter) TestemonialCreateGrpcFromSql(resp *db.Testemonial) *rmsv1.TestemonialCreateResponse {
	return &rmsv1.TestemonialCreateResponse{
		Testemonial: a.TestemonialEntityGrpcFromSql(resp),
	}
}

func (a *PublicAdapter) TestemonialUpdateSqlFromGrpc(req *rmsv1.TestemonialUpdateRequest) *db.TestemonialUpdateParams {

	return &db.TestemonialUpdateParams{
		TestemonialID:    req.TestemonialId,
		TestemonialTitle: req.TestemonialTitle,
		Breif:            req.Breif,
		TestemonialName:  req.TestemonialName,
		TestemonialImage: convertor.ToPgType(req.TestemonialImage),
	}
}
func (a *PublicAdapter) TestemonialUpdateGrpcFromSql(resp *db.Testemonial) *rmsv1.TestemonialUpdateResponse {
	return &rmsv1.TestemonialUpdateResponse{
		Testemonial: a.TestemonialEntityGrpcFromSql(resp),
	}
}
func (a *PublicAdapter) TestemonialFindForUpdateGrpcFromSql(resp *db.TestemonialFindForUpdateRow) *rmsv1.TestemonialUpdateRequest {
	return &rmsv1.TestemonialUpdateRequest{
		TestemonialTitle: resp.TestemonialTitle,
		TestemonialId:    resp.TestemonialID,
		TestemonialName:  resp.TestemonialName,
		TestemonialImage: resp.TestemonialImage.String,
		Breif:            resp.Breif,
	}

}

//  TestemonialFind(ctx context.Context, req *int32) (*db.ProductsSchemaMenuView, error)
