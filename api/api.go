package api

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)


type CoinBalanceParams struct{
	Username string
}

type CoinBalanceResponse struct{
	Code int
	Balance int64
}

type PublicRouteResponse struct{
	Code int
	Message string
}

type Error struct{
	Code int
	Message string
}

func writeError(w http.ResponseWriter, message string, code int){
	res:= Error{
		Code: code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code);
	json.NewEncoder(w).Encode(res)
}

var RequestErrorHandler = func(w http.ResponseWriter, err error){
	writeError(w, err.Error(), http.StatusBadRequest)
}
var InternalErrorHandler = func(w http.ResponseWriter, err error){
	writeError(w, "Internal error!", http.StatusInternalServerError)
}

func GlobalErrorHandler(err error, callbacks ...func(error)) error {
	if err == nil {
		return nil
	}

	if len(callbacks) > 0 && callbacks[0] != nil {
		callbacks[0](err)
	} else {
		log.Println("Erro:", err)
	}

	return err
}