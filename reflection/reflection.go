package reflection

import (
	"reflect"

	cliErrors "github.com/DerekStrickland/consul-setup-cli/errors"
)

var primitives = []string{
	"Bool",
	"Int",
	"Int8",
	"Int16",
	"Int32",
	"Int64",
	"Uint",
	"Uint8",
	"Uint16",
	"Uint32",
	"Uint64",
	"Uintptr",
	"Float32",
	"Float64",
	"Complex64",
	"Complex128",
	"String",
	"UnsafePointer",
}

// NewFromStructField returns a new reflect value from a StructField definition.
func NewFromStructField(field reflect.StructField) reflect.Value {
	return reflect.New(field.Type)
}

// GetFields provides a simple api for retrieving struct fields via reflection.
func GetFields(instance interface{}) []reflect.StructField {
	_type := reflect.TypeOf(instance)
	var fields []reflect.StructField
	for index := 0; index < _type.NumField(); index++ {
		fields = append(fields, _type.FieldByIndex([]int{index}))
	}
	return fields
}

// SetField sets a the value of an interface's member using a passed value and
// the target field's metadata.
func SetField(target *interface{}, val *interface{}, field reflect.StructField) error {
	if IsPrimitive(field) {
		return cliErrors.NewWithMessage("reflect.SetField", "Field is primitive. Use SetPrimitiveField instead.")
	}

	targetValue := reflect.ValueOf(&target)
	targetField := targetValue.Elem().FieldByName(field.Name)
	if targetField.IsValid() {
		targetField.Set(targetValue)
	}

	return nil
}

// IsPrimitive tests if a StructField is a primitive type.
func IsPrimitive(field reflect.StructField) bool {
	for _, primitive := range primitives {
		if field.Type.Name() == primitive {
			return true
		}
	}
	return false
}

// SetPrimitiveField sets primitive fields on an reflect Value for primitive types.
func SetPrimitiveField(field reflect.Value, fieldKind reflect.Kind, valueToApply interface{}) {
	switch fieldKind {
	case reflect.Bool:
		field.SetBool(valueToApply.(bool))
	case reflect.Int:
		field.SetInt(valueToApply.(int64))
	case reflect.Int8:
		field.SetInt(valueToApply.(int64))
	case reflect.Int16:
		field.SetInt(valueToApply.(int64))
	case reflect.Int32:
		field.SetInt(valueToApply.(int64))
	case reflect.Int64:
		field.SetInt(valueToApply.(int64))
	case reflect.Uint:
		field.SetUint(valueToApply.(uint64))
	case reflect.Uint8:
		field.SetUint(valueToApply.(uint64))
	case reflect.Uint16:
		field.SetUint(valueToApply.(uint64))
	case reflect.Uint32:
		field.SetUint(valueToApply.(uint64))
	case reflect.Uint64:
		field.SetUint(valueToApply.(uint64))
	case reflect.Float32:
		field.SetFloat(valueToApply.(float64))
	case reflect.Float64:
		field.SetFloat(valueToApply.(float64))
	case reflect.Complex64:
		field.SetComplex(valueToApply.(complex128))
	case reflect.Complex128:
		field.SetComplex(valueToApply.(complex128))
	case reflect.String:
		field.SetString(valueToApply.(string))
	}
}
