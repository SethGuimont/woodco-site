// go.mod: require github.com/go-chi/chi/v5
// main.go
package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type PageData struct {
	Title           string
	MetaDescription string
	Year            int
}

var tpl = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RealIP, middleware.RequestID, middleware.Logger, middleware.Recoverer)

	// Routes and pages with dynamic values
	r.Get("/", page("index.html", "Home", "Welcome to Woodco Prefinishing. "))
	r.Get("/services/", page("services.html", "Services", "Explore our prefinishing services,"))
	r.Get("/services/millwork/", page("millwork.html", "Millwork Prefinishing", "Professional Millwork finishing,"))
	r.Get("/services/doors/", page("doors.html", "Doors", "Prefinishing services for doors."))
	r.Get("/services/siding/", page("siding.html", "Siding", "Durable siding finishes."))
	r.Get("/services/paneling/", page("paneling.html", "Paneling", "Paneling prefinishing options."))
	r.Get("/services/custom-finishes/", page("custom-finishes.html", "Custom Finishes", "Custom finishing solutions."))
	r.Get("/services/custom-processes/", page("custom-processes.html", "Custom Processes", "Specialized finishing processes"))
	r.Get("/why-woodco/", page("why-woodco.html", "Why Woodco", "What makes Woodco different"))
	r.Get("/why-woodco/equipment/", page("equipment.html", "Equipment", "Our capabilities & equipment."))
	r.Get("/why-woodco/history/", page("history.html", "History", "Woodco history and heritage."))
	r.Get("/why-prefinishing/", page("why-prefinishing.html", "Why Prefinishing", "Benefits of prefinishing."))
	r.Get("/samples-brochures/", page("samples-brochures.html", "Sample & Brochures", "Request samples or brochures."))
	r.Get("/clients/", page("manufacturers.html", "Clients", "Manufactures we work with"))
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

func page(name, title, meta string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := PageData{
			Title:           title,
			MetaDescription: meta,
			Year:            time.Now().Year(),
		}
		if err := tpl.ExecuteTemplate(w, name, data); err != nil {
			log.Printf("template %s error: %v", name, err)
			http.Error(w, "template error", http.StatusInternalServerError)
		}
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
