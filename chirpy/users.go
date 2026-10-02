package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
}

type createUserPayload struct {
	Email string `json:"email"`
}

// this function returns a user, configs contains, addr and *database.Queries
func (cfg *config) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "could not read request body", err)
		return
	}

	var payload createUserPayload
	if err := json.Unmarshal(body, &payload); err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid JSON body", err)
		return
	}

	dbUser, err := cfg.queries.CreateUser(ctx, payload.Email)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not create user", err)
		return
	}
	respondWithJSON(w, http.StatusCreated, User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Email:     dbUser.Email,
	})
}

func (cfg *config) Delete(w http.ResponseWriter, r *http.Request) {
	err := cfg.queries.DeleteUser(context.Background())
	if err != nil {
		fmt.Printf("Cannot delete a user %s", err)
	}
}
