package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jbrre/workshop-docker/client"
)

func JSONError(w http.ResponseWriter, err error, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	fmt.Fprintf(w, `{"error_code": %d, "message": "%s"}`, code, err.Error())
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"message": "Hello world !"}`)
}

func getUserList(w http.ResponseWriter, r *http.Request) {
	userList, err := client.GetUserList()

	if err != nil {
		JSONError(w, err, http.StatusInternalServerError)
	}

	outputStr, err := json.Marshal(userList)
	if err != nil {
		JSONError(w, err, http.StatusInternalServerError)
	}
	fmt.Fprintln(w, string(outputStr))
}

func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/", homePage).Methods("GET")
	r.HandleFunc("/user_list", getUserList).Methods("GET")
}
