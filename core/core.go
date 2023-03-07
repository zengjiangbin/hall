package core

import (
	"github.com/zengjiangbin/driver/structs/service"
)

func NewService() service.Service {
	return new(hallCore)
}

type hallCore struct{}

func (h *hallCore) Init() {}

func (h *hallCore) Start() {}

func (h *hallCore) Name() service.Name {
	return service.HalLServiceName
}

func (h *hallCore) Version() string {
	return version
}
