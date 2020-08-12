package restful

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
)

type Controller struct{
	app     *iris.Application
	service *Service
}

func NewController(app *iris.Application, ser *Service) *Controller {
	c := &Controller{
		app:     app,
		service: ser,
	}

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"PUT", "POST", "GET"},
		AllowedHeaders:   []string{"*"},
	})

	v1 := app.Party("/api/v1", crs).AllowMethods(iris.MethodOptions)
	{
		v1.Post("/ping", c.Ping)
	}

	return c
}

func (controller *Controller) Ping(ctx iris.Context) {
	_, _ = ctx.JSON(iris.Map{
		"message": "pong",
	})
}
