package app

import (
	"testing"
	"time"

	"github.com/influx6/flux"
)

func makeMockDb() TodoDatabase {
	co := time.Now()
	return NewMockTodo([]*Todo{
		&Todo{10, "random task", "save desscriptions", co.Unix(), co.Day(), co.Month().String()},
		&Todo{1, "become frugal", "save money", co.Unix(), co.Day(), co.Month().String()},
		&Todo{2, "get shooter score", "need to save like kickass in starcraft2", co.Unix(), co.Day(), co.Month().String()},
	})
}

func TestTodoModel(t *testing.T) {
	db := makeMockDb()

	if _, err := db.FindID(1); err != nil {
		flux.FatalFailed(t, "Todo id: %d should exists", 1)
	}

	flux.LogPassed(t, "Todo id: %d was located succesfully", 1)

	if err := db.Destroy(2); err != nil {
		flux.FatalFailed(t, "Todo id: %d was not destroyed", 2)
	}

	flux.LogPassed(t, "Todo id: %d was destroyed succesfully", 2)

	if todos, _ := db.FindAll(); len(todos) > 3 {
		flux.FatalFailed(t, "Todo length is not correct, epxected 2 got %d", len(todos))
	}

	flux.LogPassed(t, "Todo length: %d is correct", 2)
}
