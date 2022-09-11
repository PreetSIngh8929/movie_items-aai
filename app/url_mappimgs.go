package app

import (
	"net/http"

	"github.com/PreetSIngh8929/movie_items-aai/controllers"
)

func mapUrls() {
	router.HandleFunc("/ping", controllers.PingController.Ping).Methods(http.MethodGet)
	router.HandleFunc("/movies", controllers.ItemsController.Create).Methods(http.MethodPost)
	router.HandleFunc("/movies/{id}", controllers.ItemsController.Get).Methods(http.MethodGet)
	router.HandleFunc("/movies/search", controllers.ItemsController.Search).Methods(http.MethodPost)
	router.HandleFunc("/movies/{id}", controllers.ItemsController.Book).Methods(http.MethodPut)
	router.HandleFunc("/movies/{id}", controllers.ItemsController.Cancel).Methods(http.MethodPut)
}
