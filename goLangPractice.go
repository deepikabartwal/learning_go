package main 
import (
	"fmt"
	"time"
	"math"
)

func add(x , y int) int {
	return x + y 
}

func split(sum int)(x,y int){
	x = sum*4/9
	y = sum - x
	return 
}

func swap(x,y string)(string, string){
	return y,x
}

var c, python, java bool

var o,p int = 1,2

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main(){
	fmt.Println("Hello,World")
	fmt.Println(add(42,24))
	fmt.Println("Welcome to the playground!")

	fmt.Println("The time is", time.Now())
	fmt.Println(split(17))
	a,b := swap("Hello","World")
	fmt.Println(a,b)
	fmt.Println(math.Pi)
	var i int
	fmt.Println(i,java,c,python)
	var l, m int = 1,2
	k:=3
	c,python,java := true,false,"no!"
	fmt.Println(l,m,k,c,python,java)
	var card, javas, pythons = true, false, "no!"
	fmt.Println(o,p,card,javas,pythons)
	// sum := 0
	// for i := 0;i<=10;i++{
	// 	sum+=1
	// }

	sum := 1
	for ; sum<1000;{
		sum+=sum
	}
	fmt.Println(sum)
	fmt.Println(sqrt(2), sqrt(-4))
}