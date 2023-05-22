package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/webcognito/kalkulacka_app/cmd/electric"
	"github.com/webcognito/kalkulacka_app/cmd/gas"
)

func main() {
	log.Print("starting server...")

	fs := http.FileServer(http.Dir("web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", index)
	http.HandleFunc("/pricelist", pricelist)

	http.HandleFunc("/einput", electric.Einput)
	http.HandleFunc("/eoutput", electric.Eoutput)

	http.HandleFunc("/ginput", gas.Ginput)
	http.HandleFunc("/goutput", gas.Goutput)

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	// Start HTTP server.
	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("web/templates/index.html", "web/templates/content_index.html")
	t.ExecuteTemplate(w, "index", "")
}

func pricelist(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("web/templates/index.html", "web/templates/content_pricelist.html")
	t.ExecuteTemplate(w, "index", "")
}

func ginput(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("web/templates/index.html", "web/templates/content_ginput.html")
	t.ExecuteTemplate(w, "index", "")
}
