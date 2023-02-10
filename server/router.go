package server

import (
	"net/http"

	"github.com/RohomRutuja/Go_API/CheckStatus"
	"github.com/gorilla/mux"
)

func InitRouter(dp *dependencies) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/websites", CheckStatus.AddWebsitesHandler(dp.httpchecker)).Methods(http.MethodPost)
	router.HandleFunc("/websites", CheckStatus.GetWebsitesHandler(dp.httpchecker)).Methods(http.MethodGet)

	return router
}
