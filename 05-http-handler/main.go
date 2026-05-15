package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RequestBody struct {
	Name string `json:"name"`
}

type ResponseBody struct {
	Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var body RequestBody
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&body); err != nil || body.Name == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ResponseBody{Message: fmt.Sprintf("Hello %s", body.Name)})
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
