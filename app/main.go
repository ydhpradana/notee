package main

import (
	"fmt"
	"log"
	_userUseCase "notee/business/users"
	_userController "notee/controllers/users"
	_userRepo "notee/drivers/databases/users"

	_categoryUseCase "notee/business/categories"
	_categoryController "notee/controllers/categories"
	_categoryRepo "notee/drivers/databases/categories"

	_noteUseCase "notee/business/notes"
	_noteController "notee/controllers/notes"
	_noteRepo "notee/drivers/databases/notes"

	_transactionUseCase "notee/business/transactions"
	_transactionController "notee/controllers/transactions"
	_transactionRepo "notee/drivers/databases/transactions"

	_dbHelper "notee/helpers/databases"

	"time"

	_middleware "notee/app/middleware"
	_routes "notee/app/routes"

	echo "github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init(){
	viper.SetConfigFile(`app/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(
		// &_newsRepo.News{},
		&_categoryRepo.Category{},
		&_userRepo.User{},
		&_noteRepo.Note{},
		&_transactionRepo.Transaction{},
	)
}


func main() {
	configdb := _dbHelper.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}
	db := configdb.InitDB()
	dbMigrate(db)
	
	configJWT := _middleware.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}

	fmt.Println(configJWT)
	
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	
	e := echo.New()

	userRepo := _userRepo.NewMysqlUserRepository(db)
	newUserUsecase := _userUseCase.NewUserUsecase(userRepo, &configJWT, timeoutContext)
	userCtrl := _userController.NewUserController(newUserUsecase)
	
	
	categoryRepo := _categoryRepo.NewMysqlCategoryRepository(db)
	newCategoryUsecase := _categoryUseCase.NewCategoryUsecase(categoryRepo, timeoutContext)
	categoryCtrl := _categoryController.NewCategoryController(newCategoryUsecase)
	
	
	noteRepo := _noteRepo.NewMysqlNoteRepository(db)
	newNoteUsecase := _noteUseCase.NewNoteUsecase(noteRepo, timeoutContext)
	noteCtrl := _noteController.NewNoteController(newNoteUsecase)
	
	
	transactionRepo := _transactionRepo.NewMysqlTransactionRepository(db)
	newTransactionUsecase := _transactionUseCase.NewTransactionUsecase(transactionRepo, timeoutContext)
	transactionCtrl := _transactionController.NewTransactionController(newTransactionUsecase)

	// eAuth := e.Group("")
	// eAuth.Use(middleware.JWT([]byte(viper.GetString(`jwt.Key`)))
	// // jwt
	// newsRepo := _newsRepo.NewMySQLNewsRepository(db)
	// newsUsecase := _newsUsecase.NewNewsUsecase(newsRepo, categoryUsecase, ipLocator, timeoutContext)
	// _newsController.NewNewsController(e, newsUsecase)

	routesInit := _routes.ControllerList{
		JWTMiddleware:      configJWT.Init(),
		UserController:     *userCtrl,
		NoteController:     *noteCtrl,
		CategoryController: *categoryCtrl,
		TransactionController: *transactionCtrl,
	}
	routesInit.RouteRegister(e)
	log.Fatal(e.Start(viper.GetString("server.address")))
}
