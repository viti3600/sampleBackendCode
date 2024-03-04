package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	router = mux.NewRouter()
	// rewardPortalRoute = mux.NewRouter()
)

func StartbackendReqAPI() {

	mapUrls()
}
func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
