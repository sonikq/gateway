package user

import "gitlab.geogracom.com/skdf/skdf-abac-go/internal/models"

type CheckOKUDRequest struct {
	UserID int `json:"user_id"`
	OKUD   int `json:"OKUD"`
}

type CheckOKUDResponse struct {
	Code   int
	Status string      `json:"status"`
	Error  *models.Err `json:"error"`
}
