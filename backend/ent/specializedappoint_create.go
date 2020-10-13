// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/jxxshark/app/ent/patient"
	"github.com/jxxshark/app/ent/specializedappoint"
	"github.com/jxxshark/app/ent/specializeddiag"
	"github.com/jxxshark/app/ent/user"
)

// SpecializedappointCreate is the builder for creating a Specializedappoint entity.
type SpecializedappointCreate struct {
	config
	mutation *SpecializedappointMutation
	hooks    []Hook
}

// SetDate sets the Date field.
func (sc *SpecializedappointCreate) SetDate(t time.Time) *SpecializedappointCreate {
	sc.mutation.SetDate(t)
	return sc
}

// SetUserID sets the user edge to User by id.
func (sc *SpecializedappointCreate) SetUserID(id int) *SpecializedappointCreate {
	sc.mutation.SetUserID(id)
	return sc
}

// SetNillableUserID sets the user edge to User by id if the given value is not nil.
func (sc *SpecializedappointCreate) SetNillableUserID(id *int) *SpecializedappointCreate {
	if id != nil {
		sc = sc.SetUserID(*id)
	}
	return sc
}

// SetUser sets the user edge to User.
func (sc *SpecializedappointCreate) SetUser(u *User) *SpecializedappointCreate {
	return sc.SetUserID(u.ID)
}

// SetPatientID sets the patient edge to Patient by id.
func (sc *SpecializedappointCreate) SetPatientID(id int) *SpecializedappointCreate {
	sc.mutation.SetPatientID(id)
	return sc
}

// SetNillablePatientID sets the patient edge to Patient by id if the given value is not nil.
func (sc *SpecializedappointCreate) SetNillablePatientID(id *int) *SpecializedappointCreate {
	if id != nil {
		sc = sc.SetPatientID(*id)
	}
	return sc
}

// SetPatient sets the patient edge to Patient.
func (sc *SpecializedappointCreate) SetPatient(p *Patient) *SpecializedappointCreate {
	return sc.SetPatientID(p.ID)
}

// SetSpecializeddiagID sets the specializeddiag edge to Specializeddiag by id.
func (sc *SpecializedappointCreate) SetSpecializeddiagID(id int) *SpecializedappointCreate {
	sc.mutation.SetSpecializeddiagID(id)
	return sc
}

// SetNillableSpecializeddiagID sets the specializeddiag edge to Specializeddiag by id if the given value is not nil.
func (sc *SpecializedappointCreate) SetNillableSpecializeddiagID(id *int) *SpecializedappointCreate {
	if id != nil {
		sc = sc.SetSpecializeddiagID(*id)
	}
	return sc
}

// SetSpecializeddiag sets the specializeddiag edge to Specializeddiag.
func (sc *SpecializedappointCreate) SetSpecializeddiag(s *Specializeddiag) *SpecializedappointCreate {
	return sc.SetSpecializeddiagID(s.ID)
}

// Mutation returns the SpecializedappointMutation object of the builder.
func (sc *SpecializedappointCreate) Mutation() *SpecializedappointMutation {
	return sc.mutation
}

// Save creates the Specializedappoint in the database.
func (sc *SpecializedappointCreate) Save(ctx context.Context) (*Specializedappoint, error) {
	if _, ok := sc.mutation.Date(); !ok {
		return nil, &ValidationError{Name: "Date", err: errors.New("ent: missing required field \"Date\"")}
	}
	var (
		err  error
		node *Specializedappoint
	)
	if len(sc.hooks) == 0 {
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SpecializedappointMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sc.mutation = mutation
			node, err = sc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			mut = sc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SpecializedappointCreate) SaveX(ctx context.Context) *Specializedappoint {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sc *SpecializedappointCreate) sqlSave(ctx context.Context) (*Specializedappoint, error) {
	s, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	s.ID = int(id)
	return s, nil
}

func (sc *SpecializedappointCreate) createSpec() (*Specializedappoint, *sqlgraph.CreateSpec) {
	var (
		s     = &Specializedappoint{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: specializedappoint.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: specializedappoint.FieldID,
			},
		}
	)
	if value, ok := sc.mutation.Date(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: specializedappoint.FieldDate,
		})
		s.Date = value
	}
	if nodes := sc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   specializedappoint.UserTable,
			Columns: []string{specializedappoint.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.PatientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   specializedappoint.PatientTable,
			Columns: []string{specializedappoint.PatientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patient.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.SpecializeddiagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   specializedappoint.SpecializeddiagTable,
			Columns: []string{specializedappoint.SpecializeddiagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: specializeddiag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return s, _spec
}
