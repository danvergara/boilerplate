package getuser

import (
	"fmt"
	"net/http"

	"github.com/danvergara/boilerplate/pkg/application"
	"github.com/julienschmidt/httprouter"
)

// Do return return the users
func Do(app *application.Application) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprintf(w, "hello")
	}
}
