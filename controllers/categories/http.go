package categories

import (
	"net/http"
	"notee/business/categories"
	"notee/controllers"
	"notee/controllers/categories/request"

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

	return controllers.NewSuccessResponse(c, category)
}

func (controller *CategoryController) GetAll(c echo.Context) error {
	category, err := controller.CategoryUseCase.GetAll(c.Request().Context())

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, category)
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
