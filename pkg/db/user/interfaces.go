package user_access

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/models/user"
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/pkg/parser"
	"strings"
)

type UserDB struct {
	*sqlx.DB
}

type IUserDB interface {
	GetForms(userID int, okud int) (user.Forms, error)
	CheckFormEmpty(userID int, formID int) (user.FormEmpty, error)
	SetFormNotEmpty(formID int) error
	GetFormData(userID int, FormID int) (user.FormData, error)
	GetFormIDByChapterID(chapterID int) (user.GetFormIDByChapterID, error)
	GetChapterData(userID int, chapterID int) (user.ChapterData, error)
	GetNSI(nsiCode string) ([]user.NSI, error)
	GetFormChapters(formID int) ([]user.FormChapter, error)
	GetActions() ([]string, error)
	SaveChapterData(chapterID int, data map[string]interface{}) error
	GetChapterFormula(chapterID int) ([]user.Formula, error)
	GetChapterTargetCells(chapterID int) ([]user.TargetCells, error)
	GetCasbinSubject(userID, formID int) (user.Subject, error)
	GetCasbinObject(formID int) (user.Object, error)
	GetTemplateJson(formID int) ([]user.DataByFormID, error)
	CloseDB() error
}

func NewUserDB(db *sqlx.DB) *UserDB {
	return &UserDB{
		DB: db,
	}
}

func (db *UserDB) GetForms(userID int, okud int) (user.Forms, error) {
	TempForms = JSONForms{}
	if err := db.QueryRow(getForms, userID, okud).Scan(&TempForms); err != nil {
		return nil, err
	}
	return user.Forms(TempForms), nil
}

func (db *UserDB) CheckFormEmpty(userID int, formID int) (user.FormEmpty, error) {
	TempFormEmpty = JSONFormEmpty{}
	err := db.QueryRow(checkFormEmpty, userID, formID).Scan(&TempFormEmpty)
	if err != nil {
		if strings.Contains(err.Error(), "no data") {
			return nil, errors.New("not authorized")
		}
		return nil, err
	}
	return user.FormEmpty(TempFormEmpty), nil
}

func (db *UserDB) SetFormNotEmpty(formID int) error {
	var count int
	err := db.QueryRow(setFormNotEmpty, formID).Scan(&count)
	if err != nil {
		if strings.Contains(err.Error(), "no data") {
			return errors.New("not authorized")
		}
		return err
	}
	return nil
}

func (db *UserDB) GetNSI(nsiCode string) ([]user.NSI, error) {
	rows, err := db.Query(getNSI, nsiCode)
	if err != nil {
		return nil, err
	}

	var nsiArr []user.NSI

	for rows.Next() {
		var (
			id   int
			name string
		)

		TempNSI = JSONNSI{}

		if err = rows.Scan(&id, &name); err != nil {
			return nil, err
		}
		if err != nil {
			return nil, err
		}

		nsiArr = append(nsiArr, user.NSI{
			ID:   id,
			Name: name,
		})
	}

	return nsiArr, nil
}

func (db *UserDB) GetFormChapters(formID int) ([]user.FormChapter, error) {
	rows, err := db.Query(getFormChapters, formID)
	if err != nil {
		return nil, err
	}

	var fc []user.FormChapter

	for rows.Next() {
		var (
			id    int
			code  string
			title string
			label string
		)

		TempFormChapters = JSONFormChapters{}

		if err = rows.Scan(&id, &code, &title, &label); err != nil {
			return nil, err
		}
		if err != nil {
			return nil, err
		}

		fc = append(fc, user.FormChapter{
			ID:    id,
			Code:  code,
			Title: title,
			Label: label,
		})
	}

	return fc, nil
}

func (db *UserDB) GetFormData(userID int, FormID int) (user.FormData, error) {
	TempFormData = JSONFormData{}
	err := db.QueryRow(getFormData, userID, FormID).Scan(&TempFormData)
	if err != nil {
		if strings.Contains(err.Error(), "no data") {
			return user.FormData{}, errors.New("not authorized")
		}
		return user.FormData{}, err
	}
	return user.FormData(TempFormData), nil
}

func (db *UserDB) GetFormIDByChapterID(chapterID int) (user.GetFormIDByChapterID, error) {
	var formID int
	err := db.QueryRow(getFormIDByChapterID, chapterID).Scan(&formID)
	if err != nil {
		if strings.Contains(err.Error(), "no data") {
			return user.GetFormIDByChapterID{}, errors.New("form id not found")
		}
		return user.GetFormIDByChapterID{}, err
	}
	return user.GetFormIDByChapterID{FormID: formID}, err
}

func (db *UserDB) GetChapterData(userID int, chapterID int) (user.ChapterData, error) {
	TempChapterData = JSONChapterData{}
	err := db.QueryRow(getChapterData, userID, chapterID).Scan(&TempChapterData)
	if err != nil {
		if strings.Contains(err.Error(), "no data") {
			return user.ChapterData{}, errors.New("not authorized")
		}
		return user.ChapterData{}, err
	}
	return user.ChapterData(TempChapterData), nil
}

func (db *UserDB) SaveChapterData(chapterID int, data map[string]interface{}) error {
	var count int
	chapterData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = db.QueryRow(saveChapterData, chapterID, chapterData).Scan(&count)
	if err != nil {
		if strings.Contains(err.Error(), "no data") {
			return errors.New("not authorized")
		}
		return err
	}
	return nil
}

func (db *UserDB) GetChapterFormula(chapterID int) ([]user.Formula, error) {
	rows, err := db.Query(getChapterFormula, chapterID)
	if err != nil {
		return nil, err
	}

	var fs []user.Formula

	for rows.Next() {
		var targetCell, source string
		if err = rows.Scan(&targetCell, &source); err != nil {
			return nil, err
		}
		f, err := parser.ParseFormula(targetCell, source)
		if err != nil {
			return nil, err
		}

		fs = append(fs, user.Formula{
			TargetCells: f.Target,
			SourceCells: f.Source,
			Layout:      f.Layout,
		})
	}

	return fs, nil
}

func (db *UserDB) GetChapterTargetCells(chapterID int) ([]user.TargetCells, error) {
	rows, err := db.Query(getChapterTargetCells, chapterID)
	if err != nil {
		return nil, err
	}

	var tc []user.TargetCells

	for rows.Next() {
		var targetCell sql.NullString
		if err = rows.Scan(&targetCell); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, errors.New("there are not formulas for this chapter in database")
			}
			return nil, err
		}

		tc = append(tc, user.TargetCells{
			TargetCell: targetCell.String,
		})
	}

	if tc == nil {
		return nil, errors.New("there are not formulas for this chapter in database")
	}
	return tc, nil
}

func (db *UserDB) GetTemplateJson(formID int) ([]user.DataByFormID, error) {
	rows, err := db.Query(getTemplateJSON, formID)
	if err != nil {
		return nil, err
	}

	var fs []user.DataByFormID

	for rows.Next() {
		var templateID int

		TemplateJSON = JSONTemplateID{}

		if err = rows.Scan(&templateID, &TemplateJSON); err != nil {
			return nil, err
		}
		if err != nil {
			return nil, err
		}

		fs = append(fs, user.DataByFormID{
			TemplateID: templateID,
			Data:       TemplateJSON,
		})
	}

	return fs, nil
}

func (db *UserDB) GetActions() ([]string, error) {
	rows, err := db.Query(getActions)
	if err != nil {
		return nil, err
	}

	var actions []string
	for rows.Next() {
		var action string
		if err := rows.Scan(&action); err != nil {
			if err == sql.ErrNoRows {
				return nil, errors.New("there are not actions in database")
			}
			return nil, err
		}
		actions = append(actions, action)
	}

	return actions, nil
}

func (db *UserDB) GetCasbinSubject(userID, formID int) (user.Subject, error) {
	var (
		value   sql.NullInt64
		_type   sql.NullString
		pStart  sql.NullString
		pFinish sql.NullString
		States  []string
		Role    sql.NullString
		OrgIds  []int64
	)
	if err := db.QueryRow(getCasbinSubject, userID, formID).Scan(&value, &_type, &pStart,
		&pFinish, (*pq.StringArray)(&States), &Role, pq.Array(&OrgIds)); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user.Subject{}, errors.New("not authorized")
		}
		return user.Subject{}, err
	}

	sub := user.Subject{
		UserID:          userID,
		ObjectValue:     value.Int64,
		ObjectType:      _type.String,
		PeriodStart:     pStart.String,
		PeriodFinish:    pFinish.String,
		States:          States,
		Role:            Role.String,
		OrganizationsID: OrgIds,
	}

	return sub, nil
}
func (db *UserDB) GetCasbinObject(formID int) (user.Object, error) {
	var (
		okud             sql.NullInt64
		period_start     sql.NullString
		period_finish    sql.NullString
		doc_state        sql.NullString
		org_responder_id sql.NullInt64
	)
	if err := db.QueryRow(getCasbinObject, formID).Scan(&okud, &period_start, &period_finish,
		&doc_state, &org_responder_id); err != nil {
		if err == sql.ErrNoRows {
			return user.Object{}, errors.New("not authorized")
		}
		return user.Object{}, err
	}

	sub := user.Object{
		OKUD:           okud.Int64,
		PeriodStart:    period_start.String,
		PeriodFinish:   period_finish.String,
		DocState:       doc_state.String,
		OrgResponderID: org_responder_id.Int64,
	}

	return sub, nil
}

func (db *UserDB) CloseDB() error {
	return db.Close()
}
