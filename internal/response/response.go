package response

import (
	"encoding/json"
	"log"
	"net/http"
)

func ResponseWithJSON(w http.ResponseWriter, code int, data interface{}){
	response, err:= json.Marshal(data)

	if err != nil{
		log.Fatalf("JSON formation failed: %v",err)
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(code)
	w.Write(response)

}

func ResponseWithError(w http.ResponseWriter, code int, msg string){
	var response = map[string]string{ "error":msg }
	ResponseWithJSON(w, code, response)
}