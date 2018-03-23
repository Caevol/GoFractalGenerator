package parfractals

import (
	"image"
	"image/color"
	"image/png"
	"os"

)

func OutputImage(width, height int, imageArr [][]float64, filename string, ColorScale [3]float64) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	
	for y := range imageArr {
		for x := range imageArr[y] {
			pixelVal := imageArr[y][x]
			img.Set(x, y, color.RGBA{
				uint8(ColorScale[0]*pixelVal), 
				uint8(ColorScale[1]*pixelVal),
				uint8(ColorScale[2]*pixelVal), 255})
		}
	}
	
	
	f, _ := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)
	
}