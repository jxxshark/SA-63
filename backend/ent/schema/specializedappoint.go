package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Specializedappoint holds the schema definition for the Specializedappoint entity.
type Specializedappoint struct {
	ent.Schema
}

// Fields of the Specializedappoint.
func (Specializedappoint) Fields() []ent.Field {
	return []ent.Field{
		field.Time("Date"),
	}
}

// Edges of the Specializedappoint.
func (Specializedappoint) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("appointment").Unique(),
		edge.From("patient", Patient.Type).Ref("appointment").Unique(),
		edge.From("specializeddiag", Specializeddiag.Type).Ref("appointment").Unique(),
	}
}
