package api

import "github.com/gorilla/mux"

type API struct {
	r *mux.Router
}

func New() *API {
	a := API{r: mux.NewRouter()}
	a.endpoints()
	return &a
}

func (api *API) Router() *mux.Router {
	return api.r
}
