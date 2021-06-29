package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/xuanmingyi/data-collection/app/file/internal/biz"
)

type fileRepo struct {
	data *Data
	log  *log.Helper
}

func NewFileRepo(data *Data, logger log.Logger) biz.FileRepo {
	return &fileRepo{data: data, log: log.NewHelper(log.With(logger, "module", "data/file"))}
}

func (r *fileRepo) CreateFile(ctx context.Context, f *biz.File) (*biz.File, error) {
	return nil, nil
}

func (r *fileRepo) GetFile(ctx context.Context, id int64) (*biz.File, error) {
	return nil, nil
}

func (r *fileRepo) UpdateFile(ctx context.Context, f *biz.File) (*biz.File, error) {
	return nil, nil
}

func (r *fileRepo) ListFile(ctx context.Context, pageNum, pageSize int64) ([]*biz.File, error) {
	return nil, nil
}
