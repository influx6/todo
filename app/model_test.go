package app

import (
	"testing"
	"time"

	"github.com/influx6/flux"
)

func makeMockDb() TodoDatabase {
	return NewMockTodo([]*Todo{
		&Todo{10, "random task", "save desscriptions", time.Now(), time.Now()},
		&Todo{1, "become frugal", "save money", time.Now(), time.Now()},
		&Todo{2, "get shooter score", "need to save like kickass in starcraft2", time.Now(), time.Now()},
	})
}

func TestTodoModel(t *testing.T) {
	db := makeMockDb()

	if _, err := db.FindID(1); err != nil {
		flux.FatalFailed(t, "Todo id: %s should exists", 1)
	}

	flux.LogPassed(t, "Todo id: %s was located succesfully", 1)

	if err := db.Destroy(2); err != nil {
		flux.FatalFailed(t, "Todo id: %s was not destroyed", 2)
	}

	flux.LogPassed(t, "Todo id: %s was destroyed succesfully", 2)

	if todos, _ := db.FindAll(); len(todos) > 3 {
		flux.FatalFailed(t, "Todo length is not correct, epxected 2 got %d", len(todos))
	}

	flux.LogPassed(t, "Todo length: %s is correct", 2)
}
