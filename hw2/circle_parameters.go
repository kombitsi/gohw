package main

import (
	"fmt"
	"math"
)

func main()  {
	var s float64
	var d float64
	fmt.Println("Введите площадь круга")
	fmt.Scanln(&s)
	d = 2*math.Sqrt(s / math.Pi)
	fmt.Println("Диаметр будет равен",d)
	fmt.Println("Длина окружности будет равена", math.Pi*d)
}





