package admin

import (
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/models"
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/models/admin"
	"gitlab.geogracom.com/skdf/skdf-abac-go/pkg/db"
	"strings"
)

type AddAdminMatrixRepo struct {
	db *db.DB
}

func NewAddAdminMatrixRepo(db *db.DB) *AddAdminMatrixRepo {
	return &AddAdminMatrixRepo{
		db: db,
	}
}

func (r *AddAdminMatrixRepo) AddUserMatrixByAdmin(request admin.AddUserMatrixByAdminRequest) admin.AddUserMatrixByAdminResponse {
	if err := r.db.AddUserMatrix(
		request.UserID, request.Value, request.Type, request.PeriodStart, request.Role, request.PeriodFinish, request.States, request.OrgsID,
	); err != nil {
		e := strings.Split(err.Error(), ": ")
		src, msg := e[0], e[1]
		return admin.AddUserMatrixByAdminResponse{
			Code:   400,
			Status: addAccessUserFail,
			Error: &models.Err{
				Source:  src,
				Message: msg,
			},
		}
	}

	return admin.AddUserMatrixByAdminResponse{
		Code:   201,
		Status: addAccessUserSuccess,
		Error:  nil,
	}
}
