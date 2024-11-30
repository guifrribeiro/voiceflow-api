package main

import (
	"log"
	"net/http"
	"voiceflow/internal/api/v1"
	"voiceflow/internal/db"
	"voiceflow/internal/repositories"
	"voiceflow/internal/services"

	"github.com/gorilla/mux"
)

func main() {
	database := db.SetupDataBase()

	userRepo := repositories.NewUserRepository(database)
	userService := services.NewUserService(userRepo)
	userHandler := v1.NewUserHandler(userService)

	router := mux.NewRouter()

	// Register routes
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	}).Methods("GET")

	router.HandleFunc("/api/v1/users", userHandler.CreateUserHandler).Methods("POST")
	router.HandleFunc("/api/v1/users", userHandler.GetUsersHandler).Methods("GET")

	log.Println("Servidor iniciado na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
