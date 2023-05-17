package main

import (
	"os"
	"log"
	"github.com/joho/godotenv"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"net/http"
)
func main()  {

	godotenv.Load()

	portString:=os.Getenv("PORT")
	if portString=="" {
		log.Fatal("PORT is not found in environment")
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
	router.Mount("/v1",router )
	
	srv:=&http.Server{
		Handler:router,
		Addr: ":"+portString,
	}
	log.Printf("Server starting on port %v",portString)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}