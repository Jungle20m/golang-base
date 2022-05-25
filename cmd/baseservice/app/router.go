package app

import (
	"encoding/json"
	"net/http"

	"github.com/Jungle20m/golang-base/cmd/baseservice/component"
	"github.com/go-chi/chi/v5"
)

type Response struct {
	Code int `json:"code"`
}

func InitRoutes(r *chi.Mux, appCtx component.AppCtx) {

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		response := Response{
			Code: 200,
		}
		body, _ := json.Marshal(response)
		w.Write(body)
	})

}
