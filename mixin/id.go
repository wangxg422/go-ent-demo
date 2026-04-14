package mixin

import (
	"go-ent-demo/schematype"
	"go-ent-demo/types"
	"go-ent-demo/util/dbutil"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type IDMixin struct {
	mixin.Schema
}

func (IDMixin) Fields() []ent.Field {
	return []ent.Field{
		// field.Int64("id").
		// 	StructTag(`json:"id,string"`).
		// 	Immutable(). // 创建后不可修改
		// 	Unique().
		// 	DefaultFunc(dbutil.IdFunc()),

		// 主键使用自定义类型，序列化时序列化为字符串，反序列化时恢复为int64
		field.Int64("id").
			GoType(types.ID(0)).
			StructTag(`json:"id"`).
			Immutable(). // 创建后不可修改
			Unique().
			SchemaType(schematype.SchemaTypeID()).
			DefaultFunc(dbutil.IDFunc),
	}
}
