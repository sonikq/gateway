package services

import (
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/models/user_access"
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/repositories"
)

type UserAccessService struct {
	repo repositories.IUserAccessRepo
}

func NewUserAccessService(repo repositories.IUserAccessRepo) *UserAccessService {
	return &UserAccessService{
		repo: repo,
	}
}

func (s *UserAccessService) AddAccessUser(request user_access.AddUserAccessRequest, response chan user_access.AddUserAccessResponse) {
	result := s.repo.AddAccessUser(request)

	response <- user_access.AddUserAccessResponse{
		Code:   result.Code,
		Status: result.Status,
		Error:  result.Error,
	}
}

func (s *UserAccessService) CheckAccessUser(request user_access.CheckUserAccessRequest, response chan user_access.CheckUserAccessResponse) {
	result := s.repo.CheckAccessUser(request)

	response <- user_access.CheckUserAccessResponse{
		Code:   result.Code,
		Status: result.Status,
		Error:  result.Error,
	}
}

func (s *UserAccessService) DeleteAccessUser(request user_access.DeleteUserAccessRequest, response chan user_access.DeleteUserAccessResponse) {
	result := s.repo.DeleteAccessUser(request)

	response <- user_access.DeleteUserAccessResponse{
		Code:   result.Code,
		Status: result.Status,
		Error:  result.Error,
	}
}

func (s *UserAccessService) UpdateAccessUser(request user_access.UpdateUserAccessRequest, response chan user_access.UpdateUserAccessResponse) {
	result := s.repo.UpdateAccessUser(request)

	response <- user_access.UpdateUserAccessResponse{
		Code:   result.Code,
		Status: result.Status,
		Error:  result.Error,
	}
}
