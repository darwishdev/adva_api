package repo

import (
	"context"

	db "github.com/darwishdev/bzns_pro_api/common/db/gen"
)

type EventsRepoInterface interface {
	EventsInputList(ctx context.Context) (*[]db.EventsInputListRow, error)
	EventsList(ctx context.Context) (*[]db.EventsListRow, error)
	EventCreate(ctx context.Context, req *db.EventCreateParams) (*db.Event, error)
	EventUpdate(ctx context.Context, req *db.EventUpdateParams) (*db.Event, error)
	EventDeleteRestore(ctx context.Context, req []int32) error
	EventRequestsList(ctx context.Context) (*[]db.EventRequestsListRow, error)
	EventRequestCreate(ctx context.Context, req *db.EventRequestCreateParams) (*db.EventRequest, error)
	EventRequestUpdate(ctx context.Context, req *db.EventRequestUpdateParams) (*db.EventRequest, error)
	EventRequestDeleteRestore(ctx context.Context, req []int32) error
	EventFindForUpdate(ctx context.Context, req *int32) (*db.EventFindForUpdateRow, error) 
}

type EventsRepo struct {
	store        db.Store
	errorHandler map[string]string
}

func NewEventsRepo(store db.Store) EventsRepoInterface {
	errorHandler := map[string]string{
		"events_event_name_key": "eventName",
	}
	return &EventsRepo{
		store:        store,
		errorHandler: errorHandler,
	}
}
