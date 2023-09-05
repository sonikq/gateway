package admin

import (
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/models"
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/models/admin"
	"gitlab.geogracom.com/skdf/skdf-abac-go/pkg/db"
	"strings"
)

type UpdateUserMatrixRepo struct {
	db *db.DB
}

func NewUpdateUserMatrixRepo(db *db.DB) *UpdateUserMatrixRepo {
	return &UpdateUserMatrixRepo{
		db: db,
	}
}

func (r *UpdateUserMatrixRepo) UpdateUserMatrixByAdmin(request admin.UpdateUserAccessRequest) admin.UpdateUserAccessResponse {
	if err := r.db.UpdateUserMatrix(
		request.UserID, request.Value, request.Type, request.PeriodStart, request.Role, request.PeriodFinish, request.States, request.OrgsID,
	); err != nil {
		e := strings.Split(err.Error(), ": ")
		src, msg := e[0], e[1]
		return admin.UpdateUserAccessResponse{
			Code:   500,
			Status: updateAccessUserFail,
			Error: &models.Err{
				Source:  src,
				Message: msg,
			},
		}
	}

	return admin.UpdateUserAccessResponse{
		Code:   200,
		Status: updateAccessUserSuccess,
		Error:  nil,
	}
}
