package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/url-shortener/internal/shortener"
)

type cache interface {
	Set(key string, value interface{}, expires time.Duration) error
	Get(key string) (string, error)
}

type server struct {
	echo  *echo.Echo
	port  string
	cache cache
}

func (server *server) RegisterRoutes() {
	api := server.echo.Group("/v1/api")

	api.POST("/url", func(c echo.Context) error {
		values, err := c.FormParams()

		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "cannot parse form data")
		}

		if values.Has("url") {
			url := values.Get("url")
			shortUrl := shortener.GenerateString()
			host := c.Request().Host
			scheme := c.Request().URL.Scheme

			server.cache.Set(shortUrl, url, time.Second*30)

			return c.String(http.StatusOK, fmt.Sprintf("%s/%s", scheme+host, shortUrl))
		}

		return echo.NewHTTPError(http.StatusBadRequest, "missing url parameter")
	})

	server.echo.GET("/:url", func(c echo.Context) error {
		shortUrl := c.Param("url")

		url, err := server.cache.Get(shortUrl)

		if err != nil {
			return echo.NotFoundHandler(c)
		}

		return c.Redirect(http.StatusPermanentRedirect, url)
	})
}

func (server *server) Start() error {
	server.RegisterRoutes()
	err := server.echo.Start("localhost" + server.port)

	if err != nil {
		server.echo.Logger.Fatal(err)
		return err
	}

	return nil
}

func New(port string, c cache) *server {
	server := &server{
		echo:  echo.New(),
		port:  port,
		cache: c,
	}

	echo.NotFoundHandler = func(c echo.Context) error {
		fmt.Println("hre")

		return c.File("static/404.html")
	}

	server.echo.Use(middleware.Logger())
	server.echo.Use(middleware.CORS())

	server.echo.Static("/", "static")

	return server
}
