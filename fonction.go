package main

import (
	"html/template"
	"net/http"
	"web/etat"
)

type Data struct {
	Output string
}

// creation d un serveur avec le port 8080
const port = ":8080"

func Index(w http.ResponseWriter, r *http.Request) {
	data := Data{}
	da := ""
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		http.ServeFile(w, r, "./templetes/404.html")
		return
	}
	testTemplet, err := template.ParseFiles("./templetes/index.html")
	if err != nil {
		http.Error(w, "Error while parsing the HTML file", http.StatusInternalServerError)
		http.ServeFile(w, r, "./templetes/500.html")
		return
	}

	if r.Method == http.MethodGet {
		testTemplet.Execute(w, data)
	} else if r.Method == http.MethodPost {
		d := r.FormValue("entree")
		banner := r.FormValue("web")
		if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
			w.WriteHeader(http.StatusInternalServerError)
			http.ServeFile(w, r, "./templetes/500.html")
			return
		}
		// Le texte ne contient pas de caractères spéciaux
		// http.Error(w, "pas de caractere special", http.StatusInternalServerError)
		if !Special(d) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "./templetes/400.html")
			return
		}
		// si la longueur de la chaine est vide
		// http.Error(w, "", http.StatusInternalServerError)
		if len(d) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "./templetes/400.html")
			return
		}

		da, err = etat.Ascifs(d, banner)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			http.ServeFile(w, r, "./templetes/500.html")
			return
		}
		data.Output = da
		testTemplet.Execute(w, data)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		http.ServeFile(w, r, "./templetes/400.html")
		return
	}
}
func Special(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] >= ' ' && s[i] <= '~' {
			return true
		}
	}
	return false
}
