package user_access

import "gitlab.geogracom.com/skdf/skdf-abac-go/internal/models"

type CheckUserAccessRequest struct {
	UserID    int    `json:"user_id"`
	OKUD      int    `json:"okud"`
	ReportID  int    `json:"report_id"`
	ChapterID int    `json:"chapter_id"`
	AllowedOP string `json:"allowed_op"`
}

type CheckUserAccessResponse struct {
	Code   int
	Status string      `json:"status"`
	Error  *models.Err `json:"error"`
}
