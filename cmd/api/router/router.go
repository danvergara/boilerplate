package router

import (
	"github.com/danvergara/boilerplate/cmd/api/handlers/getuser"
	"github.com/danvergara/boilerplate/pkg/application"
	"github.com/julienschmidt/httprouter"
)

// Get return the router (mux)
func Get(app *application.Application) *httprouter.Router {
	mux := httprouter.New()
	mux.GET("/users", getuser.Do(app))
	return mux
}
