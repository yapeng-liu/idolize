package data

//
//import (
//	"context"
//	"github.com/go-kratos/kratos/v2/log"
//	"school/internal/biz"
//)
//
//type studentRepo struct {
//	data *Data
//	log  *log.Helper
//}
//
//// NewStudentRepo 初始化 studentRepo
//func NewStudentRepo(data *Data, logger log.Logger) biz.StudentRepo {
//	return &studentRepo{
//		data: data,
//		log:  log.NewHelper(logger),
//	}
//}
//
//func (r *studentRepo) GetStudent(ctx context.Context, id int32) (*biz.Student, error) {
//	var stu biz.Student
//	r.data.gormDB.Where("id = ?", id).First(&stu) // 这里使用了 gorm
//	r.log.WithContext(ctx).Info("gormDB: GetStudent, id: ", id)
//	return &biz.Student{
//		ID:        stu.ID,
//		Name:      stu.Name,
//		Status:    stu.Status,
//		Info:      stu.Info,
//		UpdatedAt: stu.UpdatedAt,
//		CreatedAt: stu.CreatedAt,
//	}, nil
//}
