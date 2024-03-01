package main

import (
	"fmt"
	"github.com/mewzax/gocolors"
)

func main() {
    fmt.Println(ColorizeRed("gokata"))
	fmt.Println(ColorizeGreen("gokata"))
	fmt.Println(ColorizeBlue("gokata"))
	fmt.Println(ColorizeYellow("gokata"))
	fmt.Println(ColorizeMagenta("gokata"))
	fmt.Println(ColorizeCyan("gokata"))
	fmt.Println(ColorizeWhite("gokata"))
	fmt.Println(ColorizeCustom("gokata", 233, 54, 76))
}

func ColorizeRed(a string) string {
	return gocolors.Colorize(gocolors.Red, a)
}

func ColorizeGreen(a string) string {
	return gocolors.Colorize(gocolors.Green, a)
}

func ColorizeBlue(a string) string {
	return gocolors.Colorize(gocolors.Blue, a)
}

func ColorizeYellow(a string) string {
	return gocolors.Colorize(gocolors.Yellow, a)
}

func ColorizeMagenta(a string) string {
	return gocolors.Colorize(gocolors.Magenta, a)
}

func ColorizeCyan(a string) string {
	return gocolors.Colorize(gocolors.Cyan, a)
}

func ColorizeWhite(a string) string {
	return gocolors.Colorize(gocolors.White, a)
}

func ColorizeCustom(a string, r, g, b uint8) string {
	return gocolors.Colorize(gocolors.RGB(int(r), int(g), int(b)), a)
}

