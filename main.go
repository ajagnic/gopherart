package main

import (
	"bytes"
	"encoding/json"
	"fmt"
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

	img, enc := sketch.Source(rdr)

	newImg := sketch.NewSketch(img, params).Draw()

	rdr.Reset()
	sketch.Encode(rdr, newImg, enc)

	out, err := json.Marshal(rdr.Bytes())
	if err != nil {
		fmt.Println(err)
	}

	return string(out)
}
