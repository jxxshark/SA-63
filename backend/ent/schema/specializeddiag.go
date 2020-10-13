package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Specializeddiag holds the schema definition for the Specializeddiag entity.
type Specializeddiag struct {
	ent.Schema
}

// Fields of the Specializeddiag.
func (Specializeddiag) Fields() []ent.Field {
	return []ent.Field{
		field.String("specializeddiacnostictype").NotEmpty(),
	}
}

// Edges of the Specializeddiag.
func (Specializeddiag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("appointment", Specializedappoint.Type).StorageKey(edge.Column("diacnosticID")), 
		edge.From("user", User.Type).Ref("specializeddiag").Unique(),
	}
}
