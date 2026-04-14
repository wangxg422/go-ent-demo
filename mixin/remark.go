package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type RemarkMixin struct {
	mixin.Schema
}

func (RemarkMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("remark").StructTag(`json:"remark"`).Optional(),
	}
}
