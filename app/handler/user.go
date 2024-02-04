package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

const (
	ID = "id"
)

type userHandler struct{}

func NewUserHandler() *userHandler {
	return &userHandler{}
}

func (*userHandler) GetUserDetails(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(r)[ID], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad request!"))
		return
	}
	time.Sleep(100 * time.Millisecond)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("user: %v", id)))
}
