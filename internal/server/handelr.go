package server

import (
	"fmt"
	"htt-server/internal/service"
	"io"
	"net/http"
	"strconv"
)

type Handler struct {
	service *service.Hello
}

func NewHandler() *Handler {
	return &Handler{
		service: &service.Hello{},
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Hello World"))
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World"))
}

// @http://localhost:8080/application/hello?name=Arthur&typeOf=3
func (h *Handler) HelloAPPHandler(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Hello World"))

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(body) <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	fmt.Println(string(body))
	w.WriteHeader(http.StatusOK)

	name := r.URL.Query().Get("name")

	typeOf := r.URL.Query().Get("typeOf")

	w.Write([]byte("Hello " + h.service.Hello(name) + " " + strconv.Itoa(h.service.TypeOf(typeOf))))

}
