package repositories

import (
	"github.com/casbin/casbin/v2"
	admin2 "gitlab.geogracom.com/skdf/skdf-abac-go/internal/models/admin"
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/models/user"
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/models/user_access"
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/repositories/admin"
	"gitlab.geogracom.com/skdf/skdf-abac-go/pkg/db"
)

type IUserRepo interface {
	GetForms(request user.GetFormRequest) user.GetFormResponse
	CheckFormEmpty(request user.CheckFormEmptyRequest) user.CheckFormEmptyResponse
	SetFormNotEmpty(request user.SetFormNotEmptyRequest) user.SetFormNotEmptyResponse
	GetChapterData(request user.GetChapterDataRequest) user.GetChapterDataResponse
	GetFormChapters(request user.GetFormChaptersRequest) user.GetFormChaptersResponse
	GetFormData(request user.GetFormDataRequest) user.GetFormDataResponse
	GetNSI(request user.GetNSIRequest) user.GetNSIResponse
	GetTargetCells(request user.GetTargetCellsRequest) user.GetTargetCellsResponse
	SaveChapterData(request user.SaveChapterDataRequest) user.SaveChapterDataResponse
	GetUIPermissions(request user.GetUIPermissionsRequest) user.GetUIPermissionsResponse
}

type IUserAccessRepo interface {
	AddAccessUser(request user_access.AddUserAccessRequest) user_access.AddUserAccessResponse
	CheckAccessUser(request user_access.CheckUserAccessRequest) user_access.CheckUserAccessResponse
	DeleteAccessUser(request user_access.DeleteUserAccessRequest) user_access.DeleteUserAccessResponse
	UpdateAccessUser(request user_access.UpdateUserAccessRequest) user_access.UpdateUserAccessResponse
}

type IAddUserMatrixByAdminRepo interface {
	AddUserMatrixByAdmin(request admin2.AddUserMatrixByAdminRequest) admin2.AddUserMatrixByAdminResponse
}

type IDeleteUserMatrixByAdminRepo interface {
	DeleteUserMatrixByAdmin(request admin2.DeleteUserAccessRequest) admin2.DeleteUserAccessResponse
}

type IUpdateUserMatrixByAdminRepo interface {
	UpdateUserMatrixByAdmin(request admin2.UpdateUserAccessRequest) admin2.UpdateUserAccessResponse
}

type Repository struct {
	IUserRepo
	IUserAccessRepo
	IAddUserMatrixByAdminRepo
	IDeleteUserMatrixByAdminRepo
	IUpdateUserMatrixByAdminRepo
}

func NewRepository(db *db.DB, enforcer *casbin.Enforcer) *Repository {
	return &Repository{
		IUserRepo:                    NewUserRepo(db, enforcer),
		IUserAccessRepo:              NewUserAccessRepo(db),
		IAddUserMatrixByAdminRepo:    admin.NewAddAdminMatrixRepo(db),
		IDeleteUserMatrixByAdminRepo: admin.NewDeleteAccessRepo(db),
		IUpdateUserMatrixByAdminRepo: admin.NewUpdateUserMatrixRepo(db),
	}
}
