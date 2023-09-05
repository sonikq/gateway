package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/pkg/reader"
	"gitlab.geogracom.com/skdf/skdf-abac-go/pkg/logger"
)

const (
	output      = "form"
	ext         = ".xlsx"
	contentType = "application/octet-stream"
)

type excel struct {
	Error *string `json:"error,omitempty"`
}

func (h *Handler) ExportReport(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Query("user_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ErrMsgKey: "неверный тип user_id"})
		h.log.Error("Invalid request data", logger.Error(err))
		return
	}

	formID, err := strconv.Atoi(ctx.Query("form_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ErrMsgKey: "неверный тип form_id"})
		h.log.Error("Invalid request data", logger.Error(err))
		return
	}

	getFormDataURL := fmt.Sprintf("http://%s:%s/get_form_data?user_id=%d&form_id=%d",
		h.config.HTTP.Host, h.config.HTTP.Port, userID, formID)

	getFormDataReq, err := http.NewRequestWithContext(ctx, http.MethodGet, getFormDataURL, nil)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Oops, something went wrong!"})
		h.log.Error("Error in making new post request: ", logger.Error(err))
		return
	}

	getFormDataResp, err := http.DefaultClient.Do(getFormDataReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Oops, something went wrong!"})
		h.log.Error("Error in sending post request: ", logger.Error(err))
		return
	}

	getFormDataBody, err := reader.ReadBody(getFormDataResp.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Oops, something went wrong!"})
		h.log.Error("Error in reading body: ", logger.Error(err))
		return
	}

	if bytes.Contains(getFormDataBody, []byte("fail")) {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Oops, something went wrong!"})
		h.log.Error("Error in get form data: ", logger.Error(errors.New("fail")))
		return
	}

	excelReq, err := http.NewRequestWithContext(ctx, http.MethodPost, h.config.Excel.URL, bytes.NewReader(getFormDataBody))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Oops, something went wrong!"})
		h.log.Error("Error in making new post request: ", logger.Error(err))
		return
	}

	excelReq.Header.Set("Accept", "application/octet-stream")

	excelResp, err := http.DefaultClient.Do(excelReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Oops, something went wrong!"})
		h.log.Error("Error in sending post request: ", logger.Error(err))
		return
	}

	excelBody, err := reader.ReadBody(excelResp.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Oops, something went wrong!"})
		h.log.Error("Error in reading body: ", logger.Error(err))
		return
	}

	if json.Valid(excelBody) {
		var x excel
		_ = json.Unmarshal(excelBody, &x)
		if x.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Oops, something went wrong!"})
			h.log.Info("excel address", logger.String("host", h.config.Excel.URL))
			h.log.Error("Error in parsing excel body: ", logger.Error(errors.New(*x.Error)))
			return
		}
	} else {
		filename := output + ext
		ctx.Header("Content-Disposition", "attachment; filename="+filename)
		ctx.Data(http.StatusOK, contentType, excelBody)
	}
}
