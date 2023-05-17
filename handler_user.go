package main

import (
	"github.com/google/uuid"
	"fmt"
	"net/http"
	"encoding/json"	
	"time"
	"github.com/ilkerBedir/go-learning/internal/database"
)


func (apiConfig apiConfig) handlerCreateUser(w http.ResponseWriter,r *http.Request){
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params:=parameters{}
	err:=decoder.Decode(&params)
	if err!= nil{
		respondWithError(w,400,fmt.Sprintf("Error decoding parameters:%v ",err))
		return
	}

	user,err:=apiConfig.DB.Createuser(r.Context(),database.CreateuserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
        Name: params.Name,
	})
	if err!= nil{
		respondWithError(w,400,fmt.Sprintf("Couldnt create user:%v ",err))
		return
	}
	respondWithJson(w,200,databaseUserToUser(user))
	}
