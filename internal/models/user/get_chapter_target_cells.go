package user

import "gitlab.geogracom.com/skdf/skdf-abac-go/internal/models"

type GetTargetCellsRequest struct {
	ChapterID int
}

type GetTargetCellsResponse struct {
	Code     int
	Status   string      `json:"status"`
	Error    *models.Err `json:"error"`
	Response *[]TargetCells
}

type TargetCells struct {
	TargetCell string `json:"target_cell"`
}
