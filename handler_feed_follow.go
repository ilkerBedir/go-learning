package main

import (
	"github.com/google/uuid"
	"fmt"
	"net/http"
	"encoding/json"	
	"time"
	"github.com/ilkerBedir/go-learning/internal/database"
	"github.com/go-chi/chi"
)


func (apiConfig apiConfig) handlerCreateFeedFollow(w http.ResponseWriter,r *http.Request,user database.User){
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`		
	}
	decoder := json.NewDecoder(r.Body)
	params:=parameters{}
	err:=decoder.Decode(&params)
	if err!= nil{
		respondWithError(w,400,fmt.Sprintf("Error decoding parameters:%v ",err))
		return
	}
	
	feed_follow,err:=apiConfig.DB.CreateFeedFollow(r.Context(),database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID: params.FeedID,		
		UserID: user.ID,
	})
	if err!= nil{
		respondWithError(w,400,fmt.Sprintf("Couldnt create feed_follow:%v ",err))
		return
	}
	respondWithJson(w,201,databaseFeedFollowToFeedFollow(feed_follow))
}
func (cfg *apiConfig) handlerFeedFollowsGet(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := cfg.DB.GetFeedFollowsForUser(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create feed follow")
		return
	}

	respondWithJson(w, http.StatusOK, databaseFeedFollowsToFeedFollows(feedFollows))
}
func (cfg *apiConfig) handlerFeedFollowDelete(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid feed follow ID")
		return
	}

	err = cfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		ID:     feedFollowID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create feed follow")
		return
	}
	respondWithJson(w, http.StatusOK, struct{}{})
}


