package admin

import (
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/models/admin"
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/repositories"
)

type AddUserMatrixService struct {
	repo repositories.IAddUserMatrixByAdminRepo
}

func NewAddUserMatrixByAdminService(repo repositories.IAddUserMatrixByAdminRepo) *AddUserMatrixService {
	return &AddUserMatrixService{
		repo: repo,
	}
}

func (s *AddUserMatrixService) AddUserMatrixByAdmin(request admin.AddUserMatrixByAdminRequest, response chan admin.AddUserMatrixByAdminResponse) {
	result := s.repo.AddUserMatrixByAdmin(request)

	response <- admin.AddUserMatrixByAdminResponse{
		Code:   result.Code,
		Status: result.Status,
		Error:  result.Error,
	}
}
