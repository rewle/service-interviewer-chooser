package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func PanicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

func BadGateway(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadGateway)
	fmt.Fprintf(w, "{}")
}

func JSONResponse(w http.ResponseWriter, status int, data interface{}) {
	resp, err := json.Marshal(data)
	if err != nil {
		BadGateway(w)
		return
	}
	w.WriteHeader(status)
	fmt.Fprint(w, string(resp))
}
