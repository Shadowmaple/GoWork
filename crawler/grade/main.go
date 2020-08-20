package main

import (
	"fmt"
	"log"
)

func main() {
	GetGrades()
}

func GetGrades() {
	sid := "2018214830"
	psw := ""
	data, ok, err := GetGradeFromXK(sid, psw, 10)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(ok, data)
}
