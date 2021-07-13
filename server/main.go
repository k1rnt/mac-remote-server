package main

import (
	"github.com/k1rnt/mac-remote-server/route"
)

func main() {
	app := route.Init()
	app.Logger.Fatal(app.Start(":8881"))
}
