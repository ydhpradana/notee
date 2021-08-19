package routes

import (
	"notee/controllers/categories"
	"notee/controllers/notes"
	"notee/controllers/transactions"
	"notee/controllers/users"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware      middleware.JWTConfig
	UserController     users.UserController
	// NewsController     news.NewsController
	CategoryController  categories.CategoryController
	NoteController  	notes.NoteController
	TransactionController  	transactions.TransactionController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	users := e.Group("users")
	users.POST("/register", cl.UserController.Store)
	users.GET("/token", cl.UserController.CreateToken)
	users.POST("/login", cl.UserController.Login)
	users.GET("/profile/:id", cl.UserController.GetById)
	users.GET("/profile", cl.UserController.FindByToken, middleware.JWTWithConfig(cl.JWTMiddleware))

	category := e.Group("category")
	category.POST("/store", cl.CategoryController.Store, middleware.JWTWithConfig(cl.JWTMiddleware))
	category.GET("/search/:search", cl.CategoryController.GetByName, middleware.JWTWithConfig(cl.JWTMiddleware))
	category.GET("", cl.CategoryController.GetAll, middleware.JWTWithConfig(cl.JWTMiddleware))
	category.GET("/id/:id", cl.CategoryController.GetById, middleware.JWTWithConfig(cl.JWTMiddleware))
	category.PUT("/update/:id", cl.CategoryController.Update, middleware.JWTWithConfig(cl.JWTMiddleware))
	category.DELETE("/delete/:id", cl.CategoryController.Delete, middleware.JWTWithConfig(cl.JWTMiddleware))
	
	note := e.Group("note")
	note.POST("/store", cl.NoteController.Store, middleware.JWTWithConfig(cl.JWTMiddleware))
	note.GET("/search/:search", cl.NoteController.GetByName, middleware.JWTWithConfig(cl.JWTMiddleware))
	note.GET("", cl.NoteController.GetAll, middleware.JWTWithConfig(cl.JWTMiddleware))
	note.GET("/id/:id", cl.NoteController.GetById, middleware.JWTWithConfig(cl.JWTMiddleware))
	note.GET("/user/:id", cl.NoteController.GetByUserId, middleware.JWTWithConfig(cl.JWTMiddleware))
	note.GET("/cat/:id", cl.NoteController.GetByCatId, middleware.JWTWithConfig(cl.JWTMiddleware))
	note.GET("/isfree/:free", cl.NoteController.GetByIsFree, middleware.JWTWithConfig(cl.JWTMiddleware))
	note.PUT("/update/:id", cl.NoteController.Update, middleware.JWTWithConfig(cl.JWTMiddleware))
	note.DELETE("/delete/:id", cl.NoteController.Delete, middleware.JWTWithConfig(cl.JWTMiddleware))
	// category.GET("/list", cl.CategoryController.GetAll, middleware.JWTWithConfig(cl.JWTMiddleware))

	trx := e.Group("transaction")
	trx.POST("/store", cl.TransactionController.Store, middleware.JWTWithConfig(cl.JWTMiddleware))
	// trx.GET("/search/:search", cl.TranactionController.GetByName, middleware.JWTWithConfig(cl.JWTMiddleware))
	 trx.GET("", cl.TransactionController.GetTrx, middleware.JWTWithConfig(cl.JWTMiddleware))
	// trx.GET("/id/:id", cl.TranactionController.GetById, middleware.JWTWithConfig(cl.JWTMiddleware))
	// trx.GET("/user/:id", cl.TranactionController.GetByUserId, middleware.JWTWithConfig(cl.JWTMiddleware))
	// trx.GET("/cat/:id", cl.TranactionController.GetByCatId, middleware.JWTWithConfig(cl.JWTMiddleware))
	// trx.GET("/isfree/:free", cl.TranactionController.GetByIsFree, middleware.JWTWithConfig(cl.JWTMiddleware))

	// news := e.Group("news")
	// news.POST("/store", cl.NewsController.Store, middleware.JWTWithConfig(cl.JWTMiddleware))
	// news.PUT("/update", cl.NewsController.Update, middleware.JWTWithConfig(cl.JWTMiddleware))
}
