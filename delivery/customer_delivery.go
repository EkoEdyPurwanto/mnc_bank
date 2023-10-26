package delivery

import (
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
	// Parse the request body to get registration data
	var registrationData struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := ctx.Bind(&registrationData); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid request body")
	}

	// Call the Register method from the use case layer
	if err := c.customerUC.Register(registrationData.Name, registrationData.Email, registrationData.Password); err != nil {
		return ctx.JSON(http.StatusConflict, err.Error())
	}

	return ctx.JSON(http.StatusOK, "Registration successful")
}

func (c *CustomerDelivery) loginHandler(ctx echo.Context) error {
	//TODO implement login logic
	return ctx.JSON(http.StatusOK, "Login successful")
}

func (c *CustomerDelivery) logoutHandler(ctx echo.Context) error {
	//TODO implement logout logic
	return ctx.JSON(http.StatusOK, "Logout successful")
}
