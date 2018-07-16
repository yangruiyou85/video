package main

import (
	"net/http"
	"github.com/yangruiyou85/video/api/defs"
	"encoding/json"
	"io"
)

func sendErrorResponse(w http.ResponseWriter, errResp defs.ErrorResponse) {

	w.WriteHeader(errResp.HttpSC)
	resStr, _ := json.Marshal(&errResp.Error)
	io.WriteString(w, string(resStr))

}

func sendNormalResponse(w http.ResponseWriter, resp string, sc int) {

	w.WriteHeader(sc)
	io.WriteString(w, resp)
}
