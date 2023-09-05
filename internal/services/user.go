package services

import (
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/models/user"
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/repositories"
)

type UserService struct {
	repo repositories.IUserRepo
}

func NewUserService(repo repositories.IUserRepo) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetForms(request user.GetFormRequest, response chan user.GetFormResponse) {
	result := s.repo.GetForms(request)

	response <- user.GetFormResponse{
		Response: result.Response,
		Code:     result.Code,
		Status:   result.Status,
		Error:    result.Error,
	}
}

func (s *UserService) CheckFormEmpty(request user.CheckFormEmptyRequest, response chan user.CheckFormEmptyResponse) {
	result := s.repo.CheckFormEmpty(request)

	response <- user.CheckFormEmptyResponse{
		Response: result.Response,
		Code:     result.Code,
		Status:   result.Status,
		Error:    result.Error,
	}
}

func (s *UserService) SetFormNotEmpty(request user.SetFormNotEmptyRequest, response chan user.SetFormNotEmptyResponse) {
	result := s.repo.SetFormNotEmpty(request)

	response <- user.SetFormNotEmptyResponse{
		Response: result.Response,
		Code:     result.Code,
		Status:   result.Status,
		Error:    result.Error,
	}
}

func (s *UserService) GetChapterData(request user.GetChapterDataRequest, response chan user.GetChapterDataResponse) {
	result := s.repo.GetChapterData(request)

	response <- user.GetChapterDataResponse{
		Response: result.Response,
		Code:     result.Code,
		Status:   result.Status,
		Error:    result.Error,
	}
}

func (s *UserService) GetNSI(request user.GetNSIRequest, response chan user.GetNSIResponse) {
	result := s.repo.GetNSI(request)

	response <- user.GetNSIResponse{
		Response: result.Response,
		Code:     result.Code,
		Status:   result.Status,
		Error:    result.Error,
	}
}

func (s *UserService) GetChapterTargetCells(request user.GetTargetCellsRequest, response chan user.GetTargetCellsResponse) {
	result := s.repo.GetTargetCells(request)

	response <- user.GetTargetCellsResponse{
		Response: result.Response,
		Code:     result.Code,
		Status:   result.Status,
		Error:    result.Error,
	}
}

func (s *UserService) GetFormData(request user.GetFormDataRequest, response chan user.GetFormDataResponse) {
	result := s.repo.GetFormData(request)

	response <- user.GetFormDataResponse{
		Response: result.Response,
		Code:     result.Code,
		Status:   result.Status,
		Error:    result.Error,
	}
}

func (s *UserService) GetFormChapters(request user.GetFormChaptersRequest, response chan user.GetFormChaptersResponse) {
	result := s.repo.GetFormChapters(request)

	response <- user.GetFormChaptersResponse{
		Response: result.Response,
		Code:     result.Code,
		Status:   result.Status,
		Error:    result.Error,
	}
}

func (s *UserService) SaveChapterData(request user.SaveChapterDataRequest, response chan user.SaveChapterDataResponse) {
	result := s.repo.SaveChapterData(request)

	response <- user.SaveChapterDataResponse{
		Response: result.Response,
		Code:     result.Code,
		Status:   result.Status,
		Error:    result.Error,
	}
}

func (s *UserService) GetUIPermissions(request user.GetUIPermissionsRequest, response chan user.GetUIPermissionsResponse) {
	result := s.repo.GetUIPermissions(request)

	response <- user.GetUIPermissionsResponse{
		Response: result.Response,
		Code:     result.Code,
		Status:   result.Status,
		Error:    result.Error,
	}
}
