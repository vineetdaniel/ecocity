package main

import (
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/vineetdaniel/ecocity/common"
	"github.com/vineetdaniel/ecocity/routers"
)

func main() {
	//common.Startup()

	router := routers.InitRoutes()

	n := negroni.Classic()
	n.UseHandler(router)

	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: n,
	}
	log.Println("Listening....", common.AppConfig.Server)
	err := server.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
