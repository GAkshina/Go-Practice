package main

import "fmt"
// import "math"

func main() {

	//fmt.Printf("Hello world");

	//const pi = math.Pi
	//const pi = 3.14

	//fmt.Println(pi)

	var name string = "Ganesh";
	var nameaddress *string;
	nameaddress = &name
	var sss **string=&nameaddress;

	fmt.Println(*nameaddress);
	fmt.Println("*********************");
	fmt.Println(nameaddress)
	fmt.Println(sss)

	


}