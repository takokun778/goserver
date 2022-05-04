package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HelloRequest struct {
	Name string `json:"name"`
}

type HelloResponse struct {
	Message string `json:"message"`
}

func HelloHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)

		return
	}

	var req HelloRequest

	// @see https://christina04.hatenablog.com/entry/2017/01/06/190000
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	msg := fmt.Sprintf("Hello %s", req.Name)

	// @see https://konboi.hatenablog.com/entry/2014/09/23/172756
	res, err := json.Marshal(HelloResponse{Message: msg})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.Write(res)

	w.WriteHeader(http.StatusOK)
}
