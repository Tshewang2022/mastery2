package main

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/nado/chirpy/internal/auth"
)

func (cfg *config) Refresh(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Token string `json:"token"`
	}

	refreshToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Missing or malformed refresh token", err)
		return
	}

	user, err := cfg.queries.GetUserFromRefreshToken(r.Context(), refreshToken)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondWithError(w, http.StatusUnauthorized, "Invalid, expired, or revoked refresh token", err)
			return
		}
		respondWithError(w, http.StatusInternalServerError, "Couldn't look up refresh token", err)
		return
	}

	accessToken, err := auth.MakeJWT(user.ID, cfg.jwtSecret, time.Hour)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create access token", err)
		return
	}

	respondWithJSON(w, http.StatusOK, response{Token: accessToken})
}

func (cfg *config) RevokeRefreshTokens(w http.ResponseWriter, r *http.Request) {
	refreshToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Missing or malformed refresh token", err)
		return
	}

	if err := cfg.queries.RevokeRefreshToken(r.Context(), refreshToken); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't revoke refresh token", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
