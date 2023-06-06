package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Définir le gestionnaire pour les fichiers statiques
	fs := http.FileServer(http.Dir("css"))
	// Assurez-vous que le gestionnaire de fichiers statiques est utilisé pour les requêtes commençant par "/static/"
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	http.HandleFunc("/", Index)
	fmt.Println("(http://localhost:8080)--Server started on port ", port)
	//lancement du serveur
	http.ListenAndServe(port, nil)
}
