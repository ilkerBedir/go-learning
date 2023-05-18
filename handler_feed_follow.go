package main

import (
	"github.com/google/uuid"
	"fmt"
	"net/http"
	"encoding/json"	
	"time"
	"github.com/ilkerBedir/go-learning/internal/database"
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


