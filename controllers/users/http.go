package users

import (
	"fmt"
	"net/http"
	"notee/app/middleware"
	"notee/business/users"
	"notee/controllers"
	"notee/controllers/users/request"
	"notee/controllers/users/response"
	"strconv"

	echo "github.com/labstack/echo/v4"
)

type UserController struct {
	UserUseCase users.UseCase
	jwtAuth      *middleware.ConfigJWT
}

func NewUserController(userUC users.UseCase) *UserController {
	return &UserController{
		UserUseCase: userUC,
	}
}

func (controller *UserController) Login(c echo.Context) error {
	ctx := c.Request().Context()
	req := request.LoginUser{}
	if err := c.Bind(&req); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	
	//user, err := controller.UserUseCase.Login(c.Request().Context(), userLogin.Email, userLogin.Password)
	token, err := controller.UserUseCase.CreateToken(ctx, req.Email, req.Password)
	
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	response := struct {
		Token string `json:"token"`
	}{Token: token}

	return controllers.NewSuccessResponse(c, response)
}

func (controller *UserController) GetById(c echo.Context) error {
	id := c.Param("id")
	id_param, err := strconv.Atoi(id)
	user, err := controller.UserUseCase.GetById(c.Request().Context(), id_param)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, user)
}

func (ctrl *UserController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Users{}
	if err := c.Bind(&req); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err := ctrl.UserUseCase.Store(ctx, req.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, "Successfully inserted")
}

func (ctrl *UserController) CreateToken(c echo.Context) error {
	ctx := c.Request().Context()

	username := c.QueryParam("username")
	password := c.QueryParam("password")

	token, err := ctrl.UserUseCase.CreateToken(ctx, username, password)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	response := struct {
		Token string `json:"token"`
	}{Token: token}

	return controllers.NewSuccessResponse(c, response)
}

func (ctrl *UserController) FindByToken(c echo.Context) error {
	ctx := c.Request().Context()

	user, err := ctrl.jwtAuth.GetUser(c)
	fmt.Println(user)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	res, err := ctrl.UserUseCase.GetById(ctx, user.ID)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, response.FromDomain(res))
}
