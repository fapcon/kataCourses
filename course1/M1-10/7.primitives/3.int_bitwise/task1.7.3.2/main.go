package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(getFilePermissions(000))
}

func getFilePermissions(flag int) string {
	var rs [3]string
	if flag == 0 {
		rs[0] = "-,-,-"
		rs[1] = "-,-,-"
		rs[2] = "-,-,-"
	} else {
		s := strconv.Itoa(flag)
		r := strings.Split(s, "")
		fmt.Println(r)
		for i := 0; i < len(r); i++ {
			switch r[i] {
			case "0":
				rs[i] = "-,-,-"
			case "1":
				rs[i] = "-,-,Execute"
			case "2":
				rs[i] = "-,Write,-"
			case "3":
				rs[i] = "-,Write,Execute"
			case "4":
				rs[i] = "Read,-,-"
			case "5":
				rs[i] = "Read,-,Execute"
			case "6":
				rs[i] = "Read,Write,-"
			case "7":
				rs[i] = "Read,Write,Execute"
			}
		}
	}

	res := "Owner:"+rs[0]+" Group:"+rs[1]+" Other:"+rs[2]
	return res
}


