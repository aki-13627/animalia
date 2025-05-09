// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/aki-13627/animalia/backend-go/ent/dailytask"
	"github.com/aki-13627/animalia/backend-go/ent/post"
	"github.com/aki-13627/animalia/backend-go/ent/predicate"
	"github.com/aki-13627/animalia/backend-go/ent/user"
	"github.com/google/uuid"
)

// DailyTaskQuery is the builder for querying DailyTask entities.
type DailyTaskQuery struct {
	config
	ctx        *QueryContext
	order      []dailytask.OrderOption
	inters     []Interceptor
	predicates []predicate.DailyTask
	withUser   *UserQuery
	withPost   *PostQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DailyTaskQuery builder.
func (dtq *DailyTaskQuery) Where(ps ...predicate.DailyTask) *DailyTaskQuery {
	dtq.predicates = append(dtq.predicates, ps...)
	return dtq
}

// Limit the number of records to be returned by this query.
func (dtq *DailyTaskQuery) Limit(limit int) *DailyTaskQuery {
	dtq.ctx.Limit = &limit
	return dtq
}

// Offset to start from.
func (dtq *DailyTaskQuery) Offset(offset int) *DailyTaskQuery {
	dtq.ctx.Offset = &offset
	return dtq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dtq *DailyTaskQuery) Unique(unique bool) *DailyTaskQuery {
	dtq.ctx.Unique = &unique
	return dtq
}

// Order specifies how the records should be ordered.
func (dtq *DailyTaskQuery) Order(o ...dailytask.OrderOption) *DailyTaskQuery {
	dtq.order = append(dtq.order, o...)
	return dtq
}

// QueryUser chains the current query on the "user" edge.
func (dtq *DailyTaskQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: dtq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dtq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dtq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(dailytask.Table, dailytask.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, dailytask.UserTable, dailytask.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(dtq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPost chains the current query on the "post" edge.
func (dtq *DailyTaskQuery) QueryPost() *PostQuery {
	query := (&PostClient{config: dtq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dtq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dtq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(dailytask.Table, dailytask.FieldID, selector),
			sqlgraph.To(post.Table, post.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, dailytask.PostTable, dailytask.PostColumn),
		)
		fromU = sqlgraph.SetNeighbors(dtq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first DailyTask entity from the query.
// Returns a *NotFoundError when no DailyTask was found.
func (dtq *DailyTaskQuery) First(ctx context.Context) (*DailyTask, error) {
	nodes, err := dtq.Limit(1).All(setContextOp(ctx, dtq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{dailytask.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dtq *DailyTaskQuery) FirstX(ctx context.Context) *DailyTask {
	node, err := dtq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DailyTask ID from the query.
// Returns a *NotFoundError when no DailyTask ID was found.
func (dtq *DailyTaskQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = dtq.Limit(1).IDs(setContextOp(ctx, dtq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{dailytask.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dtq *DailyTaskQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := dtq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DailyTask entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one DailyTask entity is found.
// Returns a *NotFoundError when no DailyTask entities are found.
func (dtq *DailyTaskQuery) Only(ctx context.Context) (*DailyTask, error) {
	nodes, err := dtq.Limit(2).All(setContextOp(ctx, dtq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{dailytask.Label}
	default:
		return nil, &NotSingularError{dailytask.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dtq *DailyTaskQuery) OnlyX(ctx context.Context) *DailyTask {
	node, err := dtq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DailyTask ID in the query.
// Returns a *NotSingularError when more than one DailyTask ID is found.
// Returns a *NotFoundError when no entities are found.
func (dtq *DailyTaskQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = dtq.Limit(2).IDs(setContextOp(ctx, dtq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{dailytask.Label}
	default:
		err = &NotSingularError{dailytask.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dtq *DailyTaskQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := dtq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DailyTasks.
func (dtq *DailyTaskQuery) All(ctx context.Context) ([]*DailyTask, error) {
	ctx = setContextOp(ctx, dtq.ctx, ent.OpQueryAll)
	if err := dtq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*DailyTask, *DailyTaskQuery]()
	return withInterceptors[[]*DailyTask](ctx, dtq, qr, dtq.inters)
}

// AllX is like All, but panics if an error occurs.
func (dtq *DailyTaskQuery) AllX(ctx context.Context) []*DailyTask {
	nodes, err := dtq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DailyTask IDs.
func (dtq *DailyTaskQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if dtq.ctx.Unique == nil && dtq.path != nil {
		dtq.Unique(true)
	}
	ctx = setContextOp(ctx, dtq.ctx, ent.OpQueryIDs)
	if err = dtq.Select(dailytask.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dtq *DailyTaskQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := dtq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dtq *DailyTaskQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, dtq.ctx, ent.OpQueryCount)
	if err := dtq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, dtq, querierCount[*DailyTaskQuery](), dtq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (dtq *DailyTaskQuery) CountX(ctx context.Context) int {
	count, err := dtq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dtq *DailyTaskQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, dtq.ctx, ent.OpQueryExist)
	switch _, err := dtq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (dtq *DailyTaskQuery) ExistX(ctx context.Context) bool {
	exist, err := dtq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DailyTaskQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dtq *DailyTaskQuery) Clone() *DailyTaskQuery {
	if dtq == nil {
		return nil
	}
	return &DailyTaskQuery{
		config:     dtq.config,
		ctx:        dtq.ctx.Clone(),
		order:      append([]dailytask.OrderOption{}, dtq.order...),
		inters:     append([]Interceptor{}, dtq.inters...),
		predicates: append([]predicate.DailyTask{}, dtq.predicates...),
		withUser:   dtq.withUser.Clone(),
		withPost:   dtq.withPost.Clone(),
		// clone intermediate query.
		sql:  dtq.sql.Clone(),
		path: dtq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (dtq *DailyTaskQuery) WithUser(opts ...func(*UserQuery)) *DailyTaskQuery {
	query := (&UserClient{config: dtq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dtq.withUser = query
	return dtq
}

// WithPost tells the query-builder to eager-load the nodes that are connected to
// the "post" edge. The optional arguments are used to configure the query builder of the edge.
func (dtq *DailyTaskQuery) WithPost(opts ...func(*PostQuery)) *DailyTaskQuery {
	query := (&PostClient{config: dtq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dtq.withPost = query
	return dtq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.DailyTask.Query().
//		GroupBy(dailytask.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (dtq *DailyTaskQuery) GroupBy(field string, fields ...string) *DailyTaskGroupBy {
	dtq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &DailyTaskGroupBy{build: dtq}
	grbuild.flds = &dtq.ctx.Fields
	grbuild.label = dailytask.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.DailyTask.Query().
//		Select(dailytask.FieldCreatedAt).
//		Scan(ctx, &v)
func (dtq *DailyTaskQuery) Select(fields ...string) *DailyTaskSelect {
	dtq.ctx.Fields = append(dtq.ctx.Fields, fields...)
	sbuild := &DailyTaskSelect{DailyTaskQuery: dtq}
	sbuild.label = dailytask.Label
	sbuild.flds, sbuild.scan = &dtq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a DailyTaskSelect configured with the given aggregations.
func (dtq *DailyTaskQuery) Aggregate(fns ...AggregateFunc) *DailyTaskSelect {
	return dtq.Select().Aggregate(fns...)
}

func (dtq *DailyTaskQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range dtq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, dtq); err != nil {
				return err
			}
		}
	}
	for _, f := range dtq.ctx.Fields {
		if !dailytask.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dtq.path != nil {
		prev, err := dtq.path(ctx)
		if err != nil {
			return err
		}
		dtq.sql = prev
	}
	return nil
}

func (dtq *DailyTaskQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*DailyTask, error) {
	var (
		nodes       = []*DailyTask{}
		withFKs     = dtq.withFKs
		_spec       = dtq.querySpec()
		loadedTypes = [2]bool{
			dtq.withUser != nil,
			dtq.withPost != nil,
		}
	)
	if dtq.withUser != nil || dtq.withPost != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, dailytask.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*DailyTask).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &DailyTask{config: dtq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dtq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := dtq.withUser; query != nil {
		if err := dtq.loadUser(ctx, query, nodes, nil,
			func(n *DailyTask, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := dtq.withPost; query != nil {
		if err := dtq.loadPost(ctx, query, nodes, nil,
			func(n *DailyTask, e *Post) { n.Edges.Post = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (dtq *DailyTaskQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*DailyTask, init func(*DailyTask), assign func(*DailyTask, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*DailyTask)
	for i := range nodes {
		if nodes[i].user_daily_tasks == nil {
			continue
		}
		fk := *nodes[i].user_daily_tasks
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_daily_tasks" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (dtq *DailyTaskQuery) loadPost(ctx context.Context, query *PostQuery, nodes []*DailyTask, init func(*DailyTask), assign func(*DailyTask, *Post)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*DailyTask)
	for i := range nodes {
		if nodes[i].post_daily_task == nil {
			continue
		}
		fk := *nodes[i].post_daily_task
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(post.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "post_daily_task" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (dtq *DailyTaskQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dtq.querySpec()
	_spec.Node.Columns = dtq.ctx.Fields
	if len(dtq.ctx.Fields) > 0 {
		_spec.Unique = dtq.ctx.Unique != nil && *dtq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, dtq.driver, _spec)
}

func (dtq *DailyTaskQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(dailytask.Table, dailytask.Columns, sqlgraph.NewFieldSpec(dailytask.FieldID, field.TypeUUID))
	_spec.From = dtq.sql
	if unique := dtq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if dtq.path != nil {
		_spec.Unique = true
	}
	if fields := dtq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, dailytask.FieldID)
		for i := range fields {
			if fields[i] != dailytask.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dtq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dtq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dtq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dtq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dtq *DailyTaskQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dtq.driver.Dialect())
	t1 := builder.Table(dailytask.Table)
	columns := dtq.ctx.Fields
	if len(columns) == 0 {
		columns = dailytask.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dtq.sql != nil {
		selector = dtq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dtq.ctx.Unique != nil && *dtq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range dtq.predicates {
		p(selector)
	}
	for _, p := range dtq.order {
		p(selector)
	}
	if offset := dtq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dtq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DailyTaskGroupBy is the group-by builder for DailyTask entities.
type DailyTaskGroupBy struct {
	selector
	build *DailyTaskQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dtgb *DailyTaskGroupBy) Aggregate(fns ...AggregateFunc) *DailyTaskGroupBy {
	dtgb.fns = append(dtgb.fns, fns...)
	return dtgb
}

// Scan applies the selector query and scans the result into the given value.
func (dtgb *DailyTaskGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dtgb.build.ctx, ent.OpQueryGroupBy)
	if err := dtgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DailyTaskQuery, *DailyTaskGroupBy](ctx, dtgb.build, dtgb, dtgb.build.inters, v)
}

func (dtgb *DailyTaskGroupBy) sqlScan(ctx context.Context, root *DailyTaskQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(dtgb.fns))
	for _, fn := range dtgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*dtgb.flds)+len(dtgb.fns))
		for _, f := range *dtgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*dtgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dtgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// DailyTaskSelect is the builder for selecting fields of DailyTask entities.
type DailyTaskSelect struct {
	*DailyTaskQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (dts *DailyTaskSelect) Aggregate(fns ...AggregateFunc) *DailyTaskSelect {
	dts.fns = append(dts.fns, fns...)
	return dts
}

// Scan applies the selector query and scans the result into the given value.
func (dts *DailyTaskSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dts.ctx, ent.OpQuerySelect)
	if err := dts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DailyTaskQuery, *DailyTaskSelect](ctx, dts.DailyTaskQuery, dts, dts.inters, v)
}

func (dts *DailyTaskSelect) sqlScan(ctx context.Context, root *DailyTaskQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(dts.fns))
	for _, fn := range dts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*dts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
