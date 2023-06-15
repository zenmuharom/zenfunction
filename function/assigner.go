package function

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/zenmuharom/zenfunction/domain"
	"github.com/zenmuharom/zenlogger"
)

type DefaultAssigner struct {
	Logger zenlogger.Zenlogger
}

type Assigner interface {
	ReadCommand(arg string) (returnVal interface{}, err error)
	DateAdd(format string, theDate string, add int, duration string) (added string, err error)
	DateNow(format string) (generated string, err error)
	Trim(arg0, arg1 string) (trimmed string, err error)
	Substr(arg string, from int, to int) (substred string, err error)
	findArg(command string) (argument string)
}

func NewAssigner(logger zenlogger.Zenlogger) Assigner {
	return &DefaultAssigner{Logger: logger}
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

	case "arrayInteger":

	case "arrayBoolean":
	case "string":
		parent.Value = fmt.Sprintf("%v", valueToAssign.Value)
	case "integer":
		intVal, err := strconv.Atoi(fmt.Sprintf("%v", valueToAssign.Value))
		if err != nil {
			parent.Value = 0
		} else {

			parent.Value = intVal
		}
	case "boolean":
		parent.Value = valueToAssign.Value.(bool)
	default:
		parent.VarType = fmt.Sprintf("%v", valueToAssign.Value)
	}

	assigned = parent.Value

	return
}

func (assigner *DefaultAssigner) ReadCommand(str string) (arg interface{}, err error) {
	assigner.Logger.Debug("ReadCommand", zenlogger.ZenField{Key: "str", Value: str})

	// Create regular expressions to match function names and their arguments
	funcRe := regexp.MustCompile(`\b(ltrim|trim|substr|randomInt|dateNow|dateAdd)\b`)
	argRe := regexp.MustCompile(`\(([^()]|\(([^()]|\(([^()]+)\))*\))*\)`)

	// Iterate over the string and extract nested function calls
	loop := 0
	for {
		// Find the first match of a function name in the string
		funcMatch := funcRe.FindString(str)
		if funcMatch == "" {
			break
		}

		// Find the argument list for the function
		funcStart := funcRe.FindStringIndex(str)[0]
		argStart := funcStart + len(funcMatch)

		argMatches := argRe.FindStringSubmatch(str[argStart:])
		argMatch := ""
		if len(argMatches) == 0 {
			result := funcMatch
			// replace the string from raw function to its result
			str = str[:funcStart] + result
			break
		}

		// set argument part to variable
		argMatch = argMatches[0]
		argEnd := argStart + len(argMatch) - 1

		// Print the function name and argument list
		// fmt.Printf("CALL %s: %s", funcMatch, argMatch[1:len(argMatch)-1])

		// Remove the function call and its argument list from the string
		// str = str[:funcStart] + str[argEnd+1:]
		subArg := argMatch[1 : len(argMatch)-1]

		var subArgI interface{}
		if argMatch[1:len(argMatch)-1] != "" {
			assigner.Logger.Debug(fmt.Sprintf("send to ReadCommand2: %v", subArg))
			subArgI, err = assigner.ReadCommand(subArg)
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
					argArr := strings.Split(fmt.Sprintf("%v", subArg), ",")
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

				// replace the string from raw function to its result
				str = str[:funcStart] + result + str[argEnd+1:]

				assigner.Logger.Debug("execute dateNow", zenlogger.ZenField{Key: "result", Value: result}, zenlogger.ZenField{Key: "loop", Value: loop})
			case "ltrim":
				assigner.Logger.Debug("execute ltrim", zenlogger.ZenField{Key: "param", Value: subArg}, zenlogger.ZenField{Key: "loop", Value: loop})
				result := ""

				var argArr []string
				if subArg != "" {
					argArr = strings.Split(fmt.Sprintf("%v", subArg), ",")
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

				// replace the string from raw function to its result
				str = str[:funcStart] + result + str[argEnd+1:]

				assigner.Logger.Debug("execute ltrim", zenlogger.ZenField{Key: "result", Value: result}, zenlogger.ZenField{Key: "loop", Value: loop})
			case "substr":
				assigner.Logger.Debug("execute substr", zenlogger.ZenField{Key: "param", Value: subArg}, zenlogger.ZenField{Key: "loop", Value: loop})
				result := ""

				if subArg == "" {
					result = "invalid parameter"
					err = errors.New(result)
				} else {
					subArgArr := splitWithEscapedCommas(subArg)

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

				str = str[:funcStart] + result + str[argEnd+1:]
				assigner.Logger.Debug("execute substr", zenlogger.ZenField{Key: "result", Value: result}, zenlogger.ZenField{Key: "loop", Value: loop})

			case "randomInt":
				assigner.Logger.Debug("execute randomInt", zenlogger.ZenField{Key: "param", Value: subArg}, zenlogger.ZenField{Key: "loop", Value: loop})
				result := ""
				from := 0
				to := 0

				if subArg == "" {
					from = 0
					to = 17
				} else {
					subArgArr := strings.Split(fmt.Sprintf("%v", subArg), ",")

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
			case "dateNow":
				assigner.Logger.Debug("execute dateNow", zenlogger.ZenField{Key: "param", Value: subArg}, zenlogger.ZenField{Key: "loop", Value: loop})
				result, err := assigner.DateNow(subArg)
				if err != nil {
					assigner.Logger.Error("execute dateNow", zenlogger.ZenField{Key: "error", Value: err.Error()})
				} else {
					// replace the string from raw function to its result
					str = str[:funcStart] + result + str[argEnd+1:]
				}
				assigner.Logger.Debug("execute dateNow", zenlogger.ZenField{Key: "result", Value: result}, zenlogger.ZenField{Key: "loop", Value: loop})

			case "dateAdd":
				assigner.Logger.Debug("execute dateAdd", zenlogger.ZenField{Key: "param", Value: subArg}, zenlogger.ZenField{Key: "loop", Value: loop})
				result := ""
				argArr := strings.Split(fmt.Sprintf("%v", subArg), ",")

				if len(argArr) > 4 { // if parameter more than needed
					assigner.Logger.Error("execute dateAdd", zenlogger.ZenField{Key: "error", Value: "invalid parameter"})
					result = "invalid parameter"
					err = errors.New(result)
				} else if len(argArr) < 4 { // if parameter less than needed
					assigner.Logger.Error("execute dateAdd", zenlogger.ZenField{Key: "error", Value: "invalid parameter"})
					result = "invalid parameter"
					err = errors.New(result)
				} else { // if parameter match with needed
					add, errAdd := strconv.Atoi(strings.TrimSpace(argArr[2]))
					if errAdd != nil {
						assigner.Logger.Error("execute dateAdd", zenlogger.ZenField{Key: "error", Value: err.Error()})
					}
					result, err = assigner.DateAdd(strings.TrimSpace(argArr[0]), strings.TrimSpace(argArr[1]), add, strings.TrimSpace(argArr[3]))
					if err != nil {
						assigner.Logger.Error(err.Error())
					}
				}

				// replace the string from raw function to its result
				str = str[:funcStart] + result + str[argEnd+1:]

				assigner.Logger.Debug("execute dateAdd", zenlogger.ZenField{Key: "result", Value: result}, zenlogger.ZenField{Key: "loop", Value: loop})
			}
		}
		loop++

		if str == "" {
			break
		}
	}

	arg = str

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
