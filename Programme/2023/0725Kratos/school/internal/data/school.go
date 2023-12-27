package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "school/api/school/v1"
	"school/internal/biz"
)

type taskRepo struct {
	data *Data
	log  *log.Helper
}

func NewTaskRepo(data *Data, logger log.Logger) biz.TaskRepo {
	return &taskRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

type RobotTaskManage struct {
	TaskManageId      int64       `gorm:"primaryKey;comment:任务ID;index:task_manage_id"`
	ServerId          int64       `gorm:"comment:服务器ID;index:server_id"`
	TaskType          pb.TaskType `gorm:"comment:任务类型(日常任务、成长任务、限时任务);index:task_type"`
	Repeat            string      `gorm:"comment:日常任务和成长任务才需要，任务刷新时间(每日、每周、每月等等);index:repeat"`
	StartTime         string      `gorm:"comment:限时任务才需要，任务开始时间;index:start_time"`
	EndTime           string      `gorm:"comment:限时任务才需要，任务结束时间;index:end_time"`
	Participants      string      `gorm:"comment:参与人员;index:participants"`
	TriggerEvent      string      `gorm:"comment:任务触发事件,以'_'分隔;index:trigger_event"`
	TotalCount        int32       `gorm:"comment:任务完成需要触发的事件总次数;index:total_count"`
	Name              string      `gorm:"comment:任务名称;index:name"`
	Description       string      `gorm:"comment:任务描述;index:description"`
	Score             int32       `gorm:"comment:完成任务获得的积分数;index:score"`
	TaskCondition     int32       `gorm:"comment:管理任务状态，启用或未启用;index:task_condition"`
	TaskConditionTime int64       `gorm:"comment:管理任务状态修改时间，启用或未启用状态;index:task_condition_time"`
	CreatedAt         int64       `gorm:"autoCreateTime"`
	UpdatedAt         int64       `gorm:"autoUpdateTime"`
}

func (robotTaskManage *RobotTaskManage) TableName() string {
	return "robot_task_manage"
}

func (t *taskRepo) taskManageToBiz(task *RobotTaskManage) *biz.TaskManage {
	return &biz.TaskManage{
		TaskManageId:      task.TaskManageId,
		TaskType:          task.TaskType,
		Repeat:            task.Repeat,
		StartTime:         task.StartTime,
		EndTime:           task.EndTime,
		Participants:      task.Participants,
		TriggerEvent:      task.TriggerEvent,
		TotalCount:        task.TotalCount,
		Name:              task.Name,
		Description:       task.Description,
		Score:             task.Score,
		TaskCondition:     task.TaskCondition,
		TaskConditionTime: task.TaskConditionTime,
	}
}

type RobotTaskProgress struct {
	TaskProgressId  int64           `gorm:"primaryKey;comment:任务进度ID;index:task_progress_id"`
	ServerId        int64           `gorm:"comment:服务器ID;index:server_id"`
	UserId          int64           `gorm:"comment:用户ID;index:user_id"`
	TaskManageId    int64           `gorm:"comment:任务ID;index:task_manage_id"`
	CompletedCount  int32           `gorm:"comment:任务进度,已经触发的任务事件总次数;index:completed_count"`
	Status          int32           `gorm:"comment:任务状态，完成或未完成;index:status"`
	RobotTaskManage RobotTaskManage `gorm:"foreignKey:TaskManageId"`
	CreatedAt       int64           `gorm:"autoCreateTime"`
	UpdatedAt       int64           `gorm:"autoUpdateTime"`
}

func (robotTaskProgress *RobotTaskProgress) TableName() string {
	return "robot_task_progress"
}

func (t *taskRepo) taskProgressToBiz(task *RobotTaskProgress) *biz.TaskProgress {
	return &biz.TaskProgress{
		TaskProgressId: task.TaskProgressId,
		ServerId:       task.ServerId,
		UserId:         task.UserId,
		TaskManageId:   task.TaskManageId,
		CompletedCount: task.CompletedCount,
		Status:         task.Status,
	}
}

type RobotTaskUser struct {
	TaskType       int32  `gorm:"comment:服务器类型;index:task_type"`
	Name           string `gorm:"comment:任务名称;index:name"`
	Description    string `gorm:"comment:任务描述;index:description"`
	Score          int32  `gorm:"comment:完成任务获得的积分数;index:score"`
	Status         int32  `gorm:"comment:任务状态，完成或未完成;index:status"`
	CompletedCount int32  `gorm:"comment:任务进度,已经触发的任务事件总次数;index:completed_count"`
	TotalCount     int32  `gorm:"comment:任务完成需要触发的事件总次数;index:total_count"`
}

// GetTaskUserList 用户获取任务列表
func (t *taskRepo) GetTaskUserList(ctx context.Context, userId, serverId int64) ([]*biz.TaskUser, error) {
	t.log.Infof("GetTaskUserList userId= %d,serverId= %d,", userId, serverId)
	var taskUserList []*RobotTaskProgress
	err := t.data.gormDB.WithContext(ctx).Model(&RobotTaskProgress{}).Preload("RobotTaskManage").Where("user_id = ? AND server_id = ?", userId, serverId).Find(&taskUserList).Error
	if err != nil {
		return nil, err
	}
	t.log.Infof("GetTaskUserList taskUserList=%v", taskUserList)
	var ret []*biz.TaskUser
	for _, task := range taskUserList {
		tempTask := &biz.TaskUser{
			TaskType:          task.RobotTaskManage.TaskType,
			Name:              task.RobotTaskManage.Name,
			Description:       task.RobotTaskManage.Description,
			Score:             task.RobotTaskManage.Score,
			Status:            task.Status,
			CompletedCount:    task.CompletedCount,
			TotalCount:        task.RobotTaskManage.TotalCount,
			TaskConditionTime: task.RobotTaskManage.TaskConditionTime,
		}
		ret = append(ret, tempTask)
	}
	return ret, nil
}

// GetTaskManageList 管理员获取任务列表
func (t *taskRepo) GetTaskManageList(ctx context.Context, serverId int64) ([]*biz.TaskManage, error) {
	var taskManageList []*RobotTaskManage
	err := t.data.gormDB.WithContext(ctx).Model(&RobotTaskManage{}).Where("server_id = ?", serverId).Find(&taskManageList).Error
	if err != nil {
		return nil, err
	}
	var ret []*biz.TaskManage
	for _, task := range taskManageList {
		ret = append(ret, t.taskManageToBiz(task))
	}
	return ret, nil
}

// CreateTask 管理员创建任务
func (t *taskRepo) CreateTask(ctx context.Context, task *biz.TaskManage) (taskId int64, err error) {
	tempTask := &RobotTaskManage{
		ServerId:          task.ServerId,
		TaskType:          task.TaskType,
		Repeat:            task.Repeat,
		StartTime:         task.StartTime,
		EndTime:           task.EndTime,
		Participants:      task.Participants,
		TriggerEvent:      task.TriggerEvent,
		TotalCount:        task.TotalCount,
		Name:              task.Name,
		Description:       task.Description,
		Score:             task.Score,
		TaskCondition:     task.TaskCondition,
		TaskConditionTime: task.TaskConditionTime,
	}

	err = t.data.gormDB.WithContext(ctx).Create(tempTask).Error
	if err != nil {
		return 0, err
	}
	return tempTask.TaskManageId, nil
}

// DeleteTask 管理员删除任务
func (t *taskRepo) DeleteTask(ctx context.Context, taskId, serverId int64) error {
	return t.data.gormDB.WithContext(ctx).Model(RobotTaskManage{}).Where("task_id = ? and server_id = ?", taskId, serverId).Delete(RobotTaskManage{}).Error
}

// ModifyTask 管理员编辑任务
func (t *taskRepo) ModifyTask(ctx context.Context, task *biz.TaskManage) error {
	return t.data.gormDB.WithContext(ctx).Model(RobotTaskManage{}).Where("task_id = ? and server_id = ?", task.TaskManageId, task.ServerId).Updates(&task).Error
}

// SetTaskStatus 管理员设置任务启用或取消启用
func (t *taskRepo) SetTaskStatus(ctx context.Context, taskId, serverId int64, status int32) error {
	return t.data.gormDB.WithContext(ctx).Model(RobotTaskManage{}).Where("task_id = ? and server_id = ?", taskId, serverId).Updates(map[string]interface{}{
		"status": status,
	}).Error
}

//// CreateTask 创建任务
//func (r *serverRepo) CreateTask(ctx context.Context, task *biz.ServerTask) (taskId int64, err error) {
//	tempTask := &model.ServerTask{
//		ServerId:          task.ServerId,
//		TaskType:          task.TaskType,
//		TaskRepeat:        task.TaskRepeat,
//		StartTime:         task.StartTime,
//		EndTime:           task.EndTime,
//		Participants:      task.Participants,
//		TriggerEvent:      task.TriggerEvent,
//		TotalCount:        task.TotalCount,
//		Name:              task.Name,
//		Description:       task.Description,
//		Score:             task.Score,
//		TaskCondition:     task.TaskCondition,
//		TaskConditionTime: time.Now().Unix(),
//	}
//	err = r.data.DB(ctx).WithContext(ctx).Create(tempTask).Error
//	if err != nil {
//		return 0, err
//	}
//	return tempTask.Id, nil
//}

//// DeleteTask 删除任务
//func (r *serverRepo) DeleteTask(ctx context.Context, serverTaskId int64) error {
//	return r.data.DB(ctx).WithContext(ctx).Model(model.ServerTask{}).Where("id = ?", serverTaskId).Delete(model.ServerTask{}).Error
//}
//
//// ModifyTask 修改任务
//func (r *serverRepo) ModifyTask(ctx context.Context, task *biz.ServerTask) error {
//	return r.data.DB(ctx).WithContext(ctx).Model(model.ServerTask{}).Where("id = ?", task.Id).Updates(&task).Error
//}
//
//// SetTaskStatus 设置任务启用或取消启用
//func (r *serverRepo) SetTaskStatus(ctx context.Context, serverTaskId int64, status int32) error {
//	return r.data.DB(ctx).WithContext(ctx).Model(model.ServerTask{}).Where("id = ?", serverTaskId).Updates(map[string]interface{}{
//		"task_condition":      status,
//		"task_condition_time": time.Now().Unix(),
//	}).Error
//}
// GetSingleServerTaskList 获取单个服务器所有任务
//func (r *serverRepo) GetSingleServerTaskList(ctx context.Context, serverId int64) ([]*biz.ServerTask, error) {
//	var serverTasks []*model.ServerTask
//	err := r.data.DB(ctx).WithContext(ctx).Model(&model.ServerTask{}).Where("server_id = ?", serverId).Find(&serverTasks).Error
//	if err != nil {
//		return nil, err
//	}
//	//没有任务，直接返回
//	if len(serverTasks) == 0 {
//		return nil, nil
//	}
//	serverTaskList := make([]*biz.ServerTask, 0, len(serverTasks))
//	for _, task := range serverTasks {
//		serverTaskList = append(serverTaskList, ServerTaskToBiz(task))
//	}
//	return serverTaskList, nil
//}
// GetAllServerTaskConditionList 获取所有服务器启用或者未启用的任务
//func (r *serverRepo) GetAllServerTaskConditionList(ctx context.Context, taskCondition int32) ([]*biz.ServerTask, error) {
//	var serverTasks []*model.ServerTask
//	err := r.data.DB(ctx).WithContext(ctx).Model(&model.ServerTask{}).Where("task_condition = ?", taskCondition).Find(&serverTasks).Error
//	if err != nil {
//		return nil, err
//	}
//	//没有启用的任务，直接返回
//	if len(serverTasks) == 0 {
//		return nil, nil
//	}
//	serverOnlineTaskList := make([]*biz.ServerTask, 0, len(serverTasks))
//	for _, task := range serverTasks {
//		serverOnlineTaskList = append(serverOnlineTaskList, ServerTaskToBiz(task))
//	}
//	return serverOnlineTaskList, nil
//}

//userId := jwt.GetUserIdFromContext(ctx)

//ext, _ := json.Marshal(map[string]interface{}{
//"entity": map[string]interface{}{
//"action":   "1",            //
//"content":  "[视频]",         //聊天内容
//"sender":   record.GroupId, //群聊ID
//"chatType": "2",            //群聊，1为私聊
//"faceUrl":  "1",            //发送者头像
//"nickname": "1",            //发送者昵称
//},
//})
//
//apnsInfo := map[string]interface{}{
//"Title": 0, //服务器名称
//}
//
//OfflinePushInf := map[string]interface{}{
//"PushFlag": 0,
//"Desc":     "****" + "：[视频]", //服务器的用户名称：[视频]
//"Ext":      ext,
//"ApnsInfo": apnsInfo,
//}
