// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type alertRuleTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *alertRuleTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("alert_rules").
func (v *alertRuleTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *alertRuleTableType) Columns() []string {
	return []string{"template", "id", "summary", "disabled", "params", "for", "severity", "custom_labels", "filters", "channels", "created_at", "updated_at"}
}

// NewStruct makes a new struct for that view or table.
func (v *alertRuleTableType) NewStruct() reform.Struct {
	return new(alertRule)
}

// NewRecord makes a new record for that table.
func (v *alertRuleTableType) NewRecord() reform.Record {
	return new(alertRule)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *alertRuleTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// alertRuleTable represents alert_rules view or table in SQL database.
var alertRuleTable = &alertRuleTableType{
	s: parse.StructInfo{Type: "alertRule", SQLSchema: "", SQLName: "alert_rules", Fields: []parse.FieldInfo{{Name: "Template", Type: "*[]uint8", Column: "template"}, {Name: "ID", Type: "string", Column: "id"}, {Name: "Summary", Type: "string", Column: "summary"}, {Name: "Disabled", Type: "bool", Column: "disabled"}, {Name: "Params", Type: "*[]uint8", Column: "params"}, {Name: "For", Type: "string", Column: "for"}, {Name: "Severity", Type: "string", Column: "severity"}, {Name: "CustomLabels", Type: "*[]uint8", Column: "custom_labels"}, {Name: "Filters", Type: "*[]uint8", Column: "filters"}, {Name: "Channels", Type: "*[]uint8", Column: "channels"}, {Name: "CreatedAt", Type: "time.Time", Column: "created_at"}, {Name: "UpdatedAt", Type: "time.Time", Column: "updated_at"}}, PKFieldIndex: 1},
	z: new(alertRule).Values(),
}

// String returns a string representation of this struct or record.
func (s alertRule) String() string {
	res := make([]string, 12)
	res[0] = "Template: " + reform.Inspect(s.Template, true)
	res[1] = "ID: " + reform.Inspect(s.ID, true)
	res[2] = "Summary: " + reform.Inspect(s.Summary, true)
	res[3] = "Disabled: " + reform.Inspect(s.Disabled, true)
	res[4] = "Params: " + reform.Inspect(s.Params, true)
	res[5] = "For: " + reform.Inspect(s.For, true)
	res[6] = "Severity: " + reform.Inspect(s.Severity, true)
	res[7] = "CustomLabels: " + reform.Inspect(s.CustomLabels, true)
	res[8] = "Filters: " + reform.Inspect(s.Filters, true)
	res[9] = "Channels: " + reform.Inspect(s.Channels, true)
	res[10] = "CreatedAt: " + reform.Inspect(s.CreatedAt, true)
	res[11] = "UpdatedAt: " + reform.Inspect(s.UpdatedAt, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *alertRule) Values() []interface{} {
	return []interface{}{
		s.Template,
		s.ID,
		s.Summary,
		s.Disabled,
		s.Params,
		s.For,
		s.Severity,
		s.CustomLabels,
		s.Filters,
		s.Channels,
		s.CreatedAt,
		s.UpdatedAt,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *alertRule) Pointers() []interface{} {
	return []interface{}{
		&s.Template,
		&s.ID,
		&s.Summary,
		&s.Disabled,
		&s.Params,
		&s.For,
		&s.Severity,
		&s.CustomLabels,
		&s.Filters,
		&s.Channels,
		&s.CreatedAt,
		&s.UpdatedAt,
	}
}

// View returns View object for that struct.
func (s *alertRule) View() reform.View {
	return alertRuleTable
}

// Table returns Table object for that record.
func (s *alertRule) Table() reform.Table {
	return alertRuleTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *alertRule) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *alertRule) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *alertRule) HasPK() bool {
	return s.ID != alertRuleTable.z[alertRuleTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *alertRule) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.ID = string(i64)
	} else {
		s.ID = pk.(string)
	}
}

// check interfaces
var (
	_ reform.View   = alertRuleTable
	_ reform.Struct = (*alertRule)(nil)
	_ reform.Table  = alertRuleTable
	_ reform.Record = (*alertRule)(nil)
	_ fmt.Stringer  = (*alertRule)(nil)
)

func init() {
	parse.AssertUpToDate(&alertRuleTable.s, new(alertRule))
}