// Code generated by ent, DO NOT EDIT.

package filehistory

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/datumforge/enthistory"
)

const (
	// Label holds the string label denoting the filehistory type in the database.
	Label = "file_history"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldHistoryTime holds the string denoting the history_time field in the database.
	FieldHistoryTime = "history_time"
	// FieldOperation holds the string denoting the operation field in the database.
	FieldOperation = "operation"
	// FieldRef holds the string denoting the ref field in the database.
	FieldRef = "ref"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldUpdatedBy holds the string denoting the updated_by field in the database.
	FieldUpdatedBy = "updated_by"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldDeletedBy holds the string denoting the deleted_by field in the database.
	FieldDeletedBy = "deleted_by"
	// FieldMappingID holds the string denoting the mapping_id field in the database.
	FieldMappingID = "mapping_id"
	// FieldTags holds the string denoting the tags field in the database.
	FieldTags = "tags"
	// FieldFileName holds the string denoting the file_name field in the database.
	FieldFileName = "file_name"
	// FieldFileExtension holds the string denoting the file_extension field in the database.
	FieldFileExtension = "file_extension"
	// FieldFileSize holds the string denoting the file_size field in the database.
	FieldFileSize = "file_size"
	// FieldContentType holds the string denoting the content_type field in the database.
	FieldContentType = "content_type"
	// FieldStoreKey holds the string denoting the store_key field in the database.
	FieldStoreKey = "store_key"
	// FieldCategory holds the string denoting the category field in the database.
	FieldCategory = "category"
	// FieldAnnotation holds the string denoting the annotation field in the database.
	FieldAnnotation = "annotation"
	// Table holds the table name of the filehistory in the database.
	Table = "file_history"
)

// Columns holds all SQL columns for filehistory fields.
var Columns = []string{
	FieldID,
	FieldHistoryTime,
	FieldOperation,
	FieldRef,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldDeletedAt,
	FieldDeletedBy,
	FieldMappingID,
	FieldTags,
	FieldFileName,
	FieldFileExtension,
	FieldFileSize,
	FieldContentType,
	FieldStoreKey,
	FieldCategory,
	FieldAnnotation,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultHistoryTime holds the default value on creation for the "history_time" field.
	DefaultHistoryTime func() time.Time
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// DefaultMappingID holds the default value on creation for the "mapping_id" field.
	DefaultMappingID func() string
	// DefaultTags holds the default value on creation for the "tags" field.
	DefaultTags []string
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)

// OperationValidator is a validator for the "operation" field enum values. It is called by the builders before save.
func OperationValidator(o enthistory.OpType) error {
	switch o.String() {
	case "INSERT", "UPDATE", "DELETE":
		return nil
	default:
		return fmt.Errorf("filehistory: invalid enum value for operation field: %q", o)
	}
}

// OrderOption defines the ordering options for the FileHistory queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByHistoryTime orders the results by the history_time field.
func ByHistoryTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHistoryTime, opts...).ToFunc()
}

// ByOperation orders the results by the operation field.
func ByOperation(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOperation, opts...).ToFunc()
}

// ByRef orders the results by the ref field.
func ByRef(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRef, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByCreatedBy orders the results by the created_by field.
func ByCreatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedBy, opts...).ToFunc()
}

// ByUpdatedBy orders the results by the updated_by field.
func ByUpdatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedBy, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByDeletedBy orders the results by the deleted_by field.
func ByDeletedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedBy, opts...).ToFunc()
}

// ByMappingID orders the results by the mapping_id field.
func ByMappingID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMappingID, opts...).ToFunc()
}

// ByFileName orders the results by the file_name field.
func ByFileName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFileName, opts...).ToFunc()
}

// ByFileExtension orders the results by the file_extension field.
func ByFileExtension(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFileExtension, opts...).ToFunc()
}

// ByFileSize orders the results by the file_size field.
func ByFileSize(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFileSize, opts...).ToFunc()
}

// ByContentType orders the results by the content_type field.
func ByContentType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContentType, opts...).ToFunc()
}

// ByStoreKey orders the results by the store_key field.
func ByStoreKey(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStoreKey, opts...).ToFunc()
}

// ByCategory orders the results by the category field.
func ByCategory(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCategory, opts...).ToFunc()
}

// ByAnnotation orders the results by the annotation field.
func ByAnnotation(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAnnotation, opts...).ToFunc()
}

var (
	// enthistory.OpType must implement graphql.Marshaler.
	_ graphql.Marshaler = (*enthistory.OpType)(nil)
	// enthistory.OpType must implement graphql.Unmarshaler.
	_ graphql.Unmarshaler = (*enthistory.OpType)(nil)
)
