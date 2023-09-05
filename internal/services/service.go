package services

import (
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/models/admin"
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/models/user"
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/models/user_access"
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/repositories"
	admin2 "gitlab.geogracom.com/skdf/skdf-abac-go/internal/services/admin"
)

type IUserService interface {
	GetForms(request user.GetFormRequest, response chan user.GetFormResponse)
	CheckFormEmpty(request user.CheckFormEmptyRequest, response chan user.CheckFormEmptyResponse)
	SetFormNotEmpty(request user.SetFormNotEmptyRequest, response chan user.SetFormNotEmptyResponse)
	GetChapterData(request user.GetChapterDataRequest, response chan user.GetChapterDataResponse)
	GetNSI(request user.GetNSIRequest, response chan user.GetNSIResponse)
	GetChapterTargetCells(request user.GetTargetCellsRequest, response chan user.GetTargetCellsResponse)
	GetFormData(request user.GetFormDataRequest, response chan user.GetFormDataResponse)
	GetFormChapters(request user.GetFormChaptersRequest, response chan user.GetFormChaptersResponse)
	SaveChapterData(request user.SaveChapterDataRequest, response chan user.SaveChapterDataResponse)
	GetUIPermissions(request user.GetUIPermissionsRequest, response chan user.GetUIPermissionsResponse)
}

type IUserAccessService interface {
	AddAccessUser(request user_access.AddUserAccessRequest, response chan user_access.AddUserAccessResponse)
	CheckAccessUser(request user_access.CheckUserAccessRequest, response chan user_access.CheckUserAccessResponse)
	DeleteAccessUser(request user_access.DeleteUserAccessRequest, response chan user_access.DeleteUserAccessResponse)
	UpdateAccessUser(request user_access.UpdateUserAccessRequest, response chan user_access.UpdateUserAccessResponse)
}

type IAddUserMatrixByAdminService interface {
	AddUserMatrixByAdmin(request admin.AddUserMatrixByAdminRequest, response chan admin.AddUserMatrixByAdminResponse)
}

type IDeleteUserMatrixByAdminService interface {
	DeleteUserMatrixByAdmin(request admin.DeleteUserAccessRequest, response chan admin.DeleteUserAccessResponse)
}

type IUpdateUserMatrixByAdminService interface {
	UpdateUserMatrixByAdminUser(request admin.UpdateUserAccessRequest, response chan admin.UpdateUserAccessResponse)
}

type Service struct {
	IUserService
	IUserAccessService
	IAddUserMatrixByAdminService
	IDeleteUserMatrixByAdminService
	IUpdateUserMatrixByAdminService
}

func NewService(repos *repositories.Repository) *Service {
	return &Service{
		IUserService:                    NewUserService(repos.IUserRepo),
		IUserAccessService:              NewUserAccessService(repos.IUserAccessRepo),
		IAddUserMatrixByAdminService:    admin2.NewAddUserMatrixByAdminService(repos.IAddUserMatrixByAdminRepo),
		IDeleteUserMatrixByAdminService: admin2.NewDeleteAccessService(repos.IDeleteUserMatrixByAdminRepo),
		IUpdateUserMatrixByAdminService: admin2.NewUpdateUserMatrixService(repos.IUpdateUserMatrixByAdminRepo),
	}
}
