package handler

import (
	"context"

	"github.com/anacondaf/petSocialAPI/src/api/host/domain"
	"github.com/google/uuid"
)

type GetPetByIdQuery struct {
	Id uuid.UUID `validate:"required"`
}

type GetPetByIdQueryResponse struct {
	Id     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Number string    `json:"number"`
}

type GetPetByIdQueryHandler struct {
	db *domain.Queries
}

func NewGetPetByIdQueryHandler(db *domain.Queries) *GetPetByIdQueryHandler {
	return &GetPetByIdQueryHandler{
		db: db,
	}
}

func (g GetPetByIdQueryHandler) Handle(ctx context.Context, command *GetPetByIdQuery) (*GetPetByIdQueryResponse, error) {
	pet, err := g.db.GetPetById(ctx, command.Id.String())

	if err != nil {
		return nil, err
	}

	sId, _ := uuid.Parse(pet.ID)

	res := &GetPetByIdQueryResponse{
		Id:     sId,
		Name:   pet.Name.String,
		Number: pet.Number.String,
	}

	return res, nil
}
