package reflection

import (
	"fmt"
	"reflect"
)

func GetFieldTypeAndValue(model interface{}, fieldName string) (fieldType reflect.Type, fieldValue interface{}, err error) {
	e := reflect.ValueOf(model).Elem()
	for i := 0; i < e.NumField(); i++ {
		if fieldName == e.Type().Field(i).Name {
			fieldType = e.Type().Field(i).Type
			fieldValue = e.Field(i).Interface()
			return
		}
	}
	err = fmt.Errorf(
		`field "%s" not found in model "%v"`,
		fieldName,
		GetTypeName(model),
	)
	return
}

func GetID(model interface{}) (id int64, err error) {
	fieldType, fieldValue, err := GetFieldTypeAndValue(model, "ID")
	if err != nil {
		return
	}
	if reflect.TypeOf(int64(0)) != fieldType {
		err = fmt.Errorf(`the "%v" model ID field is not of type int64`, GetTypeName(model))
	}
	id = fieldValue.(int64)
	return
}

func GetTypeName(typeInterface interface{}) string {
	t := reflect.TypeOf(typeInterface)
	if t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	}
	return t.Name()
}

func CallFunc(function interface{}, inputs map[reflect.Type]reflect.Value) (results []reflect.Value) {
	funcValue := reflect.ValueOf(function)
	numberOfInputs := funcValue.Type().NumIn()

	funcInputs := make([]reflect.Value, numberOfInputs)

	for index := 0; index < numberOfInputs; index++ {
		inputType := funcValue.Type().In(index)
		funcInputs[index] = inputs[inputType]
	}
	results = funcValue.Call(funcInputs)
	return
}

func CallMethodByName(model interface{}, name string, inputs map[reflect.Type]reflect.Value) (results []reflect.Value) {
	funcValue := reflect.ValueOf(model).MethodByName(name)
	numberOfInputs := funcValue.Type().NumIn()

	funcInputs := make([]reflect.Value, numberOfInputs)

	for index := 0; index < numberOfInputs; index++ {
		inputType := funcValue.Type().In(index)
		funcInputs[index] = inputs[inputType]
	}
	results = funcValue.Call(funcInputs)
	return
}
