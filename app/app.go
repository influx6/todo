package app

import "github.com/influx6/relay/engine"

//Applo server engine
var App = engine.NewEngine(engine.NewConfig(), func(appo *engine.Engine) {

	// appo.OnInit = func(e *engine.Engine) {
	// 	e.ServeFile("/", "index.html")
	// }

})
