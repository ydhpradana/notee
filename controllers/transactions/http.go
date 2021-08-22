package transactions

import (
	"net/http"
	"notee/app/middleware"
	"notee/business/transactions"
	"notee/controllers"
	"notee/controllers/transactions/request"
	"notee/controllers/transactions/response"

	echo "github.com/labstack/echo/v4"
)

type TransactionController struct {
	TransactionUseCase transactions.UseCase
	jwtAuth        *middleware.ConfigJWT
}

func NewTransactionController(transactionUC transactions.UseCase) *TransactionController {
	return &TransactionController{
		TransactionUseCase: transactionUC,
	}
}

func (ctrl *TransactionController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Transactions{}
	if err := c.Bind(&req); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err := ctrl.TransactionUseCase.Store(ctx, req.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, "Successfully inserted")
}

func (controller *TransactionController) GetTrx(c echo.Context) error {
	ctx := c.Request().Context()
	user, err := controller.jwtAuth.GetUser(c)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	resp,  err := controller.TransactionUseCase.GetTrx(ctx, user.ID)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	responseController := []response.Transaction{}
	for _, value := range resp {
	responseController = append(responseController, response.FromDomain(value))
}

	return controllers.NewSuccessResponse(c, responseController)
}
