package mixin

import (
	"go-ent-demo/types"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type CreatedMixin struct {
	mixin.Schema
}

func (CreatedMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("created_by").StructTag(`json:"createdBy"`).Optional(),
		field.Time("created_at").StructTag(`json:"createdAt"`).
			Default(types.NewCustomTimeNow).
			UpdateDefault(types.NewCustomTimeNow).
			GoType(types.CustomTime{}).
			Optional(),
		// field.Time("created_at").StructTag(`json:"createdAt"`).Default(time.Now).Optional(),
		// 使用自定义时间类型，自定义序列化格式
		// field.Other("created_at", types.CustomTime{}).StructTag(`json:"createdAt"`).
		// 	SchemaType(map[string]string{
		// 		dialect.MySQL:    "datetime",
		// 		dialect.SQLite:   "datetime",
		// 		dialect.Postgres: "timestamp",
		// 	}).
		// 	Default(func() types.CustomTime {
		// 		return types.NewCustomTimeNow()
		// 	}).Optional(),
	}
}
