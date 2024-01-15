package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/tx7do/kratos-utils/entgo/mixin"
	"regexp"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "users",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
		schema.Comment("用户账号"),
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_name").
			Comment("用户名").
			Unique().
			MaxLen(50).
			NotEmpty().
			Immutable().
			Optional().
			Nillable().
			Match(regexp.MustCompile("^[a-zA-Z0-9]{4,16}$")),

		field.String("nick_name").
			Comment("昵称").
			MaxLen(128).
			Optional().
			Nillable(),

		field.String("password").
			Comment("登陆密码").
			MaxLen(255).
			Optional().
			Nillable().
			NotEmpty(),
	}
}

// Mixin of the User.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.Time{},
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

// Indexes of the User.
func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id", "user_name").Unique(),
	}
}
