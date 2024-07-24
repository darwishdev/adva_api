package adapter

import (
	"github.com/darwishdev/bzns_pro_api/common/convertor"
	db "github.com/darwishdev/bzns_pro_api/common/db/gen"
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/bznspro/v1"
)

func (a *PublicAdapter) servicesListRowGrpcFromSql(resp *db.ServicesListRow) *rmsv1.ServicesListRow {

	return &rmsv1.ServicesListRow{
		ServiceId:    resp.ServiceID,
		ServiceName:  resp.ServiceName,
		ServiceImage: resp.ServiceImage,
		Tags:         resp.Tags,
		Breif:        resp.Breif.String,
		CreatedAt:    resp.CreatedAt.Time.Format(a.dateFormat),
		UpdatedAt:    resp.UpdatedAt.Time.Format(a.dateFormat),
		DeletedAt:    resp.DeletedAt.Time.Format(a.dateFormat),
	}
}
func (a *PublicAdapter) ServicesListGrpcFromSql(resp *[]db.ServicesListRow) *rmsv1.ServicesListResponse {
	records := make([]*rmsv1.ServicesListRow, 0)
	deletedRecords := make([]*rmsv1.ServicesListRow, 0)
	for _, v := range *resp {
		record := a.servicesListRowGrpcFromSql(&v)
		if v.DeletedAt.Valid {
			deletedRecords = append(deletedRecords, record)
		} else {
			records = append(records, record)
		}
	}
	return &rmsv1.ServicesListResponse{
		Records:        records,
		DeletedRecords: deletedRecords,
	}
}

func (a *PublicAdapter) ServiceCreateSqlFromGrpc(req *rmsv1.ServiceCreateRequest) *db.ServiceCreateParams {

	return &db.ServiceCreateParams{
		ServiceName:  req.ServiceName,
		Breif:        convertor.ToPgType(req.Breif),
		ServiceImage: req.ServiceImage,
	}
}
func (a *PublicAdapter) ServiceEntityGrpcFromSql(resp *db.Service) *rmsv1.Service {

	return &rmsv1.Service{
		ServiceId:    int32(resp.ServiceID),
		ServiceName:  resp.ServiceName,
		ServiceImage: resp.ServiceImage,
		Tags:         resp.Tags,

		CreatedAt: resp.CreatedAt.Time.Format(a.dateFormat),
		DeletedAt: resp.DeletedAt.Time.Format(a.dateFormat),
		UpdatedAt: resp.UpdatedAt.Time.Format(a.dateFormat),
	}

}
func (a *PublicAdapter) ServiceCreateGrpcFromSql(resp *db.Service) *rmsv1.ServiceCreateResponse {
	return &rmsv1.ServiceCreateResponse{
		Service: a.ServiceEntityGrpcFromSql(resp),
	}
}

func (a *PublicAdapter) ServiceUpdateSqlFromGrpc(req *rmsv1.ServiceUpdateRequest) *db.ServiceUpdateParams {

	return &db.ServiceUpdateParams{
		ServiceID:    req.ServiceId,
		ServiceName:  req.ServiceName,
		ServiceImage: req.ServiceImage,
		Tags:         req.Tags,
	}
}
func (a *PublicAdapter) ServiceUpdateGrpcFromSql(resp *db.Service) *rmsv1.ServiceUpdateResponse {
	return &rmsv1.ServiceUpdateResponse{
		Service: a.ServiceEntityGrpcFromSql(resp),
	}
}
func (a *PublicAdapter) ServiceFindForUpdateGrpcFromSql(resp *db.ServiceFindForUpdateRow) *rmsv1.ServiceUpdateRequest {
	return &rmsv1.ServiceUpdateRequest{
		ServiceId:    resp.ServiceID,
		ServiceName:  resp.ServiceName,
		ServiceImage: resp.ServiceImage,
		Tags:         resp.Tags,
		Breif:        resp.Breif.String,
	}

}

//  ServiceFind(ctx context.Context, req *int32) (*db.ProductsSchemaMenuView, error)
