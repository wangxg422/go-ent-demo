package mixin

import (
	"go-ent-demo/types"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type UpdatedMixin struct {
	mixin.Schema
}

func (UpdatedMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("updated_by").StructTag(`json:"updatedBy"`).Optional(),
		field.Time("updated_at").StructTag(`json:"updatedAt"`).
			Default(types.NewCustomTimeNow).
			UpdateDefault(types.NewCustomTimeNow).
			GoType(types.CustomTime{}).
			Optional(),
	}
}
