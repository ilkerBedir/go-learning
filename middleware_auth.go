package main
import (
	"fmt"
	"net/http"
	"github.com/ilkerBedir/go-learning/internal/database"
	"github.com/ilkerBedir/go-learning/internal/auth"
)

type authedHandler func(http.ResponseWriter, *http.Request,database.User)

func (apiConfig *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        apikey,err:=auth.GetAPIKey(r.Header)
		if err!= nil{
			respondWithError(w,403,fmt.Sprintf("Error auth:%v ",err))
			return
		}
		user,err:=apiConfig.DB.GetUserByAPIKey(r.Context(),apikey)
		if err!= nil{
			respondWithError(w,400,fmt.Sprintf("Couldnt get user:%v ",err))
			return
		}
		handler(w,r,user)
    })
}