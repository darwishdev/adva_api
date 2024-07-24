package repo

import (
	"context"

	db "github.com/darwishdev/bzns_pro_api/common/db/gen"
)

type BlogsRepoInterface interface {
	BlogsList(ctx context.Context) (*[]db.BlogsListRow, error)
	BlogCreate(ctx context.Context, req *db.BlogCreateParams) (*db.Blog, error)
	BlogFindForUpdate(ctx context.Context, req *int32) (*db.BlogFindForUpdateRow, error)
	BlogUpdate(ctx context.Context, req *db.BlogUpdateParams) (*db.Blog, error)
	BlogDeleteRestore(ctx context.Context, req []int32) error
}

type BlogsRepo struct {
	store        db.Store
	errorHandler map[string]string
}

func NewBlogsRepo(store db.Store) BlogsRepoInterface {
	errorHandler := map[string]string{
		"events_event_name_key": "eventName",
	}
	return &BlogsRepo{
		store:        store,
		errorHandler: errorHandler,
	}
}
