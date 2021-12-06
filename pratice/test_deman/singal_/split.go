package singal_

import (
	//"fmt"
	"strings"

)

func Split(s, sep string) (result []string){
	i := strings.Index(s, sep)
	//fmt.Println("index",i)

	for i > -1{
		result = append(result,s[:i])
		s = s[i+1:]
		i = strings.Index(s,sep)
	}
	result = append(result,s)
	return result
}





