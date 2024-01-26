package refl

import "reflect"

func FindAllPaths(obj interface{}) []string {
	paths := make([]string, 0)
	FindAllPathsRecursive(reflect.ValueOf(obj), "", &paths)
	return paths
}

func FindAllPathsRecursive(val reflect.Value, currentPath string, paths *[]string) {
	// Navigate through the nested fields
	for val.Kind() == reflect.Ptr || val.Kind() == reflect.Interface {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return
	}

	// Iterate through the fields of the current struct
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)

		// Construct the path for the current field
		fieldPath := currentPath + field.Name

		// Add a dot separator if not the first field in the path
		if currentPath != "" {
			fieldPath = currentPath + "." + field.Name
		}

		// Recursively call the function for the next level of nesting
		FindAllPathsRecursive(val.Field(i), fieldPath, paths)

		// Add the path to the result if the field is not a struct
		if val.Field(i).Kind() != reflect.Struct {
			*paths = append(*paths, fieldPath)
		}
	}
}
