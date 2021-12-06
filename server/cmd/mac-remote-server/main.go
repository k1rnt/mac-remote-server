package main

import (
	"runtime"
)

func main() {
	if runtime.GOOS != "darwin" {
		panic("this application is only supported on macOS")
	}
	app := setRouter()
	app.Logger.Fatal(app.Start(":8881"))
}
