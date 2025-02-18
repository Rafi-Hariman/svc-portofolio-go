package print2

import (
	"fmt"
	"log"
	"reflect"
	"strings"
)

// Main function to print map, struct, or slice
func PrintMap(input interface{}) {
	// Convert struct or slice if necessary
	convertedInput := prepareInput(input)

	// Handle map[string]interface{}
	if m, ok := convertedInput.(map[string]interface{}); ok {
		log.Println(printMapRecursive(m, 0))
	} else if s, ok := convertedInput.([]interface{}); ok {
		log.Println(printSlice(s, 0))
	} else {
		log.Println("Unsupported input type")
	}
}

// prepareInput handles the conversion of struct or slice to the appropriate format
func prepareInput(input interface{}) interface{} {
	// Check if the input is already a map[string]interface{}
	if _, ok := input.(map[string]interface{}); ok {
		return input
	}

	// Check if the input is a struct, and convert it
	val := reflect.ValueOf(input)
	kind := val.Kind()

	switch kind {
	case reflect.Struct:
		return structToMap(input)
	case reflect.Slice, reflect.Array:
		return sliceToInterface(input)
	default:
		return input
	}
}

// Convert struct to map[string]interface{} recursively
func structToMap(data interface{}) map[string]interface{} {
	val := reflect.ValueOf(data)
	typ := reflect.TypeOf(data)
	result := make(map[string]interface{})

	// Only handle structs
	if val.Kind() == reflect.Struct {
		for i := 0; i < val.NumField(); i++ {
			field := typ.Field(i)
			tag := field.Tag.Get("json")
			if tag == "" {
				tag = field.Name
			}

			fieldValue := val.Field(i).Interface()

			// Handle nested structs by calling structToMap recursively
			if reflect.ValueOf(fieldValue).Kind() == reflect.Struct {
				result[tag] = structToMap(fieldValue)
			} else if reflect.ValueOf(fieldValue).Kind() == reflect.Slice {
				result[tag] = sliceToInterface(fieldValue)
			} else {
				result[tag] = fieldValue
			}
		}
	}
	return result
}

// Convert slice to []interface{} recursively
func sliceToInterface(slice interface{}) []interface{} {
	val := reflect.ValueOf(slice)
	result := []interface{}{}

	for i := 0; i < val.Len(); i++ {
		item := val.Index(i).Interface()
		if reflect.ValueOf(item).Kind() == reflect.Struct {
			result = append(result, structToMap(item))
		} else if reflect.ValueOf(item).Kind() == reflect.Slice {
			result = append(result, sliceToInterface(item))
		} else {
			result = append(result, item)
		}
	}

	return result
}

// Recursive function to print map[string]interface{}
func printMapRecursive(m map[string]interface{}, indent int) string {
	indentStr := strings.Repeat("    ", indent) // 4 spaces per indent level
	result := "map{\n"
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	for _, k := range keys {
		v := m[k]
		result += fmt.Sprintf("%s    \"%s\": ", indentStr, k)
		switch val := v.(type) {
		case map[string]interface{}:
			result += printMapRecursive(val, indent+1)
		case []interface{}:
			result += printSlice(val, indent+1)
		default:
			result += fmt.Sprintf("%v,\n", formatValueWithType(val))
		}
	}
	result += fmt.Sprintf("%s},\n", indentStr)
	return result
}

// Recursive function to print []interface{}
func printSlice(s []interface{}, indent int) string {
	indentStr := strings.Repeat("    ", indent)
	result := "[\n"
	for _, v := range s {
		switch val := v.(type) {
		case map[string]interface{}:
			result += printMapRecursive(val, indent+1)
		case []interface{}:
			result += printSlice(val, indent+1)
		default:
			result += fmt.Sprintf("%s    %v,\n", indentStr, formatValueWithType(val))
		}
	}
	result += fmt.Sprintf("%s],\n", indentStr)
	return result
}

// Format the value with its type
func formatValueWithType(v interface{}) string {
	var typeName string
	switch val := v.(type) {
	case string:
		typeName = "string"
		return fmt.Sprintf("%s(\"%s\")", typeName, val)
	case int, int8, int16, int32, int64:
		typeName = "int"
		return fmt.Sprintf("%s(%v)", typeName, val)
	case uint, uint8, uint16, uint32, uint64:
		typeName = "uint"
		return fmt.Sprintf("%s(%v)", typeName, val)
	case float32, float64:
		typeName = "float"
		return fmt.Sprintf("%s(%v)", typeName, val)
	case bool:
		typeName = "bool"
		return fmt.Sprintf("%s(%v)", typeName, val)
	default:
		// Use reflect for other types
		typeName = reflect.TypeOf(v).String()
		return fmt.Sprintf("%s(%v)", typeName, val)
	}
}
