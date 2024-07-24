package repo

import (
	"context"

	db "github.com/darwishdev/bzns_pro_api/common/db/gen"
)

func (repo *EventsRepo) EventRequestsList(ctx context.Context) (*[]db.EventRequestsListRow, error) {
	resp, err := repo.store.EventRequestsList(context.Background())

	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}

func (repo *EventsRepo) EventRequestCreate(ctx context.Context, req *db.EventRequestCreateParams) (*db.EventRequest, error) {
	resp, err := repo.store.EventRequestCreate(context.Background(), *req)

	if err != nil {

		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}

func (repo *EventsRepo) EventRequestUpdate(ctx context.Context, req *db.EventRequestUpdateParams) (*db.EventRequest, error) {
	resp, err := repo.store.EventRequestUpdate(context.Background(), *req)
	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}

func (repo *EventsRepo) EventRequestDeleteRestore(ctx context.Context, req []int32) error {
	err := repo.store.EventRequestDeleteRestore(context.Background(), req)
	if err != nil {
		return repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return nil
}
