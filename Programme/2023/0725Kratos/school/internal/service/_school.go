package service

//
//import (
//	"context"
//	"github.com/go-kratos/kratos/v2/log"
//	"school/internal/biz"
//
//	pb "school/api/school/v1"
//)
//
//type SchoolService struct {
//	pb.UnimplementedSchoolServer
//
//	school *biz.StudentUsecase
//	log    *log.Helper
//}
//
//func NewSchoolService(sch *biz.StudentUsecase, logger log.Logger) *SchoolService {
//	return &SchoolService{
//		school: sch,
//		log:    log.NewHelper(logger),
//	}
//}
//
//func (s *SchoolService) CreateSchool(ctx context.Context, req *pb.CreateSchoolRequest) (*pb.CreateSchoolReply, error) {
//	return &pb.CreateSchoolReply{}, nil
//}
//func (s *SchoolService) UpdateSchool(ctx context.Context, req *pb.UpdateSchoolRequest) (*pb.UpdateSchoolReply, error) {
//	return &pb.UpdateSchoolReply{}, nil
//}
//func (s *SchoolService) DeleteSchool(ctx context.Context, req *pb.DeleteSchoolRequest) (*pb.DeleteSchoolReply, error) {
//	return &pb.DeleteSchoolReply{}, nil
//}
//func (s *SchoolService) GetSchool(ctx context.Context, req *pb.GetSchoolRequest) (*pb.GetSchoolReply, error) {
//	return &pb.GetSchoolReply{}, nil
//}
//func (s *SchoolService) ListSchool(ctx context.Context, req *pb.ListSchoolRequest) (*pb.ListSchoolReply, error) {
//	return &pb.ListSchoolReply{}, nil
//}
//func (s *SchoolService) GetStudent(ctx context.Context, req *pb.GetStudentRequest) (*pb.GetStudentReply, error) {
//	sch, err := s.school.Get(ctx, req.Id)
//
//	if err != nil {
//		return nil, err
//	}
//	return &pb.GetStudentReply{
//		Id:     sch.ID,
//		Status: sch.Status,
//		Name:   sch.Name,
//	}, nil
//}
