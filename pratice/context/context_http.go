package main


import (
    "context"
    "net/http"
    "time"
)

func ContextMiddle(next http.Handler) http.Handler{
    return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
        cookie,_ :=r.Cookie("mycheck")
        if cookie !=nil{
            ctx :=context.WithValue(r.Context(),"mycheck",cookie.Value)
            next.ServeHTTP(w,r.WithContext(ctx))
        }else{
            next.ServeHTTP(w,r)
        }
    })

}
//
//// ContextMiddle是http服务中间件，统一读取通行cookie并使用ctx传递
//func ContextMiddle(next http.Handler) http.Handler {
//   return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//       cookie, _ := r.Cookie("Check")
//       if cookie != nil {
//           ctx := context.WithValue(r.Context(), "Check", cookie.Value)
//           next.ServeHTTP(w, r.WithContext(ctx))
//       } else {
//           next.ServeHTTP(w, r)
//       }
//   })
//}



func CheckHandler(w http.ResponseWriter,r *http.Request){
    expitation := time.Now().Add(24* time.Hour)
    cookie := http.Cookie{Name:"mycheck",Value:"42",Expires:expitation}
    http.SetCookie(w,&cookie)

}


func indexHandler(w http.ResponseWriter, r *http.Request){
    if chk := r.Context().Value("mycheck"); chk == "42"{
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("<h2> let is go ! </h2>"))
    }else{}
    w.WriteHeader(http.StatusNotFound)
    w.Write([]byte("<h2> No pass !</h2>"))
}

func main() {
    mux := http.NewServeMux()

    mux.HandleFunc("/", indexHandler)

    // 人为设置通行cookie
    mux.HandleFunc("/chk", CheckHandler)

    ctxMux := ContextMiddle(mux)
    http.ListenAndServe(":8080", ctxMux)
}