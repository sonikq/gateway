package user_access

import (
	"encoding/json"
	"errors"
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/models/user"
)

type JSONForms user.Forms

func (a JSONForms) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &TempForms)
}

type JSONFormEmpty user.FormEmpty

func (a JSONFormEmpty) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if b == nil {
		return errors.New("no data")
	}
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &TempFormEmpty)
}

type JSONChapterData user.ChapterData

func (a *JSONChapterData) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if b == nil {
		return errors.New("no data")
	}
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &TempChapterData)
}

type JSONFormData user.FormData

func (a *JSONFormData) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if b == nil {
		return errors.New("no data")
	}
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &TempFormData)
}

type JSONTemplateID map[string]interface{}

func (a *JSONTemplateID) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if b == nil {
		return errors.New("no data")
	}
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &TemplateJSON)
}

type JSONNSI user.NSI

func (a *JSONNSI) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if b == nil {
		return errors.New("no data")
	}
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &TempNSI)
}

type JSONFormChapters user.FormChapter

func (a *JSONFormChapters) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if b == nil {
		return errors.New("no data")
	}
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &TempFormChapters)
}
