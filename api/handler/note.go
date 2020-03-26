package handler

import (
	"encoding/json"
	"github.com/rithikjain/CleanNotesApi/api/middleware"
	"github.com/rithikjain/CleanNotesApi/api/view"
	"github.com/rithikjain/CleanNotesApi/pkg/note"
	"net/http"
)

// Protected Request
func createNote(svc note.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			view.Wrap(view.ErrMethodNotAllowed, w)
			return
		}
		note := &note.Note{}
		err := json.NewDecoder(r.Body).Decode(note)
		if err != nil {
			view.Wrap(err, w)
			return
		}
		claims, err := middleware.ValidateAndGetClaims(r.Context(), "user")
		if err != nil {
			view.Wrap(err, w)
			return
		}
		note.UserID = claims["id"].(float64)
		n, err := svc.CreateNote(note)
		if err != nil {
			view.Wrap(err, w)
			return
		}
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Note Created",
			"note":    n,
		})
	})
}

// Protected Request
func showAllNotes(svc note.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			view.Wrap(view.ErrMethodNotAllowed, w)
			return
		}
		claims, err := middleware.ValidateAndGetClaims(r.Context(), "user")
		if err != nil {
			view.Wrap(err, w)
			return
		}
		n, err := svc.ShowAllNotes(claims["id"].(float64))
		if err != nil {
			view.Wrap(err, w)
			return
		}
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Notes Retrieved",
			"notes":   n,
		})
	})
}

// Handlers
func MakeNoteHandler(r *http.ServeMux, svc note.Service) {
	r.Handle("/api/notes/create", middleware.Validate(createNote(svc)))
	r.Handle("/api/notes", middleware.Validate(showAllNotes(svc)))
}
