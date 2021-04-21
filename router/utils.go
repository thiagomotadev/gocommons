package router

import (
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
	"github.com/thiagomotadev/gocommons/reflection"
)

func routerHandler(router *Router, handleFunc interface{}) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		dependencies := router.manager.GetAll()
		dependencies[reflect.TypeOf(Vars{})] = Vars{Value: mux.Vars(r)}
		dependencies[reflect.TypeOf(OptionalVars{})] = OptionalVars{Value: r.URL.Query()}

		handlerResult := reflection.CallFunc(handleFunc, dependencies)

		result := handlerResult[0].Interface().(Result)
		routerErr := handlerResult[1].Interface().(Error)

		if routerErr.Err != nil {
			jsonModel, err := json.Marshal(jsonError{
				Status:  "Error",
				Message: routerErr.Message,
			})

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.WriteHeader(routerErr.StatusCode)
			w.Write(jsonModel)

			return
		}

		jsonModel, err := json.Marshal(result.Model)

		if err != nil {
			return
		}

		w.WriteHeader(result.StatusCode)
		w.Write(jsonModel)
	}
}
