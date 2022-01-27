package pages

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

	// get url
	url := r.Host

	// get working directory
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// set home object
	home := Home{"Welcome", url}

	// setup template
	fp := path.Join(cwd, "templates", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, home); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
