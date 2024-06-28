// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/generated/entitlementplan"
	"github.com/datumforge/datum/internal/ent/generated/entitlementplanfeature"
	"github.com/datumforge/datum/internal/ent/generated/event"
	"github.com/datumforge/datum/internal/ent/generated/feature"
	"github.com/datumforge/datum/internal/ent/generated/organization"
)

// FeatureCreate is the builder for creating a Feature entity.
type FeatureCreate struct {
	config
	mutation *FeatureMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (fc *FeatureCreate) SetCreatedAt(t time.Time) *FeatureCreate {
	fc.mutation.SetCreatedAt(t)
	return fc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fc *FeatureCreate) SetNillableCreatedAt(t *time.Time) *FeatureCreate {
	if t != nil {
		fc.SetCreatedAt(*t)
	}
	return fc
}

// SetUpdatedAt sets the "updated_at" field.
func (fc *FeatureCreate) SetUpdatedAt(t time.Time) *FeatureCreate {
	fc.mutation.SetUpdatedAt(t)
	return fc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (fc *FeatureCreate) SetNillableUpdatedAt(t *time.Time) *FeatureCreate {
	if t != nil {
		fc.SetUpdatedAt(*t)
	}
	return fc
}

// SetCreatedBy sets the "created_by" field.
func (fc *FeatureCreate) SetCreatedBy(s string) *FeatureCreate {
	fc.mutation.SetCreatedBy(s)
	return fc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (fc *FeatureCreate) SetNillableCreatedBy(s *string) *FeatureCreate {
	if s != nil {
		fc.SetCreatedBy(*s)
	}
	return fc
}

// SetUpdatedBy sets the "updated_by" field.
func (fc *FeatureCreate) SetUpdatedBy(s string) *FeatureCreate {
	fc.mutation.SetUpdatedBy(s)
	return fc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (fc *FeatureCreate) SetNillableUpdatedBy(s *string) *FeatureCreate {
	if s != nil {
		fc.SetUpdatedBy(*s)
	}
	return fc
}

// SetDeletedAt sets the "deleted_at" field.
func (fc *FeatureCreate) SetDeletedAt(t time.Time) *FeatureCreate {
	fc.mutation.SetDeletedAt(t)
	return fc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (fc *FeatureCreate) SetNillableDeletedAt(t *time.Time) *FeatureCreate {
	if t != nil {
		fc.SetDeletedAt(*t)
	}
	return fc
}

// SetDeletedBy sets the "deleted_by" field.
func (fc *FeatureCreate) SetDeletedBy(s string) *FeatureCreate {
	fc.mutation.SetDeletedBy(s)
	return fc
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (fc *FeatureCreate) SetNillableDeletedBy(s *string) *FeatureCreate {
	if s != nil {
		fc.SetDeletedBy(*s)
	}
	return fc
}

// SetMappingID sets the "mapping_id" field.
func (fc *FeatureCreate) SetMappingID(s string) *FeatureCreate {
	fc.mutation.SetMappingID(s)
	return fc
}

// SetNillableMappingID sets the "mapping_id" field if the given value is not nil.
func (fc *FeatureCreate) SetNillableMappingID(s *string) *FeatureCreate {
	if s != nil {
		fc.SetMappingID(*s)
	}
	return fc
}

// SetTags sets the "tags" field.
func (fc *FeatureCreate) SetTags(s []string) *FeatureCreate {
	fc.mutation.SetTags(s)
	return fc
}

// SetOwnerID sets the "owner_id" field.
func (fc *FeatureCreate) SetOwnerID(s string) *FeatureCreate {
	fc.mutation.SetOwnerID(s)
	return fc
}

// SetNillableOwnerID sets the "owner_id" field if the given value is not nil.
func (fc *FeatureCreate) SetNillableOwnerID(s *string) *FeatureCreate {
	if s != nil {
		fc.SetOwnerID(*s)
	}
	return fc
}

// SetName sets the "name" field.
func (fc *FeatureCreate) SetName(s string) *FeatureCreate {
	fc.mutation.SetName(s)
	return fc
}

// SetDisplayName sets the "display_name" field.
func (fc *FeatureCreate) SetDisplayName(s string) *FeatureCreate {
	fc.mutation.SetDisplayName(s)
	return fc
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (fc *FeatureCreate) SetNillableDisplayName(s *string) *FeatureCreate {
	if s != nil {
		fc.SetDisplayName(*s)
	}
	return fc
}

// SetEnabled sets the "enabled" field.
func (fc *FeatureCreate) SetEnabled(b bool) *FeatureCreate {
	fc.mutation.SetEnabled(b)
	return fc
}

// SetNillableEnabled sets the "enabled" field if the given value is not nil.
func (fc *FeatureCreate) SetNillableEnabled(b *bool) *FeatureCreate {
	if b != nil {
		fc.SetEnabled(*b)
	}
	return fc
}

// SetDescription sets the "description" field.
func (fc *FeatureCreate) SetDescription(s string) *FeatureCreate {
	fc.mutation.SetDescription(s)
	return fc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (fc *FeatureCreate) SetNillableDescription(s *string) *FeatureCreate {
	if s != nil {
		fc.SetDescription(*s)
	}
	return fc
}

// SetMetadata sets the "metadata" field.
func (fc *FeatureCreate) SetMetadata(m map[string]interface{}) *FeatureCreate {
	fc.mutation.SetMetadata(m)
	return fc
}

// SetID sets the "id" field.
func (fc *FeatureCreate) SetID(s string) *FeatureCreate {
	fc.mutation.SetID(s)
	return fc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (fc *FeatureCreate) SetNillableID(s *string) *FeatureCreate {
	if s != nil {
		fc.SetID(*s)
	}
	return fc
}

// SetOwner sets the "owner" edge to the Organization entity.
func (fc *FeatureCreate) SetOwner(o *Organization) *FeatureCreate {
	return fc.SetOwnerID(o.ID)
}

// AddPlanIDs adds the "plans" edge to the EntitlementPlan entity by IDs.
func (fc *FeatureCreate) AddPlanIDs(ids ...string) *FeatureCreate {
	fc.mutation.AddPlanIDs(ids...)
	return fc
}

// AddPlans adds the "plans" edges to the EntitlementPlan entity.
func (fc *FeatureCreate) AddPlans(e ...*EntitlementPlan) *FeatureCreate {
	ids := make([]string, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return fc.AddPlanIDs(ids...)
}

// AddEventIDs adds the "events" edge to the Event entity by IDs.
func (fc *FeatureCreate) AddEventIDs(ids ...string) *FeatureCreate {
	fc.mutation.AddEventIDs(ids...)
	return fc
}

// AddEvents adds the "events" edges to the Event entity.
func (fc *FeatureCreate) AddEvents(e ...*Event) *FeatureCreate {
	ids := make([]string, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return fc.AddEventIDs(ids...)
}

// AddFeatureIDs adds the "features" edge to the EntitlementPlanFeature entity by IDs.
func (fc *FeatureCreate) AddFeatureIDs(ids ...string) *FeatureCreate {
	fc.mutation.AddFeatureIDs(ids...)
	return fc
}

// AddFeatures adds the "features" edges to the EntitlementPlanFeature entity.
func (fc *FeatureCreate) AddFeatures(e ...*EntitlementPlanFeature) *FeatureCreate {
	ids := make([]string, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return fc.AddFeatureIDs(ids...)
}

// Mutation returns the FeatureMutation object of the builder.
func (fc *FeatureCreate) Mutation() *FeatureMutation {
	return fc.mutation
}

// Save creates the Feature in the database.
func (fc *FeatureCreate) Save(ctx context.Context) (*Feature, error) {
	if err := fc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, fc.sqlSave, fc.mutation, fc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FeatureCreate) SaveX(ctx context.Context) *Feature {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fc *FeatureCreate) Exec(ctx context.Context) error {
	_, err := fc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fc *FeatureCreate) ExecX(ctx context.Context) {
	if err := fc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fc *FeatureCreate) defaults() error {
	if _, ok := fc.mutation.CreatedAt(); !ok {
		if feature.DefaultCreatedAt == nil {
			return fmt.Errorf("generated: uninitialized feature.DefaultCreatedAt (forgotten import generated/runtime?)")
		}
		v := feature.DefaultCreatedAt()
		fc.mutation.SetCreatedAt(v)
	}
	if _, ok := fc.mutation.UpdatedAt(); !ok {
		if feature.DefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized feature.DefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := feature.DefaultUpdatedAt()
		fc.mutation.SetUpdatedAt(v)
	}
	if _, ok := fc.mutation.MappingID(); !ok {
		if feature.DefaultMappingID == nil {
			return fmt.Errorf("generated: uninitialized feature.DefaultMappingID (forgotten import generated/runtime?)")
		}
		v := feature.DefaultMappingID()
		fc.mutation.SetMappingID(v)
	}
	if _, ok := fc.mutation.Tags(); !ok {
		v := feature.DefaultTags
		fc.mutation.SetTags(v)
	}
	if _, ok := fc.mutation.Enabled(); !ok {
		v := feature.DefaultEnabled
		fc.mutation.SetEnabled(v)
	}
	if _, ok := fc.mutation.ID(); !ok {
		if feature.DefaultID == nil {
			return fmt.Errorf("generated: uninitialized feature.DefaultID (forgotten import generated/runtime?)")
		}
		v := feature.DefaultID()
		fc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (fc *FeatureCreate) check() error {
	if _, ok := fc.mutation.MappingID(); !ok {
		return &ValidationError{Name: "mapping_id", err: errors.New(`generated: missing required field "Feature.mapping_id"`)}
	}
	if v, ok := fc.mutation.OwnerID(); ok {
		if err := feature.OwnerIDValidator(v); err != nil {
			return &ValidationError{Name: "owner_id", err: fmt.Errorf(`generated: validator failed for field "Feature.owner_id": %w`, err)}
		}
	}
	if _, ok := fc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`generated: missing required field "Feature.name"`)}
	}
	if v, ok := fc.mutation.Name(); ok {
		if err := feature.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "Feature.name": %w`, err)}
		}
	}
	if _, ok := fc.mutation.Enabled(); !ok {
		return &ValidationError{Name: "enabled", err: errors.New(`generated: missing required field "Feature.enabled"`)}
	}
	return nil
}

func (fc *FeatureCreate) sqlSave(ctx context.Context) (*Feature, error) {
	if err := fc.check(); err != nil {
		return nil, err
	}
	_node, _spec := fc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Feature.ID type: %T", _spec.ID.Value)
		}
	}
	fc.mutation.id = &_node.ID
	fc.mutation.done = true
	return _node, nil
}

func (fc *FeatureCreate) createSpec() (*Feature, *sqlgraph.CreateSpec) {
	var (
		_node = &Feature{config: fc.config}
		_spec = sqlgraph.NewCreateSpec(feature.Table, sqlgraph.NewFieldSpec(feature.FieldID, field.TypeString))
	)
	_spec.Schema = fc.schemaConfig.Feature
	if id, ok := fc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := fc.mutation.CreatedAt(); ok {
		_spec.SetField(feature.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := fc.mutation.UpdatedAt(); ok {
		_spec.SetField(feature.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := fc.mutation.CreatedBy(); ok {
		_spec.SetField(feature.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := fc.mutation.UpdatedBy(); ok {
		_spec.SetField(feature.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := fc.mutation.DeletedAt(); ok {
		_spec.SetField(feature.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := fc.mutation.DeletedBy(); ok {
		_spec.SetField(feature.FieldDeletedBy, field.TypeString, value)
		_node.DeletedBy = value
	}
	if value, ok := fc.mutation.MappingID(); ok {
		_spec.SetField(feature.FieldMappingID, field.TypeString, value)
		_node.MappingID = value
	}
	if value, ok := fc.mutation.Tags(); ok {
		_spec.SetField(feature.FieldTags, field.TypeJSON, value)
		_node.Tags = value
	}
	if value, ok := fc.mutation.Name(); ok {
		_spec.SetField(feature.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := fc.mutation.DisplayName(); ok {
		_spec.SetField(feature.FieldDisplayName, field.TypeString, value)
		_node.DisplayName = value
	}
	if value, ok := fc.mutation.Enabled(); ok {
		_spec.SetField(feature.FieldEnabled, field.TypeBool, value)
		_node.Enabled = value
	}
	if value, ok := fc.mutation.Description(); ok {
		_spec.SetField(feature.FieldDescription, field.TypeString, value)
		_node.Description = &value
	}
	if value, ok := fc.mutation.Metadata(); ok {
		_spec.SetField(feature.FieldMetadata, field.TypeJSON, value)
		_node.Metadata = value
	}
	if nodes := fc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   feature.OwnerTable,
			Columns: []string{feature.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = fc.schemaConfig.Feature
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.OwnerID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fc.mutation.PlansIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   feature.PlansTable,
			Columns: feature.PlansPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(entitlementplan.FieldID, field.TypeString),
			},
		}
		edge.Schema = fc.schemaConfig.EntitlementPlanFeature
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &EntitlementPlanFeatureCreate{config: fc.config, mutation: newEntitlementPlanFeatureMutation(fc.config, OpCreate)}
		_ = createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fc.mutation.EventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   feature.EventsTable,
			Columns: feature.EventsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeString),
			},
		}
		edge.Schema = fc.schemaConfig.FeatureEvents
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fc.mutation.FeaturesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   feature.FeaturesTable,
			Columns: []string{feature.FeaturesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(entitlementplanfeature.FieldID, field.TypeString),
			},
		}
		edge.Schema = fc.schemaConfig.EntitlementPlanFeature
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// FeatureCreateBulk is the builder for creating many Feature entities in bulk.
type FeatureCreateBulk struct {
	config
	err      error
	builders []*FeatureCreate
}

// Save creates the Feature entities in the database.
func (fcb *FeatureCreateBulk) Save(ctx context.Context) ([]*Feature, error) {
	if fcb.err != nil {
		return nil, fcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(fcb.builders))
	nodes := make([]*Feature, len(fcb.builders))
	mutators := make([]Mutator, len(fcb.builders))
	for i := range fcb.builders {
		func(i int, root context.Context) {
			builder := fcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FeatureMutation)
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
					_, err = mutators[i+1].Mutate(root, fcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, fcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fcb *FeatureCreateBulk) SaveX(ctx context.Context) []*Feature {
	v, err := fcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fcb *FeatureCreateBulk) Exec(ctx context.Context) error {
	_, err := fcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fcb *FeatureCreateBulk) ExecX(ctx context.Context) {
	if err := fcb.Exec(ctx); err != nil {
		panic(err)
	}
}
