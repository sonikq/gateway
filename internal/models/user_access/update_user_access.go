package user_access

import "gitlab.geogracom.com/skdf/skdf-abac-go/internal/models"

type UpdateUserAccessRequest struct {
	UserID     int      `json:"user_id"`
	OKUDS      []int    `json:"okuds"`
	ReportIDS  []int    `json:"report_ids"`
	ChapterIDS []int    `json:"chapter_ids"`
	AllowedOPS []string `json:"allowed_ops"`
}

type UpdateUserAccessResponse struct {
	Code   int
	Status string      `json:"status"`
	Error  *models.Err `json:"error"`
}
