package api

import (
	"fmt"
	"inhouseadtest/internal/service"
	"inhouseadtest/internal/storage"
	"net/http"

	"github.com/gorilla/mux"
)

func (api *API) endpoints() {
	api.r.HandleFunc("/admin", api.adminStat).Methods("GET")
	api.r.HandleFunc("/search/{id}", api.search).Methods("GET")
	api.r.HandleFunc("/min", api.findMin).Methods("GET")
	api.r.HandleFunc("/max", api.findMax).Methods("GET")
}

func (api *API) search(w http.ResponseWriter, r *http.Request) {
	c := storage.SitesRequest
	c["/search"]++

	s := mux.Vars(r)["id"]
	store := storage.SitesStorage
	dur, ok := service.FindSiteRespTime(s, store)
	if !ok {
		ans := fmt.Sprintf("url: %s is incorrect", s)
		w.Write([]byte(ans))
		return
	}
	if dur == 0 {
		ans := fmt.Sprintf("url: %s is unavailable", s)
		w.Write([]byte(ans))
		return
	}
	ans := fmt.Sprintf("url: %s, response time: %s", s, dur)
	w.Write([]byte(ans))
}

func (api *API) findMax(w http.ResponseWriter, r *http.Request) {
	c := storage.SitesRequest
	c["/max"]++

	store := storage.SitesStorage

	s, dur := service.FindMax(store)

	ans := fmt.Sprintf("max time: %s, url: %s", dur, s)
	w.Write([]byte(ans))
}

func (api *API) findMin(w http.ResponseWriter, r *http.Request) {
	c := storage.SitesRequest
	c["/min"]++

	store := storage.SitesStorage

	s, dur := service.FindMin(store)

	ans := fmt.Sprintf("min time: %s, url: %s", dur, s)
	w.Write([]byte(ans))
}

func (api *API) adminStat(w http.ResponseWriter, r *http.Request) {
	c := storage.SitesRequest

	for k, v := range c {
		ans := fmt.Sprintf("request type: %s, request count: %d\n", k, v)
		w.Write([]byte(ans))
	}
}
