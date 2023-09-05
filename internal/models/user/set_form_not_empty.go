package user

import "gitlab.geogracom.com/skdf/skdf-abac-go/internal/models"

type SetFormNotEmptyRequest struct {
	UserID int
	FormID int
}

type SetFormNotEmptyResponse struct {
	Code     int
	Status   string      `json:"status"`
	Error    *models.Err `json:"error"`
	Response *string     `json:"message"`
}

//type FormNotEmpty struct {
//	Count int `json:"count"`
//}
