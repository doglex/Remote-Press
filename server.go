package main

import (
	"fmt"
	qrcode "github.com/skip2/go-qrcode"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var port = 9100 + rand.Intn(300)
	var host = GetOutboundIP()
	var url = fmt.Sprintf("http://%v:%v", host, port)
	fmt.Println("Run at:\n", url)

	png_file := "remote-press-gen.png"
	qrcode.WriteFile("xxxxxx", qrcode.Medium, 256, png_file)

}
