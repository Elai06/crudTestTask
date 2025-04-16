package server

import (
	"crudTestTask/env"
	"crudTestTask/internal/repository"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type UserHttpServer struct {
	repository repository.DataBaseHandler
}

func New(rep repository.DataBaseHandler) *UserHttpServer {
	return &UserHttpServer{rep}
}

func (u *UserHttpServer) Start(config env.Config) error {
	r := mux.NewRouter()
	r.HandleFunc("/users", u.createUser).Methods(http.MethodPost)
	r.HandleFunc("/users", u.getUser).Methods(http.MethodGet)
	r.HandleFunc("/users", u.updateUser).Methods(http.MethodPut)
	r.HandleFunc("/users", u.deleteUser).Methods(http.MethodDelete)

	server := &http.Server{
		Addr:         config.Port,
		Handler:      r,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
	}

	err := server.ListenAndServe()
	if err != nil {
		return err
	}

	log.Printf("Server started at port %s", config.Port)

	return nil
}

func (u *UserHttpServer) createUser(w http.ResponseWriter, r *http.Request) {
	result := repository.Data{}
	err := json.NewDecoder(r.Body).Decode(&result)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("error decoder %v", err)
		return
	}

	user, err := u.repository.Create(result)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("error creating user: %s", err)
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("error encoder %v", err)
	}
}

func (u *UserHttpServer) getUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.URL.Query().
		Get("user_id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("error parse id %v", err)
		return
	}

	data, err := u.repository.Get(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("error get user: %v", err)
		return
	}

	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("error encoding body %v", err)
		return
	}
}

func (u *UserHttpServer) updateUser(w http.ResponseWriter, r *http.Request) {
	result := repository.Data{}
	err := json.NewDecoder(r.Body).Decode(&result)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("error decoding body %v", err)
		return
	}
	data, err := u.repository.Update(result)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("error update user: %s", err)
	}

	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("error encoding body %v", err)
		return
	}
}

func (u *UserHttpServer) deleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.URL.Query().
		Get("user_id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("error parse id %v", err)
		return
	}

	err = u.repository.Delete(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("error delete user%v ,%v", id, err)
		return
	}

	err = json.NewEncoder(w).Encode(fmt.Sprintf("user deleted %v", id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("error encoding body %v", err)
		return
	}
}
