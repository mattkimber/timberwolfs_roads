package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

type Flags struct {
	Depot string
}

var flags Flags

func init() {
	flag.StringVar(&flags.Depot, "depot", "depot", "Name of the depot sprite to use (default: depot)")

	flag.Parse()
}

func main() {
	for _, a := range flag.Args() {
		// A little inefficient to do electric for everything, but
		// we can live...
		doIndexedImages(1, a, false)
		doIndexedImages(1, a, true)
		doIndexedImages(2, a, false)
		doIndexedImages(2, a, true)
	}
}

func setZero(scale int, dst *image.Paletted, x, y, w, h int) {
	for i := 0; i < w*scale; i++ {
		for j := 0; j < h*scale; j++ {
			dst.SetColorIndex(i+(x*scale), j+(y*scale), 0)
		}
	}
}

func blit8bpp(scale int, dst, src *image.Paletted, x1, y1, x2, y2, w, h int) {
	for i := 0; i < w*scale; i++ {
		for j := 0; j < h*scale; j++ {
			if src.ColorIndexAt(i+(x2*scale), j+(y2*scale)) != 0 {
				dst.SetColorIndex(i+(x1*scale), j+(y1*scale), src.ColorIndexAt(i+(x2*scale), j+(y2*scale)))
			}
		}
	}
}

func loadImage(filename string) image.Image {
	r, err := os.Open(filename)
	defer r.Close()
	if err != nil {
		panic(err)
	}

	img, err := png.Decode(r)
	if err != nil {
		panic(err)
	}

	return img
}

func doIndexedImages(scale int, name string, elec bool) {
	bounds := image.Rect(0, 0, 600*scale, 34*scale)

	var pal color.Palette
	roadImages := loadImage(fmt.Sprintf("gui_rendered/%dx/%s_8bpp.png", scale, name)).(*image.Paletted)
	if p, ok := roadImages.ColorModel().(color.Palette); ok {
		pal = p
	}

	spark := loadImage(fmt.Sprintf("gui_icons/%dx/spark_8bpp.png", scale)).(*image.Paletted)

	// Create the output image
	img := image.NewPaletted(bounds, pal)
	draw.Draw(img, bounds, &image.Uniform{C: color.White}, image.Point{}, draw.Over)

	// Draw 2x road
	setZero(scale, img, 0, 0, 20, 20)
	blit8bpp(scale, img, roadImages, 0, 3, 0, 0, 20, 14)
	if elec {
		blit8bpp(scale, img, spark, 0+9, 8, 0, 0, 12, 12)
	}

	setZero(scale, img, 25, 0, 20, 20)
	blit8bpp(scale, img, roadImages, 25, 3, 28, 0, 20, 14)
	if elec {
		blit8bpp(scale, img, spark, 25+9, 8, 0, 0, 12, 12)
	}

	// Draw 1x crossroad
	crossroad := loadImage(fmt.Sprintf("gui_rendered/%dx/%s_crossroad_8bpp.png", scale, name)).(*image.Paletted)
	setZero(scale, img, 50, 0, 20, 20)
	blit8bpp(scale, img, crossroad, 50, 3, 0, 0, 20, 12)
	if elec {
		blit8bpp(scale, img, spark, 50+9, 8, 0, 0, 12, 12)
	}

	// Draw 1x depot (depot overlaid on track)
	depot := loadImage(fmt.Sprintf("gui_rendered/%dx/%s_8bpp.png", scale, flags.Depot)).(*image.Paletted)
	setZero(scale, img, 75, 0, 20, 20)
	blit8bpp(scale, img, roadImages, 75, 5, 0, 0, 20, 14)
	blit8bpp(scale, img, depot, 75, 0, 11, 0, 20, 20)
	if elec {
		blit8bpp(scale, img, spark, 75+9, 8, 0, 0, 12, 12)
	}

	// Draw 1x tunnel (tunnel overlaid on track)
	tunnel := loadImage(fmt.Sprintf("gui_icons/%dx/tunnel_gui_8bpp.png", scale)).(*image.Paletted)
	setZero(scale, img, 100, 0, 20, 20)
	blit8bpp(scale, img, roadImages, 102, 4, 0, 0, 18, 14)
	blit8bpp(scale, img, tunnel, 100, 0, 0, 0, 20, 20)
	if elec {
		blit8bpp(scale, img, spark, 100+9, 8, 0, 0, 12, 12)
	}

	// Draw 1x convert (arrow overlaid on track)
	convertArrow := loadImage(fmt.Sprintf("gui_icons/%dx/convert_icon_8bpp.png", scale)).(*image.Paletted)
	setZero(scale, img, 125, 0, 20, 20)
	blit8bpp(scale, img, roadImages, 125, 6, 56, 0, 20, 14)
	blit8bpp(scale, img, convertArrow, 125, 1, 0, 0, 20, 20)
	if elec {
		blit8bpp(scale, img, spark, 125+9, 8, 0, 0, 12, 12)
	}

	// Vertical (|) track
	straightButton := loadImage(fmt.Sprintf("gui_icons/%dx/icon_straight_8bpp.png", scale)).(*image.Paletted)
	crossingButton := loadImage(fmt.Sprintf("gui_icons/%dx/icon_crossroad_8bpp.png", scale)).(*image.Paletted)

	// Diagonal (/) road
	setZero(scale, img, 150, 0, 46, 28)
	blit8bpp(scale, img, straightButton, 150, 0, 0, 0, 46, 28)
	blit8bpp(scale, img, roadImages, 168, 10, 0, 0, 20, 14)
	if elec {
		blit8bpp(scale, img, spark, 168+11, 10+3, 0, 0, 12, 12)
	}

	// Diagonal (\) road
	setZero(scale, img, 200, 0, 46, 28)
	blit8bpp(scale, img, straightButton, 200, 0, 0, 0, 46, 28)
	blit8bpp(scale, img, roadImages, 218, 10, 28, 0, 20, 14)
	if elec {
		blit8bpp(scale, img, spark, 218+11, 10+3, 0, 0, 12, 12)
	}
	// Crossroad
	setZero(scale, img, 250, 0, 46, 28)
	blit8bpp(scale, img, straightButton, 250, 0, 0, 0, 46, 28)
	blit8bpp(scale, img, crossroad, 268, 11, 0, 0, 20, 14)
	if elec {
		blit8bpp(scale, img, spark, 268+11, 10+3, 0, 0, 12, 12)
	}
	// Depot
	setZero(scale, img, 300, 0, 54, 34)
	blit8bpp(scale, img, crossingButton, 300, 0, 0, 0, 54, 34)
	blit8bpp(scale, img, roadImages, 300+27, 5+12, 0, 0, 20, 14)
	blit8bpp(scale, img, depot, 300+16, 12, 0, 0, 31, 20)
	if elec {
		blit8bpp(scale, img, spark, 300+27+9, 10+8, 0, 0, 12, 12)
	}

	// Tunnel
	setZero(scale, img, 400, 0, 54, 34)
	blit8bpp(scale, img, crossingButton, 400, 0, 0, 0, 54, 34)
	blit8bpp(scale, img, roadImages, 424+2, 12+4, 0, 0, 20-2, 14)
	blit8bpp(scale, img, tunnel, 424, 12, 0, 0, 20, 20)
	if elec {
		blit8bpp(scale, img, spark, 424+9, 10+8, 0, 0, 12, 12)
	}

	// Convert
	setZero(scale, img, 500, 0, 54, 34)
	blit8bpp(scale, img, crossingButton, 500, 0, 0, 0, 54, 34)
	blit8bpp(scale, img, roadImages, 523, 12+4, 56, 0, 20, 14)
	blit8bpp(scale, img, convertArrow, 523, 12, 0, 0, 20, 20)
	if elec {
		blit8bpp(scale, img, spark, 523+9, 10+8, 0, 0, 12, 12)
	}

	// Save the output image
	if elec {
		encode(img, fmt.Sprintf("gui_output/%dx/gui_%s_elec_8bpp.png", scale, name))
	} else {
		encode(img, fmt.Sprintf("gui_output/%dx/gui_%s_8bpp.png", scale, name))
	}
}

func encode(img image.Image, filename string) {
	of, err := os.Create(filename)
	defer of.Close()
	if err != nil {
		panic(err)
	}

	if err := png.Encode(of, img); err != nil {
		panic(err)
	}
}
