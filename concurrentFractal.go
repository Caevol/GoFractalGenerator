package parfractals

import (
	"time"
	"sync"
)

func GoThreadSet(width, height int, MinX, MaxX, MinY, MaxY float64, setName string) ([][]float64, time.Duration) {
	
	var transformY func(float64, float64, float64) float64
	var transformX func(float64) float64
	
	switch setName {
	case "Mandelbrot":
		transformY = MandelTransformY
		transformX = NoopTransformX
	case "Burning_Ship":
		transformY = ShipTransformY
		transformX = AbsTransformX
	case "Tricorn":
		transformY = TricornTransformY
		transformX = NoopTransformX
	default:
		return nil, 0
	}
	
	imgArr := make([][]float64, height)
	for i := range imgArr {
		imgArr[i] = make([]float64, width)
	}

	t_0 := time.Now()
	
	var wg sync.WaitGroup
	
	for y := 0; y < height; y ++ {
		for x := 0; x < width; x ++ {
			wg.Add(1)
			go func(xt, yt int) {
				defer wg.Done()
				
				imgArr[yt][xt] = float64(GetPixel(xt, yt, width, height, MinX, MaxX, MinY, MaxY, transformY, transformX))
			}(x, y)
		}
	}
	
	wg.Wait()
	
	t_1 := time.Now()
	
	diff := t_1.Sub(t_0)
	
	return imgArr, diff
}