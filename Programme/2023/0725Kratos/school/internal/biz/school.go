package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "school/api/school/v1"
)

type TaskManage struct {
	TaskManageId      int64
	ServerId          int64
	TaskType          pb.TaskType
	Repeat            string
	StartTime         string
	EndTime           string
	Participants      string
	TriggerEvent      string
	TotalCount        int32
	Name              string
	Description       string
	Score             int32
	TaskCondition     int32
	TaskConditionTime int64
}

type TaskProgress struct {
	TaskProgressId int64
	TaskManageId   int64
	ServerId       int64
	UserId         int64
	CompletedCount int32
	Status         int32
}

type TaskUser struct {
	TaskType          pb.TaskType
	Name              string
	Description       string
	Score             int32
	Status            int32
	ConditionTime     int32
	CompletedCount    int32
	TotalCount        int32
	TaskConditionTime int64
}

type TaskRepo interface {
	GetTaskUserList(ctx context.Context, userId, serverId int64) ([]*TaskUser, error)
	GetTaskManageList(ctx context.Context, serverId int64) ([]*TaskManage, error)
	CreateTask(ctx context.Context, task *TaskManage) (taskId int64, err error)
	DeleteTask(ctx context.Context, taskId int64, serverId int64) error
	ModifyTask(ctx context.Context, task *TaskManage) error
	SetTaskStatus(ctx context.Context, taskId, serverId int64, status int32) error

	//使用
	//CreateTask(ctx context.Context, task *ServerTask) (taskId int64, err error)
	//DeleteTask(ctx context.Context, serverTaskId int64) error
	//ModifyTask(ctx context.Context, task *ServerTask) error
	//SetTaskStatus(ctx context.Context, serverTaskId int64, status int32) error
	//GetAllServerTaskConditionList(ctx context.Context, taskCondition int32) ([]*ServerTask, error)
	//GetSingleServerTaskList(ctx context.Context, serverId int64) ([]*ServerTask, error)
}

type TaskUsecase struct {
	repo TaskRepo
	log  *log.Helper
}

func NewTaskUsecase(repo TaskRepo, logger log.Logger) *TaskUsecase {
	return &TaskUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *TaskUsecase) GetTaskUserList(ctx context.Context, userId, serverId int64) ([]*TaskUser, error) {
	uc.log.WithContext(ctx).Infof("GetTaskUserList userId= %d ,serverId= %d", userId, serverId)
	return uc.repo.GetTaskUserList(ctx, userId, serverId)
}
func (uc *TaskUsecase) GetTaskManageList(ctx context.Context, serverId int64) ([]*TaskManage, error) {
	uc.log.WithContext(ctx).Infof("GetTaskManageList ,serverId= %d", serverId)
	return uc.repo.GetTaskManageList(ctx, serverId)

}
func (uc *TaskUsecase) CreateTask(ctx context.Context, task *TaskManage) (taskId int64, err error) {
	uc.log.WithContext(ctx).Infof("CreateTask task= %v", task)
	return uc.repo.CreateTask(ctx, task)
}
func (uc *TaskUsecase) DeleteTask(ctx context.Context, taskId int64, serverId int64) error {
	uc.log.WithContext(ctx).Infof("DeleteTask taskId= %d ,serverId= %d", taskId, serverId)
	return uc.repo.DeleteTask(ctx, taskId, serverId)
}
func (uc *TaskUsecase) ModifyTask(ctx context.Context, task *TaskManage) error {
	uc.log.WithContext(ctx).Infof("ModifyTask task= %v", task)
	return uc.repo.ModifyTask(ctx, task)
}
func (uc *TaskUsecase) SetTaskStatus(ctx context.Context, taskId, serverId int64, status int32) error {
	uc.log.WithContext(ctx).Infof("SetStatusTask taskId= %d ,serverId= %d, status= %d", taskId, serverId, status)
	return uc.repo.SetTaskStatus(ctx, taskId, serverId, status)
}

//使用
//func (uc *ServerUsecase) GetServerTaskList(ctx context.Context, serverId int64) ([]*ServerTask, error) {
//	uc.log.WithContext(ctx).Infof("GetServerTaskList ,serverId= %d", serverId)
//	return uc.repo.GetSingleServerTaskList(ctx, serverId)
//
//}
//func (uc *ServerUsecase) CreateTask(ctx context.Context, task *ServerTask) (taskId int64, err error) {
//	uc.log.WithContext(ctx).Infof("CreateTask task= %v", task)
//	return uc.repo.CreateTask(ctx, task)
//}
//func (uc *ServerUsecase) DeleteTask(ctx context.Context, serverTaskId int64) error {
//	uc.log.WithContext(ctx).Infof("DeleteTask serverTaskId= %d", serverTaskId)
//	return uc.repo.DeleteTask(ctx, serverTaskId)
//}
//func (uc *ServerUsecase) ModifyTask(ctx context.Context, task *ServerTask) error {
//	uc.log.WithContext(ctx).Infof("ModifyTask task= %v", task)
//	return uc.repo.ModifyTask(ctx, task)
//}
//func (uc *ServerUsecase) SetTaskStatus(ctx context.Context, serverTaskId int64, status int32) error {
//	uc.log.WithContext(ctx).Infof("SetStatusTask serverTaskId= %d ,status= %d", serverTaskId, status)
//	if status != 0 && status != 1 {
//		return v1.ErrorUnKnown(gconst.ParameterError)
//	}
//	return uc.repo.SetTaskStatus(ctx, serverTaskId, status)
//}
