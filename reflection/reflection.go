package reflection

import (
	"fmt"
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

// NewZeroValue returns a new zero value for a struct field.
func NewZeroValue(target interface{}, fieldName string) interface{} {
	targetValue := reflect.ValueOf(target)
	targetField := targetValue.Elem().FieldByName(fieldName)
	return reflect.Zero(targetField.Type()).Interface()
}

// GetFieldNames provides a simple api for retrieving struct fields via reflection.
func GetFieldNames(instance interface{}) []string {
	_type := reflect.TypeOf(instance)
	var fieldNames []string
	for index := 0; index < _type.NumField(); index++ {
		fieldNames = append(fieldNames, _type.Field(index).Name)
	}

	return fieldNames
}

// SetField sets a the value of an interface's member using a passed value and
// the target field's metadata.
func SetField(target interface{}, fieldName string, value interface{}) error {
	targetValue := reflect.ValueOf(target)
	targetField := targetValue.Elem().FieldByName(fieldName)

	if !targetField.IsValid() {
		message := fmt.Sprintf("target field %s for type %s is invalid", fieldName, targetValue.Type().Name())
		return cliErrors.NewWithMessage("reflection.SetField", message)
	}

	if isPrimitive(targetField.Type().Name()) {
		err := setPrimitiveField(target, fieldName, value)
		if err != nil {
			return err
		}
		return nil
	}

	targetField.Set(reflect.ValueOf(value))

	return nil
}

// IsPrimitive tests if a interface field is a primitive.
func IsPrimitive(target interface{}, fieldName string) bool {
	return isPrimitive(reflect.ValueOf(&target).Elem().FieldByName(fieldName).Type().Name())
}

// isPrimitive tests if a StructField is a primitive.
func isPrimitive(typeName string) bool {
	for _, primitive := range primitives {
		if typeName == primitive {
			return true
		}
	}
	return false
}

func setPrimitiveField(target interface{}, fieldName string, value interface{}) error {
	targetValue := reflect.ValueOf(&target)
	field := targetValue.Elem().FieldByName(fieldName)
	fieldKind := field.Kind()

	switch fieldKind {
	case reflect.Bool:
		field.SetBool(value.(bool))
	case reflect.Int:
		field.SetInt(value.(int64))
	case reflect.Int8:
		field.SetInt(value.(int64))
	case reflect.Int16:
		field.SetInt(value.(int64))
	case reflect.Int32:
		field.SetInt(value.(int64))
	case reflect.Int64:
		field.SetInt(value.(int64))
	case reflect.Uint:
		field.SetUint(value.(uint64))
	case reflect.Uint8:
		field.SetUint(value.(uint64))
	case reflect.Uint16:
		field.SetUint(value.(uint64))
	case reflect.Uint32:
		field.SetUint(value.(uint64))
	case reflect.Uint64:
		field.SetUint(value.(uint64))
	case reflect.Float32:
		field.SetFloat(value.(float64))
	case reflect.Float64:
		field.SetFloat(value.(float64))
	case reflect.Complex64:
		field.SetComplex(value.(complex128))
	case reflect.Complex128:
		field.SetComplex(value.(complex128))
	case reflect.String:
		field.SetString(value.(string))
	default:
		message := fmt.Sprintf("No match for field %s of Kind %s", fieldName, fieldKind.String())
		return cliErrors.NewWithMessage("reflection.setPrimitiveField", message)
	}

	return nil
}
