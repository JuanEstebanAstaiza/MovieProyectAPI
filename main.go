package moviesProyectAPI

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Configurar rutas
	router.HandleFunc("/api/movie/{movie_id}", getMovieDetails).Methods("GET")
	router.HandleFunc("/api/most-viewed", getMostViewedMovies).Methods("GET")
	router.HandleFunc("/api/comment", addComment).Methods("POST")
	router.HandleFunc("/api/comment", deleteComment).Methods("DELETE")
	router.HandleFunc("/api/comment", editComment).Methods("PUT")
	router.HandleFunc("/api/user/{user_id}", modifyUserInfo).Methods("PUT")

	// Iniciar el servidor
	port := ":8080"
	fmt.Println("Servidor escuchando en el puerto", port)
	log.Fatal(http.ListenAndServe(port, router))
}

func getMovieDetails(w http.ResponseWriter, r *http.Request) {
	// Implementar lógica para obtener detalles de la película
}

func getMostViewedMovies(w http.ResponseWriter, r *http.Request) {
	// Implementar lógica para obtener las n películas más visualizadas
}

func addComment(w http.ResponseWriter, r *http.Request) {
	// Implementar lógica para agregar comentarios
}

func deleteComment(w http.ResponseWriter, r *http.Request) {
	// Implementar lógica para eliminar comentarios
}

func editComment(w http.ResponseWriter, r *http.Request) {
	// Implementar lógica para editar comentarios
}

func modifyUserInfo(w http.ResponseWriter, r *http.Request) {
	// Implementar lógica para modificar información del usuario
}
