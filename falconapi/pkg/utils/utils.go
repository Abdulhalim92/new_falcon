package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteJSONResponse(w http.ResponseWriter, httpStatus int, v interface{}) {
	body, err := json.Marshal(v)
	if err != nil {
		fmt.Println("Reponse XML Marhshal err:", err)
		return
	}
	w.WriteHeader(httpStatus)
	_, err = w.Write(body)
	if err != nil {
		fmt.Println("Response write body err:", err)
		return
	}
}
