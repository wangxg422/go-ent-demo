package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type DeletedMixin struct {
	mixin.Schema
}

func (DeletedMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("deleted_by").StructTag(`json:"-"`).Optional(),
		field.Time("deleted_at").StructTag(`json:"-"`).Optional(),
	}
}
