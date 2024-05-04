package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{color.Black, color.CMYK{74, 0, 82, 1}}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // second color in palette
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 500   // image canvas covers [-size..+size]
		nframes = 165   // number of animation frames
		delay   = 1     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*3*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+1.5), size+int(y*size+1.5),
				blackIndex)
		}
		phase += 0.029
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	// Открываем файл для записи и сохраняем анимацию
	outFile, err := os.Create("output.gif")
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer outFile.Close()

	if err := gif.EncodeAll(outFile, &anim); err != nil {
		log.Fatalf("Failed to encode GIF: %v", err)
	}
}
