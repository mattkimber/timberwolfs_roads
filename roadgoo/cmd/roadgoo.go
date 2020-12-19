package main

import (
	"flag"
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
		doIndexedImages(a, false)
		doIndexedImages(a, true)
	}
}

func setZero(dst *image.Paletted, x, y, w, h int) {
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			dst.SetColorIndex(i+x, j+y, 0)
		}
	}
}

func blit8bpp(dst, src *image.Paletted, x1, y1, x2, y2, w, h int) {
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			if src.ColorIndexAt(i+x2, j+y2) != 0 {
				dst.SetColorIndex(i+x1, j+y1, src.ColorIndexAt(i+x2, j+y2))
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

func doIndexedImages(name string, elec bool) {
	bounds := image.Rect(0, 0, 600, 34)

	var pal color.Palette
	roadImages := loadImage("gui_rendered/" + name + "_8bpp.png").(*image.Paletted)
	if p, ok := roadImages.ColorModel().(color.Palette); ok {
		pal = p
	}

	spark := loadImage("gui_icons/spark_8bpp.png").(*image.Paletted)

	// Create the output image
	img := image.NewPaletted(bounds, pal)
	draw.Draw(img, bounds, &image.Uniform{C: color.White}, image.Point{}, draw.Over)

	// Draw 2x road
	setZero(img, 0, 0, 20, 20)
	blit8bpp(img, roadImages, 0, 3, 0, 0, 20, 14)
	if elec {
		blit8bpp(img, spark, 0+9, 8, 0, 0, 12, 12)
	}


	setZero(img, 25, 0, 20, 20)
	blit8bpp(img, roadImages, 25, 3, 28, 0, 20, 14)
	if elec {
		blit8bpp(img, spark, 25+9, 8, 0, 0, 12, 12)
	}


	// Draw 1x crossroad
	crossroad := loadImage("gui_rendered/" + name + "_crossroad_8bpp.png").(*image.Paletted)
	setZero(img, 50, 0, 20, 20)
	blit8bpp(img, crossroad, 50, 3, 0, 0, 20, 12)
	if elec {
		blit8bpp(img, spark, 50+9, 8, 0, 0, 12, 12)
	}

	// Draw 1x depot (depot overlaid on track)
	depot := loadImage("gui_rendered/" + flags.Depot + "_8bpp.png").(*image.Paletted)
	setZero(img, 75, 0, 20, 20)
	blit8bpp(img, roadImages, 75, 5, 0, 0, 20, 14)
	blit8bpp(img, depot, 75, 0, 11, 0, 20, 20)
	if elec {
		blit8bpp(img, spark, 75+9, 8, 0, 0, 12, 12)
	}

	// Draw 1x tunnel (tunnel overlaid on track)
	tunnel := loadImage("gui_icons/tunnel_gui_8bpp.png").(*image.Paletted)
	setZero(img, 100, 0, 20, 20)
	blit8bpp(img, roadImages, 100, 4, 0, 0, 20, 14)
	blit8bpp(img, tunnel, 100, 0, 0, 0, 20, 20)
	if elec {
		blit8bpp(img, spark, 100+9, 8, 0, 0, 12, 12)
	}

	// Draw 1x convert (arrow overlaid on track)
	convertArrow := loadImage("gui_icons/convert_icon_8bpp.png").(*image.Paletted)
	setZero(img, 125, 0, 20, 20)
	blit8bpp(img, roadImages, 125, 6, 56, 0, 20, 14)
	blit8bpp(img, convertArrow, 125, 1, 0, 0, 20, 20)
	if elec {
		blit8bpp(img, spark, 125+9, 8, 0, 0, 12, 12)
	}

	// Vertical (|) track
	straightButton := loadImage("gui_icons/icon_straight_8bpp.png").(*image.Paletted)
	crossingButton := loadImage("gui_icons/icon_crossroad_8bpp.png").(*image.Paletted)

	// Diagonal (/) road
	setZero(img, 150, 0, 46, 28)
	blit8bpp(img, straightButton, 150, 0, 0, 0, 46, 28)
	blit8bpp(img, roadImages, 168, 10, 0, 0, 20, 14)
	if elec {
		blit8bpp(img, spark, 168+11, 10+3, 0, 0, 12, 12)
	}

	// Diagonal (\) road
	setZero(img, 200, 0, 46, 28)
	blit8bpp(img, straightButton, 200, 0, 0, 0, 46, 28)
	blit8bpp(img, roadImages, 218, 10, 28, 0, 20, 14)
	if elec {
		blit8bpp(img, spark, 218+11, 10+3, 0, 0, 12, 12)
	}
	// Crossroad
	setZero(img, 250, 0, 46, 28)
	blit8bpp(img, straightButton, 250, 0, 0, 0,46, 28)
	blit8bpp(img, crossroad, 268, 11, 0, 0, 20, 14)
	if elec {
		blit8bpp(img, spark, 268+11, 10+3, 0, 0, 12, 12)
	}
	// Depot
	setZero(img, 300, 0, 54, 34)
	blit8bpp(img, crossingButton, 300, 0, 0, 0, 54, 34)
	blit8bpp(img, roadImages, 300+27, 5+12, 0, 0, 20, 14)
	blit8bpp(img, depot, 300+16, 12, 0, 0, 31, 20)
	if elec {
		blit8bpp(img, spark, 300+27+9, 10+8, 0, 0, 12, 12)
	}

	// Tunnel
	setZero(img, 400, 0, 54, 34)
	blit8bpp(img, crossingButton, 400, 0, 0, 0, 54, 34)
	blit8bpp(img, roadImages, 424, 12+4, 0, 0, 20, 14)
	blit8bpp(img, tunnel, 424, 12, 0, 0, 20, 20)
	if elec {
		blit8bpp(img, spark, 424+9, 10+8, 0, 0, 12, 12)
	}

	// Convert
	setZero(img, 500, 0, 54, 34)
	blit8bpp(img, crossingButton, 500, 0, 0, 0, 54, 34)
	blit8bpp(img, roadImages, 523, 12+4, 56, 0, 20, 14)
	blit8bpp(img, convertArrow, 523, 12, 0, 0, 20, 20)
	if elec {
		blit8bpp(img, spark, 523+9, 10+8, 0, 0, 12, 12)
	}

	// Save the output image
	if elec {
		encode(img, "gui_output/gui_"+name+"_elec_8bpp.png")
	} else {
		encode(img, "gui_output/gui_"+name+"_8bpp.png")
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