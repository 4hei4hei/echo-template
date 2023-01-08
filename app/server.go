package main

import (
	"database/sql"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type dbCfg struct {
	dbname string
	dbuser string
	passwd string
	params string
}

func main() {
	e := echo.New()

	welcomePage(e)
	healthCheck(e)
	e.Logger.Fatal(e.Start(":8080"))
}

// Load db config from environment variables
func loadDbCfg() *dbCfg {
	cfg := new(dbCfg)
	err := godotenv.Load("./.env")
	if err != nil {
		log.Printf("Cloud not load ./.env: %v", err)
	}

	cfg.dbname = os.Getenv("MYSQL_DBNAME")
	cfg.dbuser = os.Getenv("MYSQL_DBUSER")
	cfg.passwd = os.Getenv("MYSQL_PASSWD")
	cfg.params = os.Getenv("MYSQL_PARAMS")

	return cfg
}

// Application health check
func appHealthCheck() string {
	var appHealth string
	url := "http://localhost:8080"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("%+v", err)
	}

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("%+v", err)
	}
	defer resp.Body.Close()

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("%+v", err)
		appHealth = "NG"
	} else if string(byteArray) == "Welcome to echo" {
		appHealth = "OK"
	} else {
		appHealth = "NG"
	}

	return appHealth
}

// DB health check
func dbHealthCheck() string {
	var dbHealth string
	cfg := loadDbCfg()
	sqlOpenArg := cfg.dbuser + ":" + cfg.passwd + "@(mysql:3306)/" + cfg.dbname + "?" + cfg.params

	db, err := sql.Open("mysql", sqlOpenArg)
	if err != nil {
		log.Printf("%+v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Printf("%+v", err)
		dbHealth = "NG"
	} else {
		dbHealth = "OK"
	}

	return dbHealth
}

// Welcome page
func welcomePage(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to echo")
	})
}

// Health check
func healthCheck(e *echo.Echo) {
	e.GET("/health", func(c echo.Context) error {
		appHealth := appHealthCheck()
		dbHealth := dbHealthCheck()
		return c.JSON(http.StatusOK, map[string]string{"Application status": appHealth, "Database connectivity status": dbHealth})
	})
}
