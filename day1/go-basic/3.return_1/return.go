package main 
import "fmt" 
 
 
 func main() { 
 	fmt.Println(testReturn("2")) 
 } 
 
 func testReturn(a string) (string, string) { 
 	return "3", "2" 
 } 
