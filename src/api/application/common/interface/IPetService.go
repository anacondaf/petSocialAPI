package _interface

import (
	request "github.com/anacondaf/petSocialAPI/src/api/application/request/pet"
	"github.com/google/uuid"
)

type IPetService interface {
	CreatePet(request *request.CreatePetRequest) (uuid.UUID, error)
}
