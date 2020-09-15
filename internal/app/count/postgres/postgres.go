package postgres

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/jmoiron/sqlx"

	"github.com/cage1016/ms-sample/internal/app/count/model"
)

type countRespository struct {
	db  *sqlx.DB
	log log.Logger
}

func New(db *sqlx.DB, log log.Logger) model.CountRespository {
	return &countRespository{db, log}
}

func (cr countRespository) Add(ctx context.Context) (err error) {
	_, err = cr.db.ExecContext(ctx, "update count set value=value +1 where value > 0;")
	if err != nil {
		level.Error(cr.log).Log("method", "add", "err", err)
	}
	return
}

func (cr countRespository) Get(ctx context.Context) (res int64, err error) {
	err = cr.db.GetContext(ctx, &res, "select * from count;")
	if err != nil {
		level.Error(cr.log).Log("method", "Get", "err", err)
	}
	return
}
