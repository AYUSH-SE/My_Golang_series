package main
import "fmt"

func main(){
	num:=[]int{10,20,30,40,50};
	for idx,val:=range num{
		fmt.Printf("index:%d, value:%d\n",idx,val)
	}
}