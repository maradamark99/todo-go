package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Priority uint16

const (
	LOW = iota
	MEDIUM
	HIGH
)

type TodoStorage interface {
	GetAll() []*Todo
	GetById(id *uuid.UUID)
	Store(t *Todo)
	DeleteById(id *uuid.UUID)
}

type TodoOption func(*Todo)

type Todo struct {
	Id          uuid.UUID
	Name        string
	Description string
	Priority    Priority
	Deadline    time.Time
	IsDone      bool
}

func CreateTodo(name string, options ...TodoOption) *Todo {
	t := &Todo{
		Id:       uuid.New(),
		Name:     name,
		Priority: MEDIUM,
		IsDone:   false,
	}
	for _, o := range options {
		o(t)
	}

	return t
}

func WithPriority(p Priority) func(*Todo) {
	return func(t *Todo) {
		t.Priority = p
	}
}

func WithDeadline(d time.Time) func(*Todo) {
	return func(t *Todo) {
		t.Deadline = d
	}
}

func WithDescription(s string) func(*Todo) {
	return func(t *Todo) {
		t.Description = s
	}
}

func (t Todo) String() string {
	return fmt.Sprintf("id: %s, name: %s, priority: %d", t.Id.String(), t.Name, t.Priority)
}
