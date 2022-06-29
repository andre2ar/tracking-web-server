package main

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"time"
)

func main() {
	e := echo.New()

	e.GET("/ping", ping)
	e.GET("/img", img)

	e.Logger.Fatal(e.Start(":9191"))
}

func ping(c echo.Context) error {
	_, err := os.Stat("./tmp/ok")

	if err == nil {
		return c.String(http.StatusOK, "OK")
	} else if errors.Is(err, os.ErrNotExist) {
		return c.String(http.StatusServiceUnavailable, "503 Service unavailable")
	}

	return c.String(http.StatusInternalServerError, "500 Server error")
}

func img(c echo.Context) error {
	simpleLogger("Image requested")
	return c.File("./assets/fine.gif")
}

func simpleLogger(message string) {
	currentTime := time.Now()

	fmt.Println(currentTime.Format("2006-01-02 15:04:05") + ": " + message)
}
