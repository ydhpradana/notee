package routes

import (
	"notee/controllers/categories"
	"notee/controllers/users"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware      middleware.JWTConfig
	UserController     users.UserController
	// NewsController     news.NewsController
	CategoryController  categories.CategoryController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	users := e.Group("users")
	users.POST("/register", cl.UserController.Store)
	users.GET("/token", cl.UserController.CreateToken)
	users.POST("/login", cl.UserController.Login)
	users.GET("/:id", cl.UserController.GetById)

	category := e.Group("category")
	category.POST("/store", cl.CategoryController.Store)
	category.GET("/search/:search", cl.CategoryController.GetByName)
	category.GET("", cl.CategoryController.GetAll)
	category.GET("/id/:id", cl.CategoryController.GetById)
	// category.GET("/list", cl.CategoryController.GetAll, middleware.JWTWithConfig(cl.JWTMiddleware))

	// news := e.Group("news")
	// news.POST("/store", cl.NewsController.Store, middleware.JWTWithConfig(cl.JWTMiddleware))
	// news.PUT("/update", cl.NewsController.Update, middleware.JWTWithConfig(cl.JWTMiddleware))
}
