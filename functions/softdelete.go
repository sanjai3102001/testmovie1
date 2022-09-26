package function

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

//this will cgange the status to false
func Softdelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parms := mux.Vars(r)
	var movies movie
	db.First(&movies, parms["id"])
	json.NewDecoder(r.Body).Decode(&movies)
	db.Save(&movies)
	json.NewEncoder(w).Encode(movies)
}
