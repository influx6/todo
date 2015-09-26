package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/influx6/todo/app"
)

func main() {

	if err := app.App.Load("./config/app.yaml"); err != nil {
		if app.App.Env == "dev" {
			panic(err.Error())
		}
	}

	app.App.Addr = fmt.Sprintf("%s:%s", os.Getenv("OPENSHIFT_GO_IP"), os.Getenv("OPENSHIFT_GO_PORT"))

	log.Printf("%s is now up", app.App.Name)
	app.App.Serve()
}
