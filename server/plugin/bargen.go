package plugin

import (
	"bufio"
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/png"
)

// ? we could make this configurable
const frame = 2
const width = 200
const height = 10

// ? could get this from the theme? if so how?
var barFront = color.RGBA{121, 180, 183, 0xff}
var barBG = color.RGBA{254, 251, 243, 0xff}
var barFrame = color.RGBA{157, 157, 157, 0xff}

// XXX we could cache this but do we need to?
func Genbar(progress int64) []byte {
	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{width + frame, height + frame}})

	// frame
	draw.Draw(img, image.Rect(0, 0, width+frame, height+frame), &image.Uniform{barFrame}, image.Point{0, 0}, draw.Src)
	// background
	draw.Draw(img, image.Rect(frame/2, frame/2, width, height), &image.Uniform{barBG}, image.Point{0, 0}, draw.Src)
	// forgroud
	draw.Draw(img, image.Rect(frame/2, frame/2, int(float32(progress/100.0)*float32(width)), height), &image.Uniform{barFront}, image.Point{0, 0}, draw.Src)

	// encode as PNG.
	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	// realisticly if we fail here we just return an empty byte array
	_ = png.Encode(writer, img)
	_ = writer.Flush()
	return b.Bytes()
}
