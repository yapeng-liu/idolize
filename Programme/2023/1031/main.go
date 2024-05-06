package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/metadata"
	"strings"
	"time"
)

type myTest struct {
	reason string
}

func (m *myTest) Error() string {
	return m.reason
}

func main() {
	str := "a|b|c|d"
	ids := strings.Split(str, "|")
	fmt.Println(len(ids))

	err := func() error {
		var cake bool = true
		var my myTest
		if !cake {
			fmt.Println("a")
			my.reason = "cake is false"
			fmt.Println("b")
		}
		return &my
	}()
	fmt.Println(len(ids))
	if err != nil {
		fmt.Println(err.(*myTest))
	}

	cloudDateMap := make(map[string]interface{})
	var test string
	err = json.Unmarshal([]byte(test), &cloudDateMap)
	if err != nil {
		fmt.Println(err)
	}
	now := time.Now()
	// 将时分秒和纳秒部分清零
	midnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	deadline := time.Now().AddDate(0, 0, 1)
	fmt.Println(midnight.Unix(), midnight, deadline)
	timeObj := time.Unix(deadline.Unix(), 0)
	timeObj = timeObj.AddDate(0, 0, 1)
	fmt.Println(timeObj)

	ctx := context.Background()

	ctx = metadata.AppendToClientContext(ctx, "x-md-global-ip", "192.168.1.1")

	if md, ok := metadata.FromClientContext(ctx); ok {
		ip := md.Get("x-md-global-ip")
		fmt.Println("OK ", ip)
	} else {
		fmt.Println(" Not OK ")
	}

	s := "/android/2e5e191c873413f2cbc61cfcf27875f18"
	info := strings.Split(s, "/")
	if len(info) > 2 {
		fmt.Println(info[1])
		fmt.Println(info[2])
	}

	//fmt.Println("ip: ", GetIP4Context(ctx))
}

//func GetIP4Context(ctx context.Context) string {
//	var ip string
//	if tr, ok := transport.FromServerContext(ctx); ok {
//		ip = tr.RequestHeader().Get("x-md-global-ip")
//		if ip == "" {
//			if md, ok := metadata.FromServerContext(ctx); ok {
//				ip = md.Get("x-md-global-ip")
//			}
//		}
//	}
//	return ip
//}
