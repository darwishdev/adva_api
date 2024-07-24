package repo

import (
	"context"

	db "github.com/darwishdev/bzns_pro_api/common/db/gen"
)

func (repo *BlogsRepo) BlogsList(ctx context.Context) (*[]db.BlogsListRow, error) {
	resp, err := repo.store.BlogsList(context.Background())

	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}

func (repo *BlogsRepo) BlogCreate(ctx context.Context, req *db.BlogCreateParams) (*db.Blog, error) {
	resp, err := repo.store.BlogCreate(context.Background(), *req)

	if err != nil {

		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}

func (repo *BlogsRepo) BlogFindForUpdate(ctx context.Context, req *int32) (*db.BlogFindForUpdateRow, error) {
	resp, err := repo.store.BlogFindForUpdate(context.Background(), *req)
	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}

func (repo *BlogsRepo) BlogUpdate(ctx context.Context, req *db.BlogUpdateParams) (*db.Blog, error) {
	resp, err := repo.store.BlogUpdate(context.Background(), *req)
	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}

func (repo *BlogsRepo) BlogDeleteRestore(ctx context.Context, req []int32) error {
	err := repo.store.BlogDeleteRestore(context.Background(), req)
	if err != nil {
		return repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return nil
}
