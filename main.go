package main

import (
	"fmt"
	"github.com/JuanEstebanAstaiza/MovieProyectAPI/controllers"
	"github.com/JuanEstebanAstaiza/MovieProyectAPI/utils"
	"github.com/rs/cors"
	"log"
	"net/http"

	_ "github.com/JuanEstebanAstaiza/MovieProyectAPI/controllers"
	"github.com/gorilla/mux"
)

func main() {

	err := utils.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	router := mux.NewRouter()
	router.HandleFunc("/api/movie/{movie_id}", controllers.GetMovieDetails).Methods("GET")
	router.HandleFunc("/api/movie/most-viewed/{cant}", controllers.GetMostViewedMovies).Methods("GET")
	router.HandleFunc("/api/comment", controllers.AddComment).Methods("POST")
	router.HandleFunc("/api/comment", controllers.DeleteComment).Methods("DELETE")
	router.HandleFunc("/api/comment", controllers.EditComment).Methods("PUT")
	router.HandleFunc("/api/user/{user_id}", controllers.ModifyUser).Methods("PUT")
	router.HandleFunc("/api/user/register", controllers.RegisterUser).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.LoginUser).Methods("POST")
	router.HandleFunc("/api/process-payment", controllers.ProcessPaymentHandler).Methods("POST")
	router.HandleFunc("/api/get-payments/{user_id}", controllers.GetPaymentsByUserIDHandler).Methods("GET")
	router.HandleFunc("/api/update-payment-status/{payment_id}", controllers.UpdatePaymentStatusHandler).Methods("PUT")
	router.HandleFunc("/api/get-total-payments/{user_id}", controllers.GetTotalPaymentsByUserIDHandler).Methods("GET")
	router.HandleFunc("/api/create-subscription", controllers.CreateSubscriptionHandler).Methods("POST")
	router.HandleFunc("/api/get-subscription/{user_id}", controllers.GetSubscriptionByUserIDHandler).Methods("GET")
	router.HandleFunc("/api/cancel-subscription/{user_id}", controllers.CancelSubscriptionHandler).Methods("PUT")
	// Opciones CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Permitir solicitudes desde cualquier origen
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	handler := c.Handler(router)
	// Iniciar el servidor
	port := ":8088"
	fmt.Println("Servidor escuchando en el puerto", port)
	log.Fatal(http.ListenAndServe(port, handler))
}
