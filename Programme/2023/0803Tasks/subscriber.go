package main

import "fmt"

// Subscriber 订阅者（用户）接口
type Subscriber interface {
	NewTask(string)
	GetUserId() int64
}

// ConcreteSubscriber 具体订阅者 玩家
type ConcreteSubscriber struct {
	userId int64
}

func (csr *ConcreteSubscriber) NewTask(taskName string) {
	fmt.Printf("%d 用户接受任务，任务名为：%s\n", csr.userId, taskName)
}
func (csr *ConcreteSubscriber) GetUserId() int64 {
	return csr.userId
}

func (csr *ConcreteSubscriber) DoTask() {

}
