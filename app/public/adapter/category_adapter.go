package adapter

import (
	"github.com/darwishdev/bzns_pro_api/common/convertor"
	db "github.com/darwishdev/bzns_pro_api/common/db/gen"
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/bznspro/v1"
)

func (a *PublicAdapter) categoriesListRowGrpcFromSql(resp *db.CategoriesListRow) *rmsv1.CategoriesListRow {

	return &rmsv1.CategoriesListRow{
		CategoryId:   resp.CategoryID,
		CategoryName: resp.CategoryName,
		CategoryType: resp.CategoryType,
		CreatedAt:    resp.CreatedAt.Time.Format(a.dateFormat),
		UpdatedAt:    resp.UpdatedAt.Time.Format(a.dateFormat),
		DeletedAt:    resp.DeletedAt.Time.Format(a.dateFormat),
	}
}
func (a *PublicAdapter) CategoriesListGrpcFromSql(resp *[]db.CategoriesListRow) *rmsv1.CategoriesListResponse {
	records := make([]*rmsv1.CategoriesListRow, 0)
	deletedRecords := make([]*rmsv1.CategoriesListRow, 0)
	for _, v := range *resp {
		record := a.categoriesListRowGrpcFromSql(&v)
		if v.DeletedAt.Valid {
			deletedRecords = append(deletedRecords, record)
		} else {
			records = append(records, record)
		}
	}
	return &rmsv1.CategoriesListResponse{
		Records:        records,
		DeletedRecords: deletedRecords,
	}
}

func (a *PublicAdapter) CategoryCreateSqlFromGrpc(req *rmsv1.CategoryCreateRequest) *db.CategoryCreateParams {

	return &db.CategoryCreateParams{
		CategoryName:   req.CategoryName,
		CategoryTypeID: req.CategoryTypeId,
	}
}
func (a *PublicAdapter) CategoryEntityGrpcFromSql(resp *db.Category) *rmsv1.Category {

	return &rmsv1.Category{
		CategoryId:     int32(resp.CategoryID),
		CategoryName:   resp.CategoryName,
		CategoryTypeId: resp.CategoryTypeID,
		CreatedAt:      resp.CreatedAt.Time.Format(a.dateFormat),
		DeletedAt:      resp.DeletedAt.Time.Format(a.dateFormat),
		UpdatedAt:      resp.UpdatedAt.Time.Format(a.dateFormat),
	}

}
func (a *PublicAdapter) CategoryCreateGrpcFromSql(resp *db.Category) *rmsv1.CategoryCreateResponse {
	return &rmsv1.CategoryCreateResponse{
		Category: a.CategoryEntityGrpcFromSql(resp),
	}
}

func (a *PublicAdapter) CategoryUpdateSqlFromGrpc(req *rmsv1.CategoryUpdateRequest) *db.CategoryUpdateParams {

	return &db.CategoryUpdateParams{
		CategoryID:     req.CategoryId,
		CategoryName:   req.CategoryName,
		CategoryTypeID: req.CategoryTypeId,
	}
}
func (a *PublicAdapter) CategoryUpdateGrpcFromSql(resp *db.Category) *rmsv1.CategoryUpdateResponse {
	return &rmsv1.CategoryUpdateResponse{
		Category: a.CategoryEntityGrpcFromSql(resp),
	}
}
func (a *PublicAdapter) CategoriesInputListGrpcFromSql(resp *[]db.CategoriesInputListRow) *rmsv1.CategoriesInputListResponse {
	// CategoriesInputListGrpcFromSql
	records := make([]*rmsv1.SelectInputOption, 0)
	for _, v := range *resp {
		record := convertor.ToSelectInput(v.CategoryID, v.CategoryName)
		records = append(records, record)
	}
	return &rmsv1.CategoriesInputListResponse{
		Options: records,
	}
}
func (a *PublicAdapter) CategoryFindForUpdateGrpcFromSql(resp *db.CategoryFindForUpdateRow) *rmsv1.CategoryUpdateRequest {
	return &rmsv1.CategoryUpdateRequest{
		CategoryId:     resp.CategoryID,
		CategoryName:   resp.CategoryName,
		CategoryTypeId: resp.CategoryID,
	}

}

//  CategoryFind(ctx context.Context, req *int32) (*db.ProductsSchemaMenuView, error)
