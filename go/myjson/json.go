package myjson

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// Converts a struct to JSON consisted of type,value field for each public
// member of the struct, field name is taken from the member's json tag
func StructToJSON(data interface{}) []byte {
	// a new Value initialized to the concrete value stored in the interface i
	structValue := reflect.ValueOf(data)
	// Type returns v's type, which is the representation of a Go type.
	structType := structValue.Type()

	// using buffer - strings in golang are immutable
	buf := bytes.NewBuffer([]byte{'{'})

	// iterate through struct fields
	nfields := structValue.NumField()
	for i := 0; i < nfields; i++ {
		field := fieldToJSON(
			structType.Field(i),  // StructField
			structValue.Field(i), // Value
		)
		// append field to the buffer
		buf.WriteString(field)

		// add ',' if not last element
		if i != (nfields - 1) {
			buf.WriteString(",")
		}
	}
	buf.Write([]byte{'}'}) // add closing curly '}'
	return buf.Bytes()
}

func fieldToJSON(field reflect.StructField, value reflect.Value) string {
	// name := field.Name
	name := field.Tag.Get("json")
	return fmt.Sprintf(
		`"%s":{"type":"%s","value":%v}`,
		name, strings.ToLower(field.Type.Name()), valueToString(value.Interface()),
	)
}

func valueToString(value any) string {
	switch val := value.(type) {
	case int:
		return strconv.Itoa(val)
	case string:
		return fmt.Sprintf(`"%s"`, val)
	default:
		panic(fmt.Sprintf("unsupported type: '%s'", reflect.TypeOf(val)))
	}
}
