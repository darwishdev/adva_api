package adapter

import (
	"encoding/json"

	"github.com/darwishdev/bzns_pro_api/common/convertor"
	db "github.com/darwishdev/bzns_pro_api/common/db/gen"
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/bznspro/v1"
)

//list

func (a *BlogsAdapter) BlogFindForUpdateGrpcFromSql(resp *db.BlogFindForUpdateRow) *rmsv1.BlogUpdateRequest {
	var links []*rmsv1.Link
	json.Unmarshal(resp.Links, &links)
	return &rmsv1.BlogUpdateRequest{
		BlogId:     resp.BlogID,
		BlogName:   resp.BlogName,
		BlogImage:  resp.BlogImage.String,
		Breif:      resp.Breif,
		Content:    resp.Content,
		CategoryId: resp.CategoryID,
		UserId:     resp.UserID,
		Views:      resp.Views.Int32,
		DateTime:   resp.DateTime.Time.Format("2006-01-02"),
		Links:      links,
		Tags:       resp.Tags,
	}

}

func (a *BlogsAdapter) blogsListRowGrpcFromSql(resp *db.BlogsListRow) *rmsv1.BlogsListRow {
	var links []*rmsv1.Link
	json.Unmarshal(resp.Links, &links)
	return &rmsv1.BlogsListRow{
		BlogId:       resp.BlogID,
		BlogName:     resp.BlogName,
		BlogImage:    resp.BlogImage.String,
		Breif:        resp.Breif,
		Content:      resp.Content,
		CategoryId:   resp.CategoryID,
		CategoryName: resp.CategoryName,
		UserId:       resp.UserID,
		Views:        resp.Views.Int32,
		DateTime:     resp.DateTime.Time.Format("2006-01-02"),
		Links:        links,
		Tags:         resp.Tags,
		CreatedAt:    resp.CreatedAt.Time.Format(a.dateFormat),
		ConfirmedAt:  resp.ConfirmedAt.Time.Format(a.dateFormat),
		UpdatedAt:    resp.UpdatedAt.Time.Format(a.dateFormat),
		DeletedAt:    resp.DeletedAt.Time.Format(a.dateFormat),
	}
}

func (a *BlogsAdapter) BlogsListGrpcFromSql(resp *[]db.BlogsListRow) *rmsv1.BlogsListResponse {
	records := make([]*rmsv1.BlogsListRow, 0)
	deletedRecords := make([]*rmsv1.BlogsListRow, 0)
	for _, v := range *resp {
		record := a.blogsListRowGrpcFromSql(&v)
		if v.DeletedAt.Valid {
			deletedRecords = append(deletedRecords, record)
		} else {
			records = append(records, record)
		}
	}
	return &rmsv1.BlogsListResponse{
		Records:        records,
		DeletedRecords: deletedRecords,
	}
}

func (a *BlogsAdapter) BlogCreateSqlFromGrpc(req *rmsv1.BlogCreateRequest) *db.BlogCreateParams {
	links, _ := json.Marshal(req.Links)
	return &db.BlogCreateParams{
		BlogName:   req.BlogName,
		Breif:      req.Breif,
		BlogImage:  convertor.ToPgType(req.BlogImage),
		Content:    req.Content,
		CategoryID: req.CategoryId,
		UserID:     req.UserId,
		Views:      convertor.ToPgTypeInt(req.Views),
		DateTime:   convertor.ToPgTimeStamp(req.DateTime),
		Links:      links,
		Tags:       req.Tags,
	}
}

func (a *BlogsAdapter) BlogUpdateSqlFromGrpc(req *rmsv1.BlogUpdateRequest) *db.BlogUpdateParams {
	links, _ := json.Marshal(req.Links)

	return &db.BlogUpdateParams{
		BlogID:     req.BlogId,
		BlogName:   req.BlogName,
		Breif:      req.Breif,
		BlogImage:  convertor.ToPgType(req.BlogImage),
		Content:    req.Content,
		CategoryID: req.CategoryId,
		UserID:     req.UserId,
		Views:      convertor.ToPgTypeInt(req.Views),
		DateTime:   convertor.ToPgTimeStamp(req.DateTime),
		Links:      links,
		Tags:       req.Tags,
	}
}
