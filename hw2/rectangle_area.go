package main

import (
	"fmt"
)

func main ()  {
	var j,i int32
	fmt.Println("Введите длину прямоугольника")
	fmt.Scanln(&j)
	fmt.Println("Введите ширину прямоугольника")
	fmt.Scanln(&i)
	fmt.Println("Площадь прямоугольника", j*i)
}

