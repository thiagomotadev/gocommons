package reflection

import (
	"fmt"
	"reflect"
)

func GetID(model interface{}) (id int64, err error) {
	fieldType, fieldValue, err := fieldValueOf(model, "ID")
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

func CallFunc(function interface{}, inputs map[reflect.Type]interface{}) (results []reflect.Value) {
	funcValue := reflect.ValueOf(function)
	results = callFunc(funcValue, inputs)
	return
}

func CallFuncByName(model interface{}, name string, inputs map[reflect.Type]interface{}) (results []reflect.Value) {
	funcValue := reflect.ValueOf(model).MethodByName(name)
	results = CallFunc(funcValue, inputs)
	return
}
