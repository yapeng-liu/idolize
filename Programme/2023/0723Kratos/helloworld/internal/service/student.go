package service

import (
	"context"
	"helloworld/internal/biz"

	pb "helloworld/api/helloworld/v1"
)

func (s *GreeterService) CreateStudent(ctx context.Context, req *pb.CreateStudentRequest) (*pb.CreateStudentReply, error) {
	return &pb.CreateStudentReply{}, nil
}
func (s *GreeterService) UpdateStudent(ctx context.Context, req *pb.UpdateStudentRequest) (*pb.UpdateStudentReply, error) {
	return &pb.UpdateStudentReply{}, nil
}
func (s *GreeterService) DeleteStudent(ctx context.Context, req *pb.DeleteStudentRequest) (*pb.DeleteStudentReply, error) {
	return &pb.DeleteStudentReply{}, nil
}
func (s *GreeterService) GetStudent(ctx context.Context, req *pb.GetStudentRequest) (*pb.GetStudentReply, error) {
	return &pb.GetStudentReply{}, nil
}
func (s *GreeterService) ListStudent(ctx context.Context, req *pb.ListStudentRequest) (*pb.ListStudentReply, error) {
	return &pb.ListStudentReply{}, nil
}
func (s *GreeterService) Hello(ctx context.Context, req *pb.HelloReq) (*pb.HelloResp, error) {
	g, err := s.stu.CreateStudent(ctx, &biz.Student{Name: req.Name, ID: req.Id, SayName: req.SayName})
	if err != nil {
		return nil, err
	}
	return &pb.HelloResp{Message: "Hello " + g.Name, Text: "rev " + g.ID + "rev" + g.SayName}, nil
}
