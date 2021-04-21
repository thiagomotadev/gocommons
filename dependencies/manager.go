package dependencies

import (
	"reflect"

	"github.com/thiagomotadev/gocommons/reflection"
)

type Manager struct {
	dependencies map[reflect.Type]interface{}
}

func (manager *Manager) Add(dependency interface{}) {
	reflection.CallFuncByName(dependency, "Init", manager.dependencies)
	manager.dependencies[reflect.TypeOf(dependency)] = dependency
}

func (manager Manager) Get(reflectType reflect.Type) interface{} {
	return manager.dependencies[reflectType]
}

func (manager Manager) GetAll() map[reflect.Type]interface{} {
	return manager.dependencies
}
