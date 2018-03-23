package main

import (
	"parfractals"
	"fmt"
)

func main() {

	ColorScale := [3]float64{.255, .255, .255}

	fmt.Println("Linear Times")
	
	imgArr, t_0 := parfractals.CalculateSetLinear(1750, 1000, -2.5, 1, -1, 1, "Burning_Ship")
	parfractals.OutputImage(1750, 1000, imgArr, "Burning_ship.png", ColorScale)
	fmt.Println("Burning Ship: ", t_0)
	
	imgArr, t_1 := parfractals.CalculateSetLinear(1750, 1000, -2.5, 1, -1, 1, "Mandelbrot")
	parfractals.OutputImage(1750, 1000, imgArr, "Mandelbrot.png", ColorScale)
	fmt.Println("MandelBrot: ", t_1)
	
	imgArr, t_2 := parfractals.CalculateSetLinear(1750, 1000, -2, 2, -2, 2, "Tricorn")
	parfractals.OutputImage(1750, 1000, imgArr, "Tricorn.png", ColorScale)
	fmt.Println("Tricorn: ", t_2)

	
	fmt.Println("Goroutines Times")
	
	imgArr, t_0 = parfractals.GoThreadSet(1750, 1000, -2.5, 1, -1, 1, "Burning_Ship")
	fmt.Println("Burning Ship: ", t_0)
	
	imgArr, t_1 = parfractals.GoThreadSet(1750, 1000, -2.5, 1, -1, 1, "Mandelbrot")
	fmt.Println("MandelBrot: ", t_1)
	
	imgArr, t_2 = parfractals.GoThreadSet(1750, 1000, -2, 2, -2, 2, "Tricorn")
	fmt.Println("Tricorn: ", t_2)


	
}