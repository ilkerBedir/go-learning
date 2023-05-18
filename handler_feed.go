package main

import (
	"github.com/google/uuid"
	"fmt"
	"net/http"
	"encoding/json"	
	"time"
	"github.com/ilkerBedir/go-learning/internal/database"
)


func (apiConfig apiConfig) handlerCreateFeed(w http.ResponseWriter,r *http.Request,user database.User){
	type parameters struct {
		Name string `json:"name"`
		URL string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params:=parameters{}
	err:=decoder.Decode(&params)
	if err!= nil{
		respondWithError(w,400,fmt.Sprintf("Error decoding parameters:%v ",err))
		return
	}

	feed,err:=apiConfig.DB.Createfeeds(r.Context(),database.CreatefeedsParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
        Name: params.Name,
		Url: params.URL,
		UserID: user.ID,
	})
	if err!= nil{
		respondWithError(w,400,fmt.Sprintf("Couldnt create user:%v ",err))
		return
	}
	respondWithJson(w,201,databaseFeedToFeed(feed))
}

func (apiConfig apiConfig) handlerGetFeeds(w http.ResponseWriter,r *http.Request){
	
	feeds,err:=apiConfig.DB.Getfeeds(r.Context())
	if err!= nil{
		respondWithError(w,400,fmt.Sprintf("Couldnt create user:%v ",err))
		return
	}
	respondWithJson(w,201,databaseFeedsToFeeds(feeds))
}

