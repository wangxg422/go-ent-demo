package schema

import (
	"go-ent-demo/mixin"
	"go-ent-demo/schematype"
	"go-ent-demo/types"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// SysUser holds the schema definition for the SysUser entity.
type SysUser struct {
	ent.Schema
}

// Annotations 修改表名称
func (SysUser) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_user"},
		entsql.WithComments(true),
		schema.Comment("系统用户表"),
	}
}

// Mixin 嵌入字段
func (SysUser) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IDMixin{},
		mixin.CreatedMixin{},
		mixin.UpdatedMixin{},
		mixin.DeletedMixin{},
		mixin.RemarkMixin{},
		mixin.SoftDeleteMixin{},
	}
}

// Fields of the SysUser.
func (SysUser) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_name").StructTag(`json:"userName"`).Unique().Comment("用户名"),
		field.String("nick_name").StructTag(`json:"nickName"`).Optional().Comment("用户昵称"),
		field.String("real_name").StructTag(`json:"realName"`).Comment("真实姓名"),
		field.String("gender").StructTag(`json:"gender"`).Optional().Comment("性别"),
		field.String("staff_id").StructTag(`json:"staffId"`).Comment("员工编号"),
		field.String("password").StructTag(`json:"password"`).Optional(),
		// dept_id为外键
		field.Int64("dept_id").Optional().Nillable().GoType(types.ID(0)).SchemaType(schematype.SchemaTypeID()).Comment("用户所属部门"),
		field.String("user_type").StructTag(`json:"userType"`).Optional().Comment("用户类型"),
		field.String("email").StructTag(`json:"email"`).Optional(),
		field.String("phone_number").StructTag(`json:"phoneNumber"`).Optional(),
		field.String("avatar").StructTag(`json:"avatar"`).Optional().Comment("用户头像地址"),
		field.String("user_status").StructTag(`json:"userStatus"`).SchemaType(map[string]string{
			dialect.MySQL:    "char(1)",
			dialect.Postgres: "char(1)",
		}).Optional().Comment("用户状态(0正常,1禁用,2禁止登录)"),
		field.String("address").StructTag(`json:"address"`).Optional().Comment("用户联系地址"),
	}
}

// Edges of the SysUser.
func (SysUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("sysDept", SysDept.Type).
			Ref("sysUsers").
			Field("dept_id").
			Unique(),
		edge.To("sysRoles", SysRole.Type).
			StorageKey(edge.Table("sys_user_role"), edge.Columns("user_id", "role_id")),
	}
}
