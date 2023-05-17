package main
import (
	"encoding/json"
	"net/http"
	"log"
)
func respondWithError(w http.ResponseWriter,code int, message string){
	if code > 499{
		log.Println("RespondWithError with 5xx error code",message)
	}
	type errResponse struct {
		Error string `json:"error"`
	}
	respondWithJson(w,code,errResponse{Error:message})
}
func respondWithJson(w http.ResponseWriter, code int,payload interface{}){
	dat,err := json.Marshal(payload)
	if err!= nil{
		log.Println("Failed to marshal payload:",payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type","application/json;charset=UTF-8")
	w.WriteHeader(code)
	w.Write(dat)
}