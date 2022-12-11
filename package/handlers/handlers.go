package handlers

import (
	"fmt"
	"net/http"

	"github.com/sagarshinganwade/booking_and_reservation/package/config"
	"github.com/sagarshinganwade/booking_and_reservation/package/models"
	"github.com/sagarshinganwade/booking_and_reservation/package/render"
)

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// Repo the Repository used by handlers
var Repo *Repository

// NewRepository creates new repository
func NewRepository(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for handlers.
func NewHandlers(r *Repository) {
	Repo = r
}

// Home Handler for Home Page.
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	fmt.Println("remote IP: ", remoteIP)
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderPage(w, "home.page.tmpl", &models.TemplateData{})
}

// About Handler for about page.
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	strMap := make(map[string]string)
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	strMap["remote_ip"] = remoteIP
	render.RenderPage(w, "about.page.tmpl", &models.TemplateData{
		StrMap: strMap,
	})
}

func (m *Repository) Login(w http.ResponseWriter, r *http.Request) {
	strMap := make(map[string]string)
	strMap["test"] = "Hello, again"
	render.RenderPage(w, "login.page.tmpl", &models.TemplateData{StrMap: strMap})
}
