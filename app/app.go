package app

import (
	"github.com/influx6/relay/engine"
	"github.com/jmoiron/sqlx"
)

func initApp(appo *engine.Engine) {
	appo.RedirectAll("", "/", "/todo")

	db, err := sqlx.Open(appo.Db.Type, appo.Db.Addr)

	if err != nil {
		panic(err)
	}

	dbs := NewSQLTodo(db)

	//close the db connection
	// appo.OnClose = func(e *engine.Engine) {
	// 	db.Close()
	// }

	appo.Bind(NewTodoController("todo", dbs, appo))

}

//TodoApp server engine
var App = engine.NewEngine(engine.NewConfig(), initApp)
