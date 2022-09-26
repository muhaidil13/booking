package helpers

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/muhaidil13/booking/internal/config"
)

var app *config.Appconfig

// set Helpers
func NewHelpers(a *config.Appconfig) {
	app = a
}

func ClientError(w http.ResponseWriter, status int) {
	app.InfoLog.Println("Client Status ", status)
	http.Error(w, http.StatusText(status), status)

}

func ServerError(w http.ResponseWriter, err error) {
	// print error dan trace error line yang mengalami error
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
