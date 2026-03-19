// Package survey provides functions to parse survey forms and a database interface [db.go].
package survey

import (
	"net/http"
	"strings"
)

// ParseSurvey extracts data from an r.PostForm
func ParseSurvey(r *http.Request, field string) string {
	val := r.FormValue(field)
	return strings.TrimSpace(val)
}
