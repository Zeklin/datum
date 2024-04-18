// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/generated/grouphistory"
	"github.com/datumforge/datum/internal/ent/generated/predicate"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// GroupHistoryQuery is the builder for querying GroupHistory entities.
type GroupHistoryQuery struct {
	config
	ctx        *QueryContext
	order      []grouphistory.OrderOption
	inters     []Interceptor
	predicates []predicate.GroupHistory
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*GroupHistory) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GroupHistoryQuery builder.
func (ghq *GroupHistoryQuery) Where(ps ...predicate.GroupHistory) *GroupHistoryQuery {
	ghq.predicates = append(ghq.predicates, ps...)
	return ghq
}

// Limit the number of records to be returned by this query.
func (ghq *GroupHistoryQuery) Limit(limit int) *GroupHistoryQuery {
	ghq.ctx.Limit = &limit
	return ghq
}

// Offset to start from.
func (ghq *GroupHistoryQuery) Offset(offset int) *GroupHistoryQuery {
	ghq.ctx.Offset = &offset
	return ghq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ghq *GroupHistoryQuery) Unique(unique bool) *GroupHistoryQuery {
	ghq.ctx.Unique = &unique
	return ghq
}

// Order specifies how the records should be ordered.
func (ghq *GroupHistoryQuery) Order(o ...grouphistory.OrderOption) *GroupHistoryQuery {
	ghq.order = append(ghq.order, o...)
	return ghq
}

// First returns the first GroupHistory entity from the query.
// Returns a *NotFoundError when no GroupHistory was found.
func (ghq *GroupHistoryQuery) First(ctx context.Context) (*GroupHistory, error) {
	nodes, err := ghq.Limit(1).All(setContextOp(ctx, ghq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{grouphistory.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ghq *GroupHistoryQuery) FirstX(ctx context.Context) *GroupHistory {
	node, err := ghq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GroupHistory ID from the query.
// Returns a *NotFoundError when no GroupHistory ID was found.
func (ghq *GroupHistoryQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = ghq.Limit(1).IDs(setContextOp(ctx, ghq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{grouphistory.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ghq *GroupHistoryQuery) FirstIDX(ctx context.Context) string {
	id, err := ghq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GroupHistory entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one GroupHistory entity is found.
// Returns a *NotFoundError when no GroupHistory entities are found.
func (ghq *GroupHistoryQuery) Only(ctx context.Context) (*GroupHistory, error) {
	nodes, err := ghq.Limit(2).All(setContextOp(ctx, ghq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{grouphistory.Label}
	default:
		return nil, &NotSingularError{grouphistory.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ghq *GroupHistoryQuery) OnlyX(ctx context.Context) *GroupHistory {
	node, err := ghq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GroupHistory ID in the query.
// Returns a *NotSingularError when more than one GroupHistory ID is found.
// Returns a *NotFoundError when no entities are found.
func (ghq *GroupHistoryQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = ghq.Limit(2).IDs(setContextOp(ctx, ghq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{grouphistory.Label}
	default:
		err = &NotSingularError{grouphistory.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ghq *GroupHistoryQuery) OnlyIDX(ctx context.Context) string {
	id, err := ghq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GroupHistories.
func (ghq *GroupHistoryQuery) All(ctx context.Context) ([]*GroupHistory, error) {
	ctx = setContextOp(ctx, ghq.ctx, "All")
	if err := ghq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*GroupHistory, *GroupHistoryQuery]()
	return withInterceptors[[]*GroupHistory](ctx, ghq, qr, ghq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ghq *GroupHistoryQuery) AllX(ctx context.Context) []*GroupHistory {
	nodes, err := ghq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GroupHistory IDs.
func (ghq *GroupHistoryQuery) IDs(ctx context.Context) (ids []string, err error) {
	if ghq.ctx.Unique == nil && ghq.path != nil {
		ghq.Unique(true)
	}
	ctx = setContextOp(ctx, ghq.ctx, "IDs")
	if err = ghq.Select(grouphistory.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ghq *GroupHistoryQuery) IDsX(ctx context.Context) []string {
	ids, err := ghq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ghq *GroupHistoryQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ghq.ctx, "Count")
	if err := ghq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ghq, querierCount[*GroupHistoryQuery](), ghq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ghq *GroupHistoryQuery) CountX(ctx context.Context) int {
	count, err := ghq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ghq *GroupHistoryQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ghq.ctx, "Exist")
	switch _, err := ghq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ghq *GroupHistoryQuery) ExistX(ctx context.Context) bool {
	exist, err := ghq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GroupHistoryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ghq *GroupHistoryQuery) Clone() *GroupHistoryQuery {
	if ghq == nil {
		return nil
	}
	return &GroupHistoryQuery{
		config:     ghq.config,
		ctx:        ghq.ctx.Clone(),
		order:      append([]grouphistory.OrderOption{}, ghq.order...),
		inters:     append([]Interceptor{}, ghq.inters...),
		predicates: append([]predicate.GroupHistory{}, ghq.predicates...),
		// clone intermediate query.
		sql:  ghq.sql.Clone(),
		path: ghq.path,
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
//	client.GroupHistory.Query().
//		GroupBy(grouphistory.FieldHistoryTime).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (ghq *GroupHistoryQuery) GroupBy(field string, fields ...string) *GroupHistoryGroupBy {
	ghq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &GroupHistoryGroupBy{build: ghq}
	grbuild.flds = &ghq.ctx.Fields
	grbuild.label = grouphistory.Label
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
//	client.GroupHistory.Query().
//		Select(grouphistory.FieldHistoryTime).
//		Scan(ctx, &v)
func (ghq *GroupHistoryQuery) Select(fields ...string) *GroupHistorySelect {
	ghq.ctx.Fields = append(ghq.ctx.Fields, fields...)
	sbuild := &GroupHistorySelect{GroupHistoryQuery: ghq}
	sbuild.label = grouphistory.Label
	sbuild.flds, sbuild.scan = &ghq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a GroupHistorySelect configured with the given aggregations.
func (ghq *GroupHistoryQuery) Aggregate(fns ...AggregateFunc) *GroupHistorySelect {
	return ghq.Select().Aggregate(fns...)
}

func (ghq *GroupHistoryQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ghq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ghq); err != nil {
				return err
			}
		}
	}
	for _, f := range ghq.ctx.Fields {
		if !grouphistory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if ghq.path != nil {
		prev, err := ghq.path(ctx)
		if err != nil {
			return err
		}
		ghq.sql = prev
	}
	return nil
}

func (ghq *GroupHistoryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*GroupHistory, error) {
	var (
		nodes = []*GroupHistory{}
		_spec = ghq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*GroupHistory).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &GroupHistory{config: ghq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = ghq.schemaConfig.GroupHistory
	ctx = internal.NewSchemaConfigContext(ctx, ghq.schemaConfig)
	if len(ghq.modifiers) > 0 {
		_spec.Modifiers = ghq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ghq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	for i := range ghq.loadTotal {
		if err := ghq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ghq *GroupHistoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ghq.querySpec()
	_spec.Node.Schema = ghq.schemaConfig.GroupHistory
	ctx = internal.NewSchemaConfigContext(ctx, ghq.schemaConfig)
	if len(ghq.modifiers) > 0 {
		_spec.Modifiers = ghq.modifiers
	}
	_spec.Node.Columns = ghq.ctx.Fields
	if len(ghq.ctx.Fields) > 0 {
		_spec.Unique = ghq.ctx.Unique != nil && *ghq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ghq.driver, _spec)
}

func (ghq *GroupHistoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(grouphistory.Table, grouphistory.Columns, sqlgraph.NewFieldSpec(grouphistory.FieldID, field.TypeString))
	_spec.From = ghq.sql
	if unique := ghq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ghq.path != nil {
		_spec.Unique = true
	}
	if fields := ghq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, grouphistory.FieldID)
		for i := range fields {
			if fields[i] != grouphistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ghq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ghq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ghq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ghq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ghq *GroupHistoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ghq.driver.Dialect())
	t1 := builder.Table(grouphistory.Table)
	columns := ghq.ctx.Fields
	if len(columns) == 0 {
		columns = grouphistory.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ghq.sql != nil {
		selector = ghq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ghq.ctx.Unique != nil && *ghq.ctx.Unique {
		selector.Distinct()
	}
	t1.Schema(ghq.schemaConfig.GroupHistory)
	ctx = internal.NewSchemaConfigContext(ctx, ghq.schemaConfig)
	selector.WithContext(ctx)
	for _, p := range ghq.predicates {
		p(selector)
	}
	for _, p := range ghq.order {
		p(selector)
	}
	if offset := ghq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ghq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GroupHistoryGroupBy is the group-by builder for GroupHistory entities.
type GroupHistoryGroupBy struct {
	selector
	build *GroupHistoryQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ghgb *GroupHistoryGroupBy) Aggregate(fns ...AggregateFunc) *GroupHistoryGroupBy {
	ghgb.fns = append(ghgb.fns, fns...)
	return ghgb
}

// Scan applies the selector query and scans the result into the given value.
func (ghgb *GroupHistoryGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ghgb.build.ctx, "GroupBy")
	if err := ghgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GroupHistoryQuery, *GroupHistoryGroupBy](ctx, ghgb.build, ghgb, ghgb.build.inters, v)
}

func (ghgb *GroupHistoryGroupBy) sqlScan(ctx context.Context, root *GroupHistoryQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ghgb.fns))
	for _, fn := range ghgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ghgb.flds)+len(ghgb.fns))
		for _, f := range *ghgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ghgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ghgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// GroupHistorySelect is the builder for selecting fields of GroupHistory entities.
type GroupHistorySelect struct {
	*GroupHistoryQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ghs *GroupHistorySelect) Aggregate(fns ...AggregateFunc) *GroupHistorySelect {
	ghs.fns = append(ghs.fns, fns...)
	return ghs
}

// Scan applies the selector query and scans the result into the given value.
func (ghs *GroupHistorySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ghs.ctx, "Select")
	if err := ghs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GroupHistoryQuery, *GroupHistorySelect](ctx, ghs.GroupHistoryQuery, ghs, ghs.inters, v)
}

func (ghs *GroupHistorySelect) sqlScan(ctx context.Context, root *GroupHistoryQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ghs.fns))
	for _, fn := range ghs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ghs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ghs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
