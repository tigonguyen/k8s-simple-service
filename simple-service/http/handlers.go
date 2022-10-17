package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Onefootball/simple-service/db"
)

// NewHandler returns a simpleHandler instance given a db.Postgres instance
func NewHandler(db db.Postgres) http.Handler {
	return simpleHandler{
		db: db,
	}
}

type simpleHandler struct {
	db db.Postgres
}

// ServerHTTP servers HTTP connections trying to reach a db.Postgres database.
// It tries to reach the database using its CheckConnection method, but, never
// fails:
// - If the connection with the database is successful it returns `http` 200 and a "Well done :)" message.
// - If the connection with the database fails, it returns `http` 200 and a "Running" message.
func (h simpleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.db.CheckConnection(); err != nil {
		fmt.Printf("%s - simple request\n", time.Now().Format(time.ANSIC))
		w.Write([]byte("Running"))
		return
	}

	fmt.Printf("%s - request with database\n", time.Now().Format(time.ANSIC))
	w.Write([]byte("Well done :)"))
}
