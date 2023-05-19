package domain

import "database/sql"

type ValueConfig struct {
	FieldName          sql.NullString `db:"field_name"`
	ConditionFieldId   sql.NullInt64  `db:"condition_field_id"`
	ConditionFieldName sql.NullString `db:"condition_field_name"`
	ConditionOperator  sql.NullString `db:"condition_operator"`
	ConditionValue     sql.NullString `db:"condition_value"`
}

type AssignVariableValue struct {
	Key     string
	VarType string
	Value   interface{}
}
