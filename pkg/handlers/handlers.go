package handlers

import (
	"net/http"

	"github.com/hmdrzaa11/hello-world/pkg/config"
	"github.com/hmdrzaa11/hello-world/pkg/models"
	"github.com/hmdrzaa11/hello-world/pkg/render"
)

// Repo the repository used by handlers
var Repo *Repository

// Repository its the Repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the "Repo" variable
func NewHandlers(r *Repository) {
	Repo = r
}

func (repo *Repository) Home(w http.ResponseWriter, r *http.Request) {
	ip := r.RemoteAddr
	//store it in session
	repo.App.Session.Put(r.Context(), "remote_ip", ip)

	render.RenderTemplate(w, "home.page.gohtml", &models.TemplateData{})
}

func (repo *Repository) About(w http.ResponseWriter, r *http.Request) {
	remoteIP := repo.App.Session.GetString(r.Context(), "remote_ip")
	stringMap := map[string]string{
		"test": "another silly text",
		"ip":   remoteIP,
	}
	render.RenderTemplate(w, "about.page.gohtml", &models.TemplateData{StringMap: stringMap})
}
