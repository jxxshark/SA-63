// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/jxxshark/app/ent/patient"
	"github.com/jxxshark/app/ent/specializedappoint"
)

// PatientCreate is the builder for creating a Patient entity.
type PatientCreate struct {
	config
	mutation *PatientMutation
	hooks    []Hook
}

// SetName sets the name field.
func (pc *PatientCreate) SetName(s string) *PatientCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetAge sets the age field.
func (pc *PatientCreate) SetAge(i int) *PatientCreate {
	pc.mutation.SetAge(i)
	return pc
}

// AddAppointmentIDs adds the appointment edge to Specializedappoint by ids.
func (pc *PatientCreate) AddAppointmentIDs(ids ...int) *PatientCreate {
	pc.mutation.AddAppointmentIDs(ids...)
	return pc
}

// AddAppointment adds the appointment edges to Specializedappoint.
func (pc *PatientCreate) AddAppointment(s ...*Specializedappoint) *PatientCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pc.AddAppointmentIDs(ids...)
}

// Mutation returns the PatientMutation object of the builder.
func (pc *PatientCreate) Mutation() *PatientMutation {
	return pc.mutation
}

// Save creates the Patient in the database.
func (pc *PatientCreate) Save(ctx context.Context) (*Patient, error) {
	if _, ok := pc.mutation.Name(); !ok {
		return nil, &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if v, ok := pc.mutation.Name(); ok {
		if err := patient.NameValidator(v); err != nil {
			return nil, &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if _, ok := pc.mutation.Age(); !ok {
		return nil, &ValidationError{Name: "age", err: errors.New("ent: missing required field \"age\"")}
	}
	var (
		err  error
		node *Patient
	)
	if len(pc.hooks) == 0 {
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PatientMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pc.mutation = mutation
			node, err = pc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PatientCreate) SaveX(ctx context.Context) *Patient {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pc *PatientCreate) sqlSave(ctx context.Context) (*Patient, error) {
	pa, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	pa.ID = int(id)
	return pa, nil
}

func (pc *PatientCreate) createSpec() (*Patient, *sqlgraph.CreateSpec) {
	var (
		pa    = &Patient{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: patient.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: patient.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patient.FieldName,
		})
		pa.Name = value
	}
	if value, ok := pc.mutation.Age(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: patient.FieldAge,
		})
		pa.Age = value
	}
	if nodes := pc.mutation.AppointmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.AppointmentTable,
			Columns: []string{patient.AppointmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: specializedappoint.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return pa, _spec
}
