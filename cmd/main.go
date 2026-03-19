package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"github.com/whalelogic/redis-go-quick/pkg/survey" 
)

var ctx = context.Background()
var store *survey.Store

// HTML Template for the survey
const tpl = `
<!DOCTYPE html>
<html>
<body>
	<h2>Quick Survey: What is your favorite Programming Language?</h2>
	<form method="POST" action="/submit">
		<input type="text" name="language" placeholder="e.g. Go, Python, Rust" required>
		<input type="submit" value="Vote">
	</form>
	<hr>
	<h3>Current Results:</h3>
	<ul>
		{{range .}} <li>{{.}}</li> {{end}}
	</ul>
</body>
</html>
`

func main() {
	// Initialize our reusable wrapper
	store = survey.NewStore("localhost:6379")

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/submit", submitHandler)

	fmt.Println("Server starting at :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Fetch results from Redis 
	results, _ := store.GetResults(ctx, "votes:languages")
	
	t := template.Must(template.New("web").Parse(tpl))
	t.Execute(w, results)
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
