// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/jxxshark/app/ent/patient"
	"github.com/jxxshark/app/ent/predicate"
	"github.com/jxxshark/app/ent/specializedappoint"
	"github.com/jxxshark/app/ent/specializeddiag"
	"github.com/jxxshark/app/ent/user"
)

// SpecializedappointUpdate is the builder for updating Specializedappoint entities.
type SpecializedappointUpdate struct {
	config
	hooks      []Hook
	mutation   *SpecializedappointMutation
	predicates []predicate.Specializedappoint
}

// Where adds a new predicate for the builder.
func (su *SpecializedappointUpdate) Where(ps ...predicate.Specializedappoint) *SpecializedappointUpdate {
	su.predicates = append(su.predicates, ps...)
	return su
}

// SetDate sets the Date field.
func (su *SpecializedappointUpdate) SetDate(t time.Time) *SpecializedappointUpdate {
	su.mutation.SetDate(t)
	return su
}

// SetUserID sets the user edge to User by id.
func (su *SpecializedappointUpdate) SetUserID(id int) *SpecializedappointUpdate {
	su.mutation.SetUserID(id)
	return su
}

// SetNillableUserID sets the user edge to User by id if the given value is not nil.
func (su *SpecializedappointUpdate) SetNillableUserID(id *int) *SpecializedappointUpdate {
	if id != nil {
		su = su.SetUserID(*id)
	}
	return su
}

// SetUser sets the user edge to User.
func (su *SpecializedappointUpdate) SetUser(u *User) *SpecializedappointUpdate {
	return su.SetUserID(u.ID)
}

// SetPatientID sets the patient edge to Patient by id.
func (su *SpecializedappointUpdate) SetPatientID(id int) *SpecializedappointUpdate {
	su.mutation.SetPatientID(id)
	return su
}

// SetNillablePatientID sets the patient edge to Patient by id if the given value is not nil.
func (su *SpecializedappointUpdate) SetNillablePatientID(id *int) *SpecializedappointUpdate {
	if id != nil {
		su = su.SetPatientID(*id)
	}
	return su
}

// SetPatient sets the patient edge to Patient.
func (su *SpecializedappointUpdate) SetPatient(p *Patient) *SpecializedappointUpdate {
	return su.SetPatientID(p.ID)
}

// SetSpecializeddiagID sets the specializeddiag edge to Specializeddiag by id.
func (su *SpecializedappointUpdate) SetSpecializeddiagID(id int) *SpecializedappointUpdate {
	su.mutation.SetSpecializeddiagID(id)
	return su
}

// SetNillableSpecializeddiagID sets the specializeddiag edge to Specializeddiag by id if the given value is not nil.
func (su *SpecializedappointUpdate) SetNillableSpecializeddiagID(id *int) *SpecializedappointUpdate {
	if id != nil {
		su = su.SetSpecializeddiagID(*id)
	}
	return su
}

// SetSpecializeddiag sets the specializeddiag edge to Specializeddiag.
func (su *SpecializedappointUpdate) SetSpecializeddiag(s *Specializeddiag) *SpecializedappointUpdate {
	return su.SetSpecializeddiagID(s.ID)
}

// Mutation returns the SpecializedappointMutation object of the builder.
func (su *SpecializedappointUpdate) Mutation() *SpecializedappointMutation {
	return su.mutation
}

// ClearUser clears the user edge to User.
func (su *SpecializedappointUpdate) ClearUser() *SpecializedappointUpdate {
	su.mutation.ClearUser()
	return su
}

// ClearPatient clears the patient edge to Patient.
func (su *SpecializedappointUpdate) ClearPatient() *SpecializedappointUpdate {
	su.mutation.ClearPatient()
	return su
}

// ClearSpecializeddiag clears the specializeddiag edge to Specializeddiag.
func (su *SpecializedappointUpdate) ClearSpecializeddiag() *SpecializedappointUpdate {
	su.mutation.ClearSpecializeddiag()
	return su
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (su *SpecializedappointUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(su.hooks) == 0 {
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SpecializedappointMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *SpecializedappointUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SpecializedappointUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SpecializedappointUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *SpecializedappointUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   specializedappoint.Table,
			Columns: specializedappoint.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: specializedappoint.FieldID,
			},
		},
	}
	if ps := su.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Date(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: specializedappoint.FieldDate,
		})
	}
	if su.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.PatientCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.PatientIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.SpecializeddiagCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.SpecializeddiagIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{specializedappoint.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// SpecializedappointUpdateOne is the builder for updating a single Specializedappoint entity.
type SpecializedappointUpdateOne struct {
	config
	hooks    []Hook
	mutation *SpecializedappointMutation
}

// SetDate sets the Date field.
func (suo *SpecializedappointUpdateOne) SetDate(t time.Time) *SpecializedappointUpdateOne {
	suo.mutation.SetDate(t)
	return suo
}

// SetUserID sets the user edge to User by id.
func (suo *SpecializedappointUpdateOne) SetUserID(id int) *SpecializedappointUpdateOne {
	suo.mutation.SetUserID(id)
	return suo
}

// SetNillableUserID sets the user edge to User by id if the given value is not nil.
func (suo *SpecializedappointUpdateOne) SetNillableUserID(id *int) *SpecializedappointUpdateOne {
	if id != nil {
		suo = suo.SetUserID(*id)
	}
	return suo
}

// SetUser sets the user edge to User.
func (suo *SpecializedappointUpdateOne) SetUser(u *User) *SpecializedappointUpdateOne {
	return suo.SetUserID(u.ID)
}

// SetPatientID sets the patient edge to Patient by id.
func (suo *SpecializedappointUpdateOne) SetPatientID(id int) *SpecializedappointUpdateOne {
	suo.mutation.SetPatientID(id)
	return suo
}

// SetNillablePatientID sets the patient edge to Patient by id if the given value is not nil.
func (suo *SpecializedappointUpdateOne) SetNillablePatientID(id *int) *SpecializedappointUpdateOne {
	if id != nil {
		suo = suo.SetPatientID(*id)
	}
	return suo
}

// SetPatient sets the patient edge to Patient.
func (suo *SpecializedappointUpdateOne) SetPatient(p *Patient) *SpecializedappointUpdateOne {
	return suo.SetPatientID(p.ID)
}

// SetSpecializeddiagID sets the specializeddiag edge to Specializeddiag by id.
func (suo *SpecializedappointUpdateOne) SetSpecializeddiagID(id int) *SpecializedappointUpdateOne {
	suo.mutation.SetSpecializeddiagID(id)
	return suo
}

// SetNillableSpecializeddiagID sets the specializeddiag edge to Specializeddiag by id if the given value is not nil.
func (suo *SpecializedappointUpdateOne) SetNillableSpecializeddiagID(id *int) *SpecializedappointUpdateOne {
	if id != nil {
		suo = suo.SetSpecializeddiagID(*id)
	}
	return suo
}

// SetSpecializeddiag sets the specializeddiag edge to Specializeddiag.
func (suo *SpecializedappointUpdateOne) SetSpecializeddiag(s *Specializeddiag) *SpecializedappointUpdateOne {
	return suo.SetSpecializeddiagID(s.ID)
}

// Mutation returns the SpecializedappointMutation object of the builder.
func (suo *SpecializedappointUpdateOne) Mutation() *SpecializedappointMutation {
	return suo.mutation
}

// ClearUser clears the user edge to User.
func (suo *SpecializedappointUpdateOne) ClearUser() *SpecializedappointUpdateOne {
	suo.mutation.ClearUser()
	return suo
}

// ClearPatient clears the patient edge to Patient.
func (suo *SpecializedappointUpdateOne) ClearPatient() *SpecializedappointUpdateOne {
	suo.mutation.ClearPatient()
	return suo
}

// ClearSpecializeddiag clears the specializeddiag edge to Specializeddiag.
func (suo *SpecializedappointUpdateOne) ClearSpecializeddiag() *SpecializedappointUpdateOne {
	suo.mutation.ClearSpecializeddiag()
	return suo
}

// Save executes the query and returns the updated entity.
func (suo *SpecializedappointUpdateOne) Save(ctx context.Context) (*Specializedappoint, error) {

	var (
		err  error
		node *Specializedappoint
	)
	if len(suo.hooks) == 0 {
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SpecializedappointMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			mut = suo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, suo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SpecializedappointUpdateOne) SaveX(ctx context.Context) *Specializedappoint {
	s, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return s
}

// Exec executes the query on the entity.
func (suo *SpecializedappointUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SpecializedappointUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *SpecializedappointUpdateOne) sqlSave(ctx context.Context) (s *Specializedappoint, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   specializedappoint.Table,
			Columns: specializedappoint.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: specializedappoint.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Specializedappoint.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := suo.mutation.Date(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: specializedappoint.FieldDate,
		})
	}
	if suo.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.PatientCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.PatientIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.SpecializeddiagCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.SpecializeddiagIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	s = &Specializedappoint{config: suo.config}
	_spec.Assign = s.assignValues
	_spec.ScanValues = s.scanValues()
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{specializedappoint.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return s, nil
}
