package ratings

import (
	"net/http"
	"notee/business/ratings"
	"notee/controllers"
	"notee/controllers/ratings/request"
	"notee/controllers/ratings/response"
	"strconv"

	echo "github.com/labstack/echo/v4"
)

type RatingController struct {
	RatingUseCase ratings.UseCase
}

func NewRatingController(ratingUC ratings.UseCase) *RatingController {
	return &RatingController{
		RatingUseCase: ratingUC,
	}
}

func (controller *RatingController) GetById(c echo.Context) error {
	id := c.Param("id")
	id_param, _ := strconv.Atoi(id)
	rating, err := controller.RatingUseCase.GetById(c.Request().Context(), id_param)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	responseController := []response.Rating{}
	for _, value := range rating {
		responseController = append(responseController, response.FromDomain(value))
	}

	return controllers.NewSuccessResponse(c, responseController)
}

func (ctrl *RatingController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Ratings{}
	if err := c.Bind(&req); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err := ctrl.RatingUseCase.Store(ctx, req.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, "Successfully inserted")
}

func (ctrl *RatingController) Update(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")
	req := request.Ratings{}
	if err := c.Bind(&req); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err := ctrl.RatingUseCase.Update(ctx, req.ToDomain(), id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, "Successfully Updated")
}

func (ctrl *RatingController) Delete(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")
	req := request.Ratings{}
	if err := c.Bind(&req); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err := ctrl.RatingUseCase.Delete(ctx, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, "Successfully Updated")
}
