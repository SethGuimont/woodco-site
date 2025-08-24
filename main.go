// go.mod: require github.com/go-chi/chi/v5
// main.go
package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var tpl = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RealIP, middleware.RequestID, middleware.Logger, middleware.Recoverer)
	r.Get("/", page("index.html"))
	r.Get("/services/", page("services.html"))
	r.Get("/services/millwork/", page("millwork.html"))
	r.Get("/services/doors/", page("doors.html"))
	r.Get("/services/siding/", page("siding.html"))
	r.Get("/services/paneling/", page("paneling.html"))
	r.Get("/services/custom-finishes/", page("custom-finishes.html"))
	r.Get("/services/custom-processes/", page("custom-processes.html"))
	r.Get("/why-woodco/", page("why-woodco.html"))
	r.Get("/why-woodco/equipment/", page("equipment.html"))
	r.Get("/why-woodco/history/", page("history.html"))
	r.Get("/why-prefinishing/", page("why-prefinishing.html"))
	r.Get("/samples-brochures/", page("samples-brochures.html"))
	r.Get("/clients/manufacturers/", page("clients/manufacturers.html"))
	r.Get("/contact-us/", contactGET)
	r.Post("/contact-us/", contactPOST) // HTMX-friendly
	// Legacy redirects
	r.Get("/services/index.shtml", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/services/", http.StatusMovedPermanently)
	})
	r.Get("/services/availability.shtml", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/services/", http.StatusMovedPermanently)
	})
	r.Get("/video-tour.html", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	})
	// static assets
	fs := http.FileServer(http.Dir("public"))
	r.Handle("/assets/*", http.StripPrefix("/assets/", fs))
	log.Fatal(http.ListenAndServe(":8080", r))
}
func page(name string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.ExecuteTemplate(w, name, nil)
	}
}
func contactGET(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "contact/form.html", map[string]any{"Errors": nil, "OK": false})
}
func contactPOST(w http.ResponseWriter, r *http.Request) {
	// TODO: validate, rate-limit, send email
	if r.Header.Get("HX-Request") == "true" {
		// return fragment that replaces the form
		tpl.ExecuteTemplate(w, "contact/success.partial.html", nil)
		return
	}
	http.Redirect(w, r, "/contact-us/?ok=1", http.StatusSeeOther)
}
