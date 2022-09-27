package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/muhaidil13/booking/internal/config"
	"github.com/muhaidil13/booking/internal/forms"
	"github.com/muhaidil13/booking/internal/helpers"
	"github.com/muhaidil13/booking/internal/model"
	"github.com/muhaidil13/booking/internal/render"
)

// repo use by handler
var Repo *Repository

// Repository Type
type Repository struct {
	App *config.Appconfig
}

// New Repo Create New Repository handler
func NewRepo(a *config.Appconfig) *Repository {
	return &Repository{
		App: a,
	}
}

// set repository for hanlder
func NewHandler(r *Repository) {
	Repo = r
}

// Handler Request Get
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	StringMap := make(map[string]string)
	StringMap["users"] = "hello users"
	StringMap["test"] = "hello user"
	render.RenderFile(w, r, "home.page.html", &model.TemplateData{
		StringMap: StringMap,
	})
}
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderFile(w, r, "about.page.html", &model.TemplateData{})
}
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderFile(w, r, "contact.page.html", &model.TemplateData{})
}
func (m *Repository) Room(w http.ResponseWriter, r *http.Request) {
	render.RenderFile(w, r, "room.page.html", &model.TemplateData{})
}
func (m *Repository) Pemesannan(w http.ResponseWriter, r *http.Request) {
	render.RenderFile(w, r, "pemesannan.page.html", &model.TemplateData{})
}
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderFile(w, r, "reservation.page.html", &model.TemplateData{})
}
func (m *Repository) Make_Reservation(w http.ResponseWriter, r *http.Request) {
	var emptyreservation model.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyreservation
	render.RenderFile(w, r, "make-reservation.page.html", &model.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}
func (m *Repository) PostMake_Reservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		// Logging Handing error
		helpers.ServerError(w, err)
		return
	}
	reservation := model.Reservation{
		FirstName: r.Form.Get("nama_pertama"),
		LastName:  r.Form.Get("nama_belakang"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
	}
	form := forms.New(r.PostForm)
	// form.Has("nama_pertama", r)
	form.Required("nama_pertama", "nama_belakang", "email", "phone")
	form.MinLength("nama_pertama", 5)
	form.IsEmail("email")
	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation
		render.RenderFile(w, r, "make-reservation.page.html", &model.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}
	m.App.Session.Put(r.Context(), "reservation", reservation)
	http.Redirect(w, r, "/reservation-result", http.StatusSeeOther)

}

// handler request post
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start-dates")
	end := r.Form.Get("end-dates")
	w.Write([]byte(fmt.Sprintf("Start %v, end is %v", start, end)))
}

type jsonResponse struct {
	Status  bool   `json:"Status"`
	Message string `json:"Message"`
}

// Handle Availability with json
func (m *Repository) AvailabilityJson(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		Status:  true,
		Message: "Availability",
	}
	valby, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(valby)

}

func (m *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(model.Reservation)
	if !ok {
		m.App.ErrorLog.Println("Cant Get Session form session")
		m.App.Session.Put(r.Context(), "error", "Cant Get Reservation on Session")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	// m.App.Session.Remove(r.Context(), "reservation")
	data := make(map[string]interface{})
	data["reservation"] = reservation
	render.RenderFile(w, r, "reservation.summary.page.html", &model.TemplateData{
		Data: data,
	})
}
