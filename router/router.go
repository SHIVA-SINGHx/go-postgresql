package router

import(
	"github.com/gorilla/mux"
	"go-postgrsql/middleware"
)



func Router() *mux.Router{
	router:= mux.NewRouter()

	router.HandleFunc("/api/stock{id}", middleware.GetStock).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/stock/", middleware.GetAllStock).Methods("GET", "OPTIONS")

}