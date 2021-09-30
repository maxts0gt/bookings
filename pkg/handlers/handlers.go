package handlers

import (
	"net/http"

	"github.com/maxts0gt/bookings/pkg/config"
	"github.com/maxts0gt/bookings/pkg/models"
	"github.com/maxts0gt/bookings/pkg/render"
)

//Repo repository used by the handlers
var Repo *Repository

// Repo is the repo type
type Repository struct {
	App *config.AppConfig
}

//New repo creates new repo
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//New handlers sets the repo for handlers
func NewHandlers(r *Repository) {
	Repo = r
}

//Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.gohtml", &models.TemplateData{})
}

//About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.gohtml", &models.TemplateData{
		StringMap: stringMap,
	})
}
