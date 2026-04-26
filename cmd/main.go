package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"github.com/whalelogic/redis-go-quick/pkg/survey" 
)


// Keep these outside of 'main' to re-use
var ctx = context.Background()
var store *survey.Store
var templates *template.Template

func main() {
	var err error
	templates, err = template.ParseFiles("templates/base.html", "templates/surveys/fav-language.html")
	if err != nil {
		log.Fatalf("Couldn't parse templates: %v", err)
	}



	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}

	// Initialize our reusable wrapper
	store = survey.NewStore(redisAddr)

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/submit", submitHandler)

	fmt.Println("Server starting at :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Fetch results from Redis 
	results, _ := store.GetResults(ctx, "votes:languages")
	
	err := templates.ExecuteTemplate(w, "fav-language.html", results)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Use our helper to parse the form
	favLang := survey.ParseSurvey(r, "language")

	if favLang != "" {
		// Save to Redis using our wrapper
		store.SaveResponse(ctx, "votes:languages", favLang)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
