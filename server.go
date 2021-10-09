package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/micmonay/keybd_event"
	qrcode "github.com/skip2/go-qrcode"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	//var port = 9100 + rand.Intn(300)
	var port = 9200
	var host = GetOutboundIP()
	var url = fmt.Sprintf("http://%v:%v/x", host, port)

	fmt.Println("Run at:\n", url)

	png_file := "remote-press-gen.png"
	qrcode.WriteFile(url, qrcode.Medium, 256, png_file)

	e := echo.New()
	e.Debug = true
	//e.GET("/", func(c echo.Context) error {
	//	return c.HTML(http.StatusOK, "<a href=\"/space\">\n<input type=button value=\"php中文网\" >\n</a>")
	//})
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   "static",
		Browse: true,
	}))
	e.Static("/x", "static/c.html")
	//e.GET("/", func(c echo.Context) error {
	//	return c.Static(200, "c.html")
	//})
	e.GET("/space", func(c echo.Context) error {
		fmt.Println("click space @", time.Now())
		//return c.String(http.StatusOK, "clicked")
		kb, err := keybd_event.NewKeyBonding()
		if err != nil {
			panic(err)
		}
		// For linux, it is very important to wait 2 seconds
		if runtime.GOOS == "linux" {
			time.Sleep(2 * time.Second)
		}
		// Select keys to be pressed
		kb.SetKeys(keybd_event.VK_SPACE)
		kb.Press()
		time.Sleep(10 * time.Millisecond)
		kb.Release()

		return c.JSON(200, "ok")
	})
	e.Logger.Fatal(e.Start(":" + fmt.Sprint(port)))
}
