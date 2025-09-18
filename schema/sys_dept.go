package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Dept holds the schema definition for the Dept entity.
type SysDept struct {
	ent.Schema
}

// Fields of the Dept.
func (SysDept) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").StructTag(`json:"id,string"`).Comment("部门id"),
		field.Int64("parent_id").StructTag(`json:"parentId,string"`).Comment("父级部门id").Optional(),
		field.String("ancestors").StructTag(`json:"ancestors"`).Optional().Comment("祖先部门列表").Optional(),
		field.String("dept_name").StructTag(`json:"deptName"`).Optional().Comment("部门名称").Optional(),
		field.String("dept_code").StructTag(`json:"deptCode"`).Optional().Comment("部门编码").Optional(),
		field.String("leader").StructTag(`json:"leader"`).Optional().Comment("负责人").Optional(),
		field.String("phone").StructTag(`json:"phone"`).Optional().Comment("部门联系电话").Optional(),
		field.String("email").StructTag(`json:"email"`).Optional().Comment("部门电子邮箱").Optional(),
	}
}

// Edges of the Dept.
func (SysDept) Edges() []ent.Edge {
	return []ent.Edge{
		// 部门 -> 用户 (一对多)
		edge.To("sys_users", SysUser.Type),
	}
}

func (SysDept) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_dept"},
		entsql.WithComments(true),
		schema.Comment("系统部门表"),
	}
}
