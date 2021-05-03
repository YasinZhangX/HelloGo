package main

import "fmt"

func main() {
	m1 := map[string]string{
		"name":    "yasin",
		"course":  "golang",
		"site":    "imooc",
		"quality": "ok",
	}

	m2 := make(map[string]int)

	var m3 map[string]int

	fmt.Println(m1, m2, m3)

	fmt.Println("Traversing map")
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	fmt.Println("Get values")
	courseName, status := m1["course"]
	fmt.Println(courseName, status)
	a, status := m1["a"]
	fmt.Println(a, status)
	if k, status := m1["a"]; status {
		fmt.Println(k)
	} else {
		fmt.Println("key dose not exist")
	}

	fmt.Println("Delete values")
	delete(m1, "name")
	name, ok := m1["name"]
	fmt.Println(name, ok)
}
