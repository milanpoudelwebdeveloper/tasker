package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type TasksService struct {
	store Store
}

func NewTasksService(store Store) *TasksService {
	return &TasksService{
		store,
	}
}

func (s *TasksService) RegisterRoutes(r *mux.Router) {

	r.HandleFunc("/tasks", s.handleCreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", s.handleGetTask).Methods("GET")
}

func (s *TasksService) handleCreateTask(w http.ResponseWriter, r *http.Request) {

}

func (s *TasksService) handleGetTask(w http.ResponseWriter, r *http.Request) {

}
