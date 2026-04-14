package schema

import (
	"go-ent-demo/mixin"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// SysDept holds the schema definition for the SysDept entity.
type SysDept struct {
	ent.Schema
}

// Annotations 修改表名称
func (SysDept) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_dept"},
		entsql.WithComments(true),
		schema.Comment("系统部门表"),
	}
}

// Mixin 嵌入字段
func (SysDept) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IDMixin{},
		mixin.ParentIDMixin{},
		mixin.SortMixin{},
		mixin.StatusMixin{},
		mixin.CreatedMixin{},
		mixin.UpdatedMixin{},
		mixin.DeletedMixin{},
		mixin.RemarkMixin{},
		mixin.SoftDeleteMixin{},
	}
}

// Fields of the SysDept.
func (SysDept) Fields() []ent.Field {
	return []ent.Field{
		field.String("ancestors").StructTag(`json:"ancestors"`).Optional().Comment("祖先部门列表,以','分隔"),
		field.String("dept_name").StructTag(`json:"deptName"`).Optional().Comment("部门名称"),
		field.String("dept_code").StructTag(`json:"deptCode"`).Unique().Optional().Comment("部门编码"),
		field.String("leader").StructTag(`json:"leader"`).Optional().Comment("负责人"),
		field.String("phone").StructTag(`json:"phone"`).Optional().Comment("部门联系电话"),
		field.String("email").StructTag(`json:"email"`).Optional().Comment("部门电子邮箱"),
		field.String("address").StructTag(`json:"address"`).Optional().Comment("部门地址"),
	}
}

// Edges of the SysDept.
func (SysDept) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("sysUsers", SysUser.Type),
		edge.From("sysRoles", SysRole.Type).Ref("sysDepts"),
		// 添加一对多的（O2M ）自引用，即一个节点有且仅有一个父节点，有多个子节点
		edge.To("children", SysDept.Type).
			From("parent").
			Unique().
			Field("parent_id"),
	}
}
