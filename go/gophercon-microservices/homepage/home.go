package homepage

import (
	"database/sql"
	"log"
	"net/http"
	"time"
)

type Handlers struct {
	logger *log.Logger
	db     *sql.DB
}

const message = "Hello GopherCon UK 2018!"

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	h.db.ExecContext(r.Context(), "")
	w.Header().Set("Content-Type", "text/plain; charsert=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}

func NewHandlers(logger *log.Logger, db *sql.DB) *Handlers {
	return &Handlers{
		logger: logger,
		db:     db,
	}
}

func (h *Handlers) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer h.logger.Printf("request processed in %s\n", time.Now().Sub(startTime))
		next(w, r)
	}
}

func (h *Handlers) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", h.Logger(h.Home))
}
