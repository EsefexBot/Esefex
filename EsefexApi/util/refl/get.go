package refl

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/pkg/errors"
)

func GetNestedFieldValue(obj interface{}, path string) (interface{}, error) {
	val := reflect.ValueOf(obj)

	// Navigate through the nested fields
	for val.Kind() == reflect.Ptr || val.Kind() == reflect.Interface {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("provided object is not a struct")
	}

	// Split the path into individual field names
	fieldNames := strings.Split(path, ".")

	// Call the recursive helper function
	return GetNestedFieldValueRecursive(val, fieldNames)
}

func GetNestedFieldValueRecursive(val reflect.Value, fieldNames []string) (interface{}, error) {
	// Base case: no more field names to process
	if len(fieldNames) == 0 {
		return val.Interface(), nil
	}

	// Iterate through the fields of the current struct
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)

		// Check if the field name matches the current path element
		if field.Name == fieldNames[0] {
			// Recursively call the function for the next level of nesting
			return GetNestedFieldValueRecursive(val.Field(i), fieldNames[1:])
		}
	}

	// If the field is not found, return an error
	return nil, errors.Wrap(ErrFieldNotFound, fieldNames[0])
}
