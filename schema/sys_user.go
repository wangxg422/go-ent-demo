package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type SysUser struct {
	ent.Schema
}

// Fields of the User.
func (SysUser) Fields() []ent.Field {
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
func (SysUser) Edges() []ent.Edge {
	return []ent.Edge{
		// 用户 - 部门 （多对一）
		edge.From("sys_dept", SysDept.Type).
			Ref("sys_users").
			Field("dept_id").
			Unique(), // 每个用户只属于一个部门
		edge.To("sys_roles", SysRole.Type).
			StorageKey(edge.Table("sys_user_role"), edge.Columns("user_id", "role_id"))}
}

func (SysUser) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_user"},
		entsql.WithComments(true),
		schema.Comment("系统用户表"),
	}
}
