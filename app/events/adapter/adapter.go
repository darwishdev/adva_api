package adapter

import (
	"math"

	db "github.com/darwishdev/bzns_pro_api/common/db/gen"
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/bznspro/v1"
)

type EventsAdapterInterface interface {
	eventsListRowGrpcFromSql(resp *db.EventsListRow) *rmsv1.EventsListRow
	EventsListGrpcFromSql(resp *[]db.EventsListRow) *rmsv1.EventsListResponse
	EventCreateSqlFromGrpc(req *rmsv1.EventCreateRequest) *db.EventCreateParams
	EventUpdateSqlFromGrpc(req *rmsv1.EventUpdateRequest) *db.EventUpdateParams
	EventFindForUpdateGrpcFromSql(resp *db.EventFindForUpdateRow) *rmsv1.EventUpdateRequest

	EventRequestsListGrpcFromSql(resp *[]db.EventRequestsListRow) *rmsv1.EventRequestsListResponse
	EventRequestCreateSqlFromGrpc(req *rmsv1.EventRequestCreateRequest) *db.EventRequestCreateParams
	EventRequestUpdateSqlFromGrpc(req *rmsv1.EventRequestUpdateRequest) *db.EventRequestUpdateParams
}

func Round(value float32, places int) float32 {
	shift := math.Pow(10, float64(places))
	return float32(math.Round(float64(value)*shift) / shift)
}

type EventsAdapter struct {
	dateFormat string
	round      func(value float32, places int) float32
}

func NewEventsAdapter() EventsAdapterInterface {
	return &EventsAdapter{
		dateFormat: "2006-01-02 15:04:05",
		round:      Round,
	}
}
