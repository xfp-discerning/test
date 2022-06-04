package main

import (
	"fmt"
	"log"
)

type mysql struct{}

type redis struct{}

func (r *redis) Get(name string) (*User, error) {
	return &User{Name: "Bob"}, nil
}

type User struct {
	Name string
}

type Store interface {
	Get(name string) (*User, error)
}

type Server struct {
	store Store
}

func (s *Server) User(name string) {
	user, err := s.store.Get(name)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(user)
}

func (s *Server) Run() error {
	return nil
}

func New(s Store) *Server {
	return &Server{
		store: s,
	}
}

func main() {
	store := redis{}
	s := New(&store)
	s.User("Bob")
}

// optional function
