package adapter

import (
	"github.com/darwishdev/bzns_pro_api/common/convertor"
	db "github.com/darwishdev/bzns_pro_api/common/db/gen"
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/bznspro/v1"
)

//list

func (a *EventsAdapter) eventRequestsListRowGrpcFromSql(resp *db.EventRequestsListRow) *rmsv1.EventRequestsListRow {
	finalPrice := resp.Price
	discountAmount := float32(0)
	if resp.Discount.Valid {
		disc := (resp.Price * resp.Discount.Float32 / 100)
		price := resp.Price - disc
		discountAmount = a.round(disc, 2)
		finalPrice = a.round(price, 2)

	}
	return &rmsv1.EventRequestsListRow{
		EventRequestId:  resp.EventRequestID,
		EventId:         resp.EventID,
		EventName:       resp.EventName,
		EventImage:      resp.EventImage.String,
		UserEmail:       resp.UserEmail,
		UserPhone:       resp.UserPhone,
		ConstructorName: resp.ConstructorName,
		CategoryId:      resp.CategoryID,
		CategoryName:    resp.CategoryName,
		RequestStatusId: resp.RequestStatusID,
		RequestStatus:   resp.RequestStatus,
		Price:           resp.Price,
		IsShab:          resp.IsShab.Bool,
		Discount:        resp.Discount.Float32,
		FinalPrice:      finalPrice,
		DiscountAmount:  float32(discountAmount),
		CreatedAt:       resp.CreatedAt.Time.Format(a.dateFormat),
		UpdatedAt:       resp.UpdatedAt.Time.Format(a.dateFormat),
		DeletedAt:       resp.DeletedAt.Time.Format(a.dateFormat),
	}
}

func (a *EventsAdapter) EventRequestsListGrpcFromSql(resp *[]db.EventRequestsListRow) *rmsv1.EventRequestsListResponse {
	records := make([]*rmsv1.EventRequestsListRow, 0)
	deletedRecords := make([]*rmsv1.EventRequestsListRow, 0)
	for _, v := range *resp {
		record := a.eventRequestsListRowGrpcFromSql(&v)
		if v.DeletedAt.Valid {
			deletedRecords = append(deletedRecords, record)
		} else {
			records = append(records, record)
		}
	}
	return &rmsv1.EventRequestsListResponse{
		Records:        records,
		DeletedRecords: deletedRecords,
	}
}

func (a *EventsAdapter) EventRequestCreateSqlFromGrpc(req *rmsv1.EventRequestCreateRequest) *db.EventRequestCreateParams {

	return &db.EventRequestCreateParams{
		EventID:   req.EventId,
		UserName:  req.UserName,
		UserEmail: req.UserEmail,
		UserPhone: req.UserPhone,
		IsShab:    convertor.ToPgTypeBool(req.IsShab),
		City:      convertor.ToPgType(req.City),
		Price:     req.Price,
		Discount:  convertor.ToPgTypeFloat(req.Discount),
	}
}
func (a *EventsAdapter) EventRequestUpdateSqlFromGrpc(req *rmsv1.EventRequestUpdateRequest) *db.EventRequestUpdateParams {

	return &db.EventRequestUpdateParams{
		EventRequestID:  req.EventRequestId,
		RequestStatusID: req.RequestStatusId,
	}
}
