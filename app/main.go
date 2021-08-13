package main

import (
	"log"
	_userUseCase "notee/business/users"
	_userController "notee/controllers/users"
	_userRepo "notee/drivers/databases/users"

	_dbHelper "notee/helpers/databases"

	"time"

	"github.com/labstack/echo"
	"github.com/spf13/viper"
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


func main() {
	configdb := _dbHelper.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}
	db := configdb.InitDB()
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	e := echo.New()

	userRepo := _userRepo.NewMysqlUserRepository(db)
	newUserUsecase := _userUseCase.NewUserUsecase(userRepo, timeoutContext)
	_userController.NewUserController(e, newUserUsecase)

	// eAuth := e.Group("")
	// eAuth.Use(middleware.JWT([]byte(viper.GetString(`jwt.Key`)))
	// // jwt
	// newsRepo := _newsRepo.NewMySQLNewsRepository(db)
	// newsUsecase := _newsUsecase.NewNewsUsecase(newsRepo, categoryUsecase, ipLocator, timeoutContext)
	// _newsController.NewNewsController(e, newsUsecase)
	log.Fatal(e.Start(viper.GetString("server.address")))
}
