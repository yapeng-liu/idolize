package service

//// Observer 观察者（服务器）接口（包含主题发布、取消发布和任务完成通知）
//type Observer interface {
//	Publish(Observer)
//	CancelPublish(Observer)
//	Told(taskProgressId int64) bool
//}
//
//// Subscriber 订阅者（用户）接口
//type Subscriber interface {
//	NewTask(string)
//}
//
//// ConcreteObserver 具体订阅者 玩家
//type ConcreteObserver struct {
//	userId int64
//	tasks  []Task
//}
//
//func (o *ConcreteObserver) NewTask(taskName string) {
//	fmt.Printf("%d 用户完成了一次任务，任务名为：%s\n", o.userId, taskName)
//}
//
//// Subject 主题（任务）接口(包含用户注册和取消注册)
//type Subject interface {
//	Register(Subscriber)
//	CancelRegister(Subscriber)
//	Notify()
//}
//
//// ConcreteSubject 具体主题(任务：日常任务、成长任务、限时任务) 维护一批订阅者
//type ConcreteSubject struct {
//	subscribers [][]Subscriber
//}
//
//// Register 订阅
//func (s *ConcreteSubject) Register(taskProgressId int64, obr Subscriber) {
//	s.subscribers[taskProgressId] = append(s.subscribers[taskProgressId], obr)
//}
//
//// CancelRegister 解除订阅（完成任务之后解除订阅）
//func (s *ConcreteSubject) CancelRegister(taskProgressId int64, obr Subscriber) {
//	for index, observer := range s.subscribers[taskProgressId] {
//		if observer == obr {
//			s.subscribers = append(s.subscribers[:index], s.subscribers[index+1:]...)
//			break
//		}
//	}
//}
//
//// Notify 通知 每天刷新日常任务
//func (s *ConcreteSubject) Notify(taskType int32, message string) {
//	for _, observer := range s.subscribers[taskType] {
//		observer.NewTask(message)
//	}
//}
//
//// Task 任务接口
//type Task interface {
//	TaskInit()
//	TaskUnInit()
//	Execute()
//}
//
//// ConcreteTask 具体任务
//type ConcreteTask struct {
//	TaskProgressId int64      //任务进度唯一Id
//	subject        Subject    //对应一个主题(任务)
//	TriggerEvert   []int32    //多个任务触发事件
//	TaskStatus     int32      //任务用户完成状态
//	subscriber     Subscriber //订阅者
//}
//
//type TaskManageList struct {
//	Tasks []ConcreteTask
//}
//
//// TaskInit 将用户注册到主题中
//func (t *ConcreteTask) TaskInit() {
//	t.subject.Register(t.subscriber)
//}
//
//// TaskUnInit 在主题中注销用户
//func (t *ConcreteTask) TaskUnInit() {
//	t.subject.CancelRegister(t.subscriber)
//}
//
//func (t *ConcreteTask) Execute() {
//	// 执行任务逻辑
//	// ...
//
//	// 通知观察者任务完成
//	t.subject.Notify(t.TaskProgressId)
//}
