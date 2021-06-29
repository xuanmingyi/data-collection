package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type File struct {
	Id     int64
	Path   string
	Status string
}

type FileRepo interface {
	CreateFile(ctx context.Context, f *File) (*File, error)
	UpdateFile(ctx context.Context, f *File) (*File, error)
	GetFile(ctx context.Context, id int64) (*File, error)
	ListFile(ctx context.Context, pageNum, pageSize int64) ([]*File, error)
}

type FileUseCase struct {
	repo FileRepo
	log  *log.Logger
}

func NewFileUseCase(repo FileRepo, logger log.Logger) *FileUseCase {
	return &FileUseCase{repo: repo}
}

func (uc *FileUseCase) Create(ctx context.Context, f *File) (*File, error) {
	return uc.repo.CreateFile(ctx, f)
}

func (uc *FileUseCase) UpdateFile(ctx context.Context, f *File) (*File, error) {
	return uc.repo.UpdateFile(ctx, f)
}

func (uc *FileUseCase) GetFile(ctx context.Context, id int64) (*File, error) {
	return uc.repo.GetFile(ctx, id)
}

func (uc *FileUseCase) ListFile(ctx context.Context, pageNum, pageSize int64) ([]*File, error) {
	return uc.repo.ListFile(ctx, pageNum, pageSize)
}
