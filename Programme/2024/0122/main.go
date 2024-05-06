package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func getURLPrefix(urlStr string) string {
	dir := filepath.Dir(urlStr)
	return dir
}

type Temp struct {
	OMG string `json:"omg"`
}

type Icon struct {
	temp *Temp
	AFI  string `json:"afi"`
}

func main() {
	// 示例用法
	urlStr := "https://daoke-1305543001.cos.ap-shanghai.myqcloud.com/multipleSource/avatar_frame/gongxihuode.png?T=12456"
	prefix := filepath.Base(urlStr)
	fmt.Println("URL 前缀：", prefix)

	newTime := GetTheDay(time.Now(), 2)

	fmt.Println(newTime.Format("2006-01-02 15:04:05"))

	fmt.Println(GetTodayMidnight().Format("2006-01-02 15:04:05"))

	var (
		expire  = time.Time{}
		expireB = time.Time{}
	)

	sysExpiration := GetTheDay(time.Now(), 1)

	if expire == expireB {
		fmt.Println("1")
	}
	if expire == sysExpiration {
		fmt.Println("2", expire)
	}

	//icon := &Icon{
	//	AFI: "",
	//}
	//updateIcon, err := json.Marshal(icon)
	//if err != nil {
	//	fmt.Println("routine.Go func() Marshal err", err)
	//	return
	//}
	//fmt.Println("str ", string(updateIcon), "uni ", updateIcon)
	icon := &Icon{}
	fmt.Println(icon.AFI, icon.temp)

	//var a = []int64{1, 2, 3, 4, 8}
	//fmt.Println(a[0:4])
	ff := "3,25,35,48"
	//ff2 := "25,35,48"
	//bytes := []byte(ff)
	//tempInfo := string(bytes[2:])
	//if tempInfo != ff2 {
	//	fmt.Println("ff0", tempInfo)
	//} else {
	//	fmt.Println("ff1", tempInfo)
	//}
	extras := strings.Split(ff, ",")
	firstNum, _ := strconv.Atoi(extras[0])
	firstNum = firstNum%(len(extras)-1) + 1
	extras[0] = strconv.Itoa(firstNum)
	result := strings.Join(extras, ",")
	fmt.Println("result ", result)
}

func GetTheDay(theDay time.Time, d int) time.Time {
	deadline := theDay.AddDate(0, 0, d+1)
	return time.Date(deadline.Year(), deadline.Month(), deadline.Day(), 0, 0, 0, 0, deadline.Location())
}

func GetTodayMidnight() time.Time {
	now := time.Now()
	// 将时分秒和纳秒部分清零
	midnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	return midnight
}
