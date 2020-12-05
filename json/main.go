package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string   `json:"title,omitempty"`
	Year   int      `json:"year,omitempty"`
	Color  bool     `json:"color,omitempty"`
	Actors []string `json:"actors,omitempty"`
}

func main() {
	actors := []string{"Shadow", "Nick", "Loyo", "Lawler"}
	var movie = &Movie{
		Title:  "My Shadow",
		Year:   2019,
		Color:  true,
		Actors: actors,
	}

	// json编排
	data, err := json.Marshal(movie)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	// 格式化整理过的json，便于浏览
	data2, err := json.MarshalIndent(movie, "", "----")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data2)

	// 将JSON字符串解码为GO格式
	var m Movie
	if err := json.Unmarshal(data, &m); err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Println(m)
}
