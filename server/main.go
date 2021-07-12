package main

import (
	"github.com/andybrewer/mack"
	_ "github.com/labstack/echo/v4"
	_ "github.com/progrium/macdriver/cocoa"
	_ "golang.org/x/xerrors"
)

func main() {
	mack.Say("Hello, World!")
}
