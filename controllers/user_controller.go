package Controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/your_username/your_project_name/models"
	"github.com/your_username/your_project_name/services"
)

func ModifyUserInfo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["user_id"]

	var updatedUser models.User
	err := json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = services.ModifyUserInfo(userID, updatedUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
