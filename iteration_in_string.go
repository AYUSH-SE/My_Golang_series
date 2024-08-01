package main
import "fmt"

func main(){
	str:="hello"
	for index,val:= range str{
		fmt.Printf("index:%d,val:%c\n",index,val)
	}
}