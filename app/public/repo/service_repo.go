package repo

import (
	"context"

	db "github.com/darwishdev/bzns_pro_api/common/db/gen"
)

func (repo *PublicRepo) ServiceCreate(ctx context.Context, req *db.ServiceCreateParams) (*db.Service, error) {
	category, err := repo.store.ServiceCreate(context.Background(), *req)

	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &category, nil
}

func (repo *PublicRepo) ServiceUpdate(ctx context.Context, req *db.ServiceUpdateParams) (*db.Service, error) {
	category, err := repo.store.ServiceUpdate(context.Background(), *req)

	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &category, nil
}

func (repo *PublicRepo) ServicesList(ctx context.Context) (*[]db.ServicesListRow, error) {
	resp, err := repo.store.ServicesList(context.Background())

	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}
func (repo *PublicRepo) ServiceDeleteRestore(ctx context.Context, req []int32) error {
	err := repo.store.ServiceDeleteRestore(context.Background(), req)
	if err != nil {
		return repo.store.DbErrorParser(err, repo.errorHandler)
	}
	return nil
}

func (repo *PublicRepo) ServiceFindForUpdate(ctx context.Context, req *int32) (*db.ServiceFindForUpdateRow, error) {
	resp, err := repo.store.ServiceFindForUpdate(context.Background(), *req)
	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}
