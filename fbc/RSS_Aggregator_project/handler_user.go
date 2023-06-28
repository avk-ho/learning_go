package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rss_aggregator/internal/database"
	"time"

	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	dec_err := decoder.Decode(&params)

	if dec_err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", dec_err))
		return
	}

	user, user_err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
	})
	if user_err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create user: %v", user_err))
		return
	}

	respondWithJSON(w, 200, databaseUserToUser(user))
}