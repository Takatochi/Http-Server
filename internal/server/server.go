package server

import (
	"context"
	"net/http"
)

type Http struct {
	server *http.Server
}

func NewHttp(addr string) *Http {
	mux := http.NewServeMux() // router

	handler := NewHandler()
	// https://domen/ https://rozetka.com.ua/ -> index
	mux.HandleFunc("/", HelloHandler) // endpoint
	// https://domen/application/hello https://rozetka.com.ua/application/hello -> index
	// http://localhost:8080/application/hello?name=Arthur&typeOf=3
	mux.HandleFunc("/application/hello", handler.HelloAPPHandler) // endpoint

	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	return &Http{
		server: server,
	}
}

func (h *Http) Start() error {
	return h.server.ListenAndServe()
}

func (h *Http) Stop(ctx context.Context) error {
	return h.server.Shutdown(ctx)
}
