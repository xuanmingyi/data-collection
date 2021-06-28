package biz

import "github.com/go-kratos/kratos/v2/log"

type ServiceRepo struct {
}

type ServiceUseCase struct {
	repo ServiceRepo
	log  *log.Helper
}

func NewServiceUseCase(repo ServiceRepo, logger log.Logger) *ServiceUseCase {
	return &ServiceUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usercase/service"))}
}
