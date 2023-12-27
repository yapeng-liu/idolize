package main

import "time"

const (
	taskSubjectNum = 100 * 12 //一个任务进度对应一个任务主题，总共12个任务，预估10人
)

func main() {
	//主题(任务)
	subject := &ConcreteSubject{
		subscribers: make([]Subscriber, 0, taskSubjectNum),
	}
	//订阅者
	subscriber := &ConcreteSubscriber{
		userId: 10001,
	}
	//观察者
	concreteObserver := &ConcreteObserver{}

	//具体任务
	task := &ConcreteTask{
		TaskManageId:   10,
		TaskProgressId: 1,
		subject:        subject,
		TriggerEvert:   make([]int32, 0),
		TaskStatus:     2,
		subscriber:     subscriber,
		observer:       concreteObserver,
	}
	//任务发布
	task.subject.Notify("任务1")

	time.Sleep(2 * time.Second)
}
