package schema

import (
	"go-ent-demo/mixin"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	//"entgo.io/ent/schema/field"
)

// SysRole holds the schema definition for the SysRole entity.
type SysRole struct {
	ent.Schema
}

// Annotations 修改表名称
func (SysRole) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_role"},
		entsql.WithComments(true),
		schema.Comment("系统角色表"),
	}
}

// Mixin 嵌入字段
func (SysRole) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IDMixin{},
		mixin.SortMixin{},
		mixin.StatusMixin{},
		mixin.CreatedMixin{},
		mixin.UpdatedMixin{},
		mixin.DeletedMixin{},
		mixin.RemarkMixin{},
		mixin.SoftDeleteMixin{},
	}
}

// Fields of the SysRole.
func (SysRole) Fields() []ent.Field {
	return []ent.Field{
		field.String("role_name").StructTag(`json:"roleName"`).Optional().Comment("角色名称"),
		field.String("role_code").StructTag(`json:"roleCode"`).Unique().Optional().Comment("角色编码"),
		field.Int8("menu_check_strictly").StructTag(`json:"menuCheckStrictly"`).Comment("菜单树选择项是否关联显示"),
		field.Int8("dept_check_strictly").StructTag(`json:"deptCheckStrictly"`).Comment("部门树选择项是否关联显示"),
		field.Int8("data_scope").StructTag(`json:"dataScope"`).Optional().Comment("数据权限范围(1全部数据权限 2自定义数据权限 3本部门数据权限 4本部门及以下数据权限)"),
		field.String("description").StructTag(`json:"description"`).Optional().Comment("角色描述"),
	}
}

// Edges of the SysRole.
func (SysRole) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("sysDepts", SysDept.Type).
			StorageKey(edge.Table("sys_role_dept"), edge.Columns("role_id", "dept_id")),
		edge.From("sysUsers", SysUser.Type).
			Ref("sysRoles"),
		edge.To("sysMenus", SysMenu.Type).
			StorageKey(edge.Table("sys_role_menu"), edge.Columns("role_id", "menu_id")),
	}
}
