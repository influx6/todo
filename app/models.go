package app

import (
	"encoding/json"
	"time"

	"github.com/jmoiron/sqlx"
)

// Todo defines the basic struct for the todo models
type Todo struct {
	ID          int       `db:"id" json:"id"`
	Task        string    `db:"task" json:"task"`
	Description string    `db:"desc" json:"desc"`
	Created     time.Time `db:"created_at" json:"created_at"`
	Updated     time.Time `db:"updated_at" json:"updated_at"`
}

// UnmarshalJSON loads up json data into the struct
func (t *Todo) UnmarshalJSON(bo []byte) error {
	if err := json.Unmarshal(bo, t); err != nil {
		return err
	}
	return nil
}

// TodoDatabase represents a basic
type TodoDatabase interface {
	New(string, string) error
	FindID(int) (*Todo, error)
	FindAll() ([]*Todo, error)
	Save(*Todo) error
	Update(*Todo) error
	Destroy(int) error
}

// SQLTodo provide model manager for managing Todo models
type SQLTodo struct {
	db *sqlx.DB
}

var createSchema = `
  CREATE TABLE IF NOT EXITS apps.todos (
    id not null autoincrement primarykey,
    task varchar(55) not null,
    desc text not null,
    created_at datetime not null,
    updated_at datetime not null
  )
`

// NewSQLTodo returns a new sql todo manager
func NewSQLTodo(db *sqlx.DB) *SQLTodo {
	db.MustExec(createSchema)
	return &SQLTodo{db}
}

// New creates and adds a new schema into the database
func (s *SQLTodo) New(task string, desc string) error {
	return s.Save(&Todo{
		Task:        task,
		Description: desc,
		Created:     time.Now(),
		Updated:     time.Now(),
	})
}

var selectOneSchema = `
  SELECT * from apps.todos todo
  WHERE todo.id = ?
`

// FindID returns a todo item with the specified id
func (s *SQLTodo) FindID(id int) (*Todo, error) {
	todo := Todo{}
	if err := s.db.Get(&todo, selectOneSchema, id); err != nil {
		return nil, err
	}
	return &todo, nil
}

var selectSchema = `
  SELECT * from apps.todos
`

// FindAll returns all todos in the database
func (s *SQLTodo) FindAll() ([]*Todo, error) {
	todos := []*Todo{}

	err := s.db.Select(&todos, selectSchema)

	if err != nil {
		return nil, err
	}

	return todos, nil
}

var saveSchema = `
  INSERT INTO apps.todos(task,desc,created_at,                                                                                                                           updated_at) values(?,?,?,?)
`

//Save adds a new todo into the database
func (s *SQLTodo) Save(m *Todo) error {
	if m.ID != 0 {
		return s.Update(m)
	}

	_, err := s.db.Exec(saveSchema, m.Task, m.Description, m.Created, m.Updated)

	if err != nil {
		return err
	}

	return nil
}

var updateAllSchema = `
  UPDATE apps.todos todo
  SET todo.task = ?, todo.desc = ?, todo.created_at = ?,todo.updated_at = ?
  WHERE todo.id = ?
`

var updateSchema = `
  UPDATE apps.todos todo
  SET %s
  WHERE todo.id = ?
`

// Update updates the records in the db
func (s *SQLTodo) Update(m *Todo) error {
	// if m.Updated.Before(time.Now()) {
	//   m.Updated = time.Now()
	// }
	_, err := s.db.Exec(updateAllSchema, m.Task, m.Description, m.Created, m.Updated)

	if err != nil {
		return err
	}

	return nil
}

var deleteSchema = `
  DELETE FROM apps.todos todo
  WHERE todo.id = ?
`

// Destroy deeltes a todo from the database
func (s *SQLTodo) Destroy(m int) error {
	_, err := s.db.Exec(deleteSchema, m)

	if err != nil {
		return err
	}

	return nil
}
