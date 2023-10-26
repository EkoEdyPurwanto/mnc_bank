package main

import (
	"EkoEdyPurwanto/mnc-bank/delivery"
	"EkoEdyPurwanto/mnc-bank/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	e := echo.New()
	customerUC, err := usecase.NewCustomerUseCase("../repository/customer.json")
	if err != nil {
		log.Fatal(err)
	}

	// Use middleware for logging and CORS (Cross-Origin Resource Sharing)
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	delivery.NewCustomerDelivery(customerUC, e).AuthRoute()
	log.Fatal(e.Start(":1323"))
}
