package reflection

import (
	"fmt"
	"reflect"

	"github.com/DerekStrickland/consul-setup-cli/errors"
	cliErrors "github.com/DerekStrickland/consul-setup-cli/errors"
)

var primitives = []string{
	"bool",
	"int",
	"int8",
	"int16",
	"int32",
	"int64",
	"uint",
	"uint8",
	"uint16",
	"uint32",
	"uint64",
	"uintptr",
	"float32",
	"float64",
	"complex64",
	"complex128",
	"string",
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

// GetStructField gets a struct field as an interface from another struct by interface.
func GetStructField(target interface{}, fieldName string) (interface{}, error) {
	if !isOfKind(target, []reflect.Kind{reflect.Struct, reflect.Ptr}) {
		return nil, cliErrors.NewWithMessage("reflection.GetStructField", "Cannot use GetStructField on a non-struct interface")
	}

	targetValue := reflectValue(target)
	field := targetValue.FieldByName(fieldName)
	if !field.IsValid() {
		return nil, cliErrors.NewWithMessage("reflect.GetStructField", fmt.Sprintf("No such field: %s in obj", fieldName))
	}

	return field.Interface(), nil
}

func isOfKind(target interface{}, types []reflect.Kind) bool {
	for _, _type := range types {
		if reflect.TypeOf(target).Kind() == _type {
			return true
		}
	}

	return false
}

func reflectValue(target interface{}) reflect.Value {
	var value reflect.Value

	if reflect.TypeOf(target).Kind() == reflect.Ptr {
		value = reflect.ValueOf(target).Elem()
	} else {
		value = reflect.ValueOf(target)
	}

	return value
}

// SetField sets a the value of an interface's member using a passed value and
// the target field's metadata.
func SetField(target interface{}, fieldName string, value interface{}) error {
	if isPrimitive(value) {
		err := setPrimitiveField(target, fieldName, value)
		if err != nil {
			return err
		}
		return nil
	}

	targetValue := reflect.ValueOf(target)
	targetField := targetValue.FieldByName(fieldName)

	if !targetField.IsValid() {
		message := fmt.Sprintf("target field %s for type %s is invalid", fieldName, targetValue.Type().Name())
		return cliErrors.NewWithMessage("reflection.SetField", message)
	}

	// if !targetField.CanSet() {
	// 	return cliErrors.NewWithMessage("reflection.SetField", fmt.Sprintf("Cannot set %s field value", fieldName))
	// }

	targetFieldType := targetField.Type()
	val := reflect.ValueOf(value)
	if targetFieldType != val.Type() {
		message := fmt.Sprintf("Provided value of type %s didn't match target field with name %s of type %s on interface %+v", val.Type(), fieldName, targetFieldType, target)
		err := errors.NewWithMessage("reflection.SetField", message)
		return err
	}

	targetField.Set(val)

	return nil
}

// IsPrimitive tests if a interface field is a primitive.
func IsPrimitive(target interface{}, fieldName string) bool {
	return isPrimitive(reflect.ValueOf(&target).Elem().FieldByName(fieldName).Type().Name())
}

// isPrimitive tests if a StructField is a primitive.
func isPrimitive(value interface{}) bool {
	typeName := reflect.TypeOf(value).Name()
	fmt.Println("isPrimitive: typeName is " + typeName)
	for _, primitive := range primitives {
		if typeName == primitive {
			return true
		}
	}
	return false
}

func setPrimitiveField(target interface{}, fieldName string, value interface{}) error {
	targetValue := reflect.ValueOf(target)
	fmt.Println(fmt.Sprintf("setPrimitiveField: targetValue is %+v", targetValue))
	field := targetValue.FieldByName(fieldName)
	fmt.Println(fmt.Sprintf("setPrimitiveField: field %s is %+v", fieldName, field))
	fieldKind := field.Kind()
	fmt.Println(fmt.Sprintf("setPrimitiveField: fieldKind is %+v", fieldKind))

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
