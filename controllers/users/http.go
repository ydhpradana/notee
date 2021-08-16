package users

import (
	"net/http"
	"notee/business/users"
	"notee/controllers"
	"notee/controllers/users/request"

	echo "github.com/labstack/echo/v4"
)

type UserController struct {
	UserUseCase users.UseCase
}

func NewUserController(userUC users.UseCase) *UserController {
	return &UserController{
		UserUseCase: userUC,
	}
}

func (controller *UserController) Login(c echo.Context) error {
	var userLogin request.LoginUser
	if err := c.Bind(&userLogin); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	user, err := controller.UserUseCase.Login(c.Request().Context(), userLogin.Email, userLogin.Password)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, user)
}

func (controller *UserController) GetById(c echo.Context) error {
	id := c.Param("id")
	user, err := controller.UserUseCase.GetById(c.Request().Context(), id)

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
