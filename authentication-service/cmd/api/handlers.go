package main

import (
	"errors"
	"fmt"
	"github.com/tsawler/toolbox"
	"net/http"
)

func (app *Config) Authenticate(w http.ResponseWriter, r *http.Request) {
	var tools toolbox.Tools

	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := tools.ReadJSON(w, r, &requestPayload)
	if err != nil {
		_ = tools.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	user, err := app.Models.User.GetByEmail(requestPayload.Email)
	if err != nil {
		_ = tools.ErrorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		_ = tools.ErrorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	payload := toolbox.JSONResponse{
		Error:   false,
		Message: fmt.Sprintf("Logged in %s", user.Email),
		Data:    user,
	}

	_ = tools.WriteJSON(w, http.StatusAccepted, payload)
}
