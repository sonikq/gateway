package handlers

import (
	"github.com/gin-gonic/gin"
	cfg "gitlab.geogracom.com/skdf/skdf-abac-go/configs/app"

	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/handlers/user"

	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/handlers/admin"

	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/handlers/user_access"
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/services"
	"gitlab.geogracom.com/skdf/skdf-abac-go/pkg/logger"
	"net/http"
)

type Handlers struct {
	Handler      *user_access.Handler
	UserHandler  *user.Handler
	AdminHandler *admin.Handler
}

type Option struct {
	Conf    cfg.Config
	Logger  logger.Logger
	Service *services.Service
}

func NewRouter(option Option) *gin.Engine {
	gin.SetMode(gin.TestMode)

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.MaxMultipartMemory = 8 << 20

	h := &Handlers{

		Handler: user_access.New(&user_access.HandlerConfig{
			Conf:    option.Conf,
			Logger:  option.Logger,
			Service: option.Service,
		}),

		UserHandler: user.New(&user.HandlerConfig{
			Conf:    option.Conf,
			Logger:  option.Logger,
			Service: option.Service,
		}),
		AdminHandler: admin.New(&admin.HandlerConfig{
			Conf:    option.Conf,
			Logger:  option.Logger,
			Service: option.Service,
		}),
	}

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Pong!",
		})
	})

	router.GET("/get_forms", h.UserHandler.GetForms)
	router.GET("/check_form_empty", h.UserHandler.CheckFormEmpty)
	router.GET("/get_chapter_data", h.UserHandler.GetChapterData)
	router.GET("/get_form_data", h.UserHandler.GetFormData)
	router.GET("/get_ui_permissions", h.UserHandler.GetUIPermissions)
	router.GET("/get_nsi", h.UserHandler.GetNSI)
	router.GET("/get_form_chapters", h.UserHandler.GetFormChapters)
	router.GET("/set_form_not_empty", h.UserHandler.SetFormNotEmpty)
	router.PATCH("/save_chapter_data", h.UserHandler.SaveChapterData)
	router.GET("/get_chapter_target_cells", h.UserHandler.GetChapterTargetCells)

	_user := router.Group("/user")
	{
		_user.POST("/export-report", h.UserHandler.ExportReport)
	}
	adminGroup := router.Group("/admin")
	{
		adminGroup.POST("create", h.AdminHandler.AddUserMatrixByAdmin)
		adminGroup.POST("update", h.AdminHandler.UpdateUserAccess)
		adminGroup.POST("delete", h.AdminHandler.DeleteUserAccess)

	}

	userAccess := router.Group("/user-access")
	{
		userAccess.POST("/create", h.Handler.AddUserAccess)
		userAccess.POST("/check", h.Handler.CheckUserAccess)
		userAccess.POST("/delete", h.Handler.DeleteUserAccess)
		userAccess.POST("/update", h.Handler.UpdateUserAccess)
	}

	return router
}
