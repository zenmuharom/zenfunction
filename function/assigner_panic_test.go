package function

import (
	"database/sql"
	"testing"

	"github.com/zenmuharom/zenfunction/domain"
	"github.com/zenmuharom/zenlogger"
)

// TestValidityConditionCore_TypeAssertionPanic tests the type assertion panic at line 135
func TestValidityConditionCore_TypeAssertionPanic(t *testing.T) {
	logger := zenlogger.NewZenlogger()
	assigner := NewAssigner(logger).(*DefaultAssigner)

	t.Run("Equals with integer valueToCompare - WILL PANIC", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Logf("✓ PANIC DETECTED (line 135): %v", r)
				t.Log("BUG CONFIRMED: Type assertion valueToCompare.(string) panics on non-string types")
			} else {
				t.Error("Expected panic but got none - bug may be fixed")
			}
		}()

		// This will panic with: interface conversion: interface {} is int, not string
		assigner.validityConditionCore("=", 123, "test")
	})

	t.Run("Equals with slice valueToCompare - WILL PANIC", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Logf("✓ PANIC DETECTED: %v", r)
			}
		}()

		assigner.validityConditionCore("=", []string{"test"}, "value")
	})

	t.Run("NotEquals with integer valueToCompare - WILL PANIC", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Logf("✓ PANIC DETECTED (line 279): %v", r)
				t.Log("BUG CONFIRMED: Same unsafe type assertion in != operator")
			}
		}()

		assigner.validityConditionCore("!=", 456, "test")
	})
}

// TestValidityConditionCore_ReflectPanic tests the reflect.ValueOf panic at line 149
func TestValidityConditionCore_ReflectPanic(t *testing.T) {
	logger := zenlogger.NewZenlogger()
	assigner := NewAssigner(logger).(*DefaultAssigner)

	t.Run("LessThan with nil valueToCompare - WILL PANIC", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Logf("✓ PANIC DETECTED (line 149): %v", r)
				t.Log("BUG CONFIRMED: reflect.ValueOf(nil) is called BEFORE the nil check at line 153")
			}
		}()

		// reflect.ValueOf is at line 149, but nil check is at line 153!
		assigner.validityConditionCore("<", nil, 10)
	})

	t.Run("GreaterThan with nil valueToCompare - WILL PANIC", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Logf("✓ PANIC DETECTED: %v", r)
			}
		}()

		assigner.validityConditionCore(">", nil, 10)
	})

	t.Run("LessThanEquals with nil valueToCompare - WILL PANIC", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Logf("✓ PANIC DETECTED: %v", r)
			}
		}()

		assigner.validityConditionCore("<=", nil, 10)
	})

	t.Run("GreaterThanEquals with nil valueToCompare - WILL PANIC", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Logf("✓ PANIC DETECTED: %v", r)
			}
		}()

		assigner.validityConditionCore(">=", nil, 10)
	})
}

// TestAssignValue_BooleanPanic tests the boolean type assertion panic at line 405
func TestAssignValue_BooleanPanic(t *testing.T) {
	logger := zenlogger.NewZenlogger()
	assigner := NewAssigner(logger).(*DefaultAssigner)

	t.Run("Boolean with string value - WILL PANIC", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Logf("✓ PANIC DETECTED (line 405): %v", r)
				t.Log("BUG CONFIRMED: valueToAssign.Value.(bool) panics on non-bool types")
			}
		}()

		parent := domain.AssignVariableValue{
			Key:     "flag",
			Value:   false,
			VarType: "boolean",
		}
		valueToAssign := domain.AssignVariableValue{
			Key:     "flag",
			Value:   "true", // String, not boolean
			VarType: "boolean",
		}

		assigner.AssignValue(parent, valueToAssign)
	})

	t.Run("Boolean with nil value - WILL PANIC", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Logf("✓ PANIC DETECTED: %v", r)
			}
		}()

		parent := domain.AssignVariableValue{
			Key:     "flag",
			Value:   false,
			VarType: "boolean",
		}
		valueToAssign := domain.AssignVariableValue{
			Key:     "flag",
			Value:   nil,
			VarType: "boolean",
		}

		assigner.AssignValue(parent, valueToAssign)
	})
}

// TestValidityConditionField_Integration tests the full ValidityConditionField flow
func TestValidityConditionField_Integration(t *testing.T) {
	logger := zenlogger.NewZenlogger()
	assigner := NewAssigner(logger).(*DefaultAssigner)

	t.Run("ValidityConditionField with non-string value - WILL PANIC", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Logf("✓ PANIC DETECTED: %v", r)
				t.Log("BUG CONFIRMED: ValidityConditionField calls validityConditionCore which panics")
			}
		}()

		config := domain.ValueConfig{
			ConditionFieldId:  sql.NullInt64{Int64: 1, Valid: true},
			ConditionOperator: sql.NullString{String: "=", Valid: true},
			ConditionValue:    sql.NullString{String: "test", Valid: true},
			FieldName:         sql.NullString{String: "testField", Valid: true},
		}

		// Pass a complex type - will trigger the panic in validityConditionCore
		assigner.ValidityConditionField(config, []int{1, 2, 3})
	})
}
