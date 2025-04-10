// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/aki-13627/animalia/backend-go/ent/followrelation"
	"github.com/aki-13627/animalia/backend-go/ent/user"
	"github.com/google/uuid"
)

// FollowRelationCreate is the builder for creating a FollowRelation entity.
type FollowRelationCreate struct {
	config
	mutation *FollowRelationMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (frc *FollowRelationCreate) SetCreatedAt(t time.Time) *FollowRelationCreate {
	frc.mutation.SetCreatedAt(t)
	return frc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (frc *FollowRelationCreate) SetNillableCreatedAt(t *time.Time) *FollowRelationCreate {
	if t != nil {
		frc.SetCreatedAt(*t)
	}
	return frc
}

// SetID sets the "id" field.
func (frc *FollowRelationCreate) SetID(u uuid.UUID) *FollowRelationCreate {
	frc.mutation.SetID(u)
	return frc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (frc *FollowRelationCreate) SetNillableID(u *uuid.UUID) *FollowRelationCreate {
	if u != nil {
		frc.SetID(*u)
	}
	return frc
}

// SetFromID sets the "from" edge to the User entity by ID.
func (frc *FollowRelationCreate) SetFromID(id uuid.UUID) *FollowRelationCreate {
	frc.mutation.SetFromID(id)
	return frc
}

// SetFrom sets the "from" edge to the User entity.
func (frc *FollowRelationCreate) SetFrom(u *User) *FollowRelationCreate {
	return frc.SetFromID(u.ID)
}

// SetToID sets the "to" edge to the User entity by ID.
func (frc *FollowRelationCreate) SetToID(id uuid.UUID) *FollowRelationCreate {
	frc.mutation.SetToID(id)
	return frc
}

// SetTo sets the "to" edge to the User entity.
func (frc *FollowRelationCreate) SetTo(u *User) *FollowRelationCreate {
	return frc.SetToID(u.ID)
}

// Mutation returns the FollowRelationMutation object of the builder.
func (frc *FollowRelationCreate) Mutation() *FollowRelationMutation {
	return frc.mutation
}

// Save creates the FollowRelation in the database.
func (frc *FollowRelationCreate) Save(ctx context.Context) (*FollowRelation, error) {
	frc.defaults()
	return withHooks(ctx, frc.sqlSave, frc.mutation, frc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (frc *FollowRelationCreate) SaveX(ctx context.Context) *FollowRelation {
	v, err := frc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (frc *FollowRelationCreate) Exec(ctx context.Context) error {
	_, err := frc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (frc *FollowRelationCreate) ExecX(ctx context.Context) {
	if err := frc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (frc *FollowRelationCreate) defaults() {
	if _, ok := frc.mutation.CreatedAt(); !ok {
		v := followrelation.DefaultCreatedAt()
		frc.mutation.SetCreatedAt(v)
	}
	if _, ok := frc.mutation.ID(); !ok {
		v := followrelation.DefaultID()
		frc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (frc *FollowRelationCreate) check() error {
	if _, ok := frc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "FollowRelation.created_at"`)}
	}
	if len(frc.mutation.FromIDs()) == 0 {
		return &ValidationError{Name: "from", err: errors.New(`ent: missing required edge "FollowRelation.from"`)}
	}
	if len(frc.mutation.ToIDs()) == 0 {
		return &ValidationError{Name: "to", err: errors.New(`ent: missing required edge "FollowRelation.to"`)}
	}
	return nil
}

func (frc *FollowRelationCreate) sqlSave(ctx context.Context) (*FollowRelation, error) {
	if err := frc.check(); err != nil {
		return nil, err
	}
	_node, _spec := frc.createSpec()
	if err := sqlgraph.CreateNode(ctx, frc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	frc.mutation.id = &_node.ID
	frc.mutation.done = true
	return _node, nil
}

func (frc *FollowRelationCreate) createSpec() (*FollowRelation, *sqlgraph.CreateSpec) {
	var (
		_node = &FollowRelation{config: frc.config}
		_spec = sqlgraph.NewCreateSpec(followrelation.Table, sqlgraph.NewFieldSpec(followrelation.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = frc.conflict
	if id, ok := frc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := frc.mutation.CreatedAt(); ok {
		_spec.SetField(followrelation.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := frc.mutation.FromIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   followrelation.FromTable,
			Columns: []string{followrelation.FromColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_following = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := frc.mutation.ToIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   followrelation.ToTable,
			Columns: []string{followrelation.ToColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_followers = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.FollowRelation.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FollowRelationUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (frc *FollowRelationCreate) OnConflict(opts ...sql.ConflictOption) *FollowRelationUpsertOne {
	frc.conflict = opts
	return &FollowRelationUpsertOne{
		create: frc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.FollowRelation.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (frc *FollowRelationCreate) OnConflictColumns(columns ...string) *FollowRelationUpsertOne {
	frc.conflict = append(frc.conflict, sql.ConflictColumns(columns...))
	return &FollowRelationUpsertOne{
		create: frc,
	}
}

type (
	// FollowRelationUpsertOne is the builder for "upsert"-ing
	//  one FollowRelation node.
	FollowRelationUpsertOne struct {
		create *FollowRelationCreate
	}

	// FollowRelationUpsert is the "OnConflict" setter.
	FollowRelationUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *FollowRelationUpsert) SetCreatedAt(v time.Time) *FollowRelationUpsert {
	u.Set(followrelation.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *FollowRelationUpsert) UpdateCreatedAt() *FollowRelationUpsert {
	u.SetExcluded(followrelation.FieldCreatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.FollowRelation.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(followrelation.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *FollowRelationUpsertOne) UpdateNewValues() *FollowRelationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(followrelation.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.FollowRelation.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *FollowRelationUpsertOne) Ignore() *FollowRelationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FollowRelationUpsertOne) DoNothing() *FollowRelationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FollowRelationCreate.OnConflict
// documentation for more info.
func (u *FollowRelationUpsertOne) Update(set func(*FollowRelationUpsert)) *FollowRelationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FollowRelationUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *FollowRelationUpsertOne) SetCreatedAt(v time.Time) *FollowRelationUpsertOne {
	return u.Update(func(s *FollowRelationUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *FollowRelationUpsertOne) UpdateCreatedAt() *FollowRelationUpsertOne {
	return u.Update(func(s *FollowRelationUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *FollowRelationUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for FollowRelationCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FollowRelationUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *FollowRelationUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: FollowRelationUpsertOne.ID is not supported by MySQL driver. Use FollowRelationUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *FollowRelationUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// FollowRelationCreateBulk is the builder for creating many FollowRelation entities in bulk.
type FollowRelationCreateBulk struct {
	config
	err      error
	builders []*FollowRelationCreate
	conflict []sql.ConflictOption
}

// Save creates the FollowRelation entities in the database.
func (frcb *FollowRelationCreateBulk) Save(ctx context.Context) ([]*FollowRelation, error) {
	if frcb.err != nil {
		return nil, frcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(frcb.builders))
	nodes := make([]*FollowRelation, len(frcb.builders))
	mutators := make([]Mutator, len(frcb.builders))
	for i := range frcb.builders {
		func(i int, root context.Context) {
			builder := frcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FollowRelationMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, frcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = frcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, frcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, frcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (frcb *FollowRelationCreateBulk) SaveX(ctx context.Context) []*FollowRelation {
	v, err := frcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (frcb *FollowRelationCreateBulk) Exec(ctx context.Context) error {
	_, err := frcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (frcb *FollowRelationCreateBulk) ExecX(ctx context.Context) {
	if err := frcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.FollowRelation.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FollowRelationUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (frcb *FollowRelationCreateBulk) OnConflict(opts ...sql.ConflictOption) *FollowRelationUpsertBulk {
	frcb.conflict = opts
	return &FollowRelationUpsertBulk{
		create: frcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.FollowRelation.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (frcb *FollowRelationCreateBulk) OnConflictColumns(columns ...string) *FollowRelationUpsertBulk {
	frcb.conflict = append(frcb.conflict, sql.ConflictColumns(columns...))
	return &FollowRelationUpsertBulk{
		create: frcb,
	}
}

// FollowRelationUpsertBulk is the builder for "upsert"-ing
// a bulk of FollowRelation nodes.
type FollowRelationUpsertBulk struct {
	create *FollowRelationCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.FollowRelation.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(followrelation.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *FollowRelationUpsertBulk) UpdateNewValues() *FollowRelationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(followrelation.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.FollowRelation.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *FollowRelationUpsertBulk) Ignore() *FollowRelationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FollowRelationUpsertBulk) DoNothing() *FollowRelationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FollowRelationCreateBulk.OnConflict
// documentation for more info.
func (u *FollowRelationUpsertBulk) Update(set func(*FollowRelationUpsert)) *FollowRelationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FollowRelationUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *FollowRelationUpsertBulk) SetCreatedAt(v time.Time) *FollowRelationUpsertBulk {
	return u.Update(func(s *FollowRelationUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *FollowRelationUpsertBulk) UpdateCreatedAt() *FollowRelationUpsertBulk {
	return u.Update(func(s *FollowRelationUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *FollowRelationUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the FollowRelationCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for FollowRelationCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FollowRelationUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
