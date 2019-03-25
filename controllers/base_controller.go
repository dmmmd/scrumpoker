package controllers

import (
	"fmt"
	"net/http"
)

func SendRawResponse(w http.ResponseWriter, body string) {
	_, _ = fmt.Fprintf(w, body+"\n\n")
}
