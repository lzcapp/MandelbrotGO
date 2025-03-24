package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 10000, 10000
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	//err := png.Encode(os.Stdout, img)
	//if err != nil {
	//	return
	//} // NOTE: ignoring errors

	file, err := os.Create("mandelbrot.png")
	if err != nil {
		log.Fatal(err)
	}
	// 使用png格式将数据写入文件
	err = png.Encode(file, img)
	if err != nil {
		return
	} //将image信息写入文件中
	// 关闭文件
	err = file.Close()
	if err != nil {
		return
	}
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{Y: 255 - contrast*n}
		}
	}
	return color.Black
}
