package client

import (
	"context"
	"go-ent-demo/entcore"
)

type CoreBaseDao struct {
}

// 核心库客户端
func (d CoreBaseDao) CoreDB() *entcore.Client {
	return CoreDBClient()
}

// 执行事务
func (d CoreBaseDao) CoreDBWithTx(c context.Context, fn func(tx *entcore.Tx) error) error {
	return CoreDBWithTx(c, CoreDBClient(), fn)
}
