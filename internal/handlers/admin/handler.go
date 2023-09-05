package admin

import (
	cfg "gitlab.geogracom.com/skdf/skdf-abac-go/configs/app"
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/services"
	"gitlab.geogracom.com/skdf/skdf-abac-go/pkg/logger"
)

type HandlerConfig struct {
	Conf    cfg.Config
	Logger  logger.Logger
	Service *services.Service
}

type Handler struct {
	config  cfg.Config
	log     logger.Logger
	service *services.Service
}

func New(cfg *HandlerConfig) *Handler {
	return &Handler{
		config:  cfg.Conf,
		log:     cfg.Logger,
		service: cfg.Service,
	}
}
