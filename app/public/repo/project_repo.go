package repo

import (
	"context"

	db "github.com/darwishdev/bzns_pro_api/common/db/gen"
)

func (repo *PublicRepo) ProjectCreate(ctx context.Context, req *db.ProjectCreateParams) (*db.Project, error) {
	category, err := repo.store.ProjectCreate(context.Background(), *req)

	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &category, nil
}

func (repo *PublicRepo) ProjectUpdate(ctx context.Context, req *db.ProjectUpdateParams) (*db.Project, error) {
	category, err := repo.store.ProjectUpdate(context.Background(), *req)

	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &category, nil
}

func (repo *PublicRepo) ProjectsList(ctx context.Context) (*[]db.ProjectsListRow, error) {
	resp, err := repo.store.ProjectsList(context.Background())

	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}
func (repo *PublicRepo) ProjectDeleteRestore(ctx context.Context, req []int32) error {
	err := repo.store.ProjectDeleteRestore(context.Background(), req)
	if err != nil {
		return repo.store.DbErrorParser(err, repo.errorHandler)
	}
	return nil
}

func (repo *PublicRepo) ProjectFindForUpdate(ctx context.Context, req *int32) (*db.ProjectFindForUpdateRow, error) {
	resp, err := repo.store.ProjectFindForUpdate(context.Background(), *req)
	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}
