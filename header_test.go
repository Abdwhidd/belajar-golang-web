package belajar_golang_web

import (
	"fmt"
	"net/http"
)

func RequestHeader(writer http.ResponseWriter, request *http.Request) {
	contentType := writer.Header().Get("content-type")
	fmt.Fprint(writer, contentType)
}
