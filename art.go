package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
	"github.com/olekukonko/ts"
	"github.com/inancgumus/screen"
)

func main() {
	screen.Clear()
	for {
		size, _ := ts.GetSize()
		for x := 0; x < size.Row(); x++ {
			for i := 0; i < size.Col(); i++ {
				switch rand.Intn(5) {
				case 1:
					color.Set(color.FgBlue)
				case 2:
					color.Set(color.FgCyan)
				case 3:
					color.Set(color.FgYellow)
				case 4:
					color.Set(color.FgRed)
				case 5:
					color.Set(color.FgGreen)
				case 6:
					color.Set(color.FgMagenta)
				}
				fmt.Printf("▊")
			}
		}
		color.Unset()
		time.Sleep(5 * time.Second)
	}
}
