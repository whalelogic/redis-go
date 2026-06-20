package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"github.com/whalelogic/redis-go-quick/pkg/survey" 
	"github.com/whalelogic/redis-go-quick/pkg/handlers" 
)


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

	store = survey.NewStore(redisAddr)

	handlers.SetStore(store)
	handlers.SetTemplates(templates)

	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/submit", handlers.SubmitHandler)

	fmt.Println("Server starting at :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

