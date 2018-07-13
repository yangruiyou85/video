package main

import (
	"net/http"
	"io"
)

func sendErrorResponse(w http.ResponseWriter, sc int, errMsg error) {
	w.WriteHeader(sc)
	io.WriteString(w, errMsg)
}
