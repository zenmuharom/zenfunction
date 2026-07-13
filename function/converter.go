package function

import (
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
)

func isArgumentOrganic(value interface{}) bool {
	switch value.(type) {
	case int:
		return true
	case int64:
		return true
	case float32:
		return true
	case float64:
		return true
	default:
		return false
	}
}

func convertToInt64(value interface{}) (int64, error) {
	if value == nil {
		return 0, fmt.Errorf("unsupported type: <nil>")
	}

	switch v := value.(type) {
	case string:
		trimmed := strings.TrimSpace(v)
		if trimmed == "" {
			return 0, fmt.Errorf("unsupported value: empty string")
		}
		if parsedInt, err := strconv.ParseInt(trimmed, 10, 64); err == nil {
			return parsedInt, nil
		}
		parsedFloat, err := strconv.ParseFloat(trimmed, 64)
		if err != nil || math.IsNaN(parsedFloat) || math.IsInf(parsedFloat, 0) {
			return 0, fmt.Errorf("cannot convert %q to int64", trimmed)
		}
		if math.Trunc(parsedFloat) != parsedFloat {
			return 0, fmt.Errorf("cannot convert non-integer %q to int64", trimmed)
		}
		if parsedFloat > math.MaxInt64 || parsedFloat < math.MinInt64 {
			return 0, fmt.Errorf("value out of int64 range: %q", trimmed)
		}
		return int64(parsedFloat), nil
	case json.Number:
		if parsedInt, err := v.Int64(); err == nil {
			return parsedInt, nil
		}
		parsedFloat, err := v.Float64()
		if err != nil || math.IsNaN(parsedFloat) || math.IsInf(parsedFloat, 0) {
			return 0, fmt.Errorf("cannot convert %q to int64", v.String())
		}
		if math.Trunc(parsedFloat) != parsedFloat {
			return 0, fmt.Errorf("cannot convert non-integer %q to int64", v.String())
		}
		if parsedFloat > math.MaxInt64 || parsedFloat < math.MinInt64 {
			return 0, fmt.Errorf("value out of int64 range: %q", v.String())
		}
		return int64(parsedFloat), nil
	case int:
		return int64(v), nil
	case int8:
		return int64(v), nil
	case int16:
		return int64(v), nil
	case int32:
		return int64(v), nil
	case int64:
		return v, nil
	case uint:
		if uint64(v) > math.MaxInt64 {
			return 0, fmt.Errorf("value out of int64 range")
		}
		return int64(v), nil
	case uint8:
		return int64(v), nil
	case uint16:
		return int64(v), nil
	case uint32:
		return int64(v), nil
	case uint64:
		if v > math.MaxInt64 {
			return 0, fmt.Errorf("value out of int64 range")
		}
		return int64(v), nil
	case float32:
		floatValue := float64(v)
		if math.Trunc(floatValue) != floatValue {
			return 0, fmt.Errorf("cannot convert non-integer %v to int64", v)
		}
		if floatValue > math.MaxInt64 || floatValue < math.MinInt64 {
			return 0, fmt.Errorf("value out of int64 range")
		}
		return int64(floatValue), nil
	case float64:
		if math.IsNaN(v) || math.IsInf(v, 0) {
			return 0, fmt.Errorf("cannot convert %v to int64", v)
		}
		if math.Trunc(v) != v {
			return 0, fmt.Errorf("cannot convert non-integer %v to int64", v)
		}
		if v > math.MaxInt64 || v < math.MinInt64 {
			return 0, fmt.Errorf("value out of int64 range")
		}
		return int64(v), nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, fmt.Errorf("unsupported type: %v", reflect.TypeOf(value))
	}
}

func convertToFloat64(value interface{}) (float64, error) {
	if value == nil {
		return 0, fmt.Errorf("unsupported type: <nil>")
	}

	switch v := value.(type) {
	case float64:
		if math.IsNaN(v) || math.IsInf(v, 0) {
			return 0, fmt.Errorf("cannot convert %v to float64", v)
		}
		return v, nil
	case float32:
		floatValue := float64(v)
		if math.IsNaN(floatValue) || math.IsInf(floatValue, 0) {
			return 0, fmt.Errorf("cannot convert %v to float64", v)
		}
		return floatValue, nil
	case int:
		return float64(v), nil
	case int8:
		return float64(v), nil
	case int16:
		return float64(v), nil
	case int32:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case uint:
		return float64(v), nil
	case uint8:
		return float64(v), nil
	case uint16:
		return float64(v), nil
	case uint32:
		return float64(v), nil
	case uint64:
		return float64(v), nil
	case json.Number:
		parsed, err := v.Float64()
		if err != nil || math.IsNaN(parsed) || math.IsInf(parsed, 0) {
			return 0, fmt.Errorf("cannot convert %q to float64", v.String())
		}
		return parsed, nil
	case string:
		trimmed := strings.TrimSpace(v)
		if trimmed == "" {
			return 0, fmt.Errorf("unsupported value: empty string")
		}
		parsed, err := strconv.ParseFloat(trimmed, 64)
		if err != nil || math.IsNaN(parsed) || math.IsInf(parsed, 0) {
			return 0, fmt.Errorf("cannot convert %q to float64", trimmed)
		}
		return parsed, nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, fmt.Errorf("unsupported type: %v", reflect.TypeOf(value))
	}
}

func convertToBool(value interface{}) (bool, error) {
	if value == nil {
		return false, fmt.Errorf("unsupported type: <nil>")
	}

	switch v := value.(type) {
	case bool:
		return v, nil
	case string:
		trimmed := strings.TrimSpace(v)
		if trimmed == "" {
			return false, fmt.Errorf("unsupported value: empty string")
		}
		parsed, err := strconv.ParseBool(trimmed)
		if err == nil {
			return parsed, nil
		}
		if intValue, intErr := strconv.ParseInt(trimmed, 10, 64); intErr == nil {
			return intValue != 0, nil
		}
		return false, fmt.Errorf("cannot convert %q to bool", trimmed)
	case int, int8, int16, int32, int64:
		intValue, err := convertToInt64(v)
		if err != nil {
			return false, err
		}
		return intValue != 0, nil
	case uint, uint8, uint16, uint32, uint64:
		floatValue, err := convertToFloat64(v)
		if err != nil {
			return false, err
		}
		return floatValue != 0, nil
	case float32, float64:
		floatValue, err := convertToFloat64(v)
		if err != nil {
			return false, err
		}
		return floatValue != 0, nil
	case json.Number:
		floatValue, err := convertToFloat64(v)
		if err != nil {
			return false, err
		}
		return floatValue != 0, nil
	default:
		return false, fmt.Errorf("unsupported type: %v", reflect.TypeOf(value))
	}
}

func convertToString(value interface{}) string {
	if value == nil {
		return ""
	}

	switch v := value.(type) {
	case int:
		return strconv.Itoa(v)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		if math.IsNaN(v) || math.IsInf(v, 0) {
			return fmt.Sprintf("%v", v)
		}
		return strconv.FormatFloat(v, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(v)
	case json.Number:
		return v.String()
	case []byte:
		return string(v)
	case string:
		return v
	default:
		return fmt.Sprintf("%v", v)
	}
}
