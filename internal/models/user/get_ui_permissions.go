package user

import "gitlab.geogracom.com/skdf/skdf-abac-go/internal/models"

type GetUIPermissionsRequest struct {
	UserID int
	FormID int
}

type GetUIPermissionsResponse struct {
	Code     int
	Status   string      `json:"status"`
	Error    *models.Err `json:"error"`
	Response *map[string]bool
}
