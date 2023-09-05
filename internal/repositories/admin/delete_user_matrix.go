package admin

import (
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/models"
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/models/admin"
	"gitlab.geogracom.com/skdf/skdf-abac-go/pkg/db"
	"strings"
)

type DeleteUserMatrixRepo struct {
	db *db.DB
}

func NewDeleteAccessRepo(db *db.DB) *DeleteUserMatrixRepo {
	return &DeleteUserMatrixRepo{
		db: db,
	}
}

func (r *DeleteUserMatrixRepo) DeleteUserMatrixByAdmin(request admin.DeleteUserAccessRequest) admin.DeleteUserAccessResponse {
	if err := r.db.DeleteUserMatrix(request.UserID); err != nil {
		e := strings.Split(err.Error(), ": ")
		src, msg := e[0], e[1]
		return admin.DeleteUserAccessResponse{
			Code:   500,
			Status: deleteAccessUserFail,
			Error: &models.Err{
				Source:  src,
				Message: msg,
			},
		}
	}

	return admin.DeleteUserAccessResponse{
		Code:   200,
		Status: deleteAccessUserSuccess,
		Error:  nil,
	}
}
