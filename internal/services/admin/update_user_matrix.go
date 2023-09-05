package admin

import (
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/models/admin"
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/repositories"
)

type UpdateUserMatrix struct {
	repo repositories.IUpdateUserMatrixByAdminRepo
}

func NewUpdateUserMatrixService(repo repositories.IUpdateUserMatrixByAdminRepo) *UpdateUserMatrix {
	return &UpdateUserMatrix{
		repo: repo,
	}
}

func (s *UpdateUserMatrix) UpdateUserMatrixByAdminUser(request admin.UpdateUserAccessRequest, response chan admin.UpdateUserAccessResponse) {
	result := s.repo.UpdateUserMatrixByAdmin(request)

	response <- admin.UpdateUserAccessResponse{
		Code:   result.Code,
		Status: result.Status,
		Error:  result.Error,
	}
}
