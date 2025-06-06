package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/tx7do/go-utils/entgo/mixin"
	"regexp"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "user",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").
			Comment("用户名").
			Unique().
			MaxLen(50).
			NotEmpty().
			Immutable().
			Optional().
			Nillable().
			Match(regexp.MustCompile("^[a-zA-Z0-9]{4,16}$")),

		field.String("password").
			Comment("登陆密码").
			MaxLen(255).
			Optional().
			Nillable().
			NotEmpty(),

		field.Uint32("role_id").
			Comment("角色ID").
			Optional().
			Nillable(),

		field.Uint32("org_id").
			Comment("部门ID").
			Optional().
			Nillable(),

		field.Uint32("position_id").
			Comment("职位ID").
			Optional().
			Nillable(),

		field.Uint32("work_id").
			Comment("员工工号").
			Optional().
			Nillable(),

		field.String("nick_name").
			Comment("昵称").
			MaxLen(128).
			Optional().
			Nillable(),

		field.String("real_name").
			Comment("真实名字").
			MaxLen(128).
			Optional().
			Nillable(),

		field.String("email").
			Comment("电子邮箱").
			MaxLen(127).
			Optional().
			Nillable(),

		field.String("phone").
			Comment("手机号码").
			Default("").
			MaxLen(11).
			Optional().
			Nillable(),

		field.String("avatar").
			Comment("头像").
			MaxLen(1023).
			Optional().
			Nillable(),

		field.Enum("gender").
			Comment("性别").
			Values(
				"UNKNOWN",
				"MALE",
				"FEMALE",
			).
			Optional().
			Nillable(),

		field.String("address").
			Comment("地址").
			Default("").
			MaxLen(2048).
			Optional().
			Nillable(),

		field.String("description").
			Comment("个人说明").
			MaxLen(1023).
			Optional().
			Nillable(),

		field.Enum("authority").
			Comment("授权").
			Optional().
			Nillable().
			Values(
				"SYS_ADMIN",
				"CUSTOMER_USER",
				"GUEST_USER",
				"REFRESH_TOKEN",
			).
			Default("CUSTOMER_USER"),

		field.Int64("last_login_time").
			Comment("最后一次登陆的时间").
			Optional().
			Nillable(),

		field.String("last_login_ip").
			Comment("最后一次登陆的IP").
			Default("").
			MaxLen(64).
			Optional().
			Nillable(),
	}
}

// Mixin of the User.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.CreateBy{},
		mixin.Time{},
		mixin.SwitchStatus{},
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

// Indexes of the User.
func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id", "username").Unique(),
	}
}
