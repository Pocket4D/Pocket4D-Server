package restful

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"log"
)

type ErrorRep struct {
	ErrorCode    int    `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}

type SuccessRep struct {
	AppId string `json:"app_id"`
	Name  string `json:"name"`
}

type Controller struct {
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
		v1.Post("/bundled/{name}", c.UploadMiniBundled)
		v1.Get("/bundled",c.All)
	}

	return c
}

func (controller *Controller) Ping(ctx iris.Context) {
	_, _ = ctx.JSON(iris.Map{
		"message": "pong",
	})
}

func (controller *Controller) All(ctx iris.Context) {
	err, apps := controller.service.Scan()
	if err != nil {
		_, _ = ctx.JSON(&ErrorRep{
			ErrorCode:    400,
			ErrorMessage: err.Error(),
		})
		return
	}

	_, _ = ctx.JSON(apps)
	return

}

func (controller *Controller) UploadMiniBundled(ctx iris.Context) {
	file, info, err := ctx.FormFile("bundled")
	name := ctx.Params().GetString("name")

	if err != nil {
		_, _ = ctx.JSON(&ErrorRep{
			ErrorCode:    400,
			ErrorMessage: err.Error(),
		})
		return
	}
	defer file.Close()
	buffer := make([]byte, info.Size)
	_, _ = file.Read(buffer)
	bundled := string(buffer)
	log.Printf("get mini program bundled: %s\n", bundled)

	err1, appId := controller.service.Save(name, bundled)
	if err1 != nil {
		_, _ = ctx.JSON(&ErrorRep{
			ErrorCode:    400,
			ErrorMessage: err1.Error(),
		})
		return
	}

	_, _ = ctx.JSON(&SuccessRep{
		AppId: appId,
		Name:  name,
	})
	return
}
