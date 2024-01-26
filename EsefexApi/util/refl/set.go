package refl

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/pkg/errors"
)

func SetNestedFieldValue(obj interface{}, path string, value interface{}) error {
	val := reflect.ValueOf(obj)

	// Navigate through the nested fields
	for val.Kind() == reflect.Ptr || val.Kind() == reflect.Interface {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return fmt.Errorf("provided object is not a struct")
	}

	// Split the path into individual field names
	fieldNames := strings.Split(path, ".")

	// Call the recursive helper function
	return SetNestedFieldValueRecursive(val, fieldNames, value)
}

func SetNestedFieldValueRecursive(val reflect.Value, fieldNames []string, newValue interface{}) error {
	// Base case: no more field names to process
	if len(fieldNames) == 0 {
		// Set the value at the final field
		valSettable := reflect.ValueOf(newValue)
		if val.CanSet() && valSettable.Type().AssignableTo(val.Type()) {
			val.Set(valSettable)
			return nil
		}
		return fmt.Errorf("cannot set value at path")
	}

	// Iterate through the fields of the current struct
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)

		// Check if the field name matches the current path element
		if field.Name == fieldNames[0] {
			// Recursively call the function for the next level of nesting
			return SetNestedFieldValueRecursive(val.Field(i), fieldNames[1:], newValue)
		}
	}

	// If the field is not found, return an error
	return errors.Wrap(ErrFieldNotFound, fieldNames[0])
}
