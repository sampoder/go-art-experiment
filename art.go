package main

import "fmt"

import "time"

import "github.com/fatih/color"

import "github.com/olekukonko/ts"

import "github.com/inancgumus/screen"

import "math/rand"

func main() {
	screen.Clear()
	for true {
		size, _ := ts.GetSize()
		for x := 0; x < size.Row(); x++ {
			for i := 0; i < size.Col(); i++ {
				random := rand.Intn(5)
				switch random {
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
				color.Unset()
			}
		}
		time.Sleep(5 * time.Second)
	}
}