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

type groups struct {
	Name     string
	Packages []string
}

func PackageViewHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	// database query

	audio_packages_fp, err := os.Open("C:/Users/deros/rwebserver/audio-data-packages")

	if err != nil {
		fmt.Println(err)
	}
	audio_packages, err := audio_packages_fp.ReadDir(0)
	audio_package_names := []string{}
	for p := 0; p < len(audio_packages); p++ {
		audio_package_names = append(audio_package_names, audio_packages[p].Name())
	}
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(audio_package_names)

	groups := groups{"packages", audio_package_names}

	if err := json.NewEncoder(w).Encode(&groups); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func FileViewHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	target := r.Header.Get("package-name")
	fmt.Println("target", target)
	fp, err := os.Open("C:/Users/deros/rwebserver/audio-data-packages/" + target)
	if err != nil {
		fmt.Println(err)
	}
	audio_files, err := fp.ReadDir(0)
	audio_file_names := []string{}
	for p := 0; p < len(audio_files); p++ {
		audio_file_names = append(audio_file_names, audio_files[p].Name())
	}
	if err != nil {
		fmt.Println(err)
	}

	groups := groups{"package-content", audio_file_names}
	if err := json.NewEncoder(w).Encode(&groups); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func main() {
	http.Handle("/", http.FileServer(http.Dir("js/")))
	http.HandleFunc("/packages/", PackageViewHandler)
	http.HandleFunc("/content/", FileViewHandler)
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))

}
