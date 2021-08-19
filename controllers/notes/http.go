package notes

import (
	"net/http"
	"notee/app/middleware"
	"notee/business/notes"
	"notee/controllers"
	"notee/controllers/notes/request"
	_response "notee/controllers/notes/response"

	echo "github.com/labstack/echo/v4"
)

type NoteController struct {
	NoteUseCase notes.UseCase
	jwtAuth 	*middleware.ConfigJWT
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

	return controllers.NewSuccessResponse(c, _response.FromDomain(note))
}

func (controller *NoteController) GetByCatId(c echo.Context) error {
	id := c.Param("id")
	note, err := controller.NoteUseCase.GetByCatId(c.Request().Context(), id)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	responseController := []_response.Note{}
	for _, value := range note {
		responseController = append(responseController, _response.FromDomain(value))
	}

	return controllers.NewSuccessResponse(c, responseController)
}

func (controller *NoteController) GetByUserId(c echo.Context) error {
	id := c.Param("id")
	note, err := controller.NoteUseCase.GetByUserId(c.Request().Context(), id)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, note)
}

func (controller *NoteController) GetNote(c echo.Context) error {
	ctx := c.Request().Context()
	user, err := controller.jwtAuth.GetUser(c)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	resp, err := controller.NoteUseCase.GetNote(ctx, user.ID)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	responseController := []_response.Note{}
	for _, value := range resp {
		responseController = append(responseController, _response.FromDomain(value))
	}

	return controllers.NewSuccessResponse(c, responseController)
}

func (controller *NoteController) GetByIsFree(c echo.Context) error {
	id := c.Param("free")
	note, err := controller.NoteUseCase.GetByIsFree(c.Request().Context(), id)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	responseController := []_response.Note{}
	for _, value := range note {
		responseController = append(responseController, _response.FromDomain(value))
	}

	return controllers.NewSuccessResponse(c, responseController)
}

func (controller *NoteController) GetByName(c echo.Context) error {
	search := c.Param("search")
	note, err := controller.NoteUseCase.GetByName(c.Request().Context(), search)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	responseController := []_response.Note{}
	for _, value := range note {
		responseController = append(responseController, _response.FromDomain(value))
	}

	return controllers.NewSuccessResponse(c, responseController)
}

func (controller *NoteController) GetAll(c echo.Context) error {
	note, err := controller.NoteUseCase.GetAll(c.Request().Context())

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	responseController := []_response.Note{}
	for _, value := range note {
		responseController = append(responseController, _response.FromDomain(value))
	}

	return controllers.NewSuccessResponse(c, responseController)
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
