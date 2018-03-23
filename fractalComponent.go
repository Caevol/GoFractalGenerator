package parfractals

import (
	"math"
)


const (
	MaxIterations = 1000
)

func scaleValue(x_0, originalMax, newMax, newMin float64) float64 {
	x := (x_0 / originalMax) * (newMax - newMin) + newMin
	return x
}

//implementation based on the algorithms listed on Wikipedia
//transformX and transformY are functions, Recommended pairs are listed below
func GetPixel(pixelX,
		pixelY, imageWidth, imageHeight int,
		MinX, MaxX, MinY, MaxY float64,
		transformY func(float64, float64, float64) float64,
		transformX func(float64) float64) int {
	x_0 := scaleValue(float64(pixelX), float64(imageWidth), MaxX, MinX)
	y_0 := scaleValue(float64(pixelY), float64(imageHeight), MaxY, MinY)
	
	x := x_0
	y := y_0
	
	counter := 0
	
	for ; (x*x + y*y < 4) && (counter < MaxIterations); counter ++ {
		tmp := x*x - y*y + x_0
		y = transformY(x, y, y_0) //transform y
		x = transformX(tmp) //transform x
	}
	
	return counter
	
}

//Available TransformY and TransformX functions

//Mandebrot: MandelTransformY, NoopTransformX
//BurningShip: ShipTransformY, AbsTransformX
//Tricorn: TricornTransformY, NoopTransformX
func TricornTransformY(x, y, y_0 float64) float64 {
	return -2*x*y + y_0
}

func MandelTransformY(x, y, y_0 float64) float64 {
	return 2*x*y + y_0
}

func ShipTransformY(x, y, y_0 float64) float64 {
	return math.Abs(MandelTransformY(x, y, y_0))
}

func NoopTransformX(tmp float64) float64 {
	return tmp
}

func AbsTransformX(tmp float64) float64 {
	return math.Abs(tmp)
}
