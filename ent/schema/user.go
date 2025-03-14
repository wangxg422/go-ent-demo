package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").StructTag(`json:"id,string"`).Comment("用户id"),
		field.String("user_name").StructTag(`json:"userName"`),
		field.String("nick_name").StructTag(`json:"nickName"`).Optional().Comment("用户昵称"),
		field.String("mobile").StructTag(`json:"mobile"`).Optional(),
		field.String("password").StructTag(`json:"password"`).Optional(),
		field.String("email").StructTag(`json:"email"`).Optional(),
		field.Int8("sex").StructTag(`json:"sex"`).Optional(),
		field.Int64("dept_id").StructTag(`json:"deptId,string"`).Optional().Comment("用户所属部门"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("dept", Dept.Type).
			Ref("users").
			Field("dept_id").
			Unique(),
		edge.To("roles", Role.Type).
			StorageKey(edge.Table("user_role"), edge.Columns("user_id", "role_id"))}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user"},
		entsql.WithComments(true),
		schema.Comment("系统用户表"),
	}
}
