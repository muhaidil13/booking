package render

import (
	"encoding/gob"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/muhaidil13/booking/internal/config"
	"github.com/muhaidil13/booking/internal/model"
)

var appTest config.Appconfig
var session *scs.SessionManager

func TestMain(m *testing.M) {
	// Store data di session
	gob.Register(model.Reservation{})

	// Ganti ini jika diproducttion
	appTest.InProduct = false

	session = scs.New()
	session.Lifetime = 10 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	appTest.Session = session

	app = &appTest
	os.Exit(m.Run())
}

type mywrite struct{}

func (m *mywrite) Header() http.Header {
	var h http.Header
	return h
}
func (m *mywrite) Write(a []byte) (int, error) {
	length := len(a)
	return length, nil
}
func (m *mywrite) WriteHeader(i int) {

}
