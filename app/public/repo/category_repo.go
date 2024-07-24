package repo

import (
	"context"

	db "github.com/darwishdev/bzns_pro_api/common/db/gen"
)

func (repo *PublicRepo) CategoryCreate(ctx context.Context, req *db.CategoryCreateParams) (*db.Category, error) {
	category, err := repo.store.CategoryCreate(context.Background(), *req)

	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &category, nil
}

func (repo *PublicRepo) CategoryUpdate(ctx context.Context, req *db.CategoryUpdateParams) (*db.Category, error) {
	category, err := repo.store.CategoryUpdate(context.Background(), *req)

	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &category, nil
}

func (repo *PublicRepo) CategoriesList(ctx context.Context, req int32) (*[]db.CategoriesListRow, error) {
	resp, err := repo.store.CategoriesList(context.Background(), req)

	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}
func (repo *PublicRepo) CategoriesInputList(ctx context.Context, req int32) (*[]db.CategoriesInputListRow, error) {
	resp, err := repo.store.CategoriesInputList(context.Background(), req)

	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}
func (repo *PublicRepo) CategoryDeleteRestore(ctx context.Context, req []int32) error {
	err := repo.store.CategoryDeleteRestore(context.Background(), req)
	if err != nil {
		return repo.store.DbErrorParser(err, repo.errorHandler)
	}
	return nil
}

func (repo *PublicRepo) CategoryFindForUpdate(ctx context.Context, req *int32) (*db.CategoryFindForUpdateRow, error) {
	resp, err := repo.store.CategoryFindForUpdate(context.Background(), *req)
	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}
