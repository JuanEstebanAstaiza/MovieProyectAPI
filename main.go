package moviesProyectAPI

import (
	"fmt"
	"github.com/JuanEstebanAstaiza/MovieProyectAPI/controllers"
	"log"
	"net/http"

	_ "github.com/JuanEstebanAstaiza/MovieProyectAPI/controllers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Configurar rutas
	router.HandleFunc("/api/movie/{movie_id}", controllers.GetMovieDetails()).Methods("GET")
	router.HandleFunc("/api/most-viewed", controllers.GetMostViewedMovies()).Methods("GET")
	router.HandleFunc("/api/comment", controllers.AddComment()).Methods("POST")
	router.HandleFunc("/api/comment", controllers.DeleteComment()).Methods("DELETE")
	router.HandleFunc("/api/comment", controllers.EditComment()).Methods("PUT")
	router.HandleFunc("/api/user/{user_id}", controllers.ModifyUserInfo).Methods("PUT")

	// Iniciar el servidor
	port := ":8080"
	fmt.Println("Servidor escuchando en el puerto", port)
	log.Fatal(http.ListenAndServe(port, router))
}
