package todo

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

type InMemoryTodoStorage struct {
	Todos []*Todo
}

func (i *InMemoryTodoStorage) GetAll() []*Todo {
	return i.Todos
}

func (i *InMemoryTodoStorage) Store(t *Todo) {
	i.Todos = append(i.Todos, t)
}

func (i *InMemoryTodoStorage) GetById(id uuid.UUID) *Todo {
	for _, e := range i.Todos {
		if e.Id == id {
			return e
		}
	}
	return nil
}

func (i *InMemoryTodoStorage) DeleteById(id uuid.UUID) bool {
	for idx, e := range i.Todos {
		if e.Id == id {
			i.Todos[idx] = nil
			i.Todos = append(i.Todos[:idx], i.Todos[idx+1:]...)
			return true
		}
	}
	return false
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
