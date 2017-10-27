package main

import (
	"net/http"
	"github.com/rihtim/core"
	"github.com/rihtim/core/log"
	"github.com/rihtim/core/methods"
	"github.com/rihtim/core/interceptors"
	"github.com/rihtim/samples/basic/mydataprovider"
	"github.com/rihtim/samples/basic/myinterceptors"
)

func main() {

	log.Info("Starting server")

	// set data provider
	core.DataProvider = &mydataprovider.Provider{}

	// add interceptors
	core.Interceptors.Add(interceptors.AnyPath, methods.Get, interceptors.BEFORE_EXEC, myinterceptors.CheckPath, nil)
	core.Interceptors.Add(interceptors.AnyPath, methods.Any, interceptors.BEFORE_EXEC, myinterceptors.CheckIfModifiedSinceHeader, nil)
	core.Interceptors.Add("/wrong", methods.Get, interceptors.BEFORE_EXEC, myinterceptors.CheckPath, nil)
	core.Interceptors.Add("/users", methods.Post, interceptors.BEFORE_EXEC, myinterceptors.ValidateUserInput, nil)
	core.Interceptors.Add("/users/{id}", methods.Put, interceptors.AFTER_EXEC, myinterceptors.AddChangedFieldInfoToResponse, nil)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// forward http request to rihtim core
		core.HandleHttpRequest(w, r)
	})

	http.ListenAndServe(":8080", nil)
}
