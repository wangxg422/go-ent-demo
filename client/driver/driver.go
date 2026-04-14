package driver

import (
	"context"

	"entgo.io/ent/dialect"
)

// ent driver的包装器，打印SQL和traceID等信息
type EntDriverWrapper struct {
	driver dialect.Driver
}

func NewEntDriverWrapper(driver dialect.Driver) *EntDriverWrapper {
	return &EntDriverWrapper{
		driver: driver,
	}
}

func (t *EntDriverWrapper) Query(ctx context.Context, query string, args any, v any) error {
	//log.Debugf("[ORM Query] SQL=%s, args=%v, %s", query, args, ctx)
	return t.driver.Query(ctx, query, args, v)
}

func (t *EntDriverWrapper) Exec(ctx context.Context, query string, args any, v any) error {
	//log.Debugf("[ORM Exec] SQL=%s, args=%v, %s", query, args, ctx)
	return t.driver.Exec(ctx, query, args, v)
}

func (t *EntDriverWrapper) Tx(ctx context.Context) (dialect.Tx, error) {
	// return t.driver.Tx(ctx)

	tx, err := t.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	return &EntTxWrapper{tx: tx}, nil
}

func (t *EntDriverWrapper) Close() error {
	return t.driver.Close()
}

func (t *EntDriverWrapper) Dialect() string {
	return t.driver.Dialect()
}
