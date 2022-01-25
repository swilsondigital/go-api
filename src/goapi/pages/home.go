package home

import (
	"html/template"
	"net/http"
	"os"
	"path"
)

type Home struct {
	Title   string
	SiteURL string
}

/**
* display the homepage
 */
func ShowHomePage(w http.ResponseWriter, r *http.Request) {
	// get hostname
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	// get working directory
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// set home object
	home := Home{"Welcome", hostname}

	// setup template
	fp := path.Join(cwd, "src/goapi", "templates", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, home); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
