package main

import (
	"assessment/config"
	"assessment/db/sqlc"
	"context"
	"database/sql"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var qr *db.Queries

type Err struct {
	Message string `json:"message"`
}

func InsertExpenses(c echo.Context) error {
	var ex db.InsertExpensesParams

	err := c.Bind(&ex)

	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	dbResult, err := qr.InsertExpenses(context.Background(), ex)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, dbResult)
}
func UpdateExpenses(c echo.Context) error {
	var ex db.UpdateExpensesParams

	err := c.Bind(&ex)

	log.Println(ex)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	id := c.Param("id")
	intValue := 0

	_, err = fmt.Sscan(id, &intValue)

	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	ex.ID = int32(intValue)

	dbResult, err := qr.UpdateExpenses(context.Background(), ex)

	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dbResult)
}

func GetExpenses(c echo.Context) error {

	id := c.Param("id")
	intValue := 0

	_, err := fmt.Sscan(id, &intValue)

	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	ex, err := qr.GetExpenses(context.Background(), int32(intValue))
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, ex)
}

func GetAllExpenses(c echo.Context) error {

	exList, err := qr.ListExpenses(context.Background())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, exList)
}

func main() {

	config.InitViper(".")
	cf := config.GetConfig()

	e := echo.New()
	conn, err := sql.Open(cf.DbDriver, cf.DbSource)
	if err != nil {
		log.Fatal("Can not connect to db: ", err)
	}
	qr = db.New(conn)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/expenses", InsertExpenses)
	e.GET("/expenses/:id", GetExpenses)
	e.PUT("/expenses/:id", UpdateExpenses)
	e.GET("/expenses", GetAllExpenses)

	log.Fatal(e.Start(cf.SrPort))

	go func() {
		if err := e.Start(cf.SrPort); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("Shutting down the server")
		}
	}()

	shutdown := make(chan os.Signal, 1)

	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	<-shutdown
	fmt.Println("Shutting down")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		fmt.Println("Shutting ERROR", err)
	}
	fmt.Println("Bye Bye")
}
