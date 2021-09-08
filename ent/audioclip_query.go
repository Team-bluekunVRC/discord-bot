// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Team-bluekunVRC/discord-bot/ent/audioclip"
	"github.com/Team-bluekunVRC/discord-bot/ent/predicate"
	"github.com/google/uuid"
)

// AudioClipQuery is the builder for querying AudioClip entities.
type AudioClipQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.AudioClip
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AudioClipQuery builder.
func (acq *AudioClipQuery) Where(ps ...predicate.AudioClip) *AudioClipQuery {
	acq.predicates = append(acq.predicates, ps...)
	return acq
}

// Limit adds a limit step to the query.
func (acq *AudioClipQuery) Limit(limit int) *AudioClipQuery {
	acq.limit = &limit
	return acq
}

// Offset adds an offset step to the query.
func (acq *AudioClipQuery) Offset(offset int) *AudioClipQuery {
	acq.offset = &offset
	return acq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (acq *AudioClipQuery) Unique(unique bool) *AudioClipQuery {
	acq.unique = &unique
	return acq
}

// Order adds an order step to the query.
func (acq *AudioClipQuery) Order(o ...OrderFunc) *AudioClipQuery {
	acq.order = append(acq.order, o...)
	return acq
}

// First returns the first AudioClip entity from the query.
// Returns a *NotFoundError when no AudioClip was found.
func (acq *AudioClipQuery) First(ctx context.Context) (*AudioClip, error) {
	nodes, err := acq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{audioclip.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (acq *AudioClipQuery) FirstX(ctx context.Context) *AudioClip {
	node, err := acq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AudioClip ID from the query.
// Returns a *NotFoundError when no AudioClip ID was found.
func (acq *AudioClipQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = acq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{audioclip.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (acq *AudioClipQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := acq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AudioClip entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one AudioClip entity is not found.
// Returns a *NotFoundError when no AudioClip entities are found.
func (acq *AudioClipQuery) Only(ctx context.Context) (*AudioClip, error) {
	nodes, err := acq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{audioclip.Label}
	default:
		return nil, &NotSingularError{audioclip.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (acq *AudioClipQuery) OnlyX(ctx context.Context) *AudioClip {
	node, err := acq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AudioClip ID in the query.
// Returns a *NotSingularError when exactly one AudioClip ID is not found.
// Returns a *NotFoundError when no entities are found.
func (acq *AudioClipQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = acq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{audioclip.Label}
	default:
		err = &NotSingularError{audioclip.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (acq *AudioClipQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := acq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AudioClips.
func (acq *AudioClipQuery) All(ctx context.Context) ([]*AudioClip, error) {
	if err := acq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return acq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (acq *AudioClipQuery) AllX(ctx context.Context) []*AudioClip {
	nodes, err := acq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AudioClip IDs.
func (acq *AudioClipQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := acq.Select(audioclip.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (acq *AudioClipQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := acq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (acq *AudioClipQuery) Count(ctx context.Context) (int, error) {
	if err := acq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return acq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (acq *AudioClipQuery) CountX(ctx context.Context) int {
	count, err := acq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (acq *AudioClipQuery) Exist(ctx context.Context) (bool, error) {
	if err := acq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return acq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (acq *AudioClipQuery) ExistX(ctx context.Context) bool {
	exist, err := acq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AudioClipQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (acq *AudioClipQuery) Clone() *AudioClipQuery {
	if acq == nil {
		return nil
	}
	return &AudioClipQuery{
		config:     acq.config,
		limit:      acq.limit,
		offset:     acq.offset,
		order:      append([]OrderFunc{}, acq.order...),
		predicates: append([]predicate.AudioClip{}, acq.predicates...),
		// clone intermediate query.
		sql:  acq.sql.Clone(),
		path: acq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		StoragePath string `json:"storage_path,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AudioClip.Query().
//		GroupBy(audioclip.FieldStoragePath).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (acq *AudioClipQuery) GroupBy(field string, fields ...string) *AudioClipGroupBy {
	group := &AudioClipGroupBy{config: acq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := acq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return acq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		StoragePath string `json:"storage_path,omitempty"`
//	}
//
//	client.AudioClip.Query().
//		Select(audioclip.FieldStoragePath).
//		Scan(ctx, &v)
//
func (acq *AudioClipQuery) Select(fields ...string) *AudioClipSelect {
	acq.fields = append(acq.fields, fields...)
	return &AudioClipSelect{AudioClipQuery: acq}
}

func (acq *AudioClipQuery) prepareQuery(ctx context.Context) error {
	for _, f := range acq.fields {
		if !audioclip.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if acq.path != nil {
		prev, err := acq.path(ctx)
		if err != nil {
			return err
		}
		acq.sql = prev
	}
	return nil
}

func (acq *AudioClipQuery) sqlAll(ctx context.Context) ([]*AudioClip, error) {
	var (
		nodes   = []*AudioClip{}
		withFKs = acq.withFKs
		_spec   = acq.querySpec()
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, audioclip.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &AudioClip{config: acq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, acq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (acq *AudioClipQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := acq.querySpec()
	return sqlgraph.CountNodes(ctx, acq.driver, _spec)
}

func (acq *AudioClipQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := acq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (acq *AudioClipQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   audioclip.Table,
			Columns: audioclip.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: audioclip.FieldID,
			},
		},
		From:   acq.sql,
		Unique: true,
	}
	if unique := acq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := acq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, audioclip.FieldID)
		for i := range fields {
			if fields[i] != audioclip.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := acq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := acq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := acq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := acq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (acq *AudioClipQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(acq.driver.Dialect())
	t1 := builder.Table(audioclip.Table)
	columns := acq.fields
	if len(columns) == 0 {
		columns = audioclip.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if acq.sql != nil {
		selector = acq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range acq.predicates {
		p(selector)
	}
	for _, p := range acq.order {
		p(selector)
	}
	if offset := acq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := acq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AudioClipGroupBy is the group-by builder for AudioClip entities.
type AudioClipGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (acgb *AudioClipGroupBy) Aggregate(fns ...AggregateFunc) *AudioClipGroupBy {
	acgb.fns = append(acgb.fns, fns...)
	return acgb
}

// Scan applies the group-by query and scans the result into the given value.
func (acgb *AudioClipGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := acgb.path(ctx)
	if err != nil {
		return err
	}
	acgb.sql = query
	return acgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (acgb *AudioClipGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := acgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (acgb *AudioClipGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(acgb.fields) > 1 {
		return nil, errors.New("ent: AudioClipGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := acgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (acgb *AudioClipGroupBy) StringsX(ctx context.Context) []string {
	v, err := acgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (acgb *AudioClipGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = acgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{audioclip.Label}
	default:
		err = fmt.Errorf("ent: AudioClipGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (acgb *AudioClipGroupBy) StringX(ctx context.Context) string {
	v, err := acgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (acgb *AudioClipGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(acgb.fields) > 1 {
		return nil, errors.New("ent: AudioClipGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := acgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (acgb *AudioClipGroupBy) IntsX(ctx context.Context) []int {
	v, err := acgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (acgb *AudioClipGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = acgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{audioclip.Label}
	default:
		err = fmt.Errorf("ent: AudioClipGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (acgb *AudioClipGroupBy) IntX(ctx context.Context) int {
	v, err := acgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (acgb *AudioClipGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(acgb.fields) > 1 {
		return nil, errors.New("ent: AudioClipGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := acgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (acgb *AudioClipGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := acgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (acgb *AudioClipGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = acgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{audioclip.Label}
	default:
		err = fmt.Errorf("ent: AudioClipGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (acgb *AudioClipGroupBy) Float64X(ctx context.Context) float64 {
	v, err := acgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (acgb *AudioClipGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(acgb.fields) > 1 {
		return nil, errors.New("ent: AudioClipGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := acgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (acgb *AudioClipGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := acgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (acgb *AudioClipGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = acgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{audioclip.Label}
	default:
		err = fmt.Errorf("ent: AudioClipGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (acgb *AudioClipGroupBy) BoolX(ctx context.Context) bool {
	v, err := acgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (acgb *AudioClipGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range acgb.fields {
		if !audioclip.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := acgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := acgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (acgb *AudioClipGroupBy) sqlQuery() *sql.Selector {
	selector := acgb.sql.Select()
	aggregation := make([]string, 0, len(acgb.fns))
	for _, fn := range acgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(acgb.fields)+len(acgb.fns))
		for _, f := range acgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(acgb.fields...)...)
}

// AudioClipSelect is the builder for selecting fields of AudioClip entities.
type AudioClipSelect struct {
	*AudioClipQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (acs *AudioClipSelect) Scan(ctx context.Context, v interface{}) error {
	if err := acs.prepareQuery(ctx); err != nil {
		return err
	}
	acs.sql = acs.AudioClipQuery.sqlQuery(ctx)
	return acs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (acs *AudioClipSelect) ScanX(ctx context.Context, v interface{}) {
	if err := acs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (acs *AudioClipSelect) Strings(ctx context.Context) ([]string, error) {
	if len(acs.fields) > 1 {
		return nil, errors.New("ent: AudioClipSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := acs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (acs *AudioClipSelect) StringsX(ctx context.Context) []string {
	v, err := acs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (acs *AudioClipSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = acs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{audioclip.Label}
	default:
		err = fmt.Errorf("ent: AudioClipSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (acs *AudioClipSelect) StringX(ctx context.Context) string {
	v, err := acs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (acs *AudioClipSelect) Ints(ctx context.Context) ([]int, error) {
	if len(acs.fields) > 1 {
		return nil, errors.New("ent: AudioClipSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := acs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (acs *AudioClipSelect) IntsX(ctx context.Context) []int {
	v, err := acs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (acs *AudioClipSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = acs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{audioclip.Label}
	default:
		err = fmt.Errorf("ent: AudioClipSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (acs *AudioClipSelect) IntX(ctx context.Context) int {
	v, err := acs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (acs *AudioClipSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(acs.fields) > 1 {
		return nil, errors.New("ent: AudioClipSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := acs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (acs *AudioClipSelect) Float64sX(ctx context.Context) []float64 {
	v, err := acs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (acs *AudioClipSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = acs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{audioclip.Label}
	default:
		err = fmt.Errorf("ent: AudioClipSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (acs *AudioClipSelect) Float64X(ctx context.Context) float64 {
	v, err := acs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (acs *AudioClipSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(acs.fields) > 1 {
		return nil, errors.New("ent: AudioClipSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := acs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (acs *AudioClipSelect) BoolsX(ctx context.Context) []bool {
	v, err := acs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (acs *AudioClipSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = acs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{audioclip.Label}
	default:
		err = fmt.Errorf("ent: AudioClipSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (acs *AudioClipSelect) BoolX(ctx context.Context) bool {
	v, err := acs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (acs *AudioClipSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := acs.sql.Query()
	if err := acs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}