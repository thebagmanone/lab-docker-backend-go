package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/RMMFindHolding/labdocker/routers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
		os.Exit(1)
	}
}

func main() {
	fmt.Println("RUN SOFTWARE METAFANS")

	//LoadEnv()

	router := mux.NewRouter()

	router.HandleFunc("/randomNumber", routers.RandomNumber).Methods("GET")

	/*front*/
	// router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
	// 	http.Redirect(rw, r, "/web/", http.StatusSeeOther)
	// })
	// router.PathPrefix("/web/").Handler(http.StripPrefix("/web/", http.FileServer(http.Dir("web"))))

	/*Reviso si en el sistema operativo hay un puerto asignado, si no le asigno uno*/
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "3000"
	}
	//Manejador de permisos para las transacciones, sin cors se revotarian las solicitudes externas
	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
