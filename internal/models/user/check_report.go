package user

import "gitlab.geogracom.com/skdf/skdf-abac-go/internal/models"

type CheckReportRequest struct {
	UserID   int
	ReportID int
}

type CheckReportResponse struct {
	AllowedOperations []int `json:"allowed_operations"`
	Code              int
	Status            string      `json:"status"`
	Error             *models.Err `json:"error"`
}
