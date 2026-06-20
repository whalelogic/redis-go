// Package handlers contains the HTTP handlers
package handlers

import (
	"net/http"
	"context"
	"html/template"

	"github.com/whalelogic/redis-go-quick/pkg/survey"
)


var ctx = context.Background()
var store *survey.Store
var templates *template.Template


func SetStore(s *survey.Store) {
	store = s
}

func SetTemplates(t *template.Template) {
	templates = t
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	results, _ := store.GetSurveyResults(ctx, "votes:languages")
	
	err := templates.ExecuteTemplate(w, "fav-language.html", results)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	favLang := survey.ParseSurvey(r, "language")
	if favLang != "" {
		store.SaveResponse(ctx, "votes:languages", favLang)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
