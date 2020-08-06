package main

import (
	"fmt"
	"log"
)

func main() {
	GetCourses()
}

func GetCourses() {
	sid := "2018214830"
	psw := ""

	list, err := GetSelfCoursesFromXK(sid, psw, "", "")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(list.Items)
}
