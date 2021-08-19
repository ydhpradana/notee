package categories

import (
	"net/http"
	"notee/business/categories"
	"notee/controllers"
	"notee/controllers/categories/request"
	"notee/controllers/categories/response"

	echo "github.com/labstack/echo/v4"
)

type CategoryController struct {
	CategoryUseCase categories.UseCase
}

func NewCategoryController(categoryUC categories.UseCase) *CategoryController {
	return &CategoryController{
		CategoryUseCase: categoryUC,
	}
}

func (controller *CategoryController) GetById(c echo.Context) error {
	id := c.Param("id")
	category, err := controller.CategoryUseCase.GetById(c.Request().Context(), id)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, category)
}

func (controller *CategoryController) GetByName(c echo.Context) error {
	search := c.Param("search")
	category, err := controller.CategoryUseCase.GetByName(c.Request().Context(), search)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	responseController := []response.Category{}
	for _, value := range category {
		responseController = append(responseController, response.FromDomain(value))
	}

	return controllers.NewSuccessResponse(c, responseController)
}

func (controller *CategoryController) GetAll(c echo.Context) error {
	category, err := controller.CategoryUseCase.GetAll(c.Request().Context())

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	responseController := []response.Category{}
	for _, value := range category {
		responseController = append(responseController, response.FromDomain(value))
	}

	return controllers.NewSuccessResponse(c, responseController)
}

func (ctrl *CategoryController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Categories{}
	if err := c.Bind(&req); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err := ctrl.CategoryUseCase.Store(ctx, req.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, "Successfully inserted")
}

func (ctrl *CategoryController) Update(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")
	req := request.Categories{}
	if err := c.Bind(&req); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err := ctrl.CategoryUseCase.Update(ctx, req.ToDomain(), id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, "Successfully updated")
}

func (ctrl *CategoryController) Delete(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")
	req := request.Categories{}
	if err := c.Bind(&req); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err := ctrl.CategoryUseCase.Delete(ctx, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, "Successfully deleted")
}
