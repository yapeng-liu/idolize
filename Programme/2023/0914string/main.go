package main

import "fmt"

func main() {
	str1 := "1.3.1"
	str2 := "1.3.0"
	str3 := "1.3.0.5"
	str4 := "1.4.0"
	str5 := ""
	str6 := "1.3.0.14"
	fmt.Println("1", str1 > str2)
	fmt.Println("1", str1 > str3)
	fmt.Println("1", str1 > str4)
	fmt.Println("1", str1 > str5)
	fmt.Println("1", str1 > str6)

	fmt.Println(len(str6))
}
