package main

import (
	"os"

	"github.com/kobutomo/react-catchup-server/server/adapter/controllers"
	"github.com/kobutomo/react-catchup-server/server/adapter/gateways"
	"github.com/kobutomo/react-catchup-server/server/application/usecases"
	"github.com/kobutomo/react-catchup-server/server/infrastructure/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/dig"
	"go.uber.org/multierr"
)

func main() {
	e := echo.New()
	setMiddleware(e)

	c, err := inject()
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	if err := setRouter(e, c); err != nil {
		println(err.Error())
		os.Exit(1)
	}

	e.Logger.Fatal(e.Start(":1323"))
}

func setRouter(e *echo.Echo, c *dig.Container) error {
	return c.Invoke(func(
		sampleController controllers.ISampleController,
	) {
		e.GET("/", sampleController.Get)
	})
}

func setMiddleware(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: middleware.DefaultSkipper,
		Format: `{"time":"${time_custom}","id":"${id}","remote_ip":"${remote_ip}",` +
			`"host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}",` +
			`"status":${status},"error":"${error}"` +
			`,"bytes_in":${bytes_in},"bytes_out":${bytes_out}}` + "\n",
		CustomTimeFormat: "2006-01-02 15:04:05.00000",
	}))

	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
}

func inject() (*dig.Container, error) {
	c := dig.New()

	var err error
	// controller
	multierr.Append(err, c.Provide(controllers.NewSampleController))
	// usecase
	multierr.Append(err, c.Provide(usecases.NewSampleInteractor))
	// gateway
	multierr.Append(err, c.Provide(gateways.NewUserGateway))
	//other
	multierr.Append(err, c.Provide(mysql.NewMySQL))
	if err != nil {
		return nil, err
	}

	return c, nil
}
