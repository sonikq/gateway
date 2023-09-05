package user

import "gitlab.geogracom.com/skdf/skdf-abac-go/internal/models"

type GetFormChaptersRequest struct {
	UserID int
	FormID int
}

type GetFormChaptersResponse struct {
	Code     int
	Status   string      `json:"status"`
	Error    *models.Err `json:"error"`
	Response *[]FormChapter
}

type FormChapter struct {
	ID    int    `json:"id"`
	Code  string `json:"code"`
	Title string `json:"title"`
	Label string `json:"label"`
}
