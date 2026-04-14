package mixin

import (
	"go-ent-demo/schematype"
	"go-ent-demo/types"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type ParentIDMixin struct {
	mixin.Schema
}

func (ParentIDMixin) Fields() []ent.Field {
	return []ent.Field{
		// field.Int64("parent_id").StructTag(`json:"parentId,string"`).
		// 	Optional().
		// 	Nillable(),

		// parentID作为外键使用自定义类型，序列化时序列化为字符串，反序列化时恢复为int64
		field.Int64("parent_id").
			StructTag(`json:"parentId"`).
			Nillable().Optional().
			GoType(types.ID(0)).
			SchemaType(schematype.SchemaTypeID()),
	}
}
