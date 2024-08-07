package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ronicarvalho/techweek-go-react-server/internal/store/pgstore"
)

type apiHandlers struct {
	q *pgstore.Queries
	r *chi.Mux
}

func (h apiHandlers) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.r.ServeHTTP(w, r)
}

func NewHandler(q *pgstore.Queries) http.Handler {
	a := apiHandlers{
		q: q,
	}

	r := chi.NewRouter()
	a.r = r

	return a
}
