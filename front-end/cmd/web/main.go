package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		render(w, "test.page.tmpl.html")
	})

	fmt.Println("Starting front end service on port 3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Panic(err)
	}
}

func render(w http.ResponseWriter, t string) {
	partials := []string{
		"./cmd/web/templates/base.layout.tmpl.html",
		"./cmd/web/templates/header.partial.tmpl.html",
		"./cmd/web/templates/footer.partial.tmpl.html",
	}

	var templateSlice []string
	templateSlice = append(templateSlice, fmt.Sprintf("./cmd/web/templates/%s", t))

	templateSlice = append(templateSlice, partials...)

	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// var data struct {
	// 	BrokerURL string
	// }

	// data.BrokerURL = os.Getenv("BROKER_URL")

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
