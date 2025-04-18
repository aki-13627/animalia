// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/aki-13627/animalia/backend-go/ent/enum"
	"github.com/aki-13627/animalia/backend-go/ent/tasktype"
	pgvector "github.com/pgvector/pgvector-go"
)

// TaskType is the model entity for the TaskType schema.
type TaskType struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Type holds the value of the "type" field.
	Type enum.TaskType `json:"type,omitempty"`
	// TextFeature holds the value of the "text_feature" field.
	TextFeature  pgvector.Vector `json:"text_feature,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TaskType) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case tasktype.FieldTextFeature:
			values[i] = new(pgvector.Vector)
		case tasktype.FieldID:
			values[i] = new(sql.NullInt64)
		case tasktype.FieldType:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TaskType fields.
func (tt *TaskType) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case tasktype.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			tt.ID = int(value.Int64)
		case tasktype.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				tt.Type = enum.TaskType(value.String)
			}
		case tasktype.FieldTextFeature:
			if value, ok := values[i].(*pgvector.Vector); !ok {
				return fmt.Errorf("unexpected type %T for field text_feature", values[i])
			} else if value != nil {
				tt.TextFeature = *value
			}
		default:
			tt.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the TaskType.
// This includes values selected through modifiers, order, etc.
func (tt *TaskType) Value(name string) (ent.Value, error) {
	return tt.selectValues.Get(name)
}

// Update returns a builder for updating this TaskType.
// Note that you need to call TaskType.Unwrap() before calling this method if this TaskType
// was returned from a transaction, and the transaction was committed or rolled back.
func (tt *TaskType) Update() *TaskTypeUpdateOne {
	return NewTaskTypeClient(tt.config).UpdateOne(tt)
}

// Unwrap unwraps the TaskType entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tt *TaskType) Unwrap() *TaskType {
	_tx, ok := tt.config.driver.(*txDriver)
	if !ok {
		panic("ent: TaskType is not a transactional entity")
	}
	tt.config.driver = _tx.drv
	return tt
}

// String implements the fmt.Stringer.
func (tt *TaskType) String() string {
	var builder strings.Builder
	builder.WriteString("TaskType(")
	builder.WriteString(fmt.Sprintf("id=%v, ", tt.ID))
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", tt.Type))
	builder.WriteString(", ")
	builder.WriteString("text_feature=")
	builder.WriteString(fmt.Sprintf("%v", tt.TextFeature))
	builder.WriteByte(')')
	return builder.String()
}

// TaskTypes is a parsable slice of TaskType.
type TaskTypes []*TaskType
