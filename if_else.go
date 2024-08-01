package main
import "fmt"

// func main(){
// 	var i int=10;
// 	if(i%2==0){
// 	fmt.Printf("divisible by 2")
// 	}
// }

func main() {
    fmt.Printf("Enter the number: ")
    var input int
    fmt.Scanln(&input)
    fmt.Println(input)
    if input%2 == 0 {
        fmt.Println("even")
    } else {
        fmt.Println("odd")
    }
}