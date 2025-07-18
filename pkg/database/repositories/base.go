package repositories

import (
	"fmt"
	"reflect"

	"github.com/cybercoder/restbill/pkg/database"
	"gorm.io/gorm"
)

// QueryOperator defines supported query operators
type QueryOperator string

const (
	Equal          QueryOperator = "="
	NotEqual       QueryOperator = "!="
	GreaterThan    QueryOperator = ">"
	GreaterOrEqual QueryOperator = ">="
	LessThan       QueryOperator = "<"
	LessOrEqual    QueryOperator = "<="
	Like           QueryOperator = "LIKE"
	In             QueryOperator = "IN"
	NotIn          QueryOperator = "NOT IN"
	IsNull         QueryOperator = "IS NULL"
	IsNotNull      QueryOperator = "IS NOT NULL"
)

// Condition represents a single query condition
type Condition struct {
	Field    string
	Operator QueryOperator
	Value    interface{}
}

// LogicalGroup represents a group of conditions with a logical operator
type LogicalGroup struct {
	Conditions []Condition
	Operator   string // "AND" or "OR"
}

// QueryOptions provides additional query options
type QueryOptions struct {
	Select  []string
	Preload []string
	Order   []string
	Limit   *int
	Offset  *int
}

// Repository is a generic repository implementation
type Repository[T any] struct {
	db *gorm.DB
}

// NewRepository creates a new Repository instance
func NewRepository[T any]() *Repository[T] {
	return &Repository[T]{db: database.GetDB()}
}

// ================ CRUD Operations ================

// Create inserts a new entity
func (r *Repository[T]) Create(entity *T) (*T, error) {
	if err := r.db.Create(entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

// CreateBatch inserts multiple entities
func (r *Repository[T]) CreateBatch(entities []*T) ([]*T, error) {
	if err := r.db.Create(entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

// GetByID finds an entity by its primary key
func (r *Repository[T]) GetByID(id uint, opts ...QueryOptions) (*T, error) {
	var entity T
	db := r.applyOptions(r.db, opts...)

	if err := db.First(&entity, id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// Update updates an existing entity
func (r *Repository[T]) Update(entity *T) (*T, error) {
	if err := r.db.Save(entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

// PartialUpdate updates specific fields of an entity
func (r *Repository[T]) PartialUpdate(id uint, updates map[string]interface{}) (*T, error) {
	var entity T
	if err := r.db.Model(&entity).Where("id = ?", id).Updates(updates).Error; err != nil {
		return nil, err
	}

	// Reload the entity to get all fields
	return r.GetByID(id)
}

// Delete removes an entity by its primary key
func (r *Repository[T]) Delete(id uint) error {
	var entity T
	if err := r.db.Delete(&entity, id).Error; err != nil {
		return err
	}
	return nil
}

// ================ Query Operations ================

// FindFirst finds the first record matching conditions
func (r *Repository[T]) FindFirst(conditions []Condition, opts ...QueryOptions) (*T, error) {
	var entity T
	db := r.applyConditions(r.db, conditions)
	db = r.applyOptions(db, opts...)

	if err := db.First(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// FindAll finds all records matching conditions
func (r *Repository[T]) FindAll(conditions []Condition, opts ...QueryOptions) ([]*T, error) {
	var entities []*T
	db := r.applyConditions(r.db, conditions)
	db = r.applyOptions(db, opts...)

	if err := db.Find(&entities).Error; err != nil {
		return nil, fmt.Errorf("find all failed: %w", err)
	}
	return entities, nil
}

// FindAllAdvanced finds records with complex logical conditions
func (r *Repository[T]) FindAllAdvanced(groups []LogicalGroup, opts ...QueryOptions) ([]*T, error) {
	var entities []*T
	db := r.applyLogicalGroups(r.db, groups)
	db = r.applyOptions(db, opts...)

	if err := db.Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

// Count counts records matching conditions
func (r *Repository[T]) Count(conditions []Condition) (int64, error) {
	var count int64
	var entity T
	db := r.applyConditions(r.db.Model(&entity), conditions)

	if err := db.Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// Exists checks if any record matches conditions
func (r *Repository[T]) Exists(conditions []Condition) (bool, error) {
	count, err := r.Count(conditions)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// ================ Helper Methods ================

func (r *Repository[T]) applyConditions(db *gorm.DB, conditions []Condition) *gorm.DB {
	for _, cond := range conditions {
		switch cond.Operator {
		case Equal, NotEqual, GreaterThan, GreaterOrEqual, LessThan, LessOrEqual:
			db = db.Where(fmt.Sprintf("%s %s ?", cond.Field, cond.Operator), cond.Value)
		case Like:
			db = db.Where(fmt.Sprintf("%s LIKE ?", cond.Field), cond.Value)
		case In:
			db = db.Where(fmt.Sprintf("%s IN (?)", cond.Field), cond.Value)
		case NotIn:
			db = db.Where(fmt.Sprintf("%s NOT IN (?)", cond.Field), cond.Value)
		case IsNull:
			db = db.Where(fmt.Sprintf("%s IS NULL", cond.Field))
		case IsNotNull:
			db = db.Where(fmt.Sprintf("%s IS NOT NULL", cond.Field))
		default:
			db = db.Where(fmt.Sprintf("%s = ?", cond.Field), cond.Value)
		}
	}
	return db
}

func (r *Repository[T]) applyLogicalGroups(db *gorm.DB, groups []LogicalGroup) *gorm.DB {
	for _, group := range groups {
		switch group.Operator {
		case "OR":
			db = db.Or(r.applyConditions(db, group.Conditions))
		default: // default to AND
			db = db.Where(r.applyConditions(db, group.Conditions))
		}
	}
	return db
}

func (r *Repository[T]) applyOptions(db *gorm.DB, opts ...QueryOptions) *gorm.DB {
	if len(opts) == 0 {
		return db
	}

	opt := opts[0]

	// Apply select fields
	if len(opt.Select) > 0 {
		db = db.Select(opt.Select)
	}

	// Apply preloads
	for _, preload := range opt.Preload {
		db = db.Preload(preload)
	}

	// Apply ordering
	for _, order := range opt.Order {
		db = db.Order(order)
	}

	// Apply limit
	if opt.Limit != nil {
		db = db.Limit(*opt.Limit)
	}

	// Apply offset
	if opt.Offset != nil {
		db = db.Offset(*opt.Offset)
	}

	return db
}

// GetDB returns the underlying DB instance (use with caution)
func (r *Repository[T]) GetDB() *gorm.DB {
	return r.db
}

// GetModelType returns the reflect.Type of the repository's model
func (r *Repository[T]) GetModelType() reflect.Type {
	var t T
	return reflect.TypeOf(t)
}
