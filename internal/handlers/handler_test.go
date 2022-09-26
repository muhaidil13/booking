package handlers

import (
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key   string
	value string
}

var theTest = []struct {
	name       string
	url        string
	method     string
	params     []postData
	statuscode int
}{
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"about", "/about", "GET", []postData{}, http.StatusOK},
	{"contack", "/contact", "GET", []postData{}, http.StatusOK},
	{"room", "/room", "GET", []postData{}, http.StatusOK},
	{"res-result", "/reservation-result", "GET", []postData{}, http.StatusOK},
	{"make-reservation", "/make-reservation", "GET", []postData{}, http.StatusOK},
	{"pemesanan", "/pemesannan", "GET", []postData{}, http.StatusOK},

	{"reservation", "/reservation", "POST", []postData{
		{key: "start-date1", value: "12-12-2000"},
		{key: "end-date1", value: "12-12-2001"},
	}, http.StatusOK},

	{"reservation", "/reservation-json", "POST", []postData{
		{key: "start-date1", value: "12-12-2000"},
		{key: "end-date1", value: "12-12-2001"},
	}, http.StatusOK},
	{"make-reservation", "/make-reservation", "POST", []postData{
		{key: "nama_pertama", value: "Jhony Sins"},
		{key: "nama_belakang", value: "bobbys"},
		{key: "email", value: "bobbys@gmail.com"},
		{key: "phone", value: "019191019019"},
	}, http.StatusOK},
}

func TestHandler(t *testing.T) {
	route := getRoute()
	tes := httptest.NewTLSServer(route)

	defer tes.Close()

	// test
	for _, v := range theTest {
		if v.method == "GET" {
			resp, err := tes.Client().Get(tes.URL + v.url)
			if err != nil {
				t.Error("error di", err)
				log.Fatal(err)
			}

			if resp.StatusCode != v.statuscode {
				t.Errorf("Status Tidak Sama %d yang diperlukan %d", v.statuscode, resp.StatusCode)
				log.Fatal(err)
			}
		} else {
			values := url.Values{}
			for _, x := range v.params {
				values.Add(x.key, x.value)
			}
			resp, err := tes.Client().PostForm(tes.URL+v.url, values)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			if resp.StatusCode != v.statuscode {
				t.Errorf("for %s expected %d but got %d", v.name, v.statuscode, resp.StatusCode)
			}
		}
	}
}
