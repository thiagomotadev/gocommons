package dependencies

import (
	"reflect"

	"github.com/thiagomotadev/gocommons/reflection"
)

type Manager struct {
	dependencies map[reflect.Type]reflect.Value
}

func (manager *Manager) Init() {
	manager.dependencies = make(map[reflect.Type]reflect.Value)
}

func (manager *Manager) InitWithOtherManager(otherManager *Manager) {
	manager.dependencies = otherManager.dependencies
}

func (manager *Manager) Add(dependency interface{}) {
	manager.CallMethodByName(dependency, "Init")
	manager.dependencies[reflect.TypeOf(dependency)] = reflect.ValueOf(dependency)
}

func (manager Manager) Get(reflectType reflect.Type) interface{} {
	return manager.dependencies[reflectType]
}

func (manager Manager) CallFunc(function interface{}) []reflect.Value {
	return reflection.CallFunc(function, manager.dependencies)
}

func (manager Manager) CallMethodByName(model interface{}, name string) []reflect.Value {
	return reflection.CallMethodByName(model, name, manager.dependencies)
}
