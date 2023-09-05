package user

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/models/user"
	"gitlab.geogracom.com/skdf/skdf-abac-go/pkg/logger"
	"net/http"
	"strconv"
	"time"
)

func (h *Handler) GetFormData(ctx *gin.Context) {
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

	request := user.GetFormDataRequest{
		UserID: userID,
		FormID: formID,
	}

	response := make(chan user.GetFormDataResponse, 1)

	c, cancel := context.WithTimeout(ctx, time.Second*time.Duration(h.config.CtxTimeout))
	defer cancel()

	go h.service.IUserService.GetFormData(request, response)
	defer func() {
		if r := recover(); r != nil {
			h.log.Warn("паника", logger.String("описание", "обнаружена паника"))

		}
	}()

	select {
	case <-c.Done():
		ctx.JSON(http.StatusRequestTimeout, gin.H{
			StatusKey: TimeLimitExceedErr,
		})
		h.log.Error("", logger.String(ErrMsgKey, TimeLimitExceedErr))
	case result := <-response:
		switch result.Code {
		case 200:
			jsonData, err := json.Marshal(result.Response)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{ErrMsgKey: err.Error()})
				h.log.Error(ErrMsgKey, logger.Error(err))
				return
			}
			ctx.Data(http.StatusOK, "application/json", jsonData)
		default:
			ctx.JSON(result.Code, gin.H{
				StatusKey: result.Status,
				ErrMsgKey: result.Error.Message,
			})
			h.log.Error(
				result.Status,
				logger.String(ErrSourceKey, result.Error.Source),
				logger.String(ErrMsgKey, result.Error.Message),
			)
		}
	}
}
