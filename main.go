package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// On définit la route racine "/" pour servir le fichier index.html
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	// On récupère le port (utile pour ton Procfile plus tard)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Port par défaut en local
	}

	fmt.Printf("Serveur lancé sur http://localhost:%s\n", port)

	// Lancement du serveur
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Erreur lors du lancement :", err)
	}
}
