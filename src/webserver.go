package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".html"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".html"
	body, err := os.ReadFile("./pages/" + filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "%s", p.Body)
}

type SoundGroups struct {
	Name     string
	Packages []string
}

func PackageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	// database query
	groups := SoundGroups{"Sound Groups", []string{"sample beats", "sample instrumentals"}}

	if err := json.NewEncoder(w).Encode(&groups); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("js/")))
	http.HandleFunc("/packages/", PackageHandler)
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))

}
