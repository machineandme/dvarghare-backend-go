package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	sentryecho "github.com/getsentry/sentry-go/echo"
	"github.com/labstack/echo/v4"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/sqlite3"

	_ "github.com/mattn/go-sqlite3"
)

var db *reform.DB

func init() {
	sqlDB, err := sql.Open("sqlite3", "data/fdc.db")
	if err != nil {
		log.Fatal(err)
	}
	logger := log.New(os.Stderr, "SQL: ", log.Flags())
	db = reform.NewDB(sqlDB, sqlite3.Dialect, reform.NewPrintfLogger(logger.Printf))
}

func main() {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:         "https://ad7857ded104489abd896cd353d3a91d@o423712.ingest.sentry.io/5442002",
		Release:     "0.0.1",
		DebugWriter: ioutil.Discard,
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}

	e := echo.New()
	e.Use(sentryecho.New(sentryecho.Options{
		Repanic:         false,
		WaitForDelivery: false,
		Timeout:         time.Duration(200 * time.Millisecond),
	}))

	e.GET("/", func(c echo.Context) error {
		// var ctx context.Context
		// nutrients, err := db.WithContext(ctx).SelectAllFrom(repo.Nutrient, "")
		// fmt.Printf("%v", nutrients)
		return c.String(http.StatusOK, "Hey...")
	})

	host, ok := os.LookupEnv("HOST")
	if !ok {
		host = "localhost:1324"
	}
	e.Logger.Fatal(e.Start(host))
}
