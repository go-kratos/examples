package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"

	"github.com/tx7do/go-utils/entgo/mixin"
)

// DictDetail holds the schema definition for the DictDetail entity.
type DictDetail struct {
	ent.Schema
}

func (DictDetail) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "dict_detail",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
	}
}

// Fields of the DictDetail.
func (DictDetail) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("dict_id").
			Comment("字典ID").
			Default(0).
			Optional().
			Nillable(),

		field.Int32("order_no").
			Comment("排序ID").
			Default(0).
			Optional().
			Nillable(),

		field.String("label").
			Comment("字典标签").
			Optional().
			Nillable(),

		field.String("value").
			Comment("字典值").
			Optional().
			Nillable(),
	}
}

// Mixin of the DictDetail.
func (DictDetail) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.Time{},
		mixin.CreateBy{},
	}
}
