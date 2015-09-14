package app

import (
	"fmt"
	"time"
)

// MockTodo provides a mock db model for testing the todo model database functionality
type MockTodo struct {
	db []*Todo
}

// NewMockTodo returns a new sql todo manager
func NewMockTodo(db []*Todo) *MockTodo {
	return &MockTodo{db}
}

// New creates and adds a new schema into the database
func (s *MockTodo) New(task string, desc string) error {
	return s.Save(&Todo{
		Task:        task,
		Description: desc,
		Created:     time.Now(),
		Updated:     time.Now(),
	})
}

// FindID returns a todo item with the specified id
func (s *MockTodo) FindID(id int) (*Todo, error) {
	for _, mo := range s.db {
		if mo.ID != id {
			continue
		}

		return mo, nil
	}

	return nil, fmt.Errorf("%d id does not exists", id)
}

// FindAll returns all todos in the database
func (s *MockTodo) FindAll() ([]*Todo, error) {
	return s.db, nil
}

//Save adds a new todo into the database
func (s *MockTodo) Save(m *Todo) error {
	for _, mo := range s.db {
		if mo.ID == m.ID {
			return fmt.Errorf("%d id already exists", m.ID)
		}
	}

	m.Created = time.Now()
	s.db = append(s.db, m)
	return nil
}

// Update updates the records in the db
func (s *MockTodo) Update(m *Todo) error {
	ind := -1

	for n, mo := range s.db {
		if mo.ID == m.ID {
			ind = n
			break
		}
	}

	if ind == -1 {
		return fmt.Errorf("%d id does not exists", m.ID)
	}

	m.Updated = time.Now()
	s.db[ind] = m
	return nil
}

// Destroy deeltes a todo from the database
func (s *MockTodo) Destroy(m int) error {
	ind := -1

	for n, mo := range s.db {
		if mo.ID != m {
			continue
		}

		ind = n
		break
	}

	if ind == -1 {
		return fmt.Errorf("%d id does not exists", m)
	}

	s.db = append(s.db[0:ind], s.db[ind:]...)
	return nil
}
