package repo

import (
	"context"

	db "github.com/darwishdev/bzns_pro_api/common/db/gen"
)

func (repo *PublicRepo) TeamMemberCreate(ctx context.Context, req *db.TeamMemberCreateParams) (*db.TeamMember, error) {
	category, err := repo.store.TeamMemberCreate(context.Background(), *req)

	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &category, nil
}

func (repo *PublicRepo) TeamMemberUpdate(ctx context.Context, req *db.TeamMemberUpdateParams) (*db.TeamMember, error) {
	category, err := repo.store.TeamMemberUpdate(context.Background(), *req)

	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &category, nil
}

func (repo *PublicRepo) TeamMembersList(ctx context.Context) (*[]db.TeamMembersListRow, error) {
	resp, err := repo.store.TeamMembersList(context.Background())

	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}
func (repo *PublicRepo) TeamMemberDeleteRestore(ctx context.Context, req []int32) error {
	err := repo.store.TeamMemberDeleteRestore(context.Background(), req)
	if err != nil {
		return repo.store.DbErrorParser(err, repo.errorHandler)
	}
	return nil
}

func (repo *PublicRepo) TeamMemberFindForUpdate(ctx context.Context, req *int32) (*db.TeamMemberFindForUpdateRow, error) {
	resp, err := repo.store.TeamMemberFindForUpdate(context.Background(), *req)
	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}
