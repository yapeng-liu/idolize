package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Student struct {
	ID      string
	Name    string
	SayName string
}

type StudentRepo interface {
	Save(context.Context, *Student) (*Student, error)
	Update(context.Context, *Student) (*Student, error)
}

type StudentUsercase struct {
	repo    StudentRepo
	greeter GreeterRepo
	log     *log.Helper
}

func NewStudentUsercase(repo StudentRepo, greeter GreeterRepo, logger log.Logger) *StudentUsercase {
	return &StudentUsercase{repo: repo, greeter: greeter, log: log.NewHelper(logger)}
}

func (uc *StudentUsercase) CreateStudent(ctx context.Context, stu *Student) (*Student, error) {
	uc.log.WithContext(ctx).Infof("CreateStudent: %v", stu.ID)
	return uc.repo.Save(ctx, stu)
}
