package user_access

import (
	"context"
	"github.com/gin-gonic/gin"
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/models/user_access"
	"gitlab.geogracom.com/skdf/skdf-abac-go/pkg/logger"
	"net/http"
	"time"
)

func (h *Handler) AddUserAccess(ctx *gin.Context) {

	var request user_access.AddUserAccessRequest

	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request data"})
		h.log.Error("Invalid request data", logger.Error(err))
		return
	}

	response := make(chan user_access.AddUserAccessResponse, 1)

	c, cancel := context.WithTimeout(ctx, time.Second*time.Duration(h.config.CtxTimeout))
	defer cancel()

	go h.service.IUserAccessService.AddAccessUser(request, response)
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
			ctx.JSON(http.StatusOK, gin.H{
				StatusKey: result.Status,
			})
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
