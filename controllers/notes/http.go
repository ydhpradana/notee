package notes

import (
	"net/http"
	"notee/business/notes"
	"notee/controllers"
	"notee/controllers/notes/request"

	echo "github.com/labstack/echo/v4"
)

type NoteController struct {
	NoteUseCase notes.UseCase
}

func NewNoteController(noteUC notes.UseCase) *NoteController {
	return &NoteController{
		NoteUseCase: noteUC,
	}
}

func (controller *NoteController) GetById(c echo.Context) error {
	id := c.Param("id")
	note, err := controller.NoteUseCase.GetById(c.Request().Context(), id)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, note)
}

func (controller *NoteController) GetByCatId(c echo.Context) error {
	id := c.Param("id")
	note, err := controller.NoteUseCase.GetByCatId(c.Request().Context(), id)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, note)
}

func (controller *NoteController) GetByUserId(c echo.Context) error {
	id := c.Param("id")
	note, err := controller.NoteUseCase.GetByUserId(c.Request().Context(), id)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, note)
}

func (controller *NoteController) GetByIsFree(c echo.Context) error {
	id := c.Param("free")
	note, err := controller.NoteUseCase.GetByIsFree(c.Request().Context(), id)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, note)
}

func (controller *NoteController) GetByName(c echo.Context) error {
	search := c.Param("search")
	note, err := controller.NoteUseCase.GetByName(c.Request().Context(), search)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, note)
}

func (controller *NoteController) GetAll(c echo.Context) error {
	note, err := controller.NoteUseCase.GetAll(c.Request().Context())

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, note)
}

func (ctrl *NoteController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Notes{}
	if err := c.Bind(&req); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err := ctrl.NoteUseCase.Store(ctx, req.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, "Successfully inserted")
}

func (ctrl *NoteController) Update(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")
	req := request.Notes{}
	if err := c.Bind(&req); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err := ctrl.NoteUseCase.Update(ctx, req.ToDomain(), id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, "Successfully updated")
}

func (ctrl *NoteController) Delete(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")
	req := request.Notes{}
	if err := c.Bind(&req); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err := ctrl.NoteUseCase.Delete(ctx, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, "Successfully deleted")
}
