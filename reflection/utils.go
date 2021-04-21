package reflection

import (
	"fmt"
	"reflect"
)

func fieldValueOf(model interface{}, fieldName string) (fieldType reflect.Type, fieldValue interface{}, err error) {
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

func callFunc(funcValue reflect.Value, inputs map[reflect.Type]interface{}) (results []reflect.Value) {
	numberOfInputs := funcValue.Type().NumIn()

	funcInputs := make([]reflect.Value, numberOfInputs)

	for index := 0; index < numberOfInputs; index++ {
		inputType := funcValue.Type().In(index)
		funcInputs[index] = reflect.ValueOf(inputs[inputType])
	}
	results = funcValue.Call(funcInputs)
	return
}
