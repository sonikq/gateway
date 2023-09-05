package user

import "gitlab.geogracom.com/skdf/skdf-abac-go/internal/models"

type GetNSIRequest struct {
	UserID  int
	NSICode string
}

type GetNSIResponse struct {
	Code     int
	Status   string      `json:"status"`
	Error    *models.Err `json:"error"`
	Response *[]NSI
}

type NSI struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
