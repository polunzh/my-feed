// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"polunzh/my-feed/ent/feed"
	"polunzh/my-feed/ent/group"
	"polunzh/my-feed/ent/predicate"
	"sync"
	"time"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeFeed  = "Feed"
	TypeGroup = "Group"
)

// FeedMutation represents an operation that mutates the Feed nodes in the graph.
type FeedMutation struct {
	config
	op            Op
	typ           string
	id            *int
	name          *string
	url           *string
	updated_at    *time.Time
	created_at    *time.Time
	clearedFields map[string]struct{}
	group         map[int]struct{}
	removedgroup  map[int]struct{}
	clearedgroup  bool
	done          bool
	oldValue      func(context.Context) (*Feed, error)
	predicates    []predicate.Feed
}

var _ ent.Mutation = (*FeedMutation)(nil)

// feedOption allows management of the mutation configuration using functional options.
type feedOption func(*FeedMutation)

// newFeedMutation creates new mutation for the Feed entity.
func newFeedMutation(c config, op Op, opts ...feedOption) *FeedMutation {
	m := &FeedMutation{
		config:        c,
		op:            op,
		typ:           TypeFeed,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withFeedID sets the ID field of the mutation.
func withFeedID(id int) feedOption {
	return func(m *FeedMutation) {
		var (
			err   error
			once  sync.Once
			value *Feed
		)
		m.oldValue = func(ctx context.Context) (*Feed, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Feed.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withFeed sets the old Feed of the mutation.
func withFeed(node *Feed) feedOption {
	return func(m *FeedMutation) {
		m.oldValue = func(context.Context) (*Feed, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m FeedMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m FeedMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *FeedMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *FeedMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Feed.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *FeedMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *FeedMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Feed entity.
// If the Feed object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FeedMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *FeedMutation) ResetName() {
	m.name = nil
}

// SetURL sets the "url" field.
func (m *FeedMutation) SetURL(s string) {
	m.url = &s
}

// URL returns the value of the "url" field in the mutation.
func (m *FeedMutation) URL() (r string, exists bool) {
	v := m.url
	if v == nil {
		return
	}
	return *v, true
}

// OldURL returns the old "url" field's value of the Feed entity.
// If the Feed object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FeedMutation) OldURL(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldURL is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldURL requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldURL: %w", err)
	}
	return oldValue.URL, nil
}

// ResetURL resets all changes to the "url" field.
func (m *FeedMutation) ResetURL() {
	m.url = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *FeedMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *FeedMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Feed entity.
// If the Feed object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FeedMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *FeedMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *FeedMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *FeedMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Feed entity.
// If the Feed object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FeedMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *FeedMutation) ResetCreatedAt() {
	m.created_at = nil
}

// AddGroupIDs adds the "group" edge to the Group entity by ids.
func (m *FeedMutation) AddGroupIDs(ids ...int) {
	if m.group == nil {
		m.group = make(map[int]struct{})
	}
	for i := range ids {
		m.group[ids[i]] = struct{}{}
	}
}

// ClearGroup clears the "group" edge to the Group entity.
func (m *FeedMutation) ClearGroup() {
	m.clearedgroup = true
}

// GroupCleared reports if the "group" edge to the Group entity was cleared.
func (m *FeedMutation) GroupCleared() bool {
	return m.clearedgroup
}

// RemoveGroupIDs removes the "group" edge to the Group entity by IDs.
func (m *FeedMutation) RemoveGroupIDs(ids ...int) {
	if m.removedgroup == nil {
		m.removedgroup = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.group, ids[i])
		m.removedgroup[ids[i]] = struct{}{}
	}
}

// RemovedGroup returns the removed IDs of the "group" edge to the Group entity.
func (m *FeedMutation) RemovedGroupIDs() (ids []int) {
	for id := range m.removedgroup {
		ids = append(ids, id)
	}
	return
}

// GroupIDs returns the "group" edge IDs in the mutation.
func (m *FeedMutation) GroupIDs() (ids []int) {
	for id := range m.group {
		ids = append(ids, id)
	}
	return
}

// ResetGroup resets all changes to the "group" edge.
func (m *FeedMutation) ResetGroup() {
	m.group = nil
	m.clearedgroup = false
	m.removedgroup = nil
}

// Where appends a list predicates to the FeedMutation builder.
func (m *FeedMutation) Where(ps ...predicate.Feed) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *FeedMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Feed).
func (m *FeedMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *FeedMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.name != nil {
		fields = append(fields, feed.FieldName)
	}
	if m.url != nil {
		fields = append(fields, feed.FieldURL)
	}
	if m.updated_at != nil {
		fields = append(fields, feed.FieldUpdatedAt)
	}
	if m.created_at != nil {
		fields = append(fields, feed.FieldCreatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *FeedMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case feed.FieldName:
		return m.Name()
	case feed.FieldURL:
		return m.URL()
	case feed.FieldUpdatedAt:
		return m.UpdatedAt()
	case feed.FieldCreatedAt:
		return m.CreatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *FeedMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case feed.FieldName:
		return m.OldName(ctx)
	case feed.FieldURL:
		return m.OldURL(ctx)
	case feed.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case feed.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Feed field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *FeedMutation) SetField(name string, value ent.Value) error {
	switch name {
	case feed.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case feed.FieldURL:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetURL(v)
		return nil
	case feed.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case feed.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Feed field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *FeedMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *FeedMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *FeedMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Feed numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *FeedMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *FeedMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *FeedMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Feed nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *FeedMutation) ResetField(name string) error {
	switch name {
	case feed.FieldName:
		m.ResetName()
		return nil
	case feed.FieldURL:
		m.ResetURL()
		return nil
	case feed.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case feed.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	}
	return fmt.Errorf("unknown Feed field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *FeedMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.group != nil {
		edges = append(edges, feed.EdgeGroup)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *FeedMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case feed.EdgeGroup:
		ids := make([]ent.Value, 0, len(m.group))
		for id := range m.group {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *FeedMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedgroup != nil {
		edges = append(edges, feed.EdgeGroup)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *FeedMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case feed.EdgeGroup:
		ids := make([]ent.Value, 0, len(m.removedgroup))
		for id := range m.removedgroup {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *FeedMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedgroup {
		edges = append(edges, feed.EdgeGroup)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *FeedMutation) EdgeCleared(name string) bool {
	switch name {
	case feed.EdgeGroup:
		return m.clearedgroup
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *FeedMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Feed unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *FeedMutation) ResetEdge(name string) error {
	switch name {
	case feed.EdgeGroup:
		m.ResetGroup()
		return nil
	}
	return fmt.Errorf("unknown Feed edge %s", name)
}

// GroupMutation represents an operation that mutates the Group nodes in the graph.
type GroupMutation struct {
	config
	op            Op
	typ           string
	id            *int
	name          *string
	updated_at    *time.Time
	created_at    *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Group, error)
	predicates    []predicate.Group
}

var _ ent.Mutation = (*GroupMutation)(nil)

// groupOption allows management of the mutation configuration using functional options.
type groupOption func(*GroupMutation)

// newGroupMutation creates new mutation for the Group entity.
func newGroupMutation(c config, op Op, opts ...groupOption) *GroupMutation {
	m := &GroupMutation{
		config:        c,
		op:            op,
		typ:           TypeGroup,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withGroupID sets the ID field of the mutation.
func withGroupID(id int) groupOption {
	return func(m *GroupMutation) {
		var (
			err   error
			once  sync.Once
			value *Group
		)
		m.oldValue = func(ctx context.Context) (*Group, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Group.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withGroup sets the old Group of the mutation.
func withGroup(node *Group) groupOption {
	return func(m *GroupMutation) {
		m.oldValue = func(context.Context) (*Group, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m GroupMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m GroupMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *GroupMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *GroupMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Group.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *GroupMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *GroupMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Group entity.
// If the Group object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *GroupMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *GroupMutation) ResetName() {
	m.name = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *GroupMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *GroupMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Group entity.
// If the Group object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *GroupMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *GroupMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *GroupMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *GroupMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Group entity.
// If the Group object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *GroupMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *GroupMutation) ResetCreatedAt() {
	m.created_at = nil
}

// Where appends a list predicates to the GroupMutation builder.
func (m *GroupMutation) Where(ps ...predicate.Group) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *GroupMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Group).
func (m *GroupMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *GroupMutation) Fields() []string {
	fields := make([]string, 0, 3)
	if m.name != nil {
		fields = append(fields, group.FieldName)
	}
	if m.updated_at != nil {
		fields = append(fields, group.FieldUpdatedAt)
	}
	if m.created_at != nil {
		fields = append(fields, group.FieldCreatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *GroupMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case group.FieldName:
		return m.Name()
	case group.FieldUpdatedAt:
		return m.UpdatedAt()
	case group.FieldCreatedAt:
		return m.CreatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *GroupMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case group.FieldName:
		return m.OldName(ctx)
	case group.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case group.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Group field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *GroupMutation) SetField(name string, value ent.Value) error {
	switch name {
	case group.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case group.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case group.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Group field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *GroupMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *GroupMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *GroupMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Group numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *GroupMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *GroupMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *GroupMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Group nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *GroupMutation) ResetField(name string) error {
	switch name {
	case group.FieldName:
		m.ResetName()
		return nil
	case group.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case group.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	}
	return fmt.Errorf("unknown Group field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *GroupMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *GroupMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *GroupMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *GroupMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *GroupMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *GroupMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *GroupMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Group unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *GroupMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Group edge %s", name)
}
