package main

import (
	"fmt"
	"github.com/FireStack-Lab/pocket4d-server/restful"
	"github.com/FireStack-Lab/pocket4d-server/storage"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	store := storage.NewStore("bundled")
	service := restful.NewService(store)
	restful.NewController(app, service)
	_ = app.Run(iris.Addr(fmt.Sprintf("%s:%d", "0.0.0.0", 8081)))
}
