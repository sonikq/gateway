package repositories

import (
	"github.com/casbin/casbin/v2"
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/models"
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/models/user"
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/pkg/recalculator"
	"gitlab.geogracom.com/skdf/skdf-abac-go/pkg/db"
)

type UserRepo struct {
	db       *db.DB
	enforcer *casbin.Enforcer
}

func NewUserRepo(db *db.DB, enforcer *casbin.Enforcer) *UserRepo {
	return &UserRepo{
		db:       db,
		enforcer: enforcer,
	}
}

func (r *UserRepo) GetForms(request user.GetFormRequest) user.GetFormResponse {
	res, err := r.db.IUserDB.GetForms(request.UserID, request.OKUD)
	if err != nil {
		return user.GetFormResponse{
			Response: nil,
			Code:     400,
			Status:   fail,
			Error: &models.Err{
				Source:  "pq",
				Message: err.Error(),
			},
		}
	}

	return user.GetFormResponse{
		Response: res,
		Code:     200,
		Status:   success,
		Error:    nil,
	}
}

func (r *UserRepo) CheckFormEmpty(request user.CheckFormEmptyRequest) user.CheckFormEmptyResponse {
	res, err := r.db.IUserDB.CheckFormEmpty(request.UserID, request.FormID)
	if err != nil {
		if err.Error() == "not authorized" {
			return user.CheckFormEmptyResponse{
				Response: nil,
				Code:     403,
				Status:   fail,
				Error: &models.Err{
					Source:  "pq",
					Message: err.Error(),
				},
			}
		}
		return user.CheckFormEmptyResponse{
			Response: nil,
			Code:     400,
			Status:   fail,
			Error: &models.Err{
				Source:  "pq",
				Message: err.Error(),
			},
		}
	}

	return user.CheckFormEmptyResponse{
		Response: res,
		Code:     200,
		Status:   success,
		Error:    nil,
	}
}

func (r *UserRepo) SetFormNotEmpty(request user.SetFormNotEmptyRequest) user.SetFormNotEmptyResponse {
	if err := r.db.IUserDB.SetFormNotEmpty(request.FormID); err != nil {
		if err.Error() == "not authorized" {
			return user.SetFormNotEmptyResponse{
				Response: nil,
				Code:     403,
				Status:   fail,
				Error: &models.Err{
					Source:  "pq",
					Message: err.Error(),
				},
			}
		}
		return user.SetFormNotEmptyResponse{
			Response: nil,
			Code:     400,
			Status:   fail,
			Error: &models.Err{
				Source:  "pq",
				Message: err.Error(),
			},
		}
	}

	msg := "set form not empty succeed"

	return user.SetFormNotEmptyResponse{
		Response: &msg,
		Code:     200,
		Status:   success,
		Error:    nil,
	}
}

func (r *UserRepo) GetChapterData(request user.GetChapterDataRequest) user.GetChapterDataResponse {
	res, err := r.db.IUserDB.GetChapterData(request.UserID, request.ChapterID)
	if err != nil {
		if err.Error() == "not authorized" {
			return user.GetChapterDataResponse{
				Response: nil,
				Code:     403,
				Status:   fail,
				Error: &models.Err{
					Source:  "pq",
					Message: err.Error(),
				},
			}
		}
		return user.GetChapterDataResponse{
			Response: nil,
			Code:     400,
			Status:   fail,
			Error: &models.Err{
				Source:  "pq",
				Message: err.Error(),
			},
		}
	}

	return user.GetChapterDataResponse{
		Response: &res,
		Code:     200,
		Status:   success,
		Error:    nil,
	}
}

func (r *UserRepo) GetFormChapters(request user.GetFormChaptersRequest) user.GetFormChaptersResponse {
	res, err := r.db.IUserDB.GetFormChapters(request.FormID)
	if err != nil {
		if err.Error() == "not authorized" {
			return user.GetFormChaptersResponse{
				Response: nil,
				Code:     403,
				Status:   fail,
				Error: &models.Err{
					Source:  "pq",
					Message: err.Error(),
				},
			}
		}
		return user.GetFormChaptersResponse{
			Response: nil,
			Code:     400,
			Status:   fail,
			Error: &models.Err{
				Source:  "pq",
				Message: err.Error(),
			},
		}
	}

	return user.GetFormChaptersResponse{
		Response: &res,
		Code:     200,
		Status:   success,
		Error:    nil,
	}
}

func (r *UserRepo) GetFormData(request user.GetFormDataRequest) user.GetFormDataResponse {

	res, err := r.db.IUserDB.GetFormData(request.UserID, request.FormID)

	if err != nil {
		if err.Error() == "not authorized" {
			return user.GetFormDataResponse{
				Response: nil,
				Code:     403,
				Status:   fail,
				Error: &models.Err{
					Source:  "pq",
					Message: err.Error(),
				},
			}
		}
		return user.GetFormDataResponse{
			Response: nil,
			Code:     400,
			Status:   fail,
			Error: &models.Err{
				Source:  "pq",
				Message: err.Error(),
			},
		}
	}

	return user.GetFormDataResponse{
		Response: &res,
		Code:     200,
		Status:   success,
		Error:    nil,
	}
}

func (r *UserRepo) GetNSI(request user.GetNSIRequest) user.GetNSIResponse {
	res, err := r.db.IUserDB.GetNSI(request.NSICode)
	if err != nil {
		if err.Error() == "not authorized" {
			return user.GetNSIResponse{
				Response: nil,
				Code:     403,
				Status:   fail,
				Error: &models.Err{
					Source:  "pq",
					Message: err.Error(),
				},
			}
		}
		return user.GetNSIResponse{
			Response: nil,
			Code:     400,
			Status:   fail,
			Error: &models.Err{
				Source:  "pq",
				Message: err.Error(),
			},
		}
	}

	return user.GetNSIResponse{
		Response: &res,
		Code:     200,
		Status:   success,
		Error:    nil,
	}
}

func (r *UserRepo) GetTargetCells(request user.GetTargetCellsRequest) user.GetTargetCellsResponse {
	targetCells, err := r.db.IUserDB.GetChapterTargetCells(request.ChapterID)
	if err != nil {
		var code int
		switch {
		case err.Error() == "there are not formulas for this chapter in database":
			code = 403
		default:
			code = 500
		}
		return user.GetTargetCellsResponse{
			Response: nil,
			Code:     code,
			Status:   fail,
			Error: &models.Err{
				Source:  "pq",
				Message: err.Error(),
			},
		}
	}

	return user.GetTargetCellsResponse{
		Response: &targetCells,
		Code:     200,
		Status:   success,
		Error:    nil,
	}
}

func (r *UserRepo) SaveChapterData(request user.SaveChapterDataRequest) user.SaveChapterDataResponse {
	formByChapter, err := r.db.IUserDB.GetFormIDByChapterID(request.ChapterID)
	if err != nil {
		return user.SaveChapterDataResponse{
			Response: nil,
			Code:     500,
			Status:   fail,
			Error: &models.Err{
				Source:  "pq",
				Message: err.Error(),
			},
		}
	}

	formulas, err := r.db.IUserDB.GetChapterFormula(request.ChapterID)
	if err != nil {
		return user.SaveChapterDataResponse{
			Response: nil,
			Code:     500,
			Status:   fail,
			Error: &models.Err{
				Source:  "pq",
				Message: err.Error(),
			},
		}
	}
	templateJSONS, err := r.db.IUserDB.GetTemplateJson(formByChapter.FormID)
	if err != nil || templateJSONS == nil {
		msg := "can't get json by templateID"
		return user.SaveChapterDataResponse{
			Response: &msg,
			Code:     500,
			Status:   fail,
			Error: &models.Err{
				Source:  "pq",
				Message: err.Error(),
			},
		}
	}
	if err = recalculator.ReCalculate(formulas, &request.JSON, templateJSONS); err != nil {
		return user.SaveChapterDataResponse{
			Response: nil,
			Code:     500,
			Status:   fail,
			Error: &models.Err{
				Source:  "internal",
				Message: err.Error(),
			},
		}
	}

	if err = r.db.IUserDB.SaveChapterData(request.ChapterID, request.JSON); err != nil {
		if err.Error() == "not authorized" {
			return user.SaveChapterDataResponse{
				Response: nil,
				Code:     403,
				Status:   fail,
				Error: &models.Err{
					Source:  "pq",
					Message: err.Error(),
				},
			}
		}
		return user.SaveChapterDataResponse{
			Response: nil,
			Code:     400,
			Status:   fail,
			Error: &models.Err{
				Source:  "pq",
				Message: err.Error(),
			},
		}
	}

	msg := "save succeed"

	return user.SaveChapterDataResponse{
		Response: &msg,
		Code:     200,
		Status:   success,
		Error:    nil,
	}
}

func (r *UserRepo) GetUIPermissions(request user.GetUIPermissionsRequest) user.GetUIPermissionsResponse {
	if err := r.enforcer.LoadPolicy(); err != nil {
		return user.GetUIPermissionsResponse{
			Response: nil,
			Code:     500,
			Status:   fail,
			Error: &models.Err{
				Source:  "casbin",
				Message: err.Error(),
			},
		}
	}

	actions, err := r.db.IUserDB.GetActions()
	if err != nil {
		var code int
		switch {
		case err.Error() == "there are not actions in database":
			code = 403
		default:
			code = 500
		}
		return user.GetUIPermissionsResponse{
			Response: nil,
			Code:     code,
			Status:   fail,
			Error: &models.Err{
				Source:  "pq",
				Message: err.Error(),
			},
		}
	}

	sub, err := r.db.IUserDB.GetCasbinSubject(request.UserID, request.FormID)
	if err != nil {
		var code int
		switch {
		case err.Error() == "not authorized":
			code = 403
		default:
			code = 500
		}
		return user.GetUIPermissionsResponse{
			Response: nil,
			Code:     code,
			Status:   fail,
			Error: &models.Err{
				Source:  "pq",
				Message: err.Error(),
			},
		}
	}

	obj, err := r.db.IUserDB.GetCasbinObject(request.FormID)
	if err != nil {
		var code int
		switch {
		case err.Error() == "not authorized":
			code = 403
		default:
			code = 500
		}
		return user.GetUIPermissionsResponse{
			Response: nil,
			Code:     code,
			Status:   fail,
			Error: &models.Err{
				Source:  "pq",
				Message: err.Error(),
			},
		}
	}

	allowedActions := make(map[string]bool)

	for _, action := range actions {
		ok, err := r.enforcer.Enforce(sub, obj, action)
		if err != nil {
			return user.GetUIPermissionsResponse{
				Response: nil,
				Code:     403,
				Status:   fail,
				Error: &models.Err{
					Source:  "casbin",
					Message: err.Error(),
				},
			}
		}

		allowedActions[action] = ok
	}

	return user.GetUIPermissionsResponse{
		Response: &allowedActions,
		Code:     200,
		Status:   success,
		Error:    nil,
	}
}
