// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Tabintel/invoice-system/ent/customer"
	"github.com/Tabintel/invoice-system/ent/invoice"
	"github.com/Tabintel/invoice-system/ent/invoiceitem"
	"github.com/Tabintel/invoice-system/ent/payment"
	"github.com/Tabintel/invoice-system/ent/predicate"
	"github.com/Tabintel/invoice-system/ent/user"
)

// InvoiceQuery is the builder for querying Invoice entities.
type InvoiceQuery struct {
	config
	ctx          *QueryContext
	order        []invoice.OrderOption
	inters       []Interceptor
	predicates   []predicate.Invoice
	withCreator  *UserQuery
	withItems    *InvoiceItemQuery
	withPayments *PaymentQuery
	withCustomer *CustomerQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the InvoiceQuery builder.
func (iq *InvoiceQuery) Where(ps ...predicate.Invoice) *InvoiceQuery {
	iq.predicates = append(iq.predicates, ps...)
	return iq
}

// Limit the number of records to be returned by this query.
func (iq *InvoiceQuery) Limit(limit int) *InvoiceQuery {
	iq.ctx.Limit = &limit
	return iq
}

// Offset to start from.
func (iq *InvoiceQuery) Offset(offset int) *InvoiceQuery {
	iq.ctx.Offset = &offset
	return iq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (iq *InvoiceQuery) Unique(unique bool) *InvoiceQuery {
	iq.ctx.Unique = &unique
	return iq
}

// Order specifies how the records should be ordered.
func (iq *InvoiceQuery) Order(o ...invoice.OrderOption) *InvoiceQuery {
	iq.order = append(iq.order, o...)
	return iq
}

// QueryCreator chains the current query on the "creator" edge.
func (iq *InvoiceQuery) QueryCreator() *UserQuery {
	query := (&UserClient{config: iq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(invoice.Table, invoice.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, invoice.CreatorTable, invoice.CreatorColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryItems chains the current query on the "items" edge.
func (iq *InvoiceQuery) QueryItems() *InvoiceItemQuery {
	query := (&InvoiceItemClient{config: iq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(invoice.Table, invoice.FieldID, selector),
			sqlgraph.To(invoiceitem.Table, invoiceitem.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, invoice.ItemsTable, invoice.ItemsColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPayments chains the current query on the "payments" edge.
func (iq *InvoiceQuery) QueryPayments() *PaymentQuery {
	query := (&PaymentClient{config: iq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(invoice.Table, invoice.FieldID, selector),
			sqlgraph.To(payment.Table, payment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, invoice.PaymentsTable, invoice.PaymentsColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCustomer chains the current query on the "customer" edge.
func (iq *InvoiceQuery) QueryCustomer() *CustomerQuery {
	query := (&CustomerClient{config: iq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(invoice.Table, invoice.FieldID, selector),
			sqlgraph.To(customer.Table, customer.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, invoice.CustomerTable, invoice.CustomerColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Invoice entity from the query.
// Returns a *NotFoundError when no Invoice was found.
func (iq *InvoiceQuery) First(ctx context.Context) (*Invoice, error) {
	nodes, err := iq.Limit(1).All(setContextOp(ctx, iq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{invoice.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (iq *InvoiceQuery) FirstX(ctx context.Context) *Invoice {
	node, err := iq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Invoice ID from the query.
// Returns a *NotFoundError when no Invoice ID was found.
func (iq *InvoiceQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = iq.Limit(1).IDs(setContextOp(ctx, iq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{invoice.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (iq *InvoiceQuery) FirstIDX(ctx context.Context) int {
	id, err := iq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Invoice entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Invoice entity is found.
// Returns a *NotFoundError when no Invoice entities are found.
func (iq *InvoiceQuery) Only(ctx context.Context) (*Invoice, error) {
	nodes, err := iq.Limit(2).All(setContextOp(ctx, iq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{invoice.Label}
	default:
		return nil, &NotSingularError{invoice.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (iq *InvoiceQuery) OnlyX(ctx context.Context) *Invoice {
	node, err := iq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Invoice ID in the query.
// Returns a *NotSingularError when more than one Invoice ID is found.
// Returns a *NotFoundError when no entities are found.
func (iq *InvoiceQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = iq.Limit(2).IDs(setContextOp(ctx, iq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{invoice.Label}
	default:
		err = &NotSingularError{invoice.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (iq *InvoiceQuery) OnlyIDX(ctx context.Context) int {
	id, err := iq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Invoices.
func (iq *InvoiceQuery) All(ctx context.Context) ([]*Invoice, error) {
	ctx = setContextOp(ctx, iq.ctx, ent.OpQueryAll)
	if err := iq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Invoice, *InvoiceQuery]()
	return withInterceptors[[]*Invoice](ctx, iq, qr, iq.inters)
}

// AllX is like All, but panics if an error occurs.
func (iq *InvoiceQuery) AllX(ctx context.Context) []*Invoice {
	nodes, err := iq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Invoice IDs.
func (iq *InvoiceQuery) IDs(ctx context.Context) (ids []int, err error) {
	if iq.ctx.Unique == nil && iq.path != nil {
		iq.Unique(true)
	}
	ctx = setContextOp(ctx, iq.ctx, ent.OpQueryIDs)
	if err = iq.Select(invoice.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (iq *InvoiceQuery) IDsX(ctx context.Context) []int {
	ids, err := iq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (iq *InvoiceQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, iq.ctx, ent.OpQueryCount)
	if err := iq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, iq, querierCount[*InvoiceQuery](), iq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (iq *InvoiceQuery) CountX(ctx context.Context) int {
	count, err := iq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (iq *InvoiceQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, iq.ctx, ent.OpQueryExist)
	switch _, err := iq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (iq *InvoiceQuery) ExistX(ctx context.Context) bool {
	exist, err := iq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the InvoiceQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (iq *InvoiceQuery) Clone() *InvoiceQuery {
	if iq == nil {
		return nil
	}
	return &InvoiceQuery{
		config:       iq.config,
		ctx:          iq.ctx.Clone(),
		order:        append([]invoice.OrderOption{}, iq.order...),
		inters:       append([]Interceptor{}, iq.inters...),
		predicates:   append([]predicate.Invoice{}, iq.predicates...),
		withCreator:  iq.withCreator.Clone(),
		withItems:    iq.withItems.Clone(),
		withPayments: iq.withPayments.Clone(),
		withCustomer: iq.withCustomer.Clone(),
		// clone intermediate query.
		sql:  iq.sql.Clone(),
		path: iq.path,
	}
}

// WithCreator tells the query-builder to eager-load the nodes that are connected to
// the "creator" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *InvoiceQuery) WithCreator(opts ...func(*UserQuery)) *InvoiceQuery {
	query := (&UserClient{config: iq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	iq.withCreator = query
	return iq
}

// WithItems tells the query-builder to eager-load the nodes that are connected to
// the "items" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *InvoiceQuery) WithItems(opts ...func(*InvoiceItemQuery)) *InvoiceQuery {
	query := (&InvoiceItemClient{config: iq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	iq.withItems = query
	return iq
}

// WithPayments tells the query-builder to eager-load the nodes that are connected to
// the "payments" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *InvoiceQuery) WithPayments(opts ...func(*PaymentQuery)) *InvoiceQuery {
	query := (&PaymentClient{config: iq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	iq.withPayments = query
	return iq
}

// WithCustomer tells the query-builder to eager-load the nodes that are connected to
// the "customer" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *InvoiceQuery) WithCustomer(opts ...func(*CustomerQuery)) *InvoiceQuery {
	query := (&CustomerClient{config: iq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	iq.withCustomer = query
	return iq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ReferenceNumber string `json:"reference_number,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Invoice.Query().
//		GroupBy(invoice.FieldReferenceNumber).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (iq *InvoiceQuery) GroupBy(field string, fields ...string) *InvoiceGroupBy {
	iq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &InvoiceGroupBy{build: iq}
	grbuild.flds = &iq.ctx.Fields
	grbuild.label = invoice.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ReferenceNumber string `json:"reference_number,omitempty"`
//	}
//
//	client.Invoice.Query().
//		Select(invoice.FieldReferenceNumber).
//		Scan(ctx, &v)
func (iq *InvoiceQuery) Select(fields ...string) *InvoiceSelect {
	iq.ctx.Fields = append(iq.ctx.Fields, fields...)
	sbuild := &InvoiceSelect{InvoiceQuery: iq}
	sbuild.label = invoice.Label
	sbuild.flds, sbuild.scan = &iq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a InvoiceSelect configured with the given aggregations.
func (iq *InvoiceQuery) Aggregate(fns ...AggregateFunc) *InvoiceSelect {
	return iq.Select().Aggregate(fns...)
}

func (iq *InvoiceQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range iq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, iq); err != nil {
				return err
			}
		}
	}
	for _, f := range iq.ctx.Fields {
		if !invoice.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if iq.path != nil {
		prev, err := iq.path(ctx)
		if err != nil {
			return err
		}
		iq.sql = prev
	}
	return nil
}

func (iq *InvoiceQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Invoice, error) {
	var (
		nodes       = []*Invoice{}
		withFKs     = iq.withFKs
		_spec       = iq.querySpec()
		loadedTypes = [4]bool{
			iq.withCreator != nil,
			iq.withItems != nil,
			iq.withPayments != nil,
			iq.withCustomer != nil,
		}
	)
	if iq.withCreator != nil || iq.withCustomer != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, invoice.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Invoice).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Invoice{config: iq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, iq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := iq.withCreator; query != nil {
		if err := iq.loadCreator(ctx, query, nodes, nil,
			func(n *Invoice, e *User) { n.Edges.Creator = e }); err != nil {
			return nil, err
		}
	}
	if query := iq.withItems; query != nil {
		if err := iq.loadItems(ctx, query, nodes,
			func(n *Invoice) { n.Edges.Items = []*InvoiceItem{} },
			func(n *Invoice, e *InvoiceItem) { n.Edges.Items = append(n.Edges.Items, e) }); err != nil {
			return nil, err
		}
	}
	if query := iq.withPayments; query != nil {
		if err := iq.loadPayments(ctx, query, nodes,
			func(n *Invoice) { n.Edges.Payments = []*Payment{} },
			func(n *Invoice, e *Payment) { n.Edges.Payments = append(n.Edges.Payments, e) }); err != nil {
			return nil, err
		}
	}
	if query := iq.withCustomer; query != nil {
		if err := iq.loadCustomer(ctx, query, nodes, nil,
			func(n *Invoice, e *Customer) { n.Edges.Customer = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (iq *InvoiceQuery) loadCreator(ctx context.Context, query *UserQuery, nodes []*Invoice, init func(*Invoice), assign func(*Invoice, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Invoice)
	for i := range nodes {
		if nodes[i].user_invoices == nil {
			continue
		}
		fk := *nodes[i].user_invoices
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
			return fmt.Errorf(`unexpected foreign-key "user_invoices" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (iq *InvoiceQuery) loadItems(ctx context.Context, query *InvoiceItemQuery, nodes []*Invoice, init func(*Invoice), assign func(*Invoice, *InvoiceItem)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Invoice)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.InvoiceItem(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(invoice.ItemsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.invoice_items
		if fk == nil {
			return fmt.Errorf(`foreign-key "invoice_items" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "invoice_items" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (iq *InvoiceQuery) loadPayments(ctx context.Context, query *PaymentQuery, nodes []*Invoice, init func(*Invoice), assign func(*Invoice, *Payment)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Invoice)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(invoice.PaymentsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.invoice_payments
		if fk == nil {
			return fmt.Errorf(`foreign-key "invoice_payments" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "invoice_payments" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (iq *InvoiceQuery) loadCustomer(ctx context.Context, query *CustomerQuery, nodes []*Invoice, init func(*Invoice), assign func(*Invoice, *Customer)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Invoice)
	for i := range nodes {
		if nodes[i].invoice_customer == nil {
			continue
		}
		fk := *nodes[i].invoice_customer
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(customer.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "invoice_customer" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (iq *InvoiceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := iq.querySpec()
	_spec.Node.Columns = iq.ctx.Fields
	if len(iq.ctx.Fields) > 0 {
		_spec.Unique = iq.ctx.Unique != nil && *iq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, iq.driver, _spec)
}

func (iq *InvoiceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(invoice.Table, invoice.Columns, sqlgraph.NewFieldSpec(invoice.FieldID, field.TypeInt))
	_spec.From = iq.sql
	if unique := iq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if iq.path != nil {
		_spec.Unique = true
	}
	if fields := iq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, invoice.FieldID)
		for i := range fields {
			if fields[i] != invoice.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := iq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := iq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := iq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := iq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (iq *InvoiceQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(iq.driver.Dialect())
	t1 := builder.Table(invoice.Table)
	columns := iq.ctx.Fields
	if len(columns) == 0 {
		columns = invoice.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if iq.sql != nil {
		selector = iq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if iq.ctx.Unique != nil && *iq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range iq.predicates {
		p(selector)
	}
	for _, p := range iq.order {
		p(selector)
	}
	if offset := iq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := iq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// InvoiceGroupBy is the group-by builder for Invoice entities.
type InvoiceGroupBy struct {
	selector
	build *InvoiceQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (igb *InvoiceGroupBy) Aggregate(fns ...AggregateFunc) *InvoiceGroupBy {
	igb.fns = append(igb.fns, fns...)
	return igb
}

// Scan applies the selector query and scans the result into the given value.
func (igb *InvoiceGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, igb.build.ctx, ent.OpQueryGroupBy)
	if err := igb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*InvoiceQuery, *InvoiceGroupBy](ctx, igb.build, igb, igb.build.inters, v)
}

func (igb *InvoiceGroupBy) sqlScan(ctx context.Context, root *InvoiceQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(igb.fns))
	for _, fn := range igb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*igb.flds)+len(igb.fns))
		for _, f := range *igb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*igb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := igb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// InvoiceSelect is the builder for selecting fields of Invoice entities.
type InvoiceSelect struct {
	*InvoiceQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (is *InvoiceSelect) Aggregate(fns ...AggregateFunc) *InvoiceSelect {
	is.fns = append(is.fns, fns...)
	return is
}

// Scan applies the selector query and scans the result into the given value.
func (is *InvoiceSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, is.ctx, ent.OpQuerySelect)
	if err := is.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*InvoiceQuery, *InvoiceSelect](ctx, is.InvoiceQuery, is, is.inters, v)
}

func (is *InvoiceSelect) sqlScan(ctx context.Context, root *InvoiceQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(is.fns))
	for _, fn := range is.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*is.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := is.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
