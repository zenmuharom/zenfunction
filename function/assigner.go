package function

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/zenmuharom/zenfunction/domain"
	"github.com/zenmuharom/zenfunction/variable"
	"github.com/zenmuharom/zenlogger"
)

type DefaultAssigner struct {
	Logger zenlogger.Zenlogger
}

type Assigner interface {
	ReadCommand(arg string) (returnVal interface{}, err error)
	ReadCommandV2(dType, arg string) (returnVal any, err error)
	findArg(command string) (argument string)
}

func NewAssigner(logger zenlogger.Zenlogger) Assigner {
	return &DefaultAssigner{Logger: logger}
}

func (assigner *DefaultAssigner) ReadCommand(str string) (arg interface{}, err error) {
	arg, err = assigner.coreReadCommand(str)
	if err != nil {
		assigner.Logger.Error(err.Error())
	}
	return
}

func (assigner *DefaultAssigner) ReadCommandV2(dType, str string) (result any, err error) {
	res, err_ := assigner.coreReadCommand(str)
	if err_ != nil {
		assigner.Logger.Error(err_.Error())
		err = err_
		return
	} else {

		switch dType {
		case variable.TYPE_STRING:

			switch v := res.(type) {
			case string:
				unquoted, errUnquote := strconv.Unquote(v)
				if errUnquote != nil {
					if strings.HasPrefix(v, `"`) && strings.HasSuffix(v, `"`) && len(v) >= 2 {
						// safe unwrap only outer quotes
						v = v[1 : len(v)-1]
					}
					result = v
				} else {
					result = unquoted
				}

			default:
				// for numbers, arrays, objects: convert to string (optional, sesuai kebutuhan)
				result = fmt.Sprintf("%v", v)
			}
		default:
			switch v := res.(type) {
			case string:

				if strings.HasPrefix(v, `"`) && strings.HasSuffix(v, `"`) && len(v) >= 2 {
					// safe unwrap only outer quotes
					v = v[1 : len(v)-1]
				}
				result = v
			default:
				// for numbers, arrays, objects: convert to string (optional, sesuai kebutuhan)
				result = fmt.Sprintf("%v", v)
			}
		}

	}
	return
}

func (assigner *DefaultAssigner) ValidityConditionField(config domain.ValueConfig, valueToCompare interface{}) (valid bool) {

	assigner.Logger.Debug("ValidityConditionField", zenlogger.ZenField{Key: "config", Value: config}, zenlogger.ZenField{Key: "valueToCompare", Value: valueToCompare})
	functionGenerated := ""

	if config.ConditionFieldId.Int64 == 0 {
		functionGenerated = "no condition"
		valid = true
	} else {
		functionGenerated, valid = assigner.validityConditionCore(config.ConditionOperator.String, valueToCompare, config.ConditionValue.String)
	}

	assigner.Logger.Debug("ValidityConditionField", zenlogger.ZenField{Key: "fieldName", Value: config.FieldName.String}, zenlogger.ZenField{Key: "functionGenerated", Value: functionGenerated}, zenlogger.ZenField{Key: "valid", Value: valid})

	return
}

func (assigner *DefaultAssigner) ValidityConditionValue(config domain.ValueConfig, valueToCompare interface{}) (valid bool) {
	assigner.Logger.Debug("ValidityCondition", zenlogger.ZenField{Key: "config", Value: config}, zenlogger.ZenField{Key: "valueToCompare", Value: valueToCompare})
	functionGenerated := ""

	if config.ConditionFieldId.Int64 == 0 {
		functionGenerated = "no condition"
		valid = true
	} else {
		functionGenerated, valid = assigner.validityConditionCore(config.ConditionOperator.String, valueToCompare, config.ConditionValue.String)
	}

	assigner.Logger.Debug("ValidityCondition", zenlogger.ZenField{Key: "field", Value: config.FieldName.String}, zenlogger.ZenField{Key: "functionGenerated", Value: functionGenerated}, zenlogger.ZenField{Key: "valueToCompare", Value: valueToCompare}, zenlogger.ZenField{Key: "valid", Value: valid})

	return
}

func (assigner *DefaultAssigner) validityConditionCore(operator string, valueToCompare interface{}, value interface{}) (functionGenerated string, valid bool) {
	switch operator {
	case "=":
		valueToCompareParsed := ""
		if valueToCompare != nil {
			valueToCompareParsed = valueToCompare.(string)
		}
		functionGenerated = fmt.Sprintf("%v = %v", valueToCompareParsed, value)
		valid = fmt.Sprintf("%v", valueToCompareParsed) == fmt.Sprintf("%v", value)
	case "<":
		typeValueToCompare := reflect.ValueOf(valueToCompare)
		valueToCompareInt := 0
		typeValue := reflect.ValueOf(value)
		valueInt := 0
		var e error

		// set to 0 if valueToCompare not setted
		if valueToCompare == nil {
			valueToCompare = 0
		}

		// convert valueToCompare to int
		if typeValueToCompare.Kind() != reflect.Int {
			valueToCompareInt, e = strconv.Atoi(fmt.Sprintf("%v", valueToCompare))
			if e != nil {
				assigner.Logger.Debug(e.Error())
				valueToCompareInt = len(fmt.Sprintf("%v", valueToCompare))
			}
		}

		// set to 0 if value not setted
		if value == nil {
			valueInt = 0
		}

		// convert value to int
		if typeValue.Kind() != reflect.Int {
			valueInt, e = strconv.Atoi(fmt.Sprintf("%v", value))
			if e != nil {
				assigner.Logger.Debug(e.Error())
				valueInt = len(fmt.Sprintf("%v", value))
			}
		}

		functionGenerated = fmt.Sprintf("%v < %v", valueToCompareInt, value)
		valid = valueToCompareInt < valueInt
	case "<=":
		typeValueToCompare := reflect.ValueOf(valueToCompare)
		valueToCompareInt := 0
		typeValue := reflect.ValueOf(value)
		valueInt := 0
		var e error

		// set to 0 if valueToCompare not setted
		if valueToCompare == nil {
			valueToCompare = 0
		}

		// convert valueToCompare to int
		if typeValueToCompare.Kind() != reflect.Int {
			valueToCompareInt, e = strconv.Atoi(fmt.Sprintf("%v", valueToCompare))
			if e != nil {
				assigner.Logger.Debug(e.Error())
				valueToCompareInt = len(fmt.Sprintf("%v", valueToCompare))
			}
		}

		// set to 0 if value not setted
		if value == nil {
			valueInt = 0
		}

		// convert value to int
		if typeValue.Kind() != reflect.Int {
			valueInt, e = strconv.Atoi(fmt.Sprintf("%v", value))
			if e != nil {
				assigner.Logger.Debug(e.Error())
				valueInt = len(fmt.Sprintf("%v", value))
			}
		}

		functionGenerated = fmt.Sprintf("%v <= %v", valueToCompareInt, value)
		valid = valueToCompareInt <= valueInt
	case ">":
		typeValueToCompare := reflect.ValueOf(valueToCompare)
		valueToCompareInt := 0
		typeValue := reflect.ValueOf(value)
		valueInt := 0
		var e error

		// set to 0 if valueToCompare not setted
		if valueToCompare == nil {
			valueToCompare = 0
		}

		// convert valueToCompare to int
		if typeValueToCompare.Kind() != reflect.Int {
			valueToCompareInt, e = strconv.Atoi(fmt.Sprintf("%v", valueToCompare))
			if e != nil {
				assigner.Logger.Debug(e.Error())
				valueToCompareInt = len(fmt.Sprintf("%v", valueToCompare))
			}
		}

		// set to 0 if value not setted
		if value == nil {
			valueInt = 0
		}

		// convert value to int
		if typeValue.Kind() != reflect.Int {
			valueInt, e = strconv.Atoi(fmt.Sprintf("%v", value))
			if e != nil {
				assigner.Logger.Debug(e.Error())
				valueInt = len(fmt.Sprintf("%v", value))
			}
		}

		functionGenerated = fmt.Sprintf("%v > %v", valueToCompareInt, value)
		valid = valueToCompareInt > valueInt
	case ">=":
		typeValueToCompare := reflect.ValueOf(valueToCompare)
		valueToCompareInt := 0
		typeValue := reflect.ValueOf(value)
		valueInt := 0
		var e error

		// set to 0 if valueToCompare not setted
		if valueToCompare == nil {
			valueToCompare = 0
		}

		// convert valueToCompare to int
		if typeValueToCompare.Kind() != reflect.Int {
			valueToCompareInt, e = strconv.Atoi(fmt.Sprintf("%v", valueToCompare))
			if e != nil {
				assigner.Logger.Debug(e.Error())
				valueToCompareInt = len(fmt.Sprintf("%v", valueToCompare))
			}
		}

		// set to 0 if value not setted
		if value == nil {
			valueInt = 0
		}

		// convert value to int
		if typeValue.Kind() != reflect.Int {
			valueInt, e = strconv.Atoi(fmt.Sprintf("%v", value))
			if e != nil {
				assigner.Logger.Debug(e.Error())
				valueInt = len(fmt.Sprintf("%v", value))
			}
		}

		functionGenerated = fmt.Sprintf("%v >= %v", valueToCompareInt, value)
		valid = valueToCompareInt >= valueInt
	case "!=":
		valueToCompareParsed := ""
		if valueToCompare != nil {
			valueToCompareParsed = valueToCompare.(string)
		}
		functionGenerated = fmt.Sprintf("%v != %v", valueToCompareParsed, value)
		valid = fmt.Sprintf("%v", valueToCompareParsed) != fmt.Sprintf("%v", value)
	}
	return
}

func (assigner *DefaultAssigner) AssignValue(parent domain.AssignVariableValue, valueToAssign domain.AssignVariableValue) (assigned interface{}) {

	assigner.Logger.Debug("AssignValue", zenlogger.ZenField{Key: "parent", Value: parent}, zenlogger.ZenField{Key: "valueToAssign", Value: valueToAssign})

	valueOfVariable := reflect.ValueOf(parent.Value)

	// if parent variable type is object
	switch parent.VarType {
	case "object":
		mapVariable := make(map[string]interface{}, 0)
		switch valueOfVariable.Kind() {
		case reflect.Map:
			iter := valueOfVariable.MapRange()
			for iter.Next() {
				mapVariable[iter.Key().String()] = iter.Value().Interface()
			}
			mapVariable[valueToAssign.Key] = valueToAssign.Value
			parent.Value = mapVariable
		default:
			mapVariable[valueToAssign.Key] = valueToAssign.Value
		}
		parent.Value = mapVariable
	case "arrayObject":
		// fmt.Println(fmt.Sprintf("field %v is %v to assign to parent %v which is %v", valueToAssign.Key, valueToAssign.VarType, parent.Key, parent.VarType)) // FOR DEBUG
		arrVariable := make([]interface{}, 0)
		switch valueOfVariable.Kind() {
		case reflect.Map:
			// fmt.Println("is map") // FOR DEBUG
			mapVariable := make(map[string]interface{}, 0)
			iter := valueOfVariable.MapRange()
			for iter.Next() {
				mapVariable[iter.Key().String()] = iter.Value().Interface()
			}
			mapVariable[valueToAssign.Key] = valueToAssign.Value
			parent.Value = mapVariable

		case reflect.Array, reflect.Slice:
			// fmt.Println("is array") // FOR DEBUG
			vRef := reflect.ValueOf(parent.Value)
			for i := 0; i < vRef.Len(); i++ {
				arrVariable = append(arrVariable, vRef.Index(i).Interface())
			}
			arrVariable = append(arrVariable, valueToAssign.Value)
			parent.Value = arrVariable
		default:
			// fmt.Println("default") // FOR DEBUG
			arrVariable = append(arrVariable, valueToAssign.Value)
		}
		parent.Value = arrVariable
	case "arrayString":
		arrVariable := make([]string, 0)
		switch valueOfVariable.Kind() {
		case reflect.Map:
			// fmt.Println("is map") // FOR DEBUG
			mapVariable := make([]string, 0)
			iter := valueOfVariable.MapRange()
			for iter.Next() {
				mapVariable = append(mapVariable, fmt.Sprintf("%v", iter.Value().Interface()))
			}
			mapVariable = append(mapVariable, fmt.Sprintf("%v", valueToAssign.Value))
			parent.Value = mapVariable

		case reflect.Array, reflect.Slice:
			// fmt.Println("is array") // FOR DEBUG
			vRef := reflect.ValueOf(parent.Value)
			for i := 0; i < vRef.Len(); i++ {
				arrVariable = append(arrVariable, vRef.Index(i).String())
			}
			arrVariable = append(arrVariable, fmt.Sprintf("%v", valueToAssign.Value))
			parent.Value = arrVariable
		default:
			// fmt.Println("default") // FOR DEBUG
			arrVariable = append(arrVariable, fmt.Sprintf("%v", valueToAssign.Value))
		}
		parent.Value = arrVariable
	case "arrayInteger":
	case "arrayBoolean":
	case "string":
		parent.Value = convertToString(valueToAssign.Value)
	case "integer":
		intVal, err := convertToInt64(valueToAssign.Value)
		if err != nil {
			assigner.Logger.Error("AssignValue", zenlogger.ZenField{Key: "error", Value: err.Error()})
			parent.Value = 0
		} else {
			parent.Value = intVal
		}
	case "boolean":
		parent.Value = valueToAssign.Value.(bool)
	case "float":

	default:
		parent.Value = fmt.Sprintf("%v", valueToAssign.Value)
	}

	assigned = parent.Value

	return
}

func (assigner *DefaultAssigner) coreReadCommand(funcArg any) (arg interface{}, err error) {
	str := fmt.Sprintf("%v", funcArg)
	assigner.Logger.Debug("ReadCommand", zenlogger.ZenField{Key: "str", Value: str})

	// return if there is no function used
	if isArgumentOrganic(funcArg) {
		arg = funcArg
		return
	}

	// Create regular expressions to match function names and their arguments
	// funcRe := regexp.MustCompile(`\b(json_decode|addPropertyToArray|ltrim|trim|substr|randomInt|dateFormat|dateNow|dateAdd|md5|sha1|sha256|concat|basicAuth|strtolower|lpz|rpz|lps|rps)\b`)
	// // argRe := regexp.MustCompile(`\(([^()]|\(([^()]|\(([^()]+)\))*\))*\)`)
	// argRe := regexp.MustCompile(`(\(([^()]|\(([^()]|\(([^()]+)\))*\))*\))|(\{[^{}]*\})`)

	funcRe := regexp.MustCompile(`(?:^|[^.])\b(json_decode|addPropertyToArray|lengthArray|ltrim|trim|substr|randomInt|uuid|replaceAll|dateFormat|dateNow|dateAdd|md5|sha1|sha256|hmacSha256|encryptWithPrivateKey|concat|basicAuth|strtolower|lpz|rpz|lps|rps)\b`)
	// argRe := regexp.MustCompile(`(\(([^()]|\(([^()]|\(([^()]+)\))*\))*\))|(\{[^{}]*\})`)

	// Iterate over the string and extract nested function calls
	loop := 0
	for {
		if loop > 100 {
			err = fmt.Errorf("infinite loop detected in coreReadCommand")
			assigner.Logger.Error("Infinite loop detected in coreReadCommand")
			return
		}

		// Find the first match of a function name in the string
		funcMatch := funcRe.FindString(str)
		funcMatch = strings.TrimSpace(funcMatch)
		if funcMatch == "" {
			break
		}

		// Find the argument list for the function
		funcStart := funcRe.FindStringIndex(str)[0]
		argStart := funcStart + len(funcMatch)

		argMatch := ""
		depth := 0
		inString := false
		escape := false
		for i := argStart; i < len(str); i++ {
			ch := str[i]

			if ch == '\\' && !escape {
				escape = true
				continue
			}

			if ch == '"' && !escape {
				inString = !inString
			}

			if !inString {
				if ch == '(' {
					if depth == 0 {
						argStart = i
					}
					depth++
				} else if ch == ')' {
					depth--
					if depth == 0 {
						argMatch = str[argStart : i+1]
						break
					}
				}
			}
			escape = false
		}

		if argMatch == "" {
			// fallback / error handling
			result := funcMatch
			str = str[:funcStart] + result
			break
		}

		argEnd := argStart + len(argMatch) - 1

		// Print the function name and argument list
		// fmt.Printf("CALL %s: %s", funcMatch, argMatch[1:len(argMatch)-1])

		// Remove the function call and its argument list from the string
		// str = str[:funcStart] + str[argEnd+1:]
		subArg := argMatch[1 : len(argMatch)-1]

		var subArgI interface{}

		if funcMatch == "json_decode" {
			assigner.Logger.Debug("execute json_decode", zenlogger.ZenField{Key: "param", Value: subArg}, zenlogger.ZenField{Key: "loop", Value: loop})
			result, err := assigner.JsonDecode(subArg)
			if err != nil {
				assigner.Logger.Error("execute json_decode", zenlogger.ZenField{Key: "error", Value: err.Error()})
			} else {
				arg = result
			}
			assigner.Logger.Debug("execute json_decode", zenlogger.ZenField{Key: "result", Value: result}, zenlogger.ZenField{Key: "loop", Value: loop})
			break
		}

		if argMatch[1:len(argMatch)-1] != "" {
			assigner.Logger.Debug(fmt.Sprintf("send to coreReadCommand: %q", subArg))
			subArgI, err = assigner.coreReadCommand(subArg)
			if err != nil {
				assigner.Logger.Error(err.Error())
				break
			}
			subArg = fmt.Sprintf("%v", subArgI)
			assigner.Logger.Debug(fmt.Sprintf("subArg: %v", subArg))
		}

		// Print the function name and argument list
		assigner.Logger.Debug(fmt.Sprintf("CALL %s: %s", funcMatch, subArg), zenlogger.ZenField{Key: "loop", Value: loop})

		if err == nil {
			var err error
			switch funcMatch {
			case "trim":
				assigner.Logger.Debug("execute trim", zenlogger.ZenField{Key: "param", Value: subArg}, zenlogger.ZenField{Key: "loop", Value: loop})
				result := ""

				if subArg == "" {
					result = "invalid parameter"
					err = errors.New(result)
				} else {
					argArr := splitArgs(subArg)
					lenArgArr := len(argArr)
					char := " "

					if lenArgArr != 1 && lenArgArr != 2 {
						result = "invalid parameter"
						err = errors.New(result)
					}

					if err == nil {
						if lenArgArr == 2 {
							char = argArr[1]
						}
						result, err = assigner.Trim(argArr[0], char)
						if err != nil {
							// show log error if function fail to executed
							assigner.Logger.Error("execute trim", zenlogger.ZenField{Key: "error", Value: err.Error()})
						}
					}
				}

				str = str[:funcStart] + strconv.Quote(result) + str[argEnd+1:]
				assigner.Logger.Debug("execute trim", zenlogger.ZenField{Key: "result", Value: result}, zenlogger.ZenField{Key: "loop", Value: loop})
			case "ltrim":
				assigner.Logger.Debug("execute ltrim", zenlogger.ZenField{Key: "param", Value: subArg}, zenlogger.ZenField{Key: "loop", Value: loop})
				result := ""

				var argArr []string
				if subArg != "" {
					argArr = splitArgs(subArg)
				}
				lenArgArr := len(argArr)

				// validating the parameter
				if lenArgArr == 0 {
					result = "invalid parameter"
					err = errors.New(result)
				} else if lenArgArr > 2 {
					result = "invalid parameter"
					err = errors.New(result)
				} else {
					char := " "
					if lenArgArr > 1 {
						char = argArr[1]
					}
					result, err = assigner.Ltrim(argArr[0], char)
					if err != nil {
						// show log error if function fail to executed
						assigner.Logger.Error("execute substr", zenlogger.ZenField{Key: "error", Value: err.Error()})
					}
				}

				str = str[:funcStart] + strconv.Quote(result) + str[argEnd+1:]
				assigner.Logger.Debug("execute ltrim", zenlogger.ZenField{Key: "result", Value: result}, zenlogger.ZenField{Key: "loop", Value: loop})
			case "substr":
				assigner.Logger.Debug("execute substr", zenlogger.ZenField{Key: "param", Value: subArg}, zenlogger.ZenField{Key: "loop", Value: loop})
				result := ""

				if subArg == "" {
					result = "invalid parameter"
					err = errors.New(result)
				} else {
					subArgArr := splitArgs(subArg)

					lenSubArgArr := len(subArgArr)

					if lenSubArgArr <= 1 {
						result = "invalid parameter"
						err = errors.New(result)
					}

					// set index from from arg
					from := 0

					// set default index to
					to := from

					if err == nil {

						from, err = strconv.Atoi(strings.TrimSpace(subArgArr[1]))
						if err != nil {
							assigner.Logger.Error(err.Error())
						}

						if lenSubArgArr == 3 { // if argument is 3
							if len(subArgArr) >= 2 {
								to, err = strconv.Atoi(strings.TrimSpace(subArgArr[2]))
								if err != nil {
									assigner.Logger.Error(err.Error())
								}
							}
						} else if lenSubArgArr == 2 { // if argument is 2
							to = len(subArgArr[0])
						} else { // otherwise
							result = "invalid parameter"
							err = errors.New(result)
						}
					}

					if err == nil {
						result, err = assigner.Substr(subArgArr[0], from, to)
						if err != nil {
							// show log error if function fail to executed
							assigner.Logger.Error("execute substr", zenlogger.ZenField{Key: "error", Value: err.Error()})
						}
					}

				}

				// wrap with quotes for safe splitArgs usage
				str = str[:funcStart] + strconv.Quote(result) + str[argEnd+1:]
				assigner.Logger.Debug("execute substr", zenlogger.ZenField{Key: "result", Value: fmt.Sprintf("%q", result)}, zenlogger.ZenField{Key: "loop", Value: loop})
			case "replaceAll":
				assigner.Logger.Debug("execute replaceAll", zenlogger.ZenField{Key: "param", Value: subArg}, zenlogger.ZenField{Key: "loop", Value: loop})
				result := ""

				if subArg == "" {
					result = "invalid parameter"
					err = errors.New(result)
				} else {
					subArgArr := splitArgs(subArg)

					fmt.Println(fmt.Sprintf("%q", subArgArr))

					lenSubArgArr := len(subArgArr)

					if lenSubArgArr < 2 {
						result = "invalid parameter"
						err = errors.New(result)
					}

					replaceTo := ""

					if err == nil {

						if lenSubArgArr == 3 {
							replaceTo = subArgArr[2]
						}

						result, err = assigner.ReplaceAll(subArgArr[0], subArgArr[1], replaceTo)
						if err != nil {
							// show log error if function fail to executed
							assigner.Logger.Error("execute replaceAll", zenlogger.ZenField{Key: "error", Value: err.Error()})
						}
					}

				}

				// wrap with quotes for safe splitArgs usage
				str = str[:funcStart] + strconv.Quote(result) + str[argEnd+1:]
				assigner.Logger.Debug("execute replaceAll", zenlogger.ZenField{Key: "result", Value: result}, zenlogger.ZenField{Key: "loop", Value: loop})
			case "lpz":
				assigner.Logger.Debug("execute lpz", zenlogger.ZenField{Key: "param", Value: subArg}, zenlogger.ZenField{Key: "loop", Value: loop})
				result := ""

				if subArg == "" {
					result = "invalid parameter"
					err = errors.New(result)
				} else {
					subArgArr := splitArgs(subArg)

					lenSubArgArr := len(subArgArr)

					// set index from from arg
					length := 0
					text := ""

					if err == nil {
						if lenSubArgArr < 1 {
							result = "invalid parameter"
							err = errors.New(result)
						} else if lenSubArgArr == 1 {
							text = ""
							length, err = strconv.Atoi(strings.TrimSpace(subArgArr[0]))
							if err != nil {
								assigner.Logger.Error(err.Error())
							}

						} else {
							text = subArgArr[0]
							length, err = strconv.Atoi(strings.TrimSpace(subArgArr[1]))
							if err != nil {
								assigner.Logger.Error(err.Error())
							}
						}

						if err == nil {
							result, err = assigner.Lpz(text, length)
							if err != nil {
								// show log error if function fail to executed
								assigner.Logger.Error("execute rps", zenlogger.ZenField{Key: "error", Value: err.Error()})
							}
						}
					}

				}

				// wrap with quotes for safe splitArgs usage
				str = str[:funcStart] + strconv.Quote(result) + str[argEnd+1:]
				assigner.Logger.Debug("execute lpz", zenlogger.ZenField{Key: "result", Value: result}, zenlogger.ZenField{Key: "loop", Value: loop})
			case "rpz":
				assigner.Logger.Debug("execute rpz", zenlogger.ZenField{Key: "param", Value: subArg}, zenlogger.ZenField{Key: "loop", Value: loop})
				result := ""

				if subArg == "" {
					result = "invalid parameter"
					err = errors.New(result)
				} else {
					subArgArr := splitArgs(subArg)

					lenSubArgArr := len(subArgArr)

					// set index from from arg
					length := 0
					text := ""

					if err == nil {
						if lenSubArgArr < 1 {
							result = "invalid parameter"
							err = errors.New(result)
						} else if lenSubArgArr == 1 {
							text = ""
							length, err = strconv.Atoi(strings.TrimSpace(subArgArr[0]))
							if err != nil {
								assigner.Logger.Error(err.Error())
							}

						} else {
							text = subArgArr[0]
							length, err = strconv.Atoi(strings.TrimSpace(subArgArr[1]))
							if err != nil {
								assigner.Logger.Error(err.Error())
							}
						}

						if err == nil {
							result, err = assigner.Rpz(text, length)
							if err != nil {
								// show log error if function fail to executed
								assigner.Logger.Error("execute rps", zenlogger.ZenField{Key: "error", Value: err.Error()})
							}
						}
					}

				}

				// wrap with quotes for safe splitArgs usage
				str = str[:funcStart] + strconv.Quote(result) + str[argEnd+1:]
				assigner.Logger.Debug("execute rpz", zenlogger.ZenField{Key: "result", Value: result}, zenlogger.ZenField{Key: "loop", Value: loop})
			case "lps":
				assigner.Logger.Debug("execute lps", zenlogger.ZenField{Key: "param", Value: subArg}, zenlogger.ZenField{Key: "loop", Value: loop})
				result := ""

				if subArg == "" {
					result = "invalid parameter"
					err = errors.New(result)
				} else {
					subArgArr := splitArgs(subArg)

					lenSubArgArr := len(subArgArr)

					// set index from from arg
					length := 0
					text := ""

					if err == nil {
						if lenSubArgArr < 1 {
							result = "invalid parameter"
							err = errors.New(result)
						} else if lenSubArgArr == 1 {
							text = ""
							length, err = strconv.Atoi(strings.TrimSpace(subArgArr[0]))
							if err != nil {
								assigner.Logger.Error(err.Error())
							}

						} else {
							text = subArgArr[0]
							length, err = strconv.Atoi(strings.TrimSpace(subArgArr[1]))
							if err != nil {
								assigner.Logger.Error(err.Error())
							}
						}

						if err == nil {
							result, err = assigner.Lps(text, length)
							if err != nil {
								// show log error if function fail to executed
								assigner.Logger.Error("execute rps", zenlogger.ZenField{Key: "error", Value: err.Error()})
							}
						}
					}

				}

				// wrap with quotes for safe splitArgs usage
				str = str[:funcStart] + strconv.Quote(result) + str[argEnd+1:]
				assigner.Logger.Debug("execute lps", zenlogger.ZenField{Key: "result", Value: result}, zenlogger.ZenField{Key: "str", Value: str}, zenlogger.ZenField{Key: "loop", Value: loop})
			case "rps":
				assigner.Logger.Debug("execute rps", zenlogger.ZenField{Key: "param", Value: subArg}, zenlogger.ZenField{Key: "loop", Value: loop})
				result := ""

				if subArg == "" {
					result = "invalid parameter"
					err = errors.New(result)
				} else {
					subArgArr := splitArgs(subArg)

					lenSubArgArr := len(subArgArr)

					// set index from from arg
					length := 0
					text := ""

					if err == nil {
						if lenSubArgArr < 1 {
							result = "invalid parameter"
							err = errors.New(result)
						} else if lenSubArgArr == 1 {
							text = ""
							length, err = strconv.Atoi(strings.TrimSpace(subArgArr[0]))
							if err != nil {
								assigner.Logger.Error(err.Error())
							}

						} else {
							text = subArgArr[0]
							length, err = strconv.Atoi(strings.TrimSpace(subArgArr[1]))
							if err != nil {
								assigner.Logger.Error(err.Error())
							}
						}

						if err == nil {
							result, err = assigner.Rps(text, length)
							if err != nil {
								// show log error if function fail to executed
								assigner.Logger.Error("execute rps", zenlogger.ZenField{Key: "error", Value: err.Error()})
							}
						}
					}

				}

				str = str[:funcStart] + strconv.Quote(result) + str[argEnd+1:]
				assigner.Logger.Debug("execute rps", zenlogger.ZenField{Key: "result", Value: result}, zenlogger.ZenField{Key: "str", Value: str}, zenlogger.ZenField{Key: "loop", Value: loop})
			case "randomInt":
				assigner.Logger.Debug("execute randomInt", zenlogger.ZenField{Key: "param", Value: subArg}, zenlogger.ZenField{Key: "loop", Value: loop})
				result := ""
				from := 0
				to := 0

				if subArg == "" {
					from = 0
					to = 17
				} else {
					subArgArr := splitArgs(subArg)

					if len(subArgArr) <= 2 { // if argument 1 or 2
						from, err3 := strconv.Atoi(strings.TrimSpace(subArgArr[0]))
						if err3 != nil {
							assigner.Logger.Error(err3.Error())
						}

						to = from
						if len(subArgArr) >= 2 {
							to, err3 = strconv.Atoi(strings.TrimSpace(subArgArr[1]))
							if err3 != nil {
								assigner.Logger.Error(err3.Error())
							} else if to > 17 {
								to = 17
							}
						}
					} else { // otherwise
						result = "invalid parameter"
						err = errors.New(result)
					}
				}

				// check if parameter validation is valid, if so then execute command
				if err == nil {
					result, err = assigner.RandomInt(from, to)
					if err != nil {
						assigner.Logger.Error("execute randomInt", zenlogger.ZenField{Key: "error", Value: err.Error()})
					}
				}

				// replace the string from raw function to its result
				str = str[:funcStart] + result + str[argEnd+1:]

				assigner.Logger.Debug("execute randomInt", zenlogger.ZenField{Key: "result", Value: result}, zenlogger.ZenField{Key: "loop", Value: loop})
			case "uuid":
				assigner.Logger.Debug("execute uuid", zenlogger.ZenField{Key: "param", Value: subArg}, zenlogger.ZenField{Key: "loop", Value: loop})
				result := ""

				result, err = assigner.Uuid()
				if err != nil {
					assigner.Logger.Error("execute randomInt", zenlogger.ZenField{Key: "error", Value: err.Error()})
				}

				// wrap with quotes for safe splitArgs usage
				result = fmt.Sprintf("\"%s\"", result)

				// replace the string from raw function to its result
				str = str[:funcStart] + result + str[argEnd+1:]

				assigner.Logger.Debug("execute uuid", zenlogger.ZenField{Key: "result", Value: result}, zenlogger.ZenField{Key: "loop", Value: loop})
			case "dateNow":
				assigner.Logger.Debug("execute dateNow", zenlogger.ZenField{Key: "param", Value: subArg}, zenlogger.ZenField{Key: "loop", Value: loop})
				result, err := assigner.DateNow(subArg)
				if err != nil {
					assigner.Logger.Error("execute dateNow", zenlogger.ZenField{Key: "error", Value: err.Error()})
				} else {
					str = str[:funcStart] + strconv.Quote(result) + str[argEnd+1:]
				}
				assigner.Logger.Debug("execute dateNow", zenlogger.ZenField{Key: "result", Value: result}, zenlogger.ZenField{Key: "str", Value: str}, zenlogger.ZenField{Key: "loop", Value: loop})
			case "dateAdd":
				assigner.Logger.Debug("execute dateAdd", zenlogger.ZenField{Key: "param", Value: subArg}, zenlogger.ZenField{Key: "loop", Value: loop})
				result := ""
				argArr := splitArgs(subArg)

				if len(argArr) == 4 {
					add, errAdd := strconv.Atoi(strings.TrimSpace(argArr[2]))
					if errAdd != nil {
						assigner.Logger.Error("execute dateAdd", zenlogger.ZenField{Key: "error", Value: err.Error()})
					}
					result, err = assigner.DateAdd(strings.TrimSpace(argArr[0]), strings.TrimSpace(argArr[1]), add, strings.TrimSpace(argArr[3]))
					if err != nil {
						assigner.Logger.Error(err.Error())
					}
				} else {
					assigner.Logger.Error("execute dateAdd", zenlogger.ZenField{Key: "error", Value: "invalid parameter"})
					result = "invalid parameter"
					err = errors.New(result)
				}

				str = str[:funcStart] + strconv.Quote(result) + str[argEnd+1:]
				assigner.Logger.Debug("execute dateAdd", zenlogger.ZenField{Key: "result", Value: result}, zenlogger.ZenField{Key: "str", Value: str}, zenlogger.ZenField{Key: "loop", Value: loop})
			case "dateFormat":
				assigner.Logger.Debug("execute dateFormat", zenlogger.ZenField{Key: "param", Value: subArg}, zenlogger.ZenField{Key: "loop", Value: loop})
				result := ""
				argArr := splitArgs(subArg)

				if len(argArr) > 3 { // if parameter more than needed
					assigner.Logger.Error("execute dateFormat", zenlogger.ZenField{Key: "error", Value: "invalid parameter"})
					result = "invalid parameter"
					err = errors.New(result)
				} else if len(argArr) < 3 { // if parameter less than needed
					assigner.Logger.Error("execute dateFormat", zenlogger.ZenField{Key: "error", Value: "invalid parameter"})
					result = "invalid parameter"
					err = errors.New(result)
				} else { // if parameter match with needed
					result, err = assigner.DateFormat(strings.TrimSpace(argArr[0]), strings.TrimSpace(argArr[1]), strings.TrimSpace(argArr[2]))
					if err != nil {
						assigner.Logger.Error(err.Error())
					}
				}

				str = str[:funcStart] + strconv.Quote(result) + str[argEnd+1:]
				assigner.Logger.Debug("execute dateFormat", zenlogger.ZenField{Key: "result", Value: result}, zenlogger.ZenField{Key: "str", Value: str}, zenlogger.ZenField{Key: "loop", Value: loop})
			case "md5":
				assigner.Logger.Debug("execute md5", zenlogger.ZenField{Key: "param", Value: subArg}, zenlogger.ZenField{Key: "loop", Value: loop})

				argArr := splitArgs(subArg)

				result, err := assigner.MD5(argArr[0])
				if err != nil {
					assigner.Logger.Error("execute md5", zenlogger.ZenField{Key: "error", Value: err.Error()})
				} else {
					str = str[:funcStart] + strconv.Quote(result) + str[argEnd+1:]
				}
				assigner.Logger.Debug("execute md5", zenlogger.ZenField{Key: "result", Value: result}, zenlogger.ZenField{Key: "str", Value: str}, zenlogger.ZenField{Key: "loop", Value: loop})
			case "sha1":
				assigner.Logger.Debug("execute sha1", zenlogger.ZenField{Key: "param", Value: subArg}, zenlogger.ZenField{Key: "loop", Value: loop})

				argArr := splitArgs(subArg)

				result, err := assigner.Sha1(argArr[0])
				if err != nil {
					assigner.Logger.Error("execute sha1", zenlogger.ZenField{Key: "error", Value: err.Error()})
				} else {
					str = str[:funcStart] + strconv.Quote(result) + str[argEnd+1:]
				}
				assigner.Logger.Debug("execute sha1", zenlogger.ZenField{Key: "result", Value: result}, zenlogger.ZenField{Key: "str", Value: str}, zenlogger.ZenField{Key: "loop", Value: loop})
			case "sha256":
				assigner.Logger.Debug("execute sha256", zenlogger.ZenField{Key: "param", Value: subArg}, zenlogger.ZenField{Key: "loop", Value: loop})
				args := splitArgs(subArg)
				result, err := assigner.Sha256(args...)
				if err != nil {
					assigner.Logger.Error("execute sha256", zenlogger.ZenField{Key: "error", Value: err.Error()})
				} else {
					str = str[:funcStart] + strconv.Quote(result) + str[argEnd+1:]
				}
				assigner.Logger.Debug("execute sha256", zenlogger.ZenField{Key: "result", Value: result}, zenlogger.ZenField{Key: "loop", Value: loop})
			case "hmacSha256":
				result := ""
				assigner.Logger.Debug("execute hmacSha256", zenlogger.ZenField{Key: "param", Value: subArg}, zenlogger.ZenField{Key: "loop", Value: loop})
				args := strings.Split(subArg, "|")
				if len(args) != 2 {
					err := errors.New("invalid number of arguments")
					assigner.Logger.Error("execute hmacSha256", zenlogger.ZenField{Key: "error", Value: err.Error()})
				} else {
					result, err = assigner.hmacSha256(strings.TrimSpace(args[0]), strings.TrimSpace(args[1]))
					if err != nil {
						assigner.Logger.Error("execute hmacSha256", zenlogger.ZenField{Key: "error", Value: err.Error()})
					} else {
						str = str[:funcStart] + strconv.Quote(result) + str[argEnd+1:]
					}
				}

				assigner.Logger.Debug("execute hmacSha256", zenlogger.ZenField{Key: "result", Value: result}, zenlogger.ZenField{Key: "str", Value: str}, zenlogger.ZenField{Key: "loop", Value: loop})
			case "encryptWithPrivateKey":
				result := ""
				assigner.Logger.Debug("execute EncryptWithPrivateKey", zenlogger.ZenField{Key: "param", Value: subArg}, zenlogger.ZenField{Key: "loop", Value: loop})
				subArgArr := splitArgs(subArg)

				if len(subArgArr) != 2 {
					err = errors.New("invalid number of arguments")
					assigner.Logger.Error("execute EncryptWithPrivateKey", zenlogger.ZenField{Key: "error", Value: err.Error()})
					return str, err
				}

				// Perform encryption
				privKey, ee := strconv.Unquote(subArgArr[1])
				if ee != nil {
					assigner.Logger.Debug("execute EncryptWithPrivateKey", zenlogger.ZenField{Key: "info", Value: "there is no escaped character"})
					privKey = subArgArr[1]
				}
				result, err = assigner.EncryptWithPrivateKey(strings.TrimSpace(subArgArr[0]), strings.TrimSpace(privKey))
				if err != nil {
					assigner.Logger.Error("execute EncryptWithPrivateKey", zenlogger.ZenField{Key: "error", Value: err.Error()})
					return str, err
				}

				str = str[:funcStart] + strconv.Quote(result) + str[argEnd+1:]

				assigner.Logger.Debug("execute EncryptWithPrivateKey", zenlogger.ZenField{Key: "result", Value: result}, zenlogger.ZenField{Key: "str", Value: str}, zenlogger.ZenField{Key: "loop", Value: loop})
			case "basicAuth":
				assigner.Logger.Debug("execute basicAuth", zenlogger.ZenField{Key: "param", Value: subArg}, zenlogger.ZenField{Key: "loop", Value: loop})
				subArgArr := splitArgs(subArg)
				result, err := assigner.BasicAuth(subArgArr[0])
				if err != nil {
					assigner.Logger.Error("execute basicAuth", zenlogger.ZenField{Key: "error", Value: err.Error()})
				} else {
					str = str[:funcStart] + strconv.Quote(result) + str[argEnd+1:]
				}
				assigner.Logger.Debug("execute basicAuth", zenlogger.ZenField{Key: "result", Value: result}, zenlogger.ZenField{Key: "str", Value: str}, zenlogger.ZenField{Key: "loop", Value: loop})
			case "concat":
				assigner.Logger.Debug("execute concat", zenlogger.ZenField{Key: "param", Value: subArg}, zenlogger.ZenField{Key: "loop", Value: loop})
				result := ""

				if subArg == "" {
					result = "invalid parameter"
					err = errors.New(result)
				} else {
					subArgArr := splitArgs(subArg)
					lenArgArr := len(subArgArr)

					if lenArgArr == 0 {
						result = "invalid parameter"
						err = errors.New(result)
					}

					if err == nil {
						result = assigner.Concat(subArgArr...)
					}
				}

				str = str[:funcStart] + strconv.Quote(result) + str[argEnd+1:]
				assigner.Logger.Debug("execute concat", zenlogger.ZenField{Key: "result", Value: result}, zenlogger.ZenField{Key: "str", Value: str}, zenlogger.ZenField{Key: "loop", Value: loop})
			case "strtolower":
				assigner.Logger.Debug("execute strtolower", zenlogger.ZenField{Key: "param", Value: subArg}, zenlogger.ZenField{Key: "loop", Value: loop})
				result := ""

				if subArg == "" {
					result = "invalid parameter"
					err = errors.New(result)
				} else {
					argArr := splitArgs(subArg)
					lenArgArr := len(argArr)

					if lenArgArr == 1 {
						result = assigner.Strtolower(argArr[0])
					} else {
						result = "invalid parameter"
					}
				}

				str = str[:funcStart] + strconv.Quote(result) + str[argEnd+1:]
				assigner.Logger.Debug("execute strtolower", zenlogger.ZenField{Key: "result", Value: result}, zenlogger.ZenField{Key: "str", Value: str}, zenlogger.ZenField{Key: "loop", Value: loop})
			case "addPropertyToArray":
				assigner.Logger.Debug("execute addPropertyToArray", zenlogger.ZenField{Key: "param", Value: subArg}, zenlogger.ZenField{Key: "loop", Value: loop})
				result := ""
				var jsonArrayStr, propName, funcCall string
				subArg = strings.TrimSpace(subArg)
				// If the argument starts with a JSON array, extract it properly.
				if strings.HasPrefix(subArg, "[") {
					depth := 0
					pos := -1
					for i, ch := range subArg {
						if ch == '[' {
							depth++
						} else if ch == ']' {
							depth--
							if depth == 0 {
								pos = i
								break
							}
						}
					}
					if pos == -1 {
						err = errors.New("invalid JSON array in addPropertyToArray")
						result = "invalid JSON array"
						break
					}
					jsonArrayStr = subArg[:pos+1]
					remainder := strings.TrimSpace(subArg[pos+1:])
					if strings.HasPrefix(remainder, ",") {
						remainder = strings.TrimSpace(remainder[1:])
					}
					parts := strings.SplitN(remainder, ",", 2)
					if len(parts) < 2 {
						err = errors.New("invalid parameter for addPropertyToArray, missing property name or function call")
						result = "invalid parameter"
						break
					}
					propName = strings.TrimSpace(parts[0])
					funcCall = strings.TrimSpace(parts[1])
				} else {
					// Fallback if JSON array is not recognized at start.
					argArr := splitArgs(subArg)
					if len(argArr) < 3 {
						result = "invalid parameter"
						err = errors.New(result)
						break
					}
					jsonArrayStr = strings.TrimSpace(argArr[0])
					propName = strings.TrimSpace(argArr[1])
					funcCall = strings.TrimSpace(argArr[2])
				}

				// Unmarshal the JSON array.
				var arr []map[string]interface{}
				if jsonErr := json.Unmarshal([]byte(jsonArrayStr), &arr); jsonErr != nil {
					assigner.Logger.Error("execute addPropertyToArray", zenlogger.ZenField{Key: "error", Value: jsonErr.Error()})
					result = "invalid JSON array"
					err = jsonErr
				} else {
					// Iterate through each object in the array.
					for i, item := range arr {
						var newVal interface{}
						itemFuncCall := funcCall
						// If the function call is prefixed with "item.", remove the prefix.
						if strings.HasPrefix(itemFuncCall, "item.") {
							itemFuncCall = itemFuncCall[len("item."):]
						}
						// Now call coreReadCommand recursively with the (possibly modified) function call.
						newVal, err = assigner.coreReadCommand(itemFuncCall)
						if err != nil {
							assigner.Logger.Error("execute addPropertyToArray", zenlogger.ZenField{Key: "error", Value: err.Error()})
							break
						}
						// Add the new property; we store the value as a string to match the expected output.
						item[propName] = fmt.Sprintf("%v", newVal)
						arr[i] = item
					}
					if err == nil {
						var marshalResult []byte
						marshalResult, err = json.Marshal(arr)
						if err != nil {
							assigner.Logger.Error("execute addPropertyToArray", zenlogger.ZenField{Key: "error", Value: err.Error()})
							result = "error marshalling result"
						} else {
							result = string(marshalResult)
						}
					}
				}

				str = str[:funcStart] + result + str[argEnd+1:]
				assigner.Logger.Debug("execute addPropertyToArray", zenlogger.ZenField{Key: "result", Value: result}, zenlogger.ZenField{Key: "loop", Value: loop})
			case "lengthArray":
				assigner.Logger.Debug("execute lengthArray", zenlogger.ZenField{Key: "param", Value: subArg}, zenlogger.ZenField{Key: "loop", Value: loop})
				result := "0"
				length, err := assigner.LengthArray(subArg)
				if err != nil {
					assigner.Logger.Error("execute lengthArray", zenlogger.ZenField{Key: "error", Value: err.Error()})
				} else {
					result = fmt.Sprintf("%v", length)
					// replace the string from raw function to its result
					str = str[:funcStart] + result + str[argEnd+1:]
				}
				assigner.Logger.Debug("execute lengthArray", zenlogger.ZenField{Key: "result", Value: result}, zenlogger.ZenField{Key: "loop", Value: loop})
			}
		}
		loop++

		if str == "" {
			break
		}
	}

	// set manipulated argument to arg if arg still nill
	if arg == nil {
		arg = str
	}

	return
}

func (assigner *DefaultAssigner) findArg(command string) (argument string) {
	leftIndex := 0
	rightIndex := len(command)

	// find batas kiri argument
	for i := 0; i < len(command); i++ {
		if command[i] == byte('(') {
			leftIndex = i + 1
			break
		}
	}

	// find batas kanan argument
	for i := len(command) - 1; i >= 0; i-- {
		if command[i] == byte(')') {
			rightIndex = i
			break
		}
	}

	argument = command[leftIndex:rightIndex]
	return
}
