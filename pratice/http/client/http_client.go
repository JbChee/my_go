package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
)

func test1(){
	apiurl := "http://127.0.0.1:9999/test"
	data := url.Values{}
	 data.Set("name","小王子")
	 data.Set("age","18")

	 u ,err := url.ParseRequestURI(apiurl)

	 if err !=nil{
	 	fmt.Println("failed,err",err)

	 }
	 u.RawQuery = data.Encode()  //url encode
	 fmt.Println(u.String())

	 resp ,err := http.Get(u.String())
	 defer resp.Body.Close()


}


func main() {

	resp,err := http.Get("http://www.baidu.com")
	if err !=nil{
		fmt.Println("get baidu failed ,err",err)

	}
	fmt.Println("type:",reflect.TypeOf(resp))
	//fmt.Println("resp:",resp)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("type:",reflect.TypeOf(body))
	//fmt.Println(body)
	//fmt.Println(string(body))

	test1()

}
