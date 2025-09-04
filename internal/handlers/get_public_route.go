package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/matheusgosk8/goapi/api"
)


func GetPublicRoute (w http.ResponseWriter, req *http.Request){

	var response =  api.PublicRouteResponse{
		Code: http.StatusOK,
		Message: "Server runing with full health",
	}

	w.Header().Set("Content-Type", "application/json");
	if err := json.NewEncoder(w).Encode(response); err != nil {
		api.InternalErrorHandler(w, err)
		return
	}
}
