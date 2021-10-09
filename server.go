package main

import (
	"fmt"
	qrcode "github.com/skip2/go-qrcode"
	"math/rand"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var port = 9100 + rand.Intn(300)
	var host = GetOutboundIP()
	var url = fmt.Sprintf("http://%v:%v", host, port)

	png_file := "remote-press-gen.png"
	qrcode.WriteFile(url, qrcode.Medium, 256, png_file)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":" + fmt.Sprint(port)))
	fmt.Println("Run at:\n", url)

}
