// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/generated/entitlementhistory"
	"github.com/datumforge/datum/internal/ent/generated/predicate"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// EntitlementHistoryQuery is the builder for querying EntitlementHistory entities.
type EntitlementHistoryQuery struct {
	config
	ctx        *QueryContext
	order      []entitlementhistory.OrderOption
	inters     []Interceptor
	predicates []predicate.EntitlementHistory
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*EntitlementHistory) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EntitlementHistoryQuery builder.
func (ehq *EntitlementHistoryQuery) Where(ps ...predicate.EntitlementHistory) *EntitlementHistoryQuery {
	ehq.predicates = append(ehq.predicates, ps...)
	return ehq
}

// Limit the number of records to be returned by this query.
func (ehq *EntitlementHistoryQuery) Limit(limit int) *EntitlementHistoryQuery {
	ehq.ctx.Limit = &limit
	return ehq
}

// Offset to start from.
func (ehq *EntitlementHistoryQuery) Offset(offset int) *EntitlementHistoryQuery {
	ehq.ctx.Offset = &offset
	return ehq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ehq *EntitlementHistoryQuery) Unique(unique bool) *EntitlementHistoryQuery {
	ehq.ctx.Unique = &unique
	return ehq
}

// Order specifies how the records should be ordered.
func (ehq *EntitlementHistoryQuery) Order(o ...entitlementhistory.OrderOption) *EntitlementHistoryQuery {
	ehq.order = append(ehq.order, o...)
	return ehq
}

// First returns the first EntitlementHistory entity from the query.
// Returns a *NotFoundError when no EntitlementHistory was found.
func (ehq *EntitlementHistoryQuery) First(ctx context.Context) (*EntitlementHistory, error) {
	nodes, err := ehq.Limit(1).All(setContextOp(ctx, ehq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{entitlementhistory.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ehq *EntitlementHistoryQuery) FirstX(ctx context.Context) *EntitlementHistory {
	node, err := ehq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first EntitlementHistory ID from the query.
// Returns a *NotFoundError when no EntitlementHistory ID was found.
func (ehq *EntitlementHistoryQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = ehq.Limit(1).IDs(setContextOp(ctx, ehq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{entitlementhistory.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ehq *EntitlementHistoryQuery) FirstIDX(ctx context.Context) string {
	id, err := ehq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single EntitlementHistory entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one EntitlementHistory entity is found.
// Returns a *NotFoundError when no EntitlementHistory entities are found.
func (ehq *EntitlementHistoryQuery) Only(ctx context.Context) (*EntitlementHistory, error) {
	nodes, err := ehq.Limit(2).All(setContextOp(ctx, ehq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{entitlementhistory.Label}
	default:
		return nil, &NotSingularError{entitlementhistory.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ehq *EntitlementHistoryQuery) OnlyX(ctx context.Context) *EntitlementHistory {
	node, err := ehq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only EntitlementHistory ID in the query.
// Returns a *NotSingularError when more than one EntitlementHistory ID is found.
// Returns a *NotFoundError when no entities are found.
func (ehq *EntitlementHistoryQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = ehq.Limit(2).IDs(setContextOp(ctx, ehq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{entitlementhistory.Label}
	default:
		err = &NotSingularError{entitlementhistory.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ehq *EntitlementHistoryQuery) OnlyIDX(ctx context.Context) string {
	id, err := ehq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of EntitlementHistories.
func (ehq *EntitlementHistoryQuery) All(ctx context.Context) ([]*EntitlementHistory, error) {
	ctx = setContextOp(ctx, ehq.ctx, "All")
	if err := ehq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*EntitlementHistory, *EntitlementHistoryQuery]()
	return withInterceptors[[]*EntitlementHistory](ctx, ehq, qr, ehq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ehq *EntitlementHistoryQuery) AllX(ctx context.Context) []*EntitlementHistory {
	nodes, err := ehq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of EntitlementHistory IDs.
func (ehq *EntitlementHistoryQuery) IDs(ctx context.Context) (ids []string, err error) {
	if ehq.ctx.Unique == nil && ehq.path != nil {
		ehq.Unique(true)
	}
	ctx = setContextOp(ctx, ehq.ctx, "IDs")
	if err = ehq.Select(entitlementhistory.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ehq *EntitlementHistoryQuery) IDsX(ctx context.Context) []string {
	ids, err := ehq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ehq *EntitlementHistoryQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ehq.ctx, "Count")
	if err := ehq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ehq, querierCount[*EntitlementHistoryQuery](), ehq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ehq *EntitlementHistoryQuery) CountX(ctx context.Context) int {
	count, err := ehq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ehq *EntitlementHistoryQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ehq.ctx, "Exist")
	switch _, err := ehq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ehq *EntitlementHistoryQuery) ExistX(ctx context.Context) bool {
	exist, err := ehq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EntitlementHistoryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ehq *EntitlementHistoryQuery) Clone() *EntitlementHistoryQuery {
	if ehq == nil {
		return nil
	}
	return &EntitlementHistoryQuery{
		config:     ehq.config,
		ctx:        ehq.ctx.Clone(),
		order:      append([]entitlementhistory.OrderOption{}, ehq.order...),
		inters:     append([]Interceptor{}, ehq.inters...),
		predicates: append([]predicate.EntitlementHistory{}, ehq.predicates...),
		// clone intermediate query.
		sql:  ehq.sql.Clone(),
		path: ehq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		HistoryTime time.Time `json:"history_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.EntitlementHistory.Query().
//		GroupBy(entitlementhistory.FieldHistoryTime).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (ehq *EntitlementHistoryQuery) GroupBy(field string, fields ...string) *EntitlementHistoryGroupBy {
	ehq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &EntitlementHistoryGroupBy{build: ehq}
	grbuild.flds = &ehq.ctx.Fields
	grbuild.label = entitlementhistory.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		HistoryTime time.Time `json:"history_time,omitempty"`
//	}
//
//	client.EntitlementHistory.Query().
//		Select(entitlementhistory.FieldHistoryTime).
//		Scan(ctx, &v)
func (ehq *EntitlementHistoryQuery) Select(fields ...string) *EntitlementHistorySelect {
	ehq.ctx.Fields = append(ehq.ctx.Fields, fields...)
	sbuild := &EntitlementHistorySelect{EntitlementHistoryQuery: ehq}
	sbuild.label = entitlementhistory.Label
	sbuild.flds, sbuild.scan = &ehq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a EntitlementHistorySelect configured with the given aggregations.
func (ehq *EntitlementHistoryQuery) Aggregate(fns ...AggregateFunc) *EntitlementHistorySelect {
	return ehq.Select().Aggregate(fns...)
}

func (ehq *EntitlementHistoryQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ehq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ehq); err != nil {
				return err
			}
		}
	}
	for _, f := range ehq.ctx.Fields {
		if !entitlementhistory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if ehq.path != nil {
		prev, err := ehq.path(ctx)
		if err != nil {
			return err
		}
		ehq.sql = prev
	}
	return nil
}

func (ehq *EntitlementHistoryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*EntitlementHistory, error) {
	var (
		nodes = []*EntitlementHistory{}
		_spec = ehq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*EntitlementHistory).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &EntitlementHistory{config: ehq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = ehq.schemaConfig.EntitlementHistory
	ctx = internal.NewSchemaConfigContext(ctx, ehq.schemaConfig)
	if len(ehq.modifiers) > 0 {
		_spec.Modifiers = ehq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ehq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	for i := range ehq.loadTotal {
		if err := ehq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ehq *EntitlementHistoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ehq.querySpec()
	_spec.Node.Schema = ehq.schemaConfig.EntitlementHistory
	ctx = internal.NewSchemaConfigContext(ctx, ehq.schemaConfig)
	if len(ehq.modifiers) > 0 {
		_spec.Modifiers = ehq.modifiers
	}
	_spec.Node.Columns = ehq.ctx.Fields
	if len(ehq.ctx.Fields) > 0 {
		_spec.Unique = ehq.ctx.Unique != nil && *ehq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ehq.driver, _spec)
}

func (ehq *EntitlementHistoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(entitlementhistory.Table, entitlementhistory.Columns, sqlgraph.NewFieldSpec(entitlementhistory.FieldID, field.TypeString))
	_spec.From = ehq.sql
	if unique := ehq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ehq.path != nil {
		_spec.Unique = true
	}
	if fields := ehq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, entitlementhistory.FieldID)
		for i := range fields {
			if fields[i] != entitlementhistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ehq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ehq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ehq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ehq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ehq *EntitlementHistoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ehq.driver.Dialect())
	t1 := builder.Table(entitlementhistory.Table)
	columns := ehq.ctx.Fields
	if len(columns) == 0 {
		columns = entitlementhistory.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ehq.sql != nil {
		selector = ehq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ehq.ctx.Unique != nil && *ehq.ctx.Unique {
		selector.Distinct()
	}
	t1.Schema(ehq.schemaConfig.EntitlementHistory)
	ctx = internal.NewSchemaConfigContext(ctx, ehq.schemaConfig)
	selector.WithContext(ctx)
	for _, p := range ehq.predicates {
		p(selector)
	}
	for _, p := range ehq.order {
		p(selector)
	}
	if offset := ehq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ehq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// EntitlementHistoryGroupBy is the group-by builder for EntitlementHistory entities.
type EntitlementHistoryGroupBy struct {
	selector
	build *EntitlementHistoryQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ehgb *EntitlementHistoryGroupBy) Aggregate(fns ...AggregateFunc) *EntitlementHistoryGroupBy {
	ehgb.fns = append(ehgb.fns, fns...)
	return ehgb
}

// Scan applies the selector query and scans the result into the given value.
func (ehgb *EntitlementHistoryGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ehgb.build.ctx, "GroupBy")
	if err := ehgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EntitlementHistoryQuery, *EntitlementHistoryGroupBy](ctx, ehgb.build, ehgb, ehgb.build.inters, v)
}

func (ehgb *EntitlementHistoryGroupBy) sqlScan(ctx context.Context, root *EntitlementHistoryQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ehgb.fns))
	for _, fn := range ehgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ehgb.flds)+len(ehgb.fns))
		for _, f := range *ehgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ehgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ehgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// EntitlementHistorySelect is the builder for selecting fields of EntitlementHistory entities.
type EntitlementHistorySelect struct {
	*EntitlementHistoryQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ehs *EntitlementHistorySelect) Aggregate(fns ...AggregateFunc) *EntitlementHistorySelect {
	ehs.fns = append(ehs.fns, fns...)
	return ehs
}

// Scan applies the selector query and scans the result into the given value.
func (ehs *EntitlementHistorySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ehs.ctx, "Select")
	if err := ehs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EntitlementHistoryQuery, *EntitlementHistorySelect](ctx, ehs.EntitlementHistoryQuery, ehs, ehs.inters, v)
}

func (ehs *EntitlementHistorySelect) sqlScan(ctx context.Context, root *EntitlementHistoryQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ehs.fns))
	for _, fn := range ehs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ehs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ehs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
