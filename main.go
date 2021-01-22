package main

import (
	"fmt"
	"strings"
)

//List ..
type List struct {
	Value string
}

func main() {
	scode := strings.ToTitle("a")
	Data := ""
	if len(scode) <= 3 && scode != "ZZZ" {
		if len(scode) == 1 {
			if string([]rune(scode)[0]) != "Z" {
				Data = string([]rune(scode)[0] + 1)
			} else {
				Data = "0"
			}
		} else if len(scode) == 2 {
			if string([]rune(scode)[1]) != "Z" {
				Data = string([]rune(scode)[0]) + string([]rune(scode)[1]+1)
			} else {
				Data = "1"
			}
		} else if len(scode) == 3 {
			if string([]rune(scode)[1]) == "Z" && string([]rune(scode)[2]) == "Z" {
				Data = string([]rune(scode)[0]+1) + string("A") + string("A")
			} else {
				if string([]rune(scode)[2]) != "Z" {
					Data = string([]rune(scode)[0]) + string([]rune(scode)[1]) + string([]rune(scode)[2]+1)
				} else {
					Data = "2"
				}
			}
		}
		if Data == "0" {
			Data = "AA"
		} else if Data == "1" {
			if string([]rune(scode)[0]) != "Z" {
				Data = string([]rune(scode)[0]+1) + string("A")
			}
		} else if Data == "2" {
			if string([]rune(scode)[1]) != "Z" {
				Data = string([]rune(scode)[0]) + string([]rune(scode)[1]+1) + string("A")
			}
		}

	}
	fmt.Println(Data)
}
