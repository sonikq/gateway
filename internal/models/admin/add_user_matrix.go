package admin

import "gitlab.geogracom.com/skdf/skdf-abac-go/internal/models"

type AddUserMatrixByAdminRequest struct {
	UserID       int      `json:"user_id"`
	Value        int      `json:"object_value"`
	Type         string   `json:"object_type"`
	PeriodStart  string   `json:"period_start"`
	PeriodFinish string   `json:"period_finish"`
	States       []string `json:"states"`
	Role         string   `json:"role"`
	OrgsID       []int    `json:"orgs_id"`
}

type AddUserMatrixByAdminResponse struct {
	Code   int
	Status string      `json:"status"`
	Error  *models.Err `json:"error"`
}
