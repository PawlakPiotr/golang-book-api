package main

import (
	"net/http"

	config "golang-book-api/config"

	utils "golang-book-api/utils"

	db "golang-book-api/database"

	routes "golang-book-api/routes"

	"github.com/gorilla/mux"
)

func main() {

	config, err := config.LoadConfiguration("./config/localhost.json")
	if err != nil {
		utils.Logger.Error("Loading configuration error")
		return
	}

	db.ConnectToDB(config.Database.Connection, config.Database.Name)

	router := mux.NewRouter()
	router.Use(commonMiddleware)
	routes.SetRoutes(router)

	utils.Logger.Infof("[Starting API] Listening on PORT: %s", config.Port)
	utils.Logger.Fatal(
		http.ListenAndServe(
			":"+config.Port,
			router,
		),
	)
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
