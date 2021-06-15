package main

import "fmt"

func main()  {
	var input int32
	fmt.Scanln(&input)
	fmt.Println( "сотни",input/100, "десятки",input/10%10, "единицы", input%10)
}

