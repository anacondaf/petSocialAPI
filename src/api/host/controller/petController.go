package controller

import (
	"context"
	"net/http"

	"github.com/anacondaf/petSocialAPI/src/api/infrastructure/service"

	handler "github.com/anacondaf/petSocialAPI/src/api/application/handler/pet"
	request "github.com/anacondaf/petSocialAPI/src/api/application/request/pet"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/mehdihadeli/go-mediatr"
)

func NewPetController(e *echo.Group, petService service.PetService) {
	petRoutes := e.Group("/pets")

	petRoutes.POST("", func(c echo.Context) error {
		req := new(request.CreatePetRequest)
		if err := c.Bind(req); err != nil {
			return echo.NewHTTPError(http.StatusNotImplemented, err)
		}

		id, err := petService.CreatePet(req)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotImplemented, err)
		}

		return c.JSON(http.StatusCreated, id)
	})

	petRoutes.GET("", func(c echo.Context) error {
		uuidStr, _ := uuid.Parse(c.QueryParam("id"))

		query := &handler.GetPetByIdQuery{
			Id: uuidStr,
		}

		ctx := context.Background()

		response, err := mediatr.Send[*handler.GetPetByIdQuery, *handler.GetPetByIdQueryResponse](ctx, query)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotImplemented, err)
		}

		if response == nil {
			return echo.ErrNotFound
		}

		return c.JSON(http.StatusOK, response)
	})

}
