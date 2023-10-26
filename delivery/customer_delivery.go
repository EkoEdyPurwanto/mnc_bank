package delivery

import (
	"EkoEdyPurwanto/mnc-bank/model/req"
	"EkoEdyPurwanto/mnc-bank/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CustomerDelivery struct {
	customerUC usecase.CustomerUseCase
	engine     *echo.Echo
}

func NewCustomerDelivery(customerUC usecase.CustomerUseCase, engine *echo.Echo) *CustomerDelivery {
	return &CustomerDelivery{
		customerUC: customerUC,
		engine:     engine,
	}
}

func (c *CustomerDelivery) AuthRoute() {
	rg := c.engine.Group("/api/v1")

	rg.POST("/auth/register", c.registerHandler)
	rg.POST("/auth/login", c.loginHandler)
	rg.POST("/auth/logout", c.logoutHandler)
}

func (c *CustomerDelivery) registerHandler(ctx echo.Context) error {
	// Parse the request body to get payload from dto
	var payload req.RegisterRequest
	if err := ctx.Bind(&payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid request body")
	}

	// Call the Register method from the use case layer
	if err := c.customerUC.Register(payload); err != nil {
		return ctx.JSON(http.StatusConflict, err.Error())
	}

	// Return the data payload upon successful login
	return ctx.JSON(http.StatusOK, "Registration successful")
}

func (c *CustomerDelivery) loginHandler(ctx echo.Context) error {
	// Parse the request body to get payload from dto
	var payload req.LoginRequest
	if err := ctx.Bind(&payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid request body")
	}

	// Call the Login method from the use case layer
	token, err := c.customerUC.Login(payload)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, "Unauthorized")
	}

	// Return the token upon successful login
	return ctx.JSON(http.StatusOK, map[string]string{"token": token})
}

func (c *CustomerDelivery) logoutHandler(ctx echo.Context) error {
	// Implement payment logic here
	panic("implement me")
}
