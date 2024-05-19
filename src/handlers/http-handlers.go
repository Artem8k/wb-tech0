package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rest-api/src/service"

	"github.com/gorilla/mux"
)

type Handlers struct {
	service service.Service
}

func New(service service.Service) *Handlers {
	return &Handlers{
		service: service,
	}
}

func (h *Handlers) InitHandlers() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/getOrder", h.GetOrder).Methods("GET")
	router.HandleFunc("/getTestOrder", h.GetTestOrder).Methods("GET")

	return router
}

// [status code 400] если отсутствует uid в query-параметре
// [status code 500] если ошибка при маршале
func (h *Handlers) GetOrder(w http.ResponseWriter, req *http.Request) {
	uid := req.URL.Query().Get("uid")

	if uid == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "missing uid in query", status: "400"}`))
		return
	}

	res := h.service.GetOrder(uid)

	json, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Internal Server Error", status: "500"}`))
		return
	}
	w.Write(json)
}

func (h *Handlers) GetTestOrder(w http.ResponseWriter, req *http.Request) {
	res := h.service.GetTestOrder()

	json, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
	}

	w.Write(json)
}
