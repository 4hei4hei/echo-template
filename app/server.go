package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	welcomePage(e)
	healthCheckApp(e)
	healthCheckMysql(e)

	e.Logger.Fatal(e.Start(":8080"))
}

// Welcome page
func welcomePage(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to echo")
	})
}

// Application health check
func healthCheckApp(e *echo.Echo) {
	e.GET("/health/app", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"Application status": "OK"})
	})
}

// MySQL health check
func healthCheckMysql(e *echo.Echo) {
	e.GET("/health/db", func(c echo.Context) error {
		db, err := sql.Open("mysql", "echo:echo@(mysql:3306)/echo")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		err = db.Ping()
		result := ""
		if err != nil {
			log.Fatal(err)
			result = "ERROR"
		} else {
			result = "OK"
		}
		return c.JSON(http.StatusOK, map[string]string{"Database connectivity status": result})
	})
}
