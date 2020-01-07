package main
  
import (
    "fmt"
    "time"
	"strings"
)
  func dateTime(s string){
  location:=strings.ToLower(s)
   //t := time.Now()
  switch location {
  
  case  "usa":
			loc, _ := time.LoadLocation("America/New_York")
			now := time.Now().In(loc)
			fmt.Println("USA date and time")
			fmt.Println("Location : ", location, " Time : ", now) 
  case "germany":
	         loc,_:=time.LoadLocation("Europe/Berlin")
			 now:=time.Now().In(loc)
			 fmt.Println("germany date and time")
			 fmt.Println("Location:",loc,"time:",now)
	case "japan":
			loc,_:=time.LoadLocation("Asia/Tokyo")
			now:=time.Now().In(loc)
			fmt.Println("japan date and time")
			fmt.Println("Location:",loc,"time:",now)
	case "australia":
	        loc, _ := time.LoadLocation("Australia/Eucla")
			now := time.Now().In(loc)
			fmt.Println("australia date and time")
			fmt.Println("Location : ", loc, " Time : ", now) 
	case "uk":
			loc,_:=time.LoadLocation("Europe/London")
			now:=time.Now().In(loc)
			fmt.Println("uk date and time")
			fmt.Println("Location:",loc,"time:",now)
	default :
			fmt.Println("Please enter a valid country name")
				

  }
  
  }
func main() {
   var s string 
    for{
		fmt.Println("Enter a country name of these :USA,UK,AUSTRALIA,GERMANY,JAPAN")
		fmt.Scanln(&s)
		if s!=""{
		dateTime(s)
		}else{
		fmt.Println("Please enter a valid string ")
		fmt.Scanln(&s)
		}
	}
    
}
