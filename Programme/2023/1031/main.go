package main

import (
	"fmt"
	"strings"
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
}
