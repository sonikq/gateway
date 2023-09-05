package user

import "gitlab.geogracom.com/skdf/skdf-abac-go/internal/models"

type GetFormRequest struct {
	UserID int
	OKUD   int
}

type GetFormResponse struct {
	Code     int
	Status   string      `json:"status"`
	Error    *models.Err `json:"error"`
	Response Forms       `json:"response"`
}

type Forms []struct {
	ID            int         `json:"id"`
	Period        string      `json:"period"`
	IsAnnual      bool        `json:"is_annual"`
	Name          string      `json:"name"`
	Okud          string      `json:"okud"`
	RegionGid     int         `json:"region_gid"`
	Region        string      `json:"region"`
	Caption       string      `json:"caption"`
	Year          int         `json:"year"`
	Quarter       interface{} `json:"quarter"`
	StatusCode    string      `json:"status_code"`
	Status        string      `json:"status"`
	ResponderID   int         `json:"responder_id"`
	Responder     string      `json:"responder"`
	StatusChanged string      `json:"status_changed"`
	DueDate       string      `json:"due_date"`
	Chapters      []struct {
		ID    int    `json:"id"`
		Code  string `json:"code"`
		Label string `json:"label"`
		Title string `json:"title"`
	} `json:"chapters"`
}
