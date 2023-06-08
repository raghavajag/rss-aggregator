package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/raghavajag/rss-aggregator/internal/database"
)

func (apiCfg *apiConfig) HandlerCreateFeeds(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %s", err))
		return
	}
	feeds, err := apiCfg.DB.CreateFeeds(r.Context(), database.CreateFeedsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		UserID:    user.ID,
		Url:       params.Url,
	})
	if err != nil {
		respondWithJSON(w, 403, err)
	}
	respondWithJSON(w, 201, databaseFeedsToFeeds(feeds))
}
