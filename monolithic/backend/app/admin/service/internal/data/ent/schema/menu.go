package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/tx7do/go-utils/entgo/mixin"
)

// Menu holds the schema definition for the Menu entity.
type Menu struct {
	ent.Schema
}

func (Menu) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "menu",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
	}
}

// Fields of the Menu.
func (Menu) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("id").
			Comment("id").
			StructTag(`json:"id,omitempty"`).
			Positive().
			Immutable().
			Unique(),

		field.Int32("parent_id").
			Comment("上一层菜单ID").
			Optional().
			Nillable(),

		field.Int32("order_no").
			Comment("排序ID").
			Default(0).
			Optional().
			Nillable(),

		field.String("name").
			Comment("菜单名称").
			Default("").
			MaxLen(32).
			NotEmpty().
			Optional().
			Nillable(),

		field.String("title").
			Comment("菜单标题").
			Default("").
			NotEmpty().
			Optional().
			Nillable(),

		field.Enum("type").
			Comment("菜单类型 FOLDER: 目录 MENU: 菜单 BUTTON: 按钮").
			Values(
				"FOLDER",
				"MENU",
				"BUTTON",
			).
			Optional().
			Nillable(),

		field.String("path").
			Comment("路径,当其类型为'按钮'的时候对应的数据操作名,例如:/user.service.v1.UserService/Login").
			Default("").
			Optional().
			Nillable(),

		field.String("component").
			Comment("前端页面组件").
			Default("").
			Optional().
			Nillable(),

		field.String("icon").
			Comment("图标").
			Default("").
			MaxLen(128).
			Optional().
			Nillable(),

		field.Bool("is_ext").
			Comment("是否外链").
			Default(false).
			Optional().
			Nillable(),

		field.String("ext_url").
			Comment("外链地址").
			MaxLen(255).
			Optional().
			Nillable(),

		field.Strings("permissions").
			Comment("权限代码 例如:sys:menu").
			SchemaType(map[string]string{
				dialect.MySQL:    "json",
				dialect.Postgres: "jsonb",
			}).
			Optional(),

		field.String("redirect").
			Comment("跳转路径").
			Optional().
			Nillable(),

		field.String("current_active_menu").
			Comment("当前激活路径").
			Optional().
			Nillable(),

		field.Bool("keep_alive").
			Comment("是否缓存").
			Default(false).
			Optional().
			Nillable(),

		field.Bool("show").
			Comment("是否显示").
			Default(true).
			Optional().
			Nillable(),

		field.Bool("hide_tab").
			Comment("是否显示在标签页导航").
			Default(true).
			Optional().
			Nillable(),

		field.Bool("hide_menu").
			Comment("是否显示在菜单导航").
			Default(true).
			Optional().
			Nillable(),

		field.Bool("hide_breadcrumb").
			Comment("是否显示在面包屑导航").
			Default(true).
			Optional().
			Nillable(),
	}
}

func (Menu) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.SwitchStatus{},
		mixin.Time{},
		mixin.CreateBy{},
	}
}

// Edges of the Menu.
func (Menu) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			To("children", Menu.Type).Annotations(entproto.Field(10)).
			From("parent").Unique().Field("parent_id").Annotations(entproto.Field(11)),
	}
}
