// Code generated by ent, DO NOT EDIT.

package generated

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/datumforge/datum/internal/ent/generated/filehistory"
	"github.com/datumforge/enthistory"
)

// FileHistory is the model entity for the FileHistory schema.
type FileHistory struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// HistoryTime holds the value of the "history_time" field.
	HistoryTime time.Time `json:"history_time,omitempty"`
	// Operation holds the value of the "operation" field.
	Operation enthistory.OpType `json:"operation,omitempty"`
	// Ref holds the value of the "ref" field.
	Ref string `json:"ref,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy string `json:"created_by,omitempty"`
	// UpdatedBy holds the value of the "updated_by" field.
	UpdatedBy string `json:"updated_by,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// DeletedBy holds the value of the "deleted_by" field.
	DeletedBy string `json:"deleted_by,omitempty"`
	// MappingID holds the value of the "mapping_id" field.
	MappingID string `json:"mapping_id,omitempty"`
	// tags associated with the object
	Tags []string `json:"tags,omitempty"`
	// FileName holds the value of the "file_name" field.
	FileName string `json:"file_name,omitempty"`
	// FileExtension holds the value of the "file_extension" field.
	FileExtension string `json:"file_extension,omitempty"`
	// FileSize holds the value of the "file_size" field.
	FileSize int `json:"file_size,omitempty"`
	// ContentType holds the value of the "content_type" field.
	ContentType string `json:"content_type,omitempty"`
	// StoreKey holds the value of the "store_key" field.
	StoreKey string `json:"store_key,omitempty"`
	// Category holds the value of the "category" field.
	Category string `json:"category,omitempty"`
	// Annotation holds the value of the "annotation" field.
	Annotation   string `json:"annotation,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*FileHistory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case filehistory.FieldTags:
			values[i] = new([]byte)
		case filehistory.FieldOperation:
			values[i] = new(enthistory.OpType)
		case filehistory.FieldFileSize:
			values[i] = new(sql.NullInt64)
		case filehistory.FieldID, filehistory.FieldRef, filehistory.FieldCreatedBy, filehistory.FieldUpdatedBy, filehistory.FieldDeletedBy, filehistory.FieldMappingID, filehistory.FieldFileName, filehistory.FieldFileExtension, filehistory.FieldContentType, filehistory.FieldStoreKey, filehistory.FieldCategory, filehistory.FieldAnnotation:
			values[i] = new(sql.NullString)
		case filehistory.FieldHistoryTime, filehistory.FieldCreatedAt, filehistory.FieldUpdatedAt, filehistory.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FileHistory fields.
func (fh *FileHistory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case filehistory.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				fh.ID = value.String
			}
		case filehistory.FieldHistoryTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field history_time", values[i])
			} else if value.Valid {
				fh.HistoryTime = value.Time
			}
		case filehistory.FieldOperation:
			if value, ok := values[i].(*enthistory.OpType); !ok {
				return fmt.Errorf("unexpected type %T for field operation", values[i])
			} else if value != nil {
				fh.Operation = *value
			}
		case filehistory.FieldRef:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ref", values[i])
			} else if value.Valid {
				fh.Ref = value.String
			}
		case filehistory.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				fh.CreatedAt = value.Time
			}
		case filehistory.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				fh.UpdatedAt = value.Time
			}
		case filehistory.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				fh.CreatedBy = value.String
			}
		case filehistory.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				fh.UpdatedBy = value.String
			}
		case filehistory.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				fh.DeletedAt = value.Time
			}
		case filehistory.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				fh.DeletedBy = value.String
			}
		case filehistory.FieldMappingID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mapping_id", values[i])
			} else if value.Valid {
				fh.MappingID = value.String
			}
		case filehistory.FieldTags:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field tags", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &fh.Tags); err != nil {
					return fmt.Errorf("unmarshal field tags: %w", err)
				}
			}
		case filehistory.FieldFileName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field file_name", values[i])
			} else if value.Valid {
				fh.FileName = value.String
			}
		case filehistory.FieldFileExtension:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field file_extension", values[i])
			} else if value.Valid {
				fh.FileExtension = value.String
			}
		case filehistory.FieldFileSize:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field file_size", values[i])
			} else if value.Valid {
				fh.FileSize = int(value.Int64)
			}
		case filehistory.FieldContentType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content_type", values[i])
			} else if value.Valid {
				fh.ContentType = value.String
			}
		case filehistory.FieldStoreKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field store_key", values[i])
			} else if value.Valid {
				fh.StoreKey = value.String
			}
		case filehistory.FieldCategory:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field category", values[i])
			} else if value.Valid {
				fh.Category = value.String
			}
		case filehistory.FieldAnnotation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field annotation", values[i])
			} else if value.Valid {
				fh.Annotation = value.String
			}
		default:
			fh.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the FileHistory.
// This includes values selected through modifiers, order, etc.
func (fh *FileHistory) Value(name string) (ent.Value, error) {
	return fh.selectValues.Get(name)
}

// Update returns a builder for updating this FileHistory.
// Note that you need to call FileHistory.Unwrap() before calling this method if this FileHistory
// was returned from a transaction, and the transaction was committed or rolled back.
func (fh *FileHistory) Update() *FileHistoryUpdateOne {
	return NewFileHistoryClient(fh.config).UpdateOne(fh)
}

// Unwrap unwraps the FileHistory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (fh *FileHistory) Unwrap() *FileHistory {
	_tx, ok := fh.config.driver.(*txDriver)
	if !ok {
		panic("generated: FileHistory is not a transactional entity")
	}
	fh.config.driver = _tx.drv
	return fh
}

// String implements the fmt.Stringer.
func (fh *FileHistory) String() string {
	var builder strings.Builder
	builder.WriteString("FileHistory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", fh.ID))
	builder.WriteString("history_time=")
	builder.WriteString(fh.HistoryTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("operation=")
	builder.WriteString(fmt.Sprintf("%v", fh.Operation))
	builder.WriteString(", ")
	builder.WriteString("ref=")
	builder.WriteString(fh.Ref)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(fh.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fh.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(fh.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fh.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fh.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_by=")
	builder.WriteString(fh.DeletedBy)
	builder.WriteString(", ")
	builder.WriteString("mapping_id=")
	builder.WriteString(fh.MappingID)
	builder.WriteString(", ")
	builder.WriteString("tags=")
	builder.WriteString(fmt.Sprintf("%v", fh.Tags))
	builder.WriteString(", ")
	builder.WriteString("file_name=")
	builder.WriteString(fh.FileName)
	builder.WriteString(", ")
	builder.WriteString("file_extension=")
	builder.WriteString(fh.FileExtension)
	builder.WriteString(", ")
	builder.WriteString("file_size=")
	builder.WriteString(fmt.Sprintf("%v", fh.FileSize))
	builder.WriteString(", ")
	builder.WriteString("content_type=")
	builder.WriteString(fh.ContentType)
	builder.WriteString(", ")
	builder.WriteString("store_key=")
	builder.WriteString(fh.StoreKey)
	builder.WriteString(", ")
	builder.WriteString("category=")
	builder.WriteString(fh.Category)
	builder.WriteString(", ")
	builder.WriteString("annotation=")
	builder.WriteString(fh.Annotation)
	builder.WriteByte(')')
	return builder.String()
}

// FileHistories is a parsable slice of FileHistory.
type FileHistories []*FileHistory
