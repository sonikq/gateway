package user

import (
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/models"
)

type SaveChapterDataRequest struct {
	UserID    int
	ChapterID int
	JSON      map[string]interface{} `json:"json"`
}

type SaveChapterDataResponse struct {
	Code     int
	Status   string      `json:"status"`
	Error    *models.Err `json:"error"`
	Response *string     `json:"message"`
}

type Formula struct {
	TargetCells []interface{}
	SourceCells [][]interface{}
	Layout      string
}

type DataByFormID struct {
	TemplateID int                    `json:"_form_chapter_template_id"`
	Data       map[string]interface{} `json:"_json"`
}
