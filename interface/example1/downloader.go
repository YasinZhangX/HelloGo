package main

import (
	"fmt"

	"yasinzhang.top/hellogo/interface/example1/infra"
)

func getRetriver() retriver {
	return infra.Retriever{}
}

type retriver interface {
	Get(string) string
}

func main() {
	retriver := getRetriver()
	fmt.Println(retriver.Get("https://www.baidu.com"))
}
