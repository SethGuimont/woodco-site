package main

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
	"time"
)

type PageData struct {
	Title, MetaDescription string
	Year                   int
}

func write(t *template.Template, tplName, out string, pd PageData) {
	dst := filepath.Join("dist", out)
	if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
		log.Fatal(err)
	}
	f, err := os.Create(dst)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := t.ExecuteTemplate(f, tplName, pd); err != nil {
		log.Fatal(err)
	}
}
func main() {
	t := template.Must(template.ParseGlob("templates/*.html")) // or *.gohtml
	y := time.Now().Year()
	// Home + sections (match your actual filenames)
	write(t, "index.html", "index.html", PageData{"Home", "Welcome to Woodco Prefinishing.", y})
	write(t, "services.html", "services/index.html", PageData{"Services", "Explore our prefinishing services.", y})
	write(t, "millwork.html", "services/millwork/index.html", PageData{"Millwork", "Professional millwork finishing.", y})
	write(t, "doors.html", "services/doors/index.html", PageData{"Doors", "Prefinishing services for doors.", y})
	write(t, "siding.html", "services/siding/index.html", PageData{"Siding", "Durable siding finishes.", y})
	write(t, "paneling.html", "services/paneling/index.html", PageData{"Paneling", "Paneling prefinishing options.", y})
	write(t, "custom-finishes.html", "services/custom-finishes/index.html", PageData{"Custom Finishes", "Custom finishing solutions.", y})
	write(t, "custom-processes.html", "services/custom-processes/index.html", PageData{"Custom Processes", "Specialized finishing processes.", y})
	write(t, "equipment.html", "why-woodco/equipment/index.html", PageData{"Equipment", "Our capabilities & equipment.", y})
	write(t, "history.html", "why-woodco/history/index.html", PageData{"History", "Woodco history and heritage.", y})
	write(t, "why-prefinishing.html", "why-prefinishing/index.html", PageData{"Why Prefinishing", "Benefits of prefinishing.", y})
	write(t, "samples-brochures.html", "samples-brochures/index.html", PageData{"Samples & Brochures", "Request samples or brochures.", y})
	write(t, "manufacturers.html", "clients/index.html", PageData{"Clients", "Manufacturers we work with.", y})
	// Contact & Thank You pages (static)
	write(t, "contact.html", "contact-us/index.html", PageData{"Contact Us", "Get in touch with Woodco Prefinishing.", y})
	// a simple thanks page you create at templates/thanks.html
	//write(t, "thanks.html", "thanks/index.html", PageData{"Thanks", "We received your message.", y})
}
