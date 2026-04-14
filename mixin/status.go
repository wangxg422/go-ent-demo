package mixin

import (
	"go-ent-demo/enum"
	"go-ent-demo/schematype"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type StatusMixin struct {
	mixin.Schema
}

func (StatusMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("status").SchemaType(schematype.SchemaTypeChar1()).
			StructTag(`json:"status"`).
			Default(enum.StatusValid),
	}
}
