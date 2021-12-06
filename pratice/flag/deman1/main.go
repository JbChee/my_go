package main


import (
	"fmt"
	"os"

)



func main() {

	if len(os.Args) > 0{
		for index, args := range os.Args{
			fmt.Printf("\n args[%d] = %#v\n",index,args)
		}
	}

}
