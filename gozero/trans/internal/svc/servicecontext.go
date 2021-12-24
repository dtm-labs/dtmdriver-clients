package svc

import "github.com/dtm-labs/dtmdriver-clients/gozero/trans/internal/config"

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
