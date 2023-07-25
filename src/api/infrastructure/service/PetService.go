package service

import (
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
	//pet := &domain.Pet{
	//	Id:     uuid.New(),
	//	Name:   request.Name,
	//	Number: "",
	//}
	//
	//return pet.Id, nil

	return uuid.New(), nil
}
