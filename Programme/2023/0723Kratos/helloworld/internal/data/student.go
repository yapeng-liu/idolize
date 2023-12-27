package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"helloworld/internal/biz"
)

type studentRepo struct {
	data *Data
	log  *log.Helper
}

func NewStudentRepo(data *Data, logger log.Logger) biz.StudentRepo {
	return &studentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (repo *studentRepo) Save(ctx context.Context, stu *biz.Student) (*biz.Student, error) {
	repo.log.WithContext(ctx).Infof("Save Data in database")
	return stu, nil
}

func (repo *studentRepo) Update(ctx context.Context, stu *biz.Student) (*biz.Student, error) {
	repo.log.WithContext(ctx).Infof("Update Data in database")
	return stu, nil
}
