package app

import (
	"time"

	"github.com/jmoiron/sqlx"
)

// Todo defines the basic struct for the todo models
type Todo struct {
	ID          int    `db:"id" `
	Task        string `db:"task" `
	Description string `db:"description" `
	Stamp       int64  `db:"stamp" `
	Day         int    `db:"day"`
	Month       string `db:"month"`
	Completed   int    `db:"done"`
}

// IsDone returns true/false if its done
func (t *Todo) IsDone() bool {
	return t.Completed == 1
}

// TodoDatabase represents a basic
type TodoDatabase interface {
	New(string, string, int64) error
	FindID(int) (*Todo, error)
	FindAll() ([]*Todo, error)
	Save(*Todo) error
	Update(*Todo) error
	Destroy(int) error
	Complete(int) error
	Uncomplete(int) error
}

// SQLTodo provide model manager for managing Todo models
type SQLTodo struct {
	db *sqlx.DB
}

var createSchema = `
  CREATE TABLE IF NOT EXISTS apps.todos (
    id integer not null auto_increment primary key,
    task varchar(55) not null,
    description varchar(55) not null,
    stamp double not null,
    day integer not null,
    month varchar(55) not null,
    done integer not null
  )
`

// NewSQLTodo returns a new sql todo manager
func NewSQLTodo(db *sqlx.DB) TodoDatabase {
	db.MustExec(createSchema)
	return &SQLTodo{db}
}

// New creates and adds a new schema into the database
func (s *SQLTodo) New(task, desc string, stamp int64) error {

	var co time.Time

	if stamp <= 0 {
		co = time.Now()
	} else {
		co = time.Unix(stamp, 0)
	}

	return s.Save(&Todo{
		Task:        task,
		Description: desc,
		Stamp:       co.Unix(),
		Day:         co.Day(),
		Month:       co.Month().String(),
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

	// co :=
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
  INSERT INTO apps.todos(task,description,stamp,day,month,done) values(?,?,?,?,?,?)
`

//Save adds a new todo into the database
func (s *SQLTodo) Save(m *Todo) error {
	if m.ID != 0 {
		return s.Update(m)
	}

	tx := s.db.MustBegin()

	tx.Exec(saveSchema, m.Task, m.Description, m.Stamp, m.Day, m.Month, 0)

	err := tx.Commit()

	// _, err := s.db.Exec(saveSchema, m.Task, m.Description, m.Stamp, m.Day, m.Month)

	if err != nil {
		return err
	}

	return nil
}

var updateAllSchema = `
  UPDATE apps.todos todo
  SET todo.task = ?, todo.description = ?, todo.stamp = ?,todo.day = ?,todo.month = ?,todo.done = ?
  WHERE todo.id = ?
`

// Update updates the records in the db
func (s *SQLTodo) Update(m *Todo) error {
	tx := s.db.MustBegin()
	tx.Exec(updateAllSchema, m.Task, m.Description, m.Stamp, m.Day, m.Month, m.Completed)

	err := tx.Commit()

	if err != nil {
		return err
	}

	return nil
}

var deleteSchema = `
  DELETE FROM apps.todos
  WHERE id = ?
`

// Destroy deeltes a todo from the database
func (s *SQLTodo) Destroy(m int) error {
	tx := s.db.MustBegin()
	tx.Exec(deleteSchema, m)

	err := tx.Commit()

	if err != nil {
		return err
	}

	return nil
}

var completeSchema = `
  UPDATE apps.todos todo
  SET todo.done = ?
  WHERE todo.id = ?
`

// Complete updates a todo from the database and sets its done as completed
func (s *SQLTodo) Complete(m int) error {
	tx := s.db.MustBegin()
	tx.Exec(completeSchema, 1, m)

	err := tx.Commit()

	// _, err := s.db.Exec(completeSchema, 1, m)
	// log.Printf("Complete:", err)

	if err != nil {
		return err
	}

	return nil
}

// Uncomplete updates a todo from the database and sets its done as completed
func (s *SQLTodo) Uncomplete(m int) error {
	tx := s.db.MustBegin()
	tx.Exec(completeSchema, 0, m)

	err := tx.Commit()

	if err != nil {
		return err
	}

	return nil
}
