package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"syscall/js"

	"github.com/ajagnic/gogenart/sketch"
)

func main() {
	js.Global().Set("processImage", js.FuncOf(processImage))

	<-make(chan bool)
}

func processImage(this js.Value, args []js.Value) interface{} {
	b := args[0]
	size := args[1].Int()
	paramsJSON := []byte(args[2].String())

	var params sketch.Params
	err := json.Unmarshal(paramsJSON, &params)
	if err != nil {
		fmt.Println(err)
	}

	buff := make([]byte, size)
	js.CopyBytesToGo(buff, b)
	rdr := bytes.NewBuffer(buff)

	img, enc, err := image.Decode(rdr)
	if err != nil {
		fmt.Println(err)
	}

	canvas := sketch.NewSketch(img, params)
	canvas.Draw()

	switch enc {
	case "png":
		png.Encode(rdr, canvas.Image())
	default:
		jpeg.Encode(rdr, canvas.Image(), nil)
	}

	out, err := json.Marshal(buff)
	if err != nil {
		fmt.Println(err)
	}

	return string(out)
}
