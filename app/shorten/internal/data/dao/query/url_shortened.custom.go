package query

import (
	"context"

	"gorm.io/gorm"
)

// UseQueryOption  使用QueryOption 初始化
func UseCondition(ctx context.Context, db *gorm.DB, cond *Condition) *uRLShortenedDo {
	url := Use(db).URLShortened
	do := url.WithContext(ctx)

	if cond.Where.Id > 0 {
		do = do.Where(url.ID.Eq(cond.Where.Id))
	}

	if cond.Where.Code != "" {
		do = do.Where(url.URLCode.Eq(cond.Where.Code))
	}

	if cond.Where.IsDeleted {
		do = do.Where(url.DeletedAt.Gt(0))
	} else {
		do = do.Where(url.DeletedAt.Eq(0))
	}

	if cond.Paging.Offset > 0 {
		do = do.Offset(cond.Paging.Offset)
	}

	if cond.Paging.Limit <= 0 {
		cond.Paging.Limit = DefaultPagingLimit
	}

	if cond.Paging.Limit > MaxPagingLimit {
		cond.Paging.Limit = MaxPagingLimit
	}

	return do.Limit(cond.Paging.Limit)
}
