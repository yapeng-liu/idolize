package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "school/api/school/v1"
	"school/internal/biz"
)

const (
	EventGameLogin = 1 //登陆游戏事件
	EventSignIn    = 2 //签到事件
	EventComment   = 3 //评论事件
	EventLike      = 4 //点赞事件
	EventPost      = 5 //发帖事件
)

const (
	RepeatEveryDay   = "每天"
	RepeatEveryWeek  = "每周"
	RepeatEveryMonth = "每月"
)

type SchoolService struct {
	pb.UnimplementedSchoolServer

	task *biz.TaskUsecase
	log  *log.Helper
}

func NewSchoolService(task *biz.TaskUsecase, logger log.Logger) *SchoolService {
	return &SchoolService{
		task: task,
		log:  log.NewHelper(logger),
	}
}

func (s *SchoolService) TaskUserList(ctx context.Context, req *pb.TaskUserListReq) (*pb.TaskUserListReply, error) {
	data, err := s.task.GetTaskUserList(ctx, req.UserId, req.ServerId)
	if err != nil {
		return nil, err
	}
	ret := &pb.TaskUserListReply{
		Score:        0,
		TaskUserList: make([]*pb.TaskUser, 0, len(data)),
	}
	for _, value := range data {
		ret.TaskUserList = append(ret.TaskUserList, &pb.TaskUser{
			TaskType:          value.TaskType,
			Name:              value.Name,
			Description:       value.Description,
			Score:             value.Score,
			Status:            value.Status,
			CompletedCount:    value.CompletedCount,
			TotalCount:        value.TotalCount,
			TaskConditionTime: value.TaskConditionTime,
		})
	}
	return ret, nil
}
func (s *SchoolService) TaskManageList(ctx context.Context, req *pb.TaskManageListReq) (*pb.TaskManageListReply, error) {
	data, err := s.task.GetTaskManageList(ctx, req.ServerId)
	if err != nil {
		return nil, err
	}
	ret := &pb.TaskManageListReply{
		TaskManageList: make([]*pb.TaskManage, 0, len(data)),
	}
	for _, value := range data {
		ret.TaskManageList = append(ret.TaskManageList, &pb.TaskManage{
			TaskType:      value.TaskType,
			Name:          value.Name,
			Description:   value.Description,
			Score:         value.Score,
			TaskCondition: value.TaskCondition,
		})
	}
	return ret, nil
}
func (s *SchoolService) TaskCreate(ctx context.Context, req *pb.TaskCreateReq) (*pb.TaskCreateReply, error) {
	data, err := s.task.CreateTask(ctx, &biz.TaskManage{
		ServerId:          req.ServerId,
		TaskType:          req.TaskType,
		Repeat:            req.Repeat,
		StartTime:         req.StartTime,
		EndTime:           req.EndTime,
		Participants:      req.Participants,
		TriggerEvent:      req.TriggerEvent,
		TotalCount:        req.TotalCount,
		Name:              req.Name,
		Description:       req.Description,
		Score:             req.Score,
		TaskCondition:     req.TaskCondition,
		TaskConditionTime: 1,
	})
	if err != nil {
		return nil, err
	}
	//将任务注册到事件中心，更新任务
	//
	return &pb.TaskCreateReply{TaskManageId: data}, nil
}
func (s *SchoolService) TaskDelete(ctx context.Context, req *pb.TaskDeleteReq) (*pb.TaskDeleteReply, error) {
	err := s.task.DeleteTask(ctx, req.TaskManageId, req.ServerId)
	if err != nil {
		return nil, err
	}
	return &pb.TaskDeleteReply{}, nil
}
func (s *SchoolService) TaskModify(ctx context.Context, req *pb.TaskModifyReq) (*pb.TaskModifyReply, error) {
	err := s.task.ModifyTask(ctx, &biz.TaskManage{
		TaskManageId:  req.TaskManageId,
		TaskType:      req.TaskType,
		Repeat:        req.Repeat,
		StartTime:     req.StartTime,
		EndTime:       req.EndTime,
		Participants:  req.Participants,
		TriggerEvent:  req.TriggerEvent,
		TotalCount:    req.TotalCount,
		Name:          req.Name,
		Description:   req.Description,
		Score:         req.Score,
		TaskCondition: req.TaskCondition,
	})
	if err != nil {
		return nil, err
	}
	return &pb.TaskModifyReply{}, nil
}
func (s *SchoolService) TaskStatusSet(ctx context.Context, req *pb.TaskStatusSetReq) (*pb.TaskStatusSetReply, error) {
	err := s.task.SetTaskStatus(ctx, req.TaskManageId, req.ServerId, req.TaskStatus)
	if err != nil {
		return nil, err
	}
	return &pb.TaskStatusSetReply{}, nil
}

//使用
//func (s *ServerService) GetServerTaskList(ctx context.Context, req *v1.ServerTaskListReq) (*v1.ServerTaskListReply, error) {
//	data, err := s.uc.GetServerTaskList(ctx, req.ServerId)
//	if err != nil {
//		return nil, err
//	}
//	ret := &v1.ServerTaskListReply{
//		ServerTaskList: make([]*v1.ServerTask, 0, len(data)),
//	}
//	for _, value := range data {
//		ret.ServerTaskList = append(ret.ServerTaskList, &v1.ServerTask{
//			TaskType:      value.TaskType,
//			Name:          value.Name,
//			Description:   value.Description,
//			Score:         value.Score,
//			TaskCondition: value.TaskCondition,
//		})
//	}
//	return ret, nil
//}
//func (s *ServerService) ServerTaskCreate(ctx context.Context, req *v1.ServerTaskCreateReq) (*v1.ServerTaskCreateReply, error) {
//	id, err := s.uc.CreateTask(ctx, &biz.ServerTask{
//		ServerId:      req.ServerId,
//		TaskType:      req.TaskType,
//		TaskRepeat:    req.Repeat,
//		StartTime:     req.StartTime,
//		EndTime:       req.EndTime,
//		Participants:  req.Participants,
//		TriggerEvent:  req.TriggerEvent,
//		TotalCount:    req.TotalCount,
//		Name:          req.Name,
//		Description:   req.Description,
//		Score:         req.Score,
//		TaskCondition: req.TaskCondition,
//	})
//	if err != nil {
//		return nil, err
//	}
//	return &v1.ServerTaskCreateReply{ServerTaskId: id}, nil
//}
//func (s *ServerService) ServerTaskDelete(ctx context.Context, req *v1.ServerTaskDeleteReq) (*v1.EmptyReply, error) {
//	err := s.uc.DeleteTask(ctx, req.ServerTaskId)
//	if err != nil {
//		return nil, err
//	}
//	return &v1.EmptyReply{}, nil
//}
//func (s *ServerService) ServerTaskModify(ctx context.Context, req *v1.ServerTaskModifyReq) (*v1.EmptyReply, error) {
//	err := s.uc.ModifyTask(ctx, &biz.ServerTask{
//		Id:            req.Id,
//		TaskType:      req.TaskType,
//		TaskRepeat:    req.Repeat,
//		StartTime:     req.StartTime,
//		EndTime:       req.EndTime,
//		Participants:  req.Participants,
//		TriggerEvent:  req.TriggerEvent,
//		TotalCount:    req.TotalCount,
//		Name:          req.Name,
//		Description:   req.Description,
//		Score:         req.Score,
//		TaskCondition: req.TaskCondition,
//	})
//	if err != nil {
//		return nil, err
//	}
//	return &v1.EmptyReply{}, nil
//}
//func (s *ServerService) TaskStatusSet(ctx context.Context, req *v1.ServerTaskStatusSetReq) (*v1.EmptyReply, error) {
//	err := s.uc.SetTaskStatus(ctx, req.ServerTaskId, req.TaskCondition)
//	if err != nil {
//		return nil, err
//	}
//	return &v1.EmptyReply{}, nil
//}
