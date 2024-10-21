package routers

import (
	"github.com/Djay-M/goCrawler/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	// Status Apis
	router.HandleFunc("/", controllers.FetchServerStatus).Methods("GET")
	router.HandleFunc("/status", controllers.FetchServerStatus).Methods("GET")
	router.HandleFunc("/api/v1/status", controllers.FetchServerStatus).Methods("GET")

	// Crawler Apis
	// V1
	router.HandleFunc("/api/v1/crawler", controllers.CrawlerController).Methods("POST")

	return router
}
