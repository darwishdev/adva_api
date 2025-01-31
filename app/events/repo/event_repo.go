package repo

import (
	"context"

	db "github.com/darwishdev/bzns_pro_api/common/db/gen"
)

func (repo *EventsRepo) EventsInputList(ctx context.Context) (*[]db.EventsInputListRow, error) {
	resp, err := repo.store.EventsInputList(context.Background())
	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}

func (repo *EventsRepo) EventsList(ctx context.Context) (*[]db.EventsListRow, error) {
	resp, err := repo.store.EventsList(context.Background())

	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}

func (repo *EventsRepo) EventCreate(ctx context.Context, req *db.EventCreateParams) (*db.Event, error) {
	resp, err := repo.store.EventCreate(context.Background(), *req)

	if err != nil {

		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}


func (repo *EventsRepo) EventFindForUpdate(ctx context.Context, req *int32) (*db.EventFindForUpdateRow, error) {
	resp, err := repo.store.EventFindForUpdate(context.Background(), *req)
	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}

func (repo *EventsRepo) EventUpdate(ctx context.Context, req *db.EventUpdateParams) (*db.Event, error) {
	resp, err := repo.store.EventUpdate(context.Background(), *req)
	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}

func (repo *EventsRepo) EventDeleteRestore(ctx context.Context, req []int32) error {
	err := repo.store.EventDeleteRestore(context.Background(), req)
	if err != nil {
		return repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return nil
}
