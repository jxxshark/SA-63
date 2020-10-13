// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/jxxshark/app/ent/patient"
	"github.com/jxxshark/app/ent/predicate"
	"github.com/jxxshark/app/ent/specializedappoint"
	"github.com/jxxshark/app/ent/specializeddiag"
	"github.com/jxxshark/app/ent/user"
)

// SpecializedappointQuery is the builder for querying Specializedappoint entities.
type SpecializedappointQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.Specializedappoint
	// eager-loading edges.
	withUser            *UserQuery
	withPatient         *PatientQuery
	withSpecializeddiag *SpecializeddiagQuery
	withFKs             bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (sq *SpecializedappointQuery) Where(ps ...predicate.Specializedappoint) *SpecializedappointQuery {
	sq.predicates = append(sq.predicates, ps...)
	return sq
}

// Limit adds a limit step to the query.
func (sq *SpecializedappointQuery) Limit(limit int) *SpecializedappointQuery {
	sq.limit = &limit
	return sq
}

// Offset adds an offset step to the query.
func (sq *SpecializedappointQuery) Offset(offset int) *SpecializedappointQuery {
	sq.offset = &offset
	return sq
}

// Order adds an order step to the query.
func (sq *SpecializedappointQuery) Order(o ...OrderFunc) *SpecializedappointQuery {
	sq.order = append(sq.order, o...)
	return sq
}

// QueryUser chains the current query on the user edge.
func (sq *SpecializedappointQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: sq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(specializedappoint.Table, specializedappoint.FieldID, sq.sqlQuery()),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, specializedappoint.UserTable, specializedappoint.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPatient chains the current query on the patient edge.
func (sq *SpecializedappointQuery) QueryPatient() *PatientQuery {
	query := &PatientQuery{config: sq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(specializedappoint.Table, specializedappoint.FieldID, sq.sqlQuery()),
			sqlgraph.To(patient.Table, patient.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, specializedappoint.PatientTable, specializedappoint.PatientColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySpecializeddiag chains the current query on the specializeddiag edge.
func (sq *SpecializedappointQuery) QuerySpecializeddiag() *SpecializeddiagQuery {
	query := &SpecializeddiagQuery{config: sq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(specializedappoint.Table, specializedappoint.FieldID, sq.sqlQuery()),
			sqlgraph.To(specializeddiag.Table, specializeddiag.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, specializedappoint.SpecializeddiagTable, specializedappoint.SpecializeddiagColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Specializedappoint entity in the query. Returns *NotFoundError when no specializedappoint was found.
func (sq *SpecializedappointQuery) First(ctx context.Context) (*Specializedappoint, error) {
	sSlice, err := sq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(sSlice) == 0 {
		return nil, &NotFoundError{specializedappoint.Label}
	}
	return sSlice[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sq *SpecializedappointQuery) FirstX(ctx context.Context) *Specializedappoint {
	s, err := sq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return s
}

// FirstID returns the first Specializedappoint id in the query. Returns *NotFoundError when no id was found.
func (sq *SpecializedappointQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{specializedappoint.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (sq *SpecializedappointQuery) FirstXID(ctx context.Context) int {
	id, err := sq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Specializedappoint entity in the query, returns an error if not exactly one entity was returned.
func (sq *SpecializedappointQuery) Only(ctx context.Context) (*Specializedappoint, error) {
	sSlice, err := sq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(sSlice) {
	case 1:
		return sSlice[0], nil
	case 0:
		return nil, &NotFoundError{specializedappoint.Label}
	default:
		return nil, &NotSingularError{specializedappoint.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sq *SpecializedappointQuery) OnlyX(ctx context.Context) *Specializedappoint {
	s, err := sq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return s
}

// OnlyID returns the only Specializedappoint id in the query, returns an error if not exactly one id was returned.
func (sq *SpecializedappointQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{specializedappoint.Label}
	default:
		err = &NotSingularError{specializedappoint.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sq *SpecializedappointQuery) OnlyIDX(ctx context.Context) int {
	id, err := sq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Specializedappoints.
func (sq *SpecializedappointQuery) All(ctx context.Context) ([]*Specializedappoint, error) {
	if err := sq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return sq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (sq *SpecializedappointQuery) AllX(ctx context.Context) []*Specializedappoint {
	sSlice, err := sq.All(ctx)
	if err != nil {
		panic(err)
	}
	return sSlice
}

// IDs executes the query and returns a list of Specializedappoint ids.
func (sq *SpecializedappointQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := sq.Select(specializedappoint.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sq *SpecializedappointQuery) IDsX(ctx context.Context) []int {
	ids, err := sq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sq *SpecializedappointQuery) Count(ctx context.Context) (int, error) {
	if err := sq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return sq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (sq *SpecializedappointQuery) CountX(ctx context.Context) int {
	count, err := sq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sq *SpecializedappointQuery) Exist(ctx context.Context) (bool, error) {
	if err := sq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return sq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (sq *SpecializedappointQuery) ExistX(ctx context.Context) bool {
	exist, err := sq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sq *SpecializedappointQuery) Clone() *SpecializedappointQuery {
	return &SpecializedappointQuery{
		config:     sq.config,
		limit:      sq.limit,
		offset:     sq.offset,
		order:      append([]OrderFunc{}, sq.order...),
		unique:     append([]string{}, sq.unique...),
		predicates: append([]predicate.Specializedappoint{}, sq.predicates...),
		// clone intermediate query.
		sql:  sq.sql.Clone(),
		path: sq.path,
	}
}

//  WithUser tells the query-builder to eager-loads the nodes that are connected to
// the "user" edge. The optional arguments used to configure the query builder of the edge.
func (sq *SpecializedappointQuery) WithUser(opts ...func(*UserQuery)) *SpecializedappointQuery {
	query := &UserQuery{config: sq.config}
	for _, opt := range opts {
		opt(query)
	}
	sq.withUser = query
	return sq
}

//  WithPatient tells the query-builder to eager-loads the nodes that are connected to
// the "patient" edge. The optional arguments used to configure the query builder of the edge.
func (sq *SpecializedappointQuery) WithPatient(opts ...func(*PatientQuery)) *SpecializedappointQuery {
	query := &PatientQuery{config: sq.config}
	for _, opt := range opts {
		opt(query)
	}
	sq.withPatient = query
	return sq
}

//  WithSpecializeddiag tells the query-builder to eager-loads the nodes that are connected to
// the "specializeddiag" edge. The optional arguments used to configure the query builder of the edge.
func (sq *SpecializedappointQuery) WithSpecializeddiag(opts ...func(*SpecializeddiagQuery)) *SpecializedappointQuery {
	query := &SpecializeddiagQuery{config: sq.config}
	for _, opt := range opts {
		opt(query)
	}
	sq.withSpecializeddiag = query
	return sq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Date time.Time `json:"Date,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Specializedappoint.Query().
//		GroupBy(specializedappoint.FieldDate).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (sq *SpecializedappointQuery) GroupBy(field string, fields ...string) *SpecializedappointGroupBy {
	group := &SpecializedappointGroupBy{config: sq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return sq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Date time.Time `json:"Date,omitempty"`
//	}
//
//	client.Specializedappoint.Query().
//		Select(specializedappoint.FieldDate).
//		Scan(ctx, &v)
//
func (sq *SpecializedappointQuery) Select(field string, fields ...string) *SpecializedappointSelect {
	selector := &SpecializedappointSelect{config: sq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return sq.sqlQuery(), nil
	}
	return selector
}

func (sq *SpecializedappointQuery) prepareQuery(ctx context.Context) error {
	if sq.path != nil {
		prev, err := sq.path(ctx)
		if err != nil {
			return err
		}
		sq.sql = prev
	}
	return nil
}

func (sq *SpecializedappointQuery) sqlAll(ctx context.Context) ([]*Specializedappoint, error) {
	var (
		nodes       = []*Specializedappoint{}
		withFKs     = sq.withFKs
		_spec       = sq.querySpec()
		loadedTypes = [3]bool{
			sq.withUser != nil,
			sq.withPatient != nil,
			sq.withSpecializeddiag != nil,
		}
	)
	if sq.withUser != nil || sq.withPatient != nil || sq.withSpecializeddiag != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, specializedappoint.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &Specializedappoint{config: sq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		if withFKs {
			values = append(values, node.fkValues()...)
		}
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, sq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := sq.withUser; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Specializedappoint)
		for i := range nodes {
			if fk := nodes[i].userid; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(user.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "userid" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.User = n
			}
		}
	}

	if query := sq.withPatient; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Specializedappoint)
		for i := range nodes {
			if fk := nodes[i].patientid; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(patient.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "patientid" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Patient = n
			}
		}
	}

	if query := sq.withSpecializeddiag; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Specializedappoint)
		for i := range nodes {
			if fk := nodes[i].diacnosticID; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(specializeddiag.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "diacnosticID" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Specializeddiag = n
			}
		}
	}

	return nodes, nil
}

func (sq *SpecializedappointQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sq.querySpec()
	return sqlgraph.CountNodes(ctx, sq.driver, _spec)
}

func (sq *SpecializedappointQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := sq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (sq *SpecializedappointQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   specializedappoint.Table,
			Columns: specializedappoint.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: specializedappoint.FieldID,
			},
		},
		From:   sq.sql,
		Unique: true,
	}
	if ps := sq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sq *SpecializedappointQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(sq.driver.Dialect())
	t1 := builder.Table(specializedappoint.Table)
	selector := builder.Select(t1.Columns(specializedappoint.Columns...)...).From(t1)
	if sq.sql != nil {
		selector = sq.sql
		selector.Select(selector.Columns(specializedappoint.Columns...)...)
	}
	for _, p := range sq.predicates {
		p(selector)
	}
	for _, p := range sq.order {
		p(selector)
	}
	if offset := sq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SpecializedappointGroupBy is the builder for group-by Specializedappoint entities.
type SpecializedappointGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sgb *SpecializedappointGroupBy) Aggregate(fns ...AggregateFunc) *SpecializedappointGroupBy {
	sgb.fns = append(sgb.fns, fns...)
	return sgb
}

// Scan applies the group-by query and scan the result into the given value.
func (sgb *SpecializedappointGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := sgb.path(ctx)
	if err != nil {
		return err
	}
	sgb.sql = query
	return sgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (sgb *SpecializedappointGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := sgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (sgb *SpecializedappointGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(sgb.fields) > 1 {
		return nil, errors.New("ent: SpecializedappointGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := sgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (sgb *SpecializedappointGroupBy) StringsX(ctx context.Context) []string {
	v, err := sgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (sgb *SpecializedappointGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = sgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{specializedappoint.Label}
	default:
		err = fmt.Errorf("ent: SpecializedappointGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (sgb *SpecializedappointGroupBy) StringX(ctx context.Context) string {
	v, err := sgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (sgb *SpecializedappointGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(sgb.fields) > 1 {
		return nil, errors.New("ent: SpecializedappointGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := sgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (sgb *SpecializedappointGroupBy) IntsX(ctx context.Context) []int {
	v, err := sgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (sgb *SpecializedappointGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = sgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{specializedappoint.Label}
	default:
		err = fmt.Errorf("ent: SpecializedappointGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (sgb *SpecializedappointGroupBy) IntX(ctx context.Context) int {
	v, err := sgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (sgb *SpecializedappointGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(sgb.fields) > 1 {
		return nil, errors.New("ent: SpecializedappointGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := sgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (sgb *SpecializedappointGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := sgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (sgb *SpecializedappointGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = sgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{specializedappoint.Label}
	default:
		err = fmt.Errorf("ent: SpecializedappointGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (sgb *SpecializedappointGroupBy) Float64X(ctx context.Context) float64 {
	v, err := sgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (sgb *SpecializedappointGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(sgb.fields) > 1 {
		return nil, errors.New("ent: SpecializedappointGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := sgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (sgb *SpecializedappointGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := sgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (sgb *SpecializedappointGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = sgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{specializedappoint.Label}
	default:
		err = fmt.Errorf("ent: SpecializedappointGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (sgb *SpecializedappointGroupBy) BoolX(ctx context.Context) bool {
	v, err := sgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sgb *SpecializedappointGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := sgb.sqlQuery().Query()
	if err := sgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sgb *SpecializedappointGroupBy) sqlQuery() *sql.Selector {
	selector := sgb.sql
	columns := make([]string, 0, len(sgb.fields)+len(sgb.fns))
	columns = append(columns, sgb.fields...)
	for _, fn := range sgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(sgb.fields...)
}

// SpecializedappointSelect is the builder for select fields of Specializedappoint entities.
type SpecializedappointSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (ss *SpecializedappointSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := ss.path(ctx)
	if err != nil {
		return err
	}
	ss.sql = query
	return ss.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ss *SpecializedappointSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ss.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (ss *SpecializedappointSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ss.fields) > 1 {
		return nil, errors.New("ent: SpecializedappointSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ss *SpecializedappointSelect) StringsX(ctx context.Context) []string {
	v, err := ss.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (ss *SpecializedappointSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ss.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{specializedappoint.Label}
	default:
		err = fmt.Errorf("ent: SpecializedappointSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ss *SpecializedappointSelect) StringX(ctx context.Context) string {
	v, err := ss.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (ss *SpecializedappointSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ss.fields) > 1 {
		return nil, errors.New("ent: SpecializedappointSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ss *SpecializedappointSelect) IntsX(ctx context.Context) []int {
	v, err := ss.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (ss *SpecializedappointSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ss.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{specializedappoint.Label}
	default:
		err = fmt.Errorf("ent: SpecializedappointSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ss *SpecializedappointSelect) IntX(ctx context.Context) int {
	v, err := ss.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (ss *SpecializedappointSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ss.fields) > 1 {
		return nil, errors.New("ent: SpecializedappointSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ss *SpecializedappointSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ss.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (ss *SpecializedappointSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ss.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{specializedappoint.Label}
	default:
		err = fmt.Errorf("ent: SpecializedappointSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ss *SpecializedappointSelect) Float64X(ctx context.Context) float64 {
	v, err := ss.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (ss *SpecializedappointSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ss.fields) > 1 {
		return nil, errors.New("ent: SpecializedappointSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ss *SpecializedappointSelect) BoolsX(ctx context.Context) []bool {
	v, err := ss.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (ss *SpecializedappointSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ss.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{specializedappoint.Label}
	default:
		err = fmt.Errorf("ent: SpecializedappointSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ss *SpecializedappointSelect) BoolX(ctx context.Context) bool {
	v, err := ss.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ss *SpecializedappointSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ss.sqlQuery().Query()
	if err := ss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ss *SpecializedappointSelect) sqlQuery() sql.Querier {
	selector := ss.sql
	selector.Select(selector.Columns(ss.fields...)...)
	return selector
}
