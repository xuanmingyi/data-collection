package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/xuanmingyi/data-collection/app/file/internal/biz"
)

var ProviderSet = wire.NewSet(NewFileService)

type FileService struct {
	bc  *biz.FileUseCase
	log *log.Helper
}

func NewFileService(bc *biz.FileUseCase, logger log.Logger) *FileService {
	return &FileService{bc: bc}
}
