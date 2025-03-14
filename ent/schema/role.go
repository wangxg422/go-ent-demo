package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Role holds the schema definition for the Role entity.
type Role struct {
	ent.Schema
}

// Fields of the Role.
func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").StructTag(`json:"id,string"`).Comment("角色id"),
		field.String("role_name").StructTag(`json:"roleName"`).Optional().Comment("角色名称"),
		field.String("role_code").StructTag(`json:"roleCode"`).Optional().Comment("角色编码"),
	}
}

// Edges of the Role.
func (Role) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("depts", Dept.Type).
			StorageKey(edge.Table("role_dept"), edge.Columns("role_id", "dept_id")),
		edge.From("users", User.Type).
			Ref("roles"),
	}
}

func (Role) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "role"},
		entsql.WithComments(true),
		schema.Comment("系统角色表"),
	}
}

