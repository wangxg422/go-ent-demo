package driver

import (
	"context"

	"entgo.io/ent/dialect"
)

// ent driver的事务包装器，打印SQL和traceID等信息
type EntTxWrapper struct {
	tx dialect.Tx
}

func (t *EntTxWrapper) Exec(ctx context.Context, query string, args, v any) error {
	//log.Debugf("[ORM Tx Exec] SQL=%s, args=%v, %s", query, args, ctx)
	return t.tx.Exec(ctx, query, args, v)
}

func (t *EntTxWrapper) Query(ctx context.Context, query string, args, v any) error {
	//log.Debugf("[ORM Tx Query] SQL=%s, args=%v, %s", query, args, ctx)
	return t.tx.Query(ctx, query, args, v)
}

func (t *EntTxWrapper) Commit() error {
	return t.tx.Commit()
}

func (t *EntTxWrapper) Rollback() error {
	return t.tx.Rollback()
}
