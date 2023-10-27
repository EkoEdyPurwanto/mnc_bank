package main

import (
	"EkoEdyPurwanto/mnc-bank/delivery"
	"EkoEdyPurwanto/mnc-bank/model"
	"EkoEdyPurwanto/mnc-bank/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	e := echo.New()
	var cus []*model.Customer
	customerUC := usecase.NewCustomerUseCase(cus)

	// Use middleware for logging and CORS (Cross-Origin Resource Sharing)
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	delivery.NewCustomerDelivery(customerUC, e).AuthRoute()
	log.Fatal(e.Start(":1323"))
}
