package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image/jpeg"
	"syscall/js"

	_ "github.com/ajagnic/gogenart/sketch"
)

func main() {
	fmt.Println("Go WASM")
	js.Global().Set("processImage", js.FuncOf(processImage))
	<-make(chan bool)
}

func processImage(this js.Value, args []js.Value) interface{} {
	fmt.Println("processImage")

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

	err = jpeg.Encode(rdr, img, nil)
	if err != nil {
		fmt.Println(err)
	}

	jb, err := json.Marshal(b)
	if err != nil {
		fmt.Println(err)
	}

	return string(jb)
}
