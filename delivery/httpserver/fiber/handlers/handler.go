package handlers

import "github.com/techerfan/2DCH7-20059/pkg/logger"

type Handler struct {
	logger logger.Logger
}

func NewHandler(logger logger.Logger) *Handler {
	return &Handler{logger: logger}
}
