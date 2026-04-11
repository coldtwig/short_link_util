package stat

import (
	"go/http-api/configs"
	"go/http-api/pkg/middleware"
	"go/http-api/pkg/res"
	"net/http"
	"time"
)

const (
	GroupByDay   = "day"
	GroupByMonth = "month"
)

type StatHandler struct {
	StatRepository *StatRepository
	Config         *configs.Config
}

type StatHandlerDeps struct {
	StatRepository *StatRepository
	Config         *configs.Config
}

func NewStatHandler(router *http.ServeMux, deps StatHandlerDeps) *StatHandler {
	handler := &StatHandler{
		StatRepository: deps.StatRepository,
	}

	router.Handle("/stat", middleware.IsAuthed(handler.getStat(), deps.Config))

	return handler
}

func (handler *StatHandler) getStat() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		from, err := time.Parse(time.DateOnly, r.URL.Query().Get("from"))
		if err != nil {
			res.Json(w, err, http.StatusBadRequest)
		}
		to, err := time.Parse(time.DateOnly, r.URL.Query().Get("to"))
		if err != nil {
			res.Json(w, err, http.StatusBadRequest)
		}
		by := r.URL.Query().Get("by")
		if by != GroupByDay && by != GroupByMonth {
			http.Error(w, "Invalid parameter by", http.StatusBadRequest)
			return
		}

		stats := handler.StatRepository.GetStats(by, from, to)

		res.Json(w, stats, http.StatusOK)
	}
}
