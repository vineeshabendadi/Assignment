package main
 
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)
 func find(file []string,a string) int {
		for line,s:=range file{
			if strings.Contains(s,a){
				return line+1
			}
		}
		return -1
	}
func main() {
	file, err := os.Open("sample.txt")
 
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
 
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string
 
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	var a string
	fmt.Println("enter a string to search")
	fmt.Scanln(&a)
	if a!=""{
 result:=find(txtlines,a)
 if result==-1{
 fmt.Println("String not found")
 }else{
 fmt.Println("String found at line ",result)
 }}else{
 fmt.Println("please enter valid string ")
 }
 //fmt.Println(result)
	file.Close()
}
 