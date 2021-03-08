package main

import (
	"fmt"

	"yasinzhang.top/hellogo/interface/example2/mock"
	"yasinzhang.top/hellogo/interface/example2/real"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

const url = "https://www.baidu.com"

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url,
		map[string]string{
			"name":   "yasin",
			"course": "golang",
		})
}

type RetriverPoster interface {
	Retriever
	Poster
}

func session(s RetriverPoster) string {
	s.Post(url, map[string]string{
		"contents": "post",
	})
	return s.Get(url)
}

func main() {
	var r Retriever
	r = &mock.Retriver{Contents: "mock retrive"}
	inspect(r)

	// Type assertion
	if mockRetriver, ok := r.(*mock.Retriver); ok {
		fmt.Println(mockRetriver.Contents)
	} else {
		fmt.Println("not mock retriver")
	}

	fmt.Println("session")
	s := &mock.Retriver{Contents: "mock retrive"}
	fmt.Println(session(s))
}

func inspect(r Retriever) {
	fmt.Println(r)
	fmt.Printf(" > %T %v\r\n", r, r)
	fmt.Print(" > Type switch: ")
	switch v := r.(type) {
	case *mock.Retriver:
		fmt.Println("Content: ", v.Contents)
	case *real.Retriver:
		fmt.Println("UserAgent: ", v.UserAgent)
	}
}
