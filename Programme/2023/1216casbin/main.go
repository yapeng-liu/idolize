package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

func check(enforcer *casbin.Enforcer, sub, obj, act string) bool {
	ok, err := enforcer.Enforce(sub, obj, act)
	if err != nil {
		log.Fatal("err ", err)
		return false
	}
	if ok {
		fmt.Printf("%s CAN %s %s\n", sub, act, obj)
	} else {
		fmt.Printf("%s CANNOT %s %s\n", sub, act, obj)
	}
	return ok
}

func main() {
	enforcer, err := casbin.NewEnforcer("./model.conf", "./policy.csv")
	if err != nil {
		log.Fatal("err ", err)
	}
	go func() {
		time.Sleep(13 * time.Second)
		ok, err := enforcer.AddPolicy("dajun", "data1", "write")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(ok)
	}()
	for {
		time.Sleep(3 * time.Second)
		check(enforcer, "dajun", "data1", "read")
		time.Sleep(3 * time.Second)
		check(enforcer, "lizi", "data2", "write")
		time.Sleep(3 * time.Second)
		check(enforcer, "dajun", "data1", "write")
		time.Sleep(3 * time.Second)
		check(enforcer, "lizi", "data2", "read")
	}

}
