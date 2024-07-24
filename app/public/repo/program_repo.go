package repo

import (
	"context"

	db "github.com/darwishdev/bzns_pro_api/common/db/gen"
)

func (repo *PublicRepo) ProgramCreate(ctx context.Context, req *db.ProgramCreateParams) (*db.Program, error) {
	category, err := repo.store.ProgramCreate(context.Background(), *req)

	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &category, nil
}

func (repo *PublicRepo) ProgramUpdate(ctx context.Context, req *db.ProgramUpdateParams) (*db.Program, error) {
	category, err := repo.store.ProgramUpdate(context.Background(), *req)

	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &category, nil
}

func (repo *PublicRepo) ProgramsList(ctx context.Context) (*[]db.ProgramsListRow, error) {
	resp, err := repo.store.ProgramsList(context.Background())

	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}
func (repo *PublicRepo) ProgramDeleteRestore(ctx context.Context, req []int32) error {
	err := repo.store.ProgramDeleteRestore(context.Background(), req)
	if err != nil {
		return repo.store.DbErrorParser(err, repo.errorHandler)
	}
	return nil
}

func (repo *PublicRepo) ProgramFindForUpdate(ctx context.Context, req *int32) (*db.ProgramFindForUpdateRow, error) {
	resp, err := repo.store.ProgramFindForUpdate(context.Background(), *req)
	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}
