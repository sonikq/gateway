package admin

import (
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/models/admin"
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/repositories"
)

type DeleteUserMatrixByAdminService struct {
	repo repositories.IDeleteUserMatrixByAdminRepo
}

func NewDeleteAccessService(repo repositories.IDeleteUserMatrixByAdminRepo) *DeleteUserMatrixByAdminService {
	return &DeleteUserMatrixByAdminService{
		repo: repo,
	}
}

func (s *DeleteUserMatrixByAdminService) DeleteUserMatrixByAdmin(request admin.DeleteUserAccessRequest, response chan admin.DeleteUserAccessResponse) {
	result := s.repo.DeleteUserMatrixByAdmin(request)

	response <- admin.DeleteUserAccessResponse{
		Code:   result.Code,
		Status: result.Status,
		Error:  result.Error,
	}
}
