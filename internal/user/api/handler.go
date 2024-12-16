package api

import (
	"encoding/json"
	"net/http"

	"CodeLabs/internal/user/service"
)

var userService = service.NewUserService()

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	success, err := userService.Signup(req.Username, req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	resp := map[string]interface{}{"success": success, "message": "Signup successful"}
	json.NewEncoder(w).Encode(resp)
}
