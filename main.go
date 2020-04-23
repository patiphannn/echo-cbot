package main

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/polnoy/echo-cbot/common"
	"github.com/polnoy/echo-cbot/controllers"
	"github.com/spf13/viper"
)

// CustomValidator validator struct
type CustomValidator struct {
	validator *validator.Validate
}

// Validate handle validator
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func init() {
	common.ConnectDb()
}

func main() {
	// Echo instance
	e := echo.New()

	// Validator
	e.Validator = &CustomValidator{validator: validator.New()}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	// Auth
	{
		auth := new(controllers.Auth)
		g := e.Group("/auth")
		{
			g.POST("/signup", auth.Signup)
			g.POST("/signin", auth.Signin)
		}
	}

	// Project
	{
		project := new(controllers.Project)
		g := e.Group("/projects")
		g.Use(middleware.JWT([]byte(viper.GetString("access_key"))))
		{
			g.GET("", project.Gets)
			g.GET("/:id", project.Get)
			g.POST("", project.Create)
			g.PUT("/:id", project.Update)
			g.DELETE("/:id", project.Delete)
		}
	}

	// Intent
	{
		intent := new(controllers.Intent)
		g := e.Group("/intents")
		g.Use(middleware.JWT([]byte(viper.GetString("access_key"))))
		{
			g.GET("", intent.Gets)
			g.GET("/:id", intent.Get)
			g.POST("", intent.Create)
			g.PUT("/:id", intent.Update)
			g.DELETE("/:id", intent.Delete)
		}
	}

	// Reponse
	{
		response := new(controllers.Response)
		g := e.Group("/reponses")
		g.Use(middleware.JWT([]byte(viper.GetString("access_key"))))
		{
			g.POST("", response.Answer)
		}
	}

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
