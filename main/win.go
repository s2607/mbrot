// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// +build ignore
//
// This build tag means that "go install golang.org/x/exp/shiny/..." doesn't
// install this example program. Use "go run main.go" to run it.
// Basic is a basic example of a graphical application.
package main

import (
	"fmt"
	"github.com/s2607/mandelbrot"
	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
	//	"golang.org/x/image/math/f64"
	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
	"image"
	"image/color"
	"image/draw"
	"log"
	//	"math"
)

var (
	blue0 = color.RGBA{0x00, 0x00, 0x1f, 0xff}
	blue1 = color.RGBA{0x00, 0x00, 0x3f, 0xff}
)

func main() {
	driver.Main(func(s screen.Screen) {
		w, err := s.NewWindow(nil)
		if err != nil {
			log.Fatal(err)
		}
		defer w.Release()
		winSize := image.Point{1000, 1000}
		b, err := s.NewBuffer(winSize)
		if err != nil {
			log.Fatal(err)
		}
		defer b.Release()
		drawGradient(b.RGBA())
		t, err := s.NewTexture(winSize)
		if err != nil {
			log.Fatal(err)
		}
		defer t.Release()
		t.Upload(image.Point{}, b, b.Bounds())
		var sz size.Event
		for e := range w.Events() {
			// This print message is to help programmers learn what events this
			// example program generates. A real program shouldn't print such
			// messages; they're not important to end users.
			format := "got %#v\n"
			if _, ok := e.(fmt.Stringer); ok {
				format = "got %v\n"
			}
			fmt.Printf(format, e)
			switch e := e.(type) {
			case lifecycle.Event:
				if e.To == lifecycle.StageDead {
					return
				}
			case key.Event:
				if e.Code == key.CodeEscape {
					return
				}
			case paint.Event:
				w.Fill(sz.Bounds(), blue0, draw.Src)
				w.Fill(sz.Bounds().Inset(10), blue1, draw.Src)
				w.Upload(image.Point{}, b, b.Bounds())
				w.Publish()
			case size.Event:
				sz = e
			case error:
				log.Print(e)
			}
		}
	})
}
func drawGradient(m *image.RGBA) {
	b := m.Bounds()
	var scale = (b.Min.X - b.Max.X) / 2
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			if mandelbrot.Mbi(float64(-x-scale), float64(y+scale), float64(scale)) {
				m.SetRGBA(x, y, color.RGBA{0xff, 0xff, 0xff, 0xff})
			} else {
				m.SetRGBA(x, y, color.RGBA{0x00, 0x00, 0xff, 0xff})
			}
		}
	}
}
