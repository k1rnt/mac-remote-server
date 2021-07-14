package main

import (
	"runtime"

	"github.com/k1rnt/mac-remote-server/route"
)

func main() {
	if runtime.GOOS != "darwin" {
		panic("this application is only supported on macOS")
	}
	app := route.Init()
	app.Logger.Fatal(app.Start(":8881"))
}
