package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image/jpeg"
	"syscall/js"

	"github.com/ajagnic/gogenart/sketch"
)

func main() {
	js.Global().Set("processImage", js.FuncOf(processImage))

	<-make(chan bool)
}

func processImage(this js.Value, args []js.Value) interface{} {
	var values []int
	err := json.Unmarshal([]byte(args[0].String()), &values)
	if err != nil {
		fmt.Println(err)
	}

	var b []byte
	for _, v := range values {
		b = append(b, byte(v))
	}
	rdr := bytes.NewBuffer(b)

	img, err := jpeg.Decode(rdr)
	if err != nil {
		fmt.Println(err)
	}

	canvas := sketch.NewSketch(img, sketch.Params{
		Iterations:         10000,
		Width:              600,
		Height:             600,
		PolygonSidesMin:    3,
		PolygonSidesMax:    5,
		PolygonFillChance:  1.0,
		PolygonColorChance: 0.0,
		PolygonSizeRatio:   0.1,
		PixelShake:         0.0,
		Greyscale:          false,
	})
	canvas.Draw()

	err = jpeg.Encode(rdr, canvas.Image(), nil)
	if err != nil {
		fmt.Println(err)
	}

	out, err := json.Marshal(b)
	if err != nil {
		fmt.Println(err)
	}

	return string(out)
}
