
package main
import "fmt"
func fibonocci(x int){
	a:=0
	b:=1
	fmt.Print(a,",")
	fmt.Print(b)
	i:=2
	for i<=x{
		c:=a+b
		if c<x{
			fmt.Print(",")
			fmt.Print(c)
		
		}
		a=b
		b=c
		i++
		
		}
		
	
	}

func main(){
	var n int
	for {
	
	fmt.Println("enter n value")
	fmt.Scanln(&n)
	
	if n<=0 {
	fmt.Println("please enter  positive number ")
	fmt.Scanln(&n)
	fibonocci(n)
	}else{fibonocci(n)
	}
	}
}