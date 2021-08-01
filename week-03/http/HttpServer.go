package http

import (
	"context"
	//"context"
	"net/http"
	"strconv"
)

func NewHttp(ctx context.Context,port int) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		request = request.WithContext(ctx)
		writer.Write([]byte("hello world http"))
	})
	return &http.Server{
		Addr:    ":"+strconv.Itoa(port),
		Handler: mux,
	}
}