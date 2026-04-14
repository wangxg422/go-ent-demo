package schema

import (
	"go-ent-demo/mixin"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// SysMenu holds the schema definition for the SysMenu entity.
type SysMenu struct {
	ent.Schema
}

// Annotations 修改表名称
func (SysMenu) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_menu"},
		entsql.WithComments(true),
		schema.Comment("系统菜单权限表"),
	}
}

// Mixin 嵌入字段
func (SysMenu) Mixin() []ent.Mixin {
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

// Fields of the SysMenu.
func (SysMenu) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").StructTag(`json:"name"`).Comment("菜单名称"),
		field.String("path").StructTag(`json:"path"`).Optional().Comment("路由地址"),
		field.Int8("type").StructTag(`json:"type"`).Comment("菜单类型 (1: 目录, 2: 菜单, 3: 按钮)"),
		field.String("component").StructTag(`json:"component"`).Default("").Comment("组件路径"),
		field.String("permission").StructTag(`json:"permission"`).Optional().Default("").Comment("权限编码, 多个之间以','分隔"),
		field.String("redirect").StructTag(`json:"redirect"`).Default("").Comment("如果没有特殊情况，父级路由的 redirect 属性，不需要指定，默认会指向第一个子路由"),
		field.String("meta_title").StructTag(`json:"metaTitle"`).Comment("用于配置页面的标题，会在菜单和标签页中显示。一般会配合国际化使用"),
		field.String("meta_icon").StructTag(`json:"metaIcon"`).Default("").Comment("菜单和标签页的图标，一般会配合图标库使用，如果是http链接，会自动加载图片"),
		field.String("meta_active_icon").StructTag(`json:"metaActiveIcon"`).Default("").Comment("菜单和标签页的图标，一般会配合图标库使用，如果是http链接，会自动加载图片"),
		field.Int8("meta_keep_alive").StructTag(`json:"metaKeepAlive"`).Default(2).Comment("页面是否开启缓存(1: 开启, 2: 不开启, 默认值: 2)，开启后页面会缓存，不会重新加载，仅在标签页启用时有效"),
		field.Int8("meta_hide_in_menu").StructTag(`json:"metaHideInMenu"`).Default(2).Comment("配置页面是否在菜单中隐藏(1: 隐藏, 2: 不隐藏, 默认值: 2)，隐藏后页面不会在菜单中显示"),
		field.Int8("meta_hide_in_tab").StructTag(`json:"metaHideInTab"`).Default(2).Comment("配置页面是否在标签页中隐藏 (1: 隐藏, 2: 不隐藏, 默认值: 2)，隐藏后页面不会在标签页中显示"),
		field.Int8("meta_hide_in_breadcrumb").StructTag(`json:"metaHideInBreadcrumb"`).Default(2).Comment("配置页面是否在面包屑中隐藏 (1: 隐藏, 2: 不隐藏， 默认值2)，隐藏后页面不会在面包屑中显示"),
		field.Int8("meta_hide_children_in_menu").StructTag(`json:"metaHideChildrenInMenu"`).Default(2).Comment("配置页面的子页面是否在菜单中隐藏 (1: 隐藏 2:不隐藏, 默认值: 2)，隐藏后子页面不会在菜单中显示"),
		field.String("meta_authority").StructTag(`json:"metaAuthority"`).Default("").Comment("配置页面的权限，只有拥有对应权限的用户才能访问页面，不配置则不需要权限。默认值: NULL。当前为保留字段"),
		field.String("meta_badge").StructTag(`json:"metaBadge"`).Default("").Comment("配置页面的徽标，会在菜单显示。默认值''"),
		field.Int8("meta_badge_type").StructTag(`json:"metaBadgeType"`).Default(2).Comment("配置页面的徽标类型 (1: 小红点dot, 2: 文本normal, 默认值: 2)"),
		field.Int8("meta_badge_variants").StructTag(`json:"metaBadgeVariants"`).Default(4).Comment("配置页面的徽标颜色 (1：''default''，2： ''destructive'', 3: ''primary'', 4: ''success'', 5: ''warning'', 6: string, 默认值：4)"),
		field.Int8("meta_full_path_key").StructTag(`json:"metaFullPathKey"`).Default(1).Comment("是否将路由的完整路径作为tab key (1: 是, 2: 否, 默认值: 1)"),
		field.String("meta_active_path").StructTag(`json:"metaActivePath"`).Default("").Comment("配置当前激活的菜单，有时候页面没有显示在菜单内，需要激活父级菜单时使用"),
		field.Int8("meta_affix_tab").StructTag(`json:"metaAffixTab"`).Default(2).Comment("配置页面是否固定标签页(1: 是, 2: 否, 默认值: 2), 固定后页面不可关闭"),
		field.Int("meta_affix_tab_order").StructTag(`json:"metaAffixTabOrder"`).Default(0).Comment("配置页面固定标签页的排序, 采用升序排序。"),
		field.String("meta_iframe_src").StructTag(`json:"metaIframeSrc"`).Default("").Comment("配置内嵌页面的 iframe 地址，设置后会在当前页面内嵌对应的页面"),
		field.Int8("meta_ignore_access").StructTag(`json:"metaIgnoreAccess"`).Default(2).Comment("配置页面是否忽略权限 (1: 忽略权限, 2: 不忽略权限, 默认值: 2)，直接可以访问"),
		field.String("meta_link").StructTag(`json:"metaLink"`).Default("").Comment("配置外链跳转路径，会在新窗口打开"),
		field.Int("meta_max_num_of_open_tab").StructTag(`json:"metaMaxNumOfOpenTab"`).Default(-1).Comment("用于配置标签页最大打开数量，设置后会在打开新标签页时自动关闭最早打开的标签页(仅在打开同名标签页时生效)"),
		field.Int8("meta_menu_visible_with_forbidden").StructTag(`json:"metaMenuVisibleWithForbidden"`).Default(2).Comment("配置页面在菜单可以看到，但是访问会被重定向到403, (1: 启用, 2: 不启用, 默认值: 2)"),
		field.Int8("meta_open_in_new_window").StructTag(`json:"metaOpenInNewWindow"`).Default(2).Comment("菜单是否在新窗口打开页面(1: 是, 2: 否, 默认值: 2)"),
		field.Int("meta_order").StructTag(`json:"metaOrder"`).Default(0).Comment("配置页面的排序，用于路由到菜单排序。注意: 排序仅针对一级菜单有效，二级菜单的排序需要在对应的一级菜单中按代码顺序设置"),
		field.String("meta_query").StructTag(`json:"metaQuery"`).Default("").Comment("配置页面的菜单参数，会在菜单中传递给页面"),
		field.Int8("meta_no_basic_layout").StructTag(`json:"metaNoBasicLayout"`).Default(2).Comment("配置当前路由不使用基础布局，仅在顶级时生效。默认情况下，所有的路由都会被包裹在基础布局中（包含顶部以及侧边等导航部件），如果你的页面不需要这些部件，可以设置为true。(1: 	不使用基础布局 2: 使用基础布局, 默认值: 2)"),
	}
}

// Edges of the SysMenu.
func (SysMenu) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("sysRoles", SysRole.Type).
			Ref("sysMenus"),
		// 添加一对多的（O2M ）自引用，即一个节点有且仅有一个父节点，有多个子节点
		edge.To("children", SysMenu.Type).
			From("parent").
			Unique().
			Field("parent_id"),
	}
}
