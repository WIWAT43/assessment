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
		return c.JSON(http.StatusBadRequest, err)
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

	dbResult, err := qr.UpdateExpenses(context.Background(), ex)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, dbResult)
}

func GetExpenses(c echo.Context) error {

	id := c.Param("id")
	intValue := 0

	_, err := fmt.Sscan(id, &intValue)

	if err != nil {
		return c.JSON(http.StatusBadRequest, id)
	}

	ex, err := qr.GetExpenses(context.Background(), int32(intValue))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, ex)
}

func GetAllExpenses(c echo.Context) error {

	exList, err := qr.ListExpenses(context.Background())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, exList)
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

}
