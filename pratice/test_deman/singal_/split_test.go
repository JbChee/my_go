package singal_


import (
	//"fmt"

	//"reflect"
	"testing"
)
//
//func TestSplit(t *testing.T){
//	got := Split("a:b:c:d",":")
//
//	want := []string{"a","b","c"}
//	if !reflect.DeepEqual(want,got){    //slice 不能直接比较
//
//		t.Errorf("expected:%v,got:%v",want,got)
//	}
//
//}


//func TestMoreSplit(t *testing.T){
//	got := Split("chen#jing#bo","#")
//
//	fmt.Println(got)
//
//	want := []string{"chen","jing","bo"}
//	if !reflect.DeepEqual(want,got){    //slice 不能直接比较
//
//		t.Errorf("expected:%#v, got:%#v",want,got)  //按go相应的语法打印值
//	}
//
//}

func BenchmarkSplit(b *testing.B){


	for i:=0;i<b.N;i++{
		Split("中国好声音，中国新世代","中国")

	}

}





