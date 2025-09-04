package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/matheusgosk8/goapi/api"
	"github.com/matheusgosk8/goapi/internal/handlers"
	log "github.com/sirupsen/logrus"
)


func main(){

	log.SetReportCaller(true);
	var r *chi.Mux = chi.NewRouter();
	handlers.Handler(r);

	fmt.Println("Starting GO API ...");

	fmt.Println("GO api started on port: http://localhost:8000");

	err:= http.ListenAndServe("localhost:8000", r)
	api.GlobalErrorHandler(err)

}