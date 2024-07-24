package adapter

import (
	"encoding/json"

	"github.com/darwishdev/bzns_pro_api/common/convertor"
	db "github.com/darwishdev/bzns_pro_api/common/db/gen"
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/bznspro/v1"
)

//list

func (a *EventsAdapter) EventFindForUpdateGrpcFromSql(resp *db.EventFindForUpdateRow) *rmsv1.EventUpdateRequest {

	var eventPlan []*rmsv1.EventPlanRow
	if len(resp.EventPlan) > 0 {
		for _, planRow := range resp.EventPlan {
			var parsedPlanRow rmsv1.EventPlanRow
			_ = json.Unmarshal(planRow, &parsedPlanRow)

			eventPlan = append(eventPlan, &parsedPlanRow)

		}
	}
	return &rmsv1.EventUpdateRequest{
		EventId:          resp.EventID,
		EventName:        resp.EventName,
		EventLocation:    resp.EventLocation,
		EventLocationUrl: resp.EventLocationUrl,
		ConstructorName:  resp.ConstructorName,
		ConstructorTitle: resp.ConstructorTitle,
		ConstructorImage: resp.ConstructorImage,
		EventPlan:        eventPlan,
		Price:            resp.Price,
		Discount:         resp.Discount.Float32,
		EventGoals:       resp.EventGoals,
		EventBreif:       resp.EventBreif.String,
		EventDescription: resp.EventDescription.String,
		ShabDiscount:     resp.ShabDiscount.Float32,
		EventVideo:       resp.EventVideo.String,
		Tags:             resp.Tags,
		EventDate:        resp.EventDate.Time.Format("2006-01-02"),
		EventStartTime:   convertor.ToPgTimeString(resp.EventStartTime),
		EventEndTime:     convertor.ToPgTimeString(resp.EventEndTime),
		EventHours:       resp.EventHours.Int32,
		CategoryId:       resp.CategoryID,
		EventImage:       resp.EventImage.String,
	}

}

func (a *EventsAdapter) eventsListRowGrpcFromSql(resp *db.EventsListRow) *rmsv1.EventsListRow {

	var eventPlan []*rmsv1.EventPlanRow
	if len(resp.EventPlan) > 0 {
		for _, planRow := range resp.EventPlan {
			var parsedPlanRow rmsv1.EventPlanRow
			_ = json.Unmarshal(planRow, &parsedPlanRow)

			eventPlan = append(eventPlan, &parsedPlanRow)

		}
	}

	finalPrice := resp.Price
	shabDiscountfinalPrice := resp.Price

	var discountAmount float32 = 0
	var shabDiscountAmount float32 = 0

	if resp.Discount.Valid {
		disc := (resp.Price * resp.Discount.Float32 / 100)
		price := resp.Price - disc
		discountAmount = a.round(disc, 2)
		finalPrice = a.round(price, 2)
	}
	if resp.ShabDiscount.Valid {
		disc := (resp.Price * resp.ShabDiscount.Float32 / 100)
		price := resp.Price - disc
		shabDiscountAmount = a.round(disc, 2)
		shabDiscountfinalPrice = a.round(price, 2)
	}
	return &rmsv1.EventsListRow{
		EventId:                resp.EventID,
		EventName:              resp.EventName,
		EventLocation:          resp.EventLocation,
		EventLocationUrl:       resp.EventLocationUrl,
		ConstructorName:        resp.ConstructorName,
		ConstructorTitle:       resp.ConstructorTitle,
		ConstructorImage:       resp.ConstructorImage,
		EventPlan:              eventPlan,
		Price:                  resp.Price,
		Discount:               resp.Discount.Float32,
		ShabDiscount:           resp.ShabDiscount.Float32,
		EventGoals:             resp.EventGoals,
		EventBreif:             resp.EventBreif.String,
		EventDescription:       resp.EventDescription.String,
		EventVideo:             resp.EventVideo.String,
		Tags:                   resp.Tags,
		EventDate:              resp.EventDate.Time.Format("2006-01-02"),
		EventStartTime:         convertor.ToPgTimeString(resp.EventStartTime),
		EventEndTime:           convertor.ToPgTimeString(resp.EventEndTime),
		EventHours:             resp.EventHours.Int32,
		CreatedAt:              resp.CreatedAt.Time.Format(a.dateFormat),
		DeletedAt:              resp.DeletedAt.Time.Format(a.dateFormat),
		CategoryId:             resp.CategoryID,
		EventImage:             resp.EventImage.String,
		FinalPrice:             finalPrice,
		DiscountAmount:         discountAmount,
		ShabDiscountFinalPrice: shabDiscountfinalPrice,
		ShabDiscountAmount:     shabDiscountAmount,
	}
}

func (a *EventsAdapter) EventsListGrpcFromSql(resp *[]db.EventsListRow) *rmsv1.EventsListResponse {
	records := make([]*rmsv1.EventsListRow, 0)
	deletedRecords := make([]*rmsv1.EventsListRow, 0)
	for _, v := range *resp {
		record := a.eventsListRowGrpcFromSql(&v)
		if v.DeletedAt.Valid {
			deletedRecords = append(deletedRecords, record)
		} else {
			records = append(records, record)
		}
	}
	return &rmsv1.EventsListResponse{
		Records:        records,
		DeletedRecords: deletedRecords,
	}
}

func (a *EventsAdapter) EventCreateSqlFromGrpc(req *rmsv1.EventCreateRequest) *db.EventCreateParams {
	var eventPlan [][]byte
	for _, row := range req.EventPlan {
		ser, _ := json.Marshal(row)
		eventPlan = append(eventPlan, ser)
	}
	return &db.EventCreateParams{
		EventName:        req.EventName,
		EventLocation:    req.EventLocation,
		EventLocationUrl: req.EventLocationUrl,
		ConstructorName:  req.ConstructorName,
		ConstructorTitle: req.ConstructorTitle,
		ConstructorImage: req.ConstructorImage,
		EventPlan:        eventPlan,
		Tags:             req.Tags,
		EventGoals:       req.EventGoals,
		EventBreif:       convertor.ToPgType(req.EventBreif),
		EventDescription: convertor.ToPgType(req.EventDescription),
		EventVideo:       convertor.ToPgType(req.EventVideo),
		EventDate:        convertor.ToPgDate(req.EventDate),
		EventStartTime:   convertor.ToPgTime(req.EventStartTime),
		EventEndTime:     convertor.ToPgTime(req.EventEndTime),
		EventHours:       convertor.ToPgTypeInt(req.EventHours),
		CategoryID:       req.CategoryId,
		EventImage:       convertor.ToPgType(req.EventImage),
		Price:            req.Price,
		Discount:         convertor.ToPgTypeFloat(req.Discount),
		ShabDiscount:     convertor.ToPgTypeFloat(req.ShabDiscount),
	}
}

func (a *EventsAdapter) EventUpdateSqlFromGrpc(req *rmsv1.EventUpdateRequest) *db.EventUpdateParams {
	var eventPlan [][]byte
	for _, row := range req.EventPlan {
		ser, _ := json.Marshal(row)
		eventPlan = append(eventPlan, ser)
	}

	return &db.EventUpdateParams{
		EventID:          req.EventId,
		EventName:        req.EventName,
		EventLocation:    req.EventLocation,
		EventLocationUrl: req.EventLocationUrl,
		ConstructorName:  req.ConstructorName,
		ConstructorTitle: req.ConstructorTitle,
		ConstructorImage: req.ConstructorImage,
		Tags:             req.Tags,
		EventPlan:        eventPlan,
		EventGoals:       req.EventGoals,
		EventBreif:       convertor.ToPgType(req.EventBreif),
		EventDescription: convertor.ToPgType(req.EventDescription),
		EventVideo:       convertor.ToPgType(req.EventVideo),
		EventDate:        convertor.ToPgDate(req.EventDate),
		EventStartTime:   convertor.ToPgTime(req.EventStartTime),
		EventEndTime:     convertor.ToPgTime(req.EventEndTime),
		EventHours:       convertor.ToPgTypeInt(req.EventHours),
		CategoryID:       req.CategoryId,
		EventImage:       convertor.ToPgType(req.EventImage),
		Price:            req.Price,
		ShabDiscount:     convertor.ToPgTypeFloat(req.ShabDiscount),
		Discount:         convertor.ToPgTypeFloat(req.Discount),
	}
}
