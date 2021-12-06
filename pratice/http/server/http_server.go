package main


import (
	"fmt"
	"net/http"
	//"time"
)

func test(w http.ResponseWriter, r *http.Request){
	str := "hello world"
	data := r.URL.Query()
	fmt.Println("url",data)

	fmt.Println(data.Get("name"))
	fmt.Println(data.Get("age"))

	urlmethod := r.Method
	fmt.Println("method",urlmethod)

	w.Write([]byte(str))
}


func main() {
	//自定义 server
	//s := &http.Server{
	//	Addr: ":9999",
	//	Handler: test,
	//	ReadTimeout: 10 * time.Second,
	//	WriteTimeout: 10 * time.Second,
	//	MaxHeaderBytes: 1 << 20,
	//
	//
	//}
	//s.ListenAndServe()

	http.HandleFunc("/test",test)
	http.ListenAndServe("127.0.0.1:9999",nil)
}
