package user

import "gitlab.geogracom.com/skdf/skdf-abac-go/internal/models"

type CheckFormEmptyRequest struct {
	UserID int
	FormID int
}

type CheckFormEmptyResponse struct {
	Code     int
	Status   string      `json:"status"`
	Error    *models.Err `json:"error"`
	Response FormEmpty
}

type FormEmpty []struct {
	IsEmpty bool `json:"is_empty"`
}
