package zenfunction

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetAvailableFunctions(t *testing.T) {
	functions := GetAvailableFunctions()

	t.Run("Should return functions", func(t *testing.T) {
		assert.Greater(t, len(functions), 0, "Should have at least one function")
	})

	t.Run("All functions should have required fields", func(t *testing.T) {
		for _, fn := range functions {
			assert.NotEmpty(t, fn.Name, "Function name should not be empty")
			assert.NotEmpty(t, fn.Description, "Function description should not be empty")
			assert.NotEmpty(t, fn.Category, "Function category should not be empty")
			assert.NotEmpty(t, fn.ReturnType, "Function return type should not be empty")
		}
	})

	t.Run("All parameters should have required fields", func(t *testing.T) {
		for _, fn := range functions {
			for _, param := range fn.Parameters {
				assert.NotEmpty(t, param.Name, "Parameter name should not be empty for function: "+fn.Name)
				assert.NotEmpty(t, param.Type, "Parameter type should not be empty for function: "+fn.Name)
				assert.NotEmpty(t, param.Description, "Parameter description should not be empty for function: "+fn.Name)
			}
		}
	})

	t.Run("Should have examples", func(t *testing.T) {
		for _, fn := range functions {
			assert.Greater(t, len(fn.Examples), 0, "Function should have at least one example: "+fn.Name)
		}
	})
}

func TestGetFunctionInfo(t *testing.T) {
	t.Run("Should return function info for valid function", func(t *testing.T) {
		info := GetFunctionInfo("md5")
		require.NotNil(t, info)
		assert.Equal(t, "md5", info.Name)
		assert.Equal(t, "crypto", info.Category)
		assert.Equal(t, TypeString, info.ReturnType)
		assert.Greater(t, len(info.Parameters), 0)
	})

	t.Run("Should return nil for invalid function", func(t *testing.T) {
		info := GetFunctionInfo("nonexistent")
		assert.Nil(t, info)
	})

	t.Run("Should work for all registered functions", func(t *testing.T) {
		allFunctions := GetAvailableFunctions()
		for _, fn := range allFunctions {
			info := GetFunctionInfo(fn.Name)
			assert.NotNil(t, info, "Should find function: "+fn.Name)
			assert.Equal(t, fn.Name, info.Name)
		}
	})
}

func TestGetFunctionsByCategory(t *testing.T) {
	t.Run("Should return functions for string category", func(t *testing.T) {
		functions := GetFunctionsByCategory("string")
		assert.Greater(t, len(functions), 0)
		for _, fn := range functions {
			assert.Equal(t, "string", fn.Category)
		}
	})

	t.Run("Should return functions for crypto category", func(t *testing.T) {
		functions := GetFunctionsByCategory("crypto")
		assert.Greater(t, len(functions), 0)
		for _, fn := range functions {
			assert.Equal(t, "crypto", fn.Category)
		}
	})

	t.Run("Should return empty for invalid category", func(t *testing.T) {
		functions := GetFunctionsByCategory("nonexistent")
		assert.Equal(t, 0, len(functions))
	})
}

func TestGetCategories(t *testing.T) {
	categories := GetCategories()

	t.Run("Should return all categories", func(t *testing.T) {
		assert.Greater(t, len(categories), 0)

		// Check for expected categories
		categoryMap := make(map[string]bool)
		for _, cat := range categories {
			categoryMap[cat] = true
		}

		assert.True(t, categoryMap["string"], "Should have string category")
		assert.True(t, categoryMap["crypto"], "Should have crypto category")
		assert.True(t, categoryMap["date"], "Should have date category")
		assert.True(t, categoryMap["utility"], "Should have utility category")
		assert.True(t, categoryMap["json"], "Should have json category")
	})
}

func TestGetFunctionNames(t *testing.T) {
	names := GetFunctionNames()

	t.Run("Should return all function names", func(t *testing.T) {
		assert.Greater(t, len(names), 0)

		// Check for some expected functions
		expectedFunctions := []string{"md5", "concat", "dateNow", "trim", "substr"}
		for _, expected := range expectedFunctions {
			assert.Contains(t, names, expected, "Should contain function: "+expected)
		}
	})
}

func TestParameterTypes(t *testing.T) {
	t.Run("Parameter type constants should be defined", func(t *testing.T) {
		assert.Equal(t, ParameterType("string"), TypeString)
		assert.Equal(t, ParameterType("int"), TypeInt)
		assert.Equal(t, ParameterType("float"), TypeFloat)
		assert.Equal(t, ParameterType("bool"), TypeBool)
		assert.Equal(t, ParameterType("any"), TypeAny)
		assert.Equal(t, ParameterType("array"), TypeArray)
		assert.Equal(t, ParameterType("object"), TypeObject)
	})
}

func TestJSONSerialization(t *testing.T) {
	t.Run("FunctionInfo should be JSON serializable", func(t *testing.T) {
		info := GetFunctionInfo("md5")
		require.NotNil(t, info)

		jsonData, err := json.Marshal(info)
		require.NoError(t, err)
		assert.NotEmpty(t, jsonData)

		var decoded FunctionInfo
		err = json.Unmarshal(jsonData, &decoded)
		require.NoError(t, err)
		assert.Equal(t, info.Name, decoded.Name)
		assert.Equal(t, info.Description, decoded.Description)
		assert.Equal(t, info.Category, decoded.Category)
	})

	t.Run("All functions should be JSON serializable", func(t *testing.T) {
		functions := GetAvailableFunctions()
		jsonData, err := json.Marshal(functions)
		require.NoError(t, err)
		assert.NotEmpty(t, jsonData)

		var decoded []FunctionInfo
		err = json.Unmarshal(jsonData, &decoded)
		require.NoError(t, err)
		assert.Equal(t, len(functions), len(decoded))
	})
}

func TestSpecificFunctions(t *testing.T) {
	t.Run("md5 function should have correct metadata", func(t *testing.T) {
		info := GetFunctionInfo("md5")
		require.NotNil(t, info)
		assert.Equal(t, "md5", info.Name)
		assert.Equal(t, "crypto", info.Category)
		assert.Equal(t, TypeString, info.ReturnType)
		assert.Equal(t, 1, len(info.Parameters))
		assert.Equal(t, "text", info.Parameters[0].Name)
		assert.True(t, info.Parameters[0].Required)
	})

	t.Run("concat function should support variadic parameters", func(t *testing.T) {
		info := GetFunctionInfo("concat")
		require.NotNil(t, info)
		assert.Greater(t, len(info.Parameters), 2, "Should have multiple parameters including variadic")

		// Check for variadic parameter
		hasVariadic := false
		for _, param := range info.Parameters {
			if param.Name == "..." {
				hasVariadic = true
				assert.False(t, param.Required, "Variadic parameter should be optional")
			}
		}
		assert.True(t, hasVariadic, "concat should have variadic parameter")
	})

	t.Run("dateAdd function should have 4 parameters", func(t *testing.T) {
		info := GetFunctionInfo("dateAdd")
		require.NotNil(t, info)
		assert.Equal(t, 4, len(info.Parameters))
		assert.Equal(t, "format", info.Parameters[0].Name)
		assert.Equal(t, "date", info.Parameters[1].Name)
		assert.Equal(t, "amount", info.Parameters[2].Name)
		assert.Equal(t, "unit", info.Parameters[3].Name)
	})

	t.Run("uuid function should have no parameters", func(t *testing.T) {
		info := GetFunctionInfo("uuid")
		require.NotNil(t, info)
		assert.Equal(t, 0, len(info.Parameters))
	})
}
