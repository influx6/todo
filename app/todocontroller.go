package app

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/influx6/relay"
	"github.com/influx6/relay/engine"
)

type jsonTask struct {
	ID        int    `json:"id"`
	Task      string `json:"task"`
	Desc      string `json:"desc"`
	Stamp     int64  `json:"stamp"`
	Day       int    `json:"day"`
	Month     string `json:"month"`
	Completed int    `json:"month"`
}

// TodoController provides a controller for the todo service
type TodoController struct {
	*relay.Controller
	db   TodoDatabase
	home *relay.AssetTemplate
	edit *relay.AssetTemplate
}

// NewTodoController returns a new TodoController
func NewTodoController(path string, db TodoDatabase, app *engine.Engine) *TodoController {
	tc := &TodoController{
		Controller: relay.NewController(path),
		db:         db,
	}

	tc.home = app.Template.MustCreate("home.render", []string{"./home", "./layout"}, nil)

	tc.register()
	return tc
}

// Index handles the request for /
func (c *TodoController) Index(req *relay.HTTPRequest) {
	todos, err := c.db.FindAll()

	if err != nil {
		req.Res.WriteHeader(504)
		return
	}

	_, err = req.Write(relay.HTMLRender(200, "layout", todos, c.home.Tmpl))

	if err != nil {
		req.Res.WriteHeader(504)
	}
}

//Save handles save requests
func (c *TodoController) Save(req *relay.HTTPRequest) {
	message, err := req.Message()

	if err != nil {
		req.Res.WriteHeader(404)
		req.Write("Err decoding body!")
	}

	if message.MessageType == "body" {
		todo := jsonTask{}

		if err := json.NewEncoder(bytes.NewBuffer(message.Payload)).Encode(&todo); err != nil {
			req.Res.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := c.db.New(todo.Task, todo.Desc, todo.Stamp); err != nil {
			req.Res.WriteHeader(http.StatusExpectationFailed)
			return
		}

		req.Res.WriteHeader(http.StatusCreated)
		return
	}

	task := message.Form.Get("task")
	desc := message.Form.Get("description")

	if task == "" || desc == "" {
		req.Res.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Printf("Submitting: %s %s", task, desc)
	if err := c.db.New(task, desc, 0); err != nil {
		req.Res.WriteHeader(http.StatusExpectationFailed)
		return
	}

	c.Render("/todo", req.Res, req.Req, req.Params)
}

//UpdateAsComplete handles save requests
func (c *TodoController) UpdateAsComplete(req *relay.HTTPRequest) {
	id, err := strconv.Atoi(req.Params.Get("id").(string))

	if err != nil {
		req.Res.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = c.db.Complete(id); err != nil {
		req.Res.WriteHeader(http.StatusExpectationFailed)
		return
	}

	req.Res.WriteHeader(http.StatusOK)
}

//UpdateAsUncomplete handles save requests
func (c *TodoController) UpdateAsUncomplete(req *relay.HTTPRequest) {
	id, err := strconv.Atoi(req.Params.Get("id").(string))

	if err != nil {
		req.Res.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = c.db.Uncomplete(id); err != nil {
		req.Res.WriteHeader(http.StatusExpectationFailed)
		return
	}

	req.Res.WriteHeader(http.StatusOK)
}

//Delete handles save requests
func (c *TodoController) Delete(req *relay.HTTPRequest) {
	id, err := strconv.Atoi(req.Params.Get("id").(string))

	if err != nil {
		req.Res.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = c.db.Destroy(id); err != nil {
		req.Res.WriteHeader(http.StatusExpectationFailed)
		return
	}

	req.Res.WriteHeader(http.StatusOK)
}

func (c *TodoController) register() {
	c.BindHTTP("post", `/`, c.Save, relay.BasicHTTPCodec)
	c.BindHTTP("get head", "/", c.Index, relay.UseHTTPEncoder(relay.HTMLEncoder))
	c.BindHTTP("post", `/{id:[\d+]}`, c.UpdateAsComplete, relay.BasicHTTPCodec)
	c.BindHTTP("put", `/{id:[\d+]}`, c.UpdateAsUncomplete, relay.BasicHTTPCodec)
	c.BindHTTP("delete", `/{id:[\d+]}`, c.Delete, relay.BasicHTTPCodec)
}
