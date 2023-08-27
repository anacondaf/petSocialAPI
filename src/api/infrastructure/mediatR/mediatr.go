package mediatR

import (
	handler "github.com/anacondaf/petSocialAPI/src/api/application/handler/pet"
	"github.com/anacondaf/petSocialAPI/src/api/host/domain"
	"github.com/mehdihadeli/go-mediatr"
	"go.uber.org/zap"
)

func RegisterHandler(db *domain.Queries, logger *zap.Logger) {
	getPetByIdQueryHandler := handler.NewGetPetByIdQueryHandler(db)

	err := mediatr.RegisterRequestHandler[*handler.GetPetByIdQuery, *handler.GetPetByIdQueryResponse](getPetByIdQueryHandler)
	if err != nil {
		logger.Fatal(err.Error())
	}
}
