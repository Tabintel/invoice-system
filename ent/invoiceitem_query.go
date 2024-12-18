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
	"github.com/Tabintel/invoice-system/ent/invoiceitem"
	"github.com/Tabintel/invoice-system/ent/predicate"
)

// InvoiceItemQuery is the builder for querying InvoiceItem entities.
type InvoiceItemQuery struct {
	config
	ctx        *QueryContext
	order      []invoiceitem.OrderOption
	inters     []Interceptor
	predicates []predicate.InvoiceItem
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the InvoiceItemQuery builder.
func (iiq *InvoiceItemQuery) Where(ps ...predicate.InvoiceItem) *InvoiceItemQuery {
	iiq.predicates = append(iiq.predicates, ps...)
	return iiq
}

// Limit the number of records to be returned by this query.
func (iiq *InvoiceItemQuery) Limit(limit int) *InvoiceItemQuery {
	iiq.ctx.Limit = &limit
	return iiq
}

// Offset to start from.
func (iiq *InvoiceItemQuery) Offset(offset int) *InvoiceItemQuery {
	iiq.ctx.Offset = &offset
	return iiq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (iiq *InvoiceItemQuery) Unique(unique bool) *InvoiceItemQuery {
	iiq.ctx.Unique = &unique
	return iiq
}

// Order specifies how the records should be ordered.
func (iiq *InvoiceItemQuery) Order(o ...invoiceitem.OrderOption) *InvoiceItemQuery {
	iiq.order = append(iiq.order, o...)
	return iiq
}

// First returns the first InvoiceItem entity from the query.
// Returns a *NotFoundError when no InvoiceItem was found.
func (iiq *InvoiceItemQuery) First(ctx context.Context) (*InvoiceItem, error) {
	nodes, err := iiq.Limit(1).All(setContextOp(ctx, iiq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{invoiceitem.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (iiq *InvoiceItemQuery) FirstX(ctx context.Context) *InvoiceItem {
	node, err := iiq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first InvoiceItem ID from the query.
// Returns a *NotFoundError when no InvoiceItem ID was found.
func (iiq *InvoiceItemQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = iiq.Limit(1).IDs(setContextOp(ctx, iiq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{invoiceitem.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (iiq *InvoiceItemQuery) FirstIDX(ctx context.Context) int {
	id, err := iiq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single InvoiceItem entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one InvoiceItem entity is found.
// Returns a *NotFoundError when no InvoiceItem entities are found.
func (iiq *InvoiceItemQuery) Only(ctx context.Context) (*InvoiceItem, error) {
	nodes, err := iiq.Limit(2).All(setContextOp(ctx, iiq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{invoiceitem.Label}
	default:
		return nil, &NotSingularError{invoiceitem.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (iiq *InvoiceItemQuery) OnlyX(ctx context.Context) *InvoiceItem {
	node, err := iiq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only InvoiceItem ID in the query.
// Returns a *NotSingularError when more than one InvoiceItem ID is found.
// Returns a *NotFoundError when no entities are found.
func (iiq *InvoiceItemQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = iiq.Limit(2).IDs(setContextOp(ctx, iiq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{invoiceitem.Label}
	default:
		err = &NotSingularError{invoiceitem.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (iiq *InvoiceItemQuery) OnlyIDX(ctx context.Context) int {
	id, err := iiq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of InvoiceItems.
func (iiq *InvoiceItemQuery) All(ctx context.Context) ([]*InvoiceItem, error) {
	ctx = setContextOp(ctx, iiq.ctx, ent.OpQueryAll)
	if err := iiq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*InvoiceItem, *InvoiceItemQuery]()
	return withInterceptors[[]*InvoiceItem](ctx, iiq, qr, iiq.inters)
}

// AllX is like All, but panics if an error occurs.
func (iiq *InvoiceItemQuery) AllX(ctx context.Context) []*InvoiceItem {
	nodes, err := iiq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of InvoiceItem IDs.
func (iiq *InvoiceItemQuery) IDs(ctx context.Context) (ids []int, err error) {
	if iiq.ctx.Unique == nil && iiq.path != nil {
		iiq.Unique(true)
	}
	ctx = setContextOp(ctx, iiq.ctx, ent.OpQueryIDs)
	if err = iiq.Select(invoiceitem.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (iiq *InvoiceItemQuery) IDsX(ctx context.Context) []int {
	ids, err := iiq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (iiq *InvoiceItemQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, iiq.ctx, ent.OpQueryCount)
	if err := iiq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, iiq, querierCount[*InvoiceItemQuery](), iiq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (iiq *InvoiceItemQuery) CountX(ctx context.Context) int {
	count, err := iiq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (iiq *InvoiceItemQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, iiq.ctx, ent.OpQueryExist)
	switch _, err := iiq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (iiq *InvoiceItemQuery) ExistX(ctx context.Context) bool {
	exist, err := iiq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the InvoiceItemQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (iiq *InvoiceItemQuery) Clone() *InvoiceItemQuery {
	if iiq == nil {
		return nil
	}
	return &InvoiceItemQuery{
		config:     iiq.config,
		ctx:        iiq.ctx.Clone(),
		order:      append([]invoiceitem.OrderOption{}, iiq.order...),
		inters:     append([]Interceptor{}, iiq.inters...),
		predicates: append([]predicate.InvoiceItem{}, iiq.predicates...),
		// clone intermediate query.
		sql:  iiq.sql.Clone(),
		path: iiq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (iiq *InvoiceItemQuery) GroupBy(field string, fields ...string) *InvoiceItemGroupBy {
	iiq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &InvoiceItemGroupBy{build: iiq}
	grbuild.flds = &iiq.ctx.Fields
	grbuild.label = invoiceitem.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (iiq *InvoiceItemQuery) Select(fields ...string) *InvoiceItemSelect {
	iiq.ctx.Fields = append(iiq.ctx.Fields, fields...)
	sbuild := &InvoiceItemSelect{InvoiceItemQuery: iiq}
	sbuild.label = invoiceitem.Label
	sbuild.flds, sbuild.scan = &iiq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a InvoiceItemSelect configured with the given aggregations.
func (iiq *InvoiceItemQuery) Aggregate(fns ...AggregateFunc) *InvoiceItemSelect {
	return iiq.Select().Aggregate(fns...)
}

func (iiq *InvoiceItemQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range iiq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, iiq); err != nil {
				return err
			}
		}
	}
	for _, f := range iiq.ctx.Fields {
		if !invoiceitem.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if iiq.path != nil {
		prev, err := iiq.path(ctx)
		if err != nil {
			return err
		}
		iiq.sql = prev
	}
	return nil
}

func (iiq *InvoiceItemQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*InvoiceItem, error) {
	var (
		nodes   = []*InvoiceItem{}
		withFKs = iiq.withFKs
		_spec   = iiq.querySpec()
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, invoiceitem.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*InvoiceItem).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &InvoiceItem{config: iiq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, iiq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (iiq *InvoiceItemQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := iiq.querySpec()
	_spec.Node.Columns = iiq.ctx.Fields
	if len(iiq.ctx.Fields) > 0 {
		_spec.Unique = iiq.ctx.Unique != nil && *iiq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, iiq.driver, _spec)
}

func (iiq *InvoiceItemQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(invoiceitem.Table, invoiceitem.Columns, sqlgraph.NewFieldSpec(invoiceitem.FieldID, field.TypeInt))
	_spec.From = iiq.sql
	if unique := iiq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if iiq.path != nil {
		_spec.Unique = true
	}
	if fields := iiq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, invoiceitem.FieldID)
		for i := range fields {
			if fields[i] != invoiceitem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := iiq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := iiq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := iiq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := iiq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (iiq *InvoiceItemQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(iiq.driver.Dialect())
	t1 := builder.Table(invoiceitem.Table)
	columns := iiq.ctx.Fields
	if len(columns) == 0 {
		columns = invoiceitem.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if iiq.sql != nil {
		selector = iiq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if iiq.ctx.Unique != nil && *iiq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range iiq.predicates {
		p(selector)
	}
	for _, p := range iiq.order {
		p(selector)
	}
	if offset := iiq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := iiq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// InvoiceItemGroupBy is the group-by builder for InvoiceItem entities.
type InvoiceItemGroupBy struct {
	selector
	build *InvoiceItemQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (iigb *InvoiceItemGroupBy) Aggregate(fns ...AggregateFunc) *InvoiceItemGroupBy {
	iigb.fns = append(iigb.fns, fns...)
	return iigb
}

// Scan applies the selector query and scans the result into the given value.
func (iigb *InvoiceItemGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, iigb.build.ctx, ent.OpQueryGroupBy)
	if err := iigb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*InvoiceItemQuery, *InvoiceItemGroupBy](ctx, iigb.build, iigb, iigb.build.inters, v)
}

func (iigb *InvoiceItemGroupBy) sqlScan(ctx context.Context, root *InvoiceItemQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(iigb.fns))
	for _, fn := range iigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*iigb.flds)+len(iigb.fns))
		for _, f := range *iigb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*iigb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := iigb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// InvoiceItemSelect is the builder for selecting fields of InvoiceItem entities.
type InvoiceItemSelect struct {
	*InvoiceItemQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (iis *InvoiceItemSelect) Aggregate(fns ...AggregateFunc) *InvoiceItemSelect {
	iis.fns = append(iis.fns, fns...)
	return iis
}

// Scan applies the selector query and scans the result into the given value.
func (iis *InvoiceItemSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, iis.ctx, ent.OpQuerySelect)
	if err := iis.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*InvoiceItemQuery, *InvoiceItemSelect](ctx, iis.InvoiceItemQuery, iis, iis.inters, v)
}

func (iis *InvoiceItemSelect) sqlScan(ctx context.Context, root *InvoiceItemQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(iis.fns))
	for _, fn := range iis.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*iis.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := iis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
