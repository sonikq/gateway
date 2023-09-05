package repositories

import (
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/models"
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/models/user_access"
	"gitlab.geogracom.com/skdf/skdf-abac-go/pkg/db"
	"strings"
)

type UserAccessRepo struct {
	db *db.DB
}

func NewUserAccessRepo(db *db.DB) *UserAccessRepo {
	return &UserAccessRepo{
		db: db,
	}
}

func (r *UserAccessRepo) AddAccessUser(request user_access.AddUserAccessRequest) user_access.AddUserAccessResponse {
	if err := r.db.IAccessUserDB.AddUserAccess(
		request.UserID, request.OKUDS, request.ReportIDS, request.ChapterIDS, request.AllowedOPS,
	); err != nil {
		e := strings.Split(err.Error(), ": ")
		src, msg := e[0], e[1]
		return user_access.AddUserAccessResponse{
			Code:   400,
			Status: addAccessUserFail,
			Error: &models.Err{
				Source:  src,
				Message: msg,
			},
		}
	}

	return user_access.AddUserAccessResponse{
		Code:   201,
		Status: addAccessUserSuccess,
		Error:  nil,
	}
}

func (r *UserAccessRepo) CheckAccessUser(request user_access.CheckUserAccessRequest) user_access.CheckUserAccessResponse {
	exists, err := r.db.IAccessUserDB.CheckUserAccess(
		request.UserID, request.OKUD, request.ReportID, request.ChapterID, request.AllowedOP)
	if err != nil {
		e := strings.Split(err.Error(), ": ")
		src, msg := e[0], e[1]
		return user_access.CheckUserAccessResponse{
			Code:   400,
			Status: checkAccessUserFail,
			Error: &models.Err{
				Source:  src,
				Message: msg,
			},
		}
	}

	if !exists {
		return user_access.CheckUserAccessResponse{
			Code:   404,
			Status: userNotFoundMsg,
			Error: &models.Err{
				Source:  "pq",
				Message: userNotFoundMsg,
			},
		}
	}

	return user_access.CheckUserAccessResponse{
		Code:   200,
		Status: checkAccessUserSuccess,
		Error:  nil,
	}
}

func (r *UserAccessRepo) DeleteAccessUser(request user_access.DeleteUserAccessRequest) user_access.DeleteUserAccessResponse {
	if err := r.db.IAccessUserDB.DeleteUserAccess(request.UserID); err != nil {
		e := strings.Split(err.Error(), ": ")
		src, msg := e[0], e[1]
		return user_access.DeleteUserAccessResponse{
			Code:   500,
			Status: deleteAccessUserFail,
			Error: &models.Err{
				Source:  src,
				Message: msg,
			},
		}
	}

	return user_access.DeleteUserAccessResponse{
		Code:   200,
		Status: deleteAccessUserSuccess,
		Error:  nil,
	}
}

func (r *UserAccessRepo) UpdateAccessUser(request user_access.UpdateUserAccessRequest) user_access.UpdateUserAccessResponse {
	if err := r.db.IAccessUserDB.UpdateUserAccess(
		request.UserID, request.OKUDS, request.ReportIDS, request.ChapterIDS, request.AllowedOPS,
	); err != nil {
		e := strings.Split(err.Error(), ": ")
		src, msg := e[0], e[1]
		return user_access.UpdateUserAccessResponse{
			Code:   500,
			Status: updateAccessUserFail,
			Error: &models.Err{
				Source:  src,
				Message: msg,
			},
		}
	}

	return user_access.UpdateUserAccessResponse{
		Code:   200,
		Status: updateAccessUserSuccess,
		Error:  nil,
	}
}
