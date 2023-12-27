package main

// Subject 主题（任务）接口(包含用户注册和取消注册)
type Subject interface {
	Register(Subscriber)
	CancelRegister(Subscriber)
	Notify(message string)
}

// ConcreteSubject 具体主题(任务) 维护一批订阅者
type ConcreteSubject struct {
	subscribers []Subscriber
}

// Register 订阅
func (cst *ConcreteSubject) Register(obr Subscriber) {
	cst.subscribers = append(cst.subscribers, obr)
}

// CancelRegister 解除订阅（完成任务之后解除订阅）
func (cst *ConcreteSubject) CancelRegister(obr Subscriber) {
	for index, observer := range cst.subscribers {
		if observer == obr {
			cst.subscribers = append(cst.subscribers[:index], cst.subscribers[index+1:]...)
			break
		}
	}
}

// Notify 通知 每天刷新日常任务
func (cst *ConcreteSubject) Notify(message string) {
	for _, observer := range cst.subscribers {
		observer.NewTask(message)
	}
}
