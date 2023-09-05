package admin

import "gitlab.geogracom.com/skdf/skdf-abac-go/internal/models"

type DeleteUserAccessRequest struct {
	UserID int `json:"user_id"`
}

type DeleteUserAccessResponse struct {
	Code   int
	Status string      `json:"status"`
	Error  *models.Err `json:"error"`
}
