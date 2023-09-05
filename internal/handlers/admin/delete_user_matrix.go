package admin

import (
	"context"
	"github.com/gin-gonic/gin"
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/models/admin"
	"gitlab.geogracom.com/skdf/skdf-abac-go/pkg/logger"
	"net/http"
	"time"
)

func (h *Handler) DeleteUserAccess(ctx *gin.Context) {

	var request admin.DeleteUserAccessRequest

	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request data"})
		h.log.Error("Invalid request data", logger.Error(err))
		return
	}

	response := make(chan admin.DeleteUserAccessResponse, 1)

	c, cancel := context.WithTimeout(ctx, time.Second*time.Duration(h.config.CtxTimeout))
	defer cancel()

	go h.service.IDeleteUserMatrixByAdminService.DeleteUserMatrixByAdmin(request, response)

	select {
	case <-c.Done():
		ctx.JSON(http.StatusRequestTimeout, gin.H{
			statusKey: timeLimitExceedErr,
		})
		h.log.Error("", logger.String(errMsgKey, timeLimitExceedErr))
	case result := <-response:
		switch result.Code {
		case 200:
			ctx.JSON(http.StatusOK, gin.H{
				statusKey: result.Status,
			})
		default:
			ctx.JSON(result.Code, gin.H{
				statusKey: result.Status,
				errMsgKey: result.Error.Message,
			})
			h.log.Error(
				result.Status,
				logger.String(errSourceKey, result.Error.Source),
				logger.String(errMsgKey, result.Error.Message),
			)
		}
	}

}
