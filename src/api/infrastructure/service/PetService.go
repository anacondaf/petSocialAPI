package service

import (
	"context"
	"database/sql"
	request "github.com/anacondaf/petSocialAPI/src/api/application/request/pet"
	"github.com/anacondaf/petSocialAPI/src/api/host/domain"
	"github.com/google/uuid"
)

type PetService struct {
	db *domain.Queries
}

func NewPetService(db *domain.Queries) *PetService {
	return &PetService{
		db: db,
	}
}

func (p PetService) CreatePet(request *request.CreatePetRequest) (uuid.UUID, error) {
	id := uuid.New()

	_, err := p.db.CreatePetAndReturnId(context.Background(), domain.CreatePetAndReturnIdParams{
		ID: id.String(),
		Name: sql.NullString{
			String: request.Name,
			Valid:  true,
		},
	})
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}
