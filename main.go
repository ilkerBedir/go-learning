package main

import (
	"os"
	"log"
	"github.com/joho/godotenv"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"net/http"
	"github.com/ilkerBedir/go-learning/internal/database"
	"database/sql"
	_ "github.com/lib/pq"
)
type apiConfig struct {
	DB *database.Queries
}
func main()  {

	godotenv.Load()

	portString:=os.Getenv("PORT")
	if portString=="" {
		log.Fatal("PORT is not found in environment")
	}
	dbURL:=os.Getenv("DB_URL")
	if dbURL=="" {
		log.Fatal("DB_URL is not found in environment")
	}
	conn,err:=sql.Open("postgres",dbURL)
	apiCnfg:=apiConfig{
		DB: database.New(conn),
	}
	if(err != nil){
		log.Fatal("Can't connect to database : ",err)
	}
	router:=chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*","http://*"},
		AllowedMethods: []string{"GET","POST","PUT","DELETE","OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	router.Get("/healthz",handlerReadiness)
	router.Get("/err",handlerError)
	router.Post("/users",apiCnfg.handlerCreateUser)
	router.Get("/users",apiCnfg.middlewareAuth(apiCnfg.handlerGetUserByAPIKey))
	router.Post("/feeds",apiCnfg.middlewareAuth(apiCnfg.handlerCreateFeed))
	router.Get("/feeds",(apiCnfg.handlerGetFeeds))
	router.Post("/feed-follows",apiCnfg.middlewareAuth(apiCnfg.handlerCreateFeedFollow))
	router.Mount("/v1",router )
	
	srv:=&http.Server{
		Handler:router,
		Addr: ":"+portString,
	}
	log.Printf("Server starting on port %v",portString)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}