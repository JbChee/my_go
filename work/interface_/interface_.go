package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	Name string
}

//interface{} to struct{}
func test(value interface{}) {

	switch v := value.(type) {
	case string:
		fmt.Println(v)
	case *User:
		if resp, ok := value.(User); ok {
			fmt.Printf("Name = %s", resp.Name)
		}
	default:
		fmt.Println("default")
	}
}

//Go 默认会将数值当做 float64 处理
func unmarsh_interface() {
	var data = []byte(`{"status": 200}`)
	var result map[string]interface{}

	if err := json.Unmarshal(data, &result); err != nil {
		log.Fatalln(err)
	}

	var status = int64(result["status"].(float64))
	fmt.Println("Status value: ", status)
}

func main() {

	any := User{
		Name: "chen",
	}
	test(any)

	unmarsh_interface()

}
