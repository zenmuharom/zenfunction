package function

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zenmuharom/zenfunction/domain"
	"github.com/zenmuharom/zenlogger"
)

// TestValidityConditionCore_NoMorePanic tests that the fixes prevent panics
func TestValidityConditionCore_NoMorePanic(t *testing.T) {
	logger := zenlogger.NewZenlogger()
	assigner := NewAssigner(logger).(*DefaultAssigner)

	t.Run("Equals with integer valueToCompare - NO PANIC", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("❌ UNEXPECTED PANIC: %v - Fix did not work!", r)
			}
		}()

		// This should NOT panic anymore
		functionGenerated, valid := assigner.validityConditionCore("=", 123, "test")
		t.Logf("✓ NO PANIC - Function executed successfully")
		t.Logf("Result: %s, valid: %v", functionGenerated, valid)
		assert.False(t, valid)
		assert.Contains(t, functionGenerated, "123")
	})

	t.Run("Equals with slice valueToCompare - NO PANIC", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("❌ UNEXPECTED PANIC: %v", r)
			}
		}()

		functionGenerated, valid := assigner.validityConditionCore("=", []string{"test"}, "value")
		t.Logf("✓ NO PANIC - Function executed successfully")
		t.Logf("Result: %s, valid: %v", functionGenerated, valid)
		assert.False(t, valid)
	})

	t.Run("Equals with map valueToCompare - NO PANIC", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("❌ UNEXPECTED PANIC: %v", r)
			}
		}()

		functionGenerated, valid := assigner.validityConditionCore("=", map[string]string{"key": "value"}, "test")
		t.Logf("✓ NO PANIC - Function executed successfully")
		t.Logf("Result: %s, valid: %v", functionGenerated, valid)
		assert.False(t, valid)
	})

	t.Run("NotEquals with integer valueToCompare - NO PANIC", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("❌ UNEXPECTED PANIC: %v", r)
			}
		}()

		functionGenerated, valid := assigner.validityConditionCore("!=", 456, "test")
		t.Logf("✓ NO PANIC - Function executed successfully")
		t.Logf("Result: %s, valid: %v", functionGenerated, valid)
		assert.True(t, valid) // 456 != "test"
		assert.Contains(t, functionGenerated, "456")
	})

	t.Run("NotEquals with float valueToCompare - NO PANIC", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("❌ UNEXPECTED PANIC: %v", r)
			}
		}()

		functionGenerated, valid := assigner.validityConditionCore("!=", 3.14, "test")
		t.Logf("✓ NO PANIC - Function executed successfully")
		t.Logf("Result: %s, valid: %v", functionGenerated, valid)
		assert.True(t, valid)
	})
}

// TestAssignValue_BooleanNoMorePanic tests that the boolean fix prevents panics
func TestAssignValue_BooleanNoMorePanic(t *testing.T) {
	logger := zenlogger.NewZenlogger()
	assigner := NewAssigner(logger).(*DefaultAssigner)

	t.Run("Boolean with string value - NO PANIC", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("❌ UNEXPECTED PANIC: %v - Fix did not work!", r)
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

		result := assigner.AssignValue(parent, valueToAssign)
		t.Logf("✓ NO PANIC - Function executed successfully")
		t.Logf("Result: %v (type: %T)", result, result)

		// Should default to false when type mismatch
		assert.Equal(t, false, result)
	})

	t.Run("Boolean with integer value - NO PANIC", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("❌ UNEXPECTED PANIC: %v", r)
			}
		}()

		parent := domain.AssignVariableValue{
			Key:     "flag",
			Value:   true,
			VarType: "boolean",
		}
		valueToAssign := domain.AssignVariableValue{
			Key:     "flag",
			Value:   1, // Integer, not boolean
			VarType: "boolean",
		}

		result := assigner.AssignValue(parent, valueToAssign)
		t.Logf("✓ NO PANIC - Function executed successfully")
		t.Logf("Result: %v", result)
		assert.Equal(t, false, result) // Should default to false
	})

	t.Run("Boolean with nil value - NO PANIC", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("❌ UNEXPECTED PANIC: %v", r)
			}
		}()

		parent := domain.AssignVariableValue{
			Key:     "flag",
			Value:   true,
			VarType: "boolean",
		}
		valueToAssign := domain.AssignVariableValue{
			Key:     "flag",
			Value:   nil,
			VarType: "boolean",
		}

		result := assigner.AssignValue(parent, valueToAssign)
		t.Logf("✓ NO PANIC - Function executed successfully")
		t.Logf("Result: %v", result)
		assert.Equal(t, false, result) // Should default to false
	})

	t.Run("Boolean with actual bool value - WORKS CORRECTLY", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("❌ UNEXPECTED PANIC: %v", r)
			}
		}()

		parent := domain.AssignVariableValue{
			Key:     "flag",
			Value:   false,
			VarType: "boolean",
		}
		valueToAssign := domain.AssignVariableValue{
			Key:     "flag",
			Value:   true, // Actual boolean
			VarType: "boolean",
		}

		result := assigner.AssignValue(parent, valueToAssign)
		t.Logf("✓ NO PANIC - Function executed successfully")
		t.Logf("Result: %v", result)
		assert.Equal(t, true, result) // Should work correctly
	})
}

// TestValidityConditionField_IntegrationNoMorePanic tests the full flow without panics
func TestValidityConditionField_IntegrationNoMorePanic(t *testing.T) {
	logger := zenlogger.NewZenlogger()
	assigner := NewAssigner(logger).(*DefaultAssigner)

	t.Run("ValidityConditionField with integer array - NO PANIC", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("❌ UNEXPECTED PANIC: %v - Fix did not work!", r)
			}
		}()

		config := domain.ValueConfig{
			ConditionFieldId:  sql.NullInt64{Int64: 1, Valid: true},
			ConditionOperator: sql.NullString{String: "=", Valid: true},
			ConditionValue:    sql.NullString{String: "test", Valid: true},
			FieldName:         sql.NullString{String: "testField", Valid: true},
		}

		// This should NOT panic anymore
		valid := assigner.ValidityConditionField(config, []int{1, 2, 3})
		t.Logf("✓ NO PANIC - Function executed successfully")
		t.Logf("Valid: %v", valid)
		assert.False(t, valid)
	})

	t.Run("ValidityConditionField with struct - NO PANIC", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("❌ UNEXPECTED PANIC: %v", r)
			}
		}()

		config := domain.ValueConfig{
			ConditionFieldId:  sql.NullInt64{Int64: 1, Valid: true},
			ConditionOperator: sql.NullString{String: "=", Valid: true},
			ConditionValue:    sql.NullString{String: "test", Valid: true},
			FieldName:         sql.NullString{String: "testField", Valid: true},
		}

		type TestStruct struct {
			Name string
		}

		valid := assigner.ValidityConditionField(config, TestStruct{Name: "test"})
		t.Logf("✓ NO PANIC - Function executed successfully")
		t.Logf("Valid: %v", valid)
	})

	t.Run("ValidityConditionField with nil - NO PANIC", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("❌ UNEXPECTED PANIC: %v", r)
			}
		}()

		config := domain.ValueConfig{
			ConditionFieldId:  sql.NullInt64{Int64: 1, Valid: true},
			ConditionOperator: sql.NullString{String: "=", Valid: true},
			ConditionValue:    sql.NullString{String: "test", Valid: true},
			FieldName:         sql.NullString{String: "testField", Valid: true},
		}

		valid := assigner.ValidityConditionField(config, nil)
		t.Logf("✓ NO PANIC - Function executed successfully")
		t.Logf("Valid: %v", valid)
		assert.False(t, valid)
	})
}

// TestValidityConditionValue_NoMorePanic tests ValidityConditionValue without panics
func TestValidityConditionValue_NoMorePanic(t *testing.T) {
	logger := zenlogger.NewZenlogger()
	assigner := NewAssigner(logger).(*DefaultAssigner)

	t.Run("ValidityConditionValue with various types - NO PANIC", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("❌ UNEXPECTED PANIC: %v", r)
			}
		}()

		config := domain.ValueConfig{
			ConditionFieldId:  sql.NullInt64{Int64: 1, Valid: true},
			ConditionOperator: sql.NullString{String: "!=", Valid: true},
			ConditionValue:    sql.NullString{String: "test", Valid: true},
			FieldName:         sql.NullString{String: "testField", Valid: true},
		}

		// Test with different types
		testCases := []interface{}{
			123,
			3.14,
			[]string{"a", "b"},
			map[string]int{"count": 5},
			nil,
		}

		for i, testCase := range testCases {
			valid := assigner.ValidityConditionValue(config, testCase)
			t.Logf("✓ Test case %d - NO PANIC with type %T: valid=%v", i+1, testCase, valid)
		}
	})
}
