package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_ttf"
	"strconv"
	"fmt"
	"os"
)

type Label struct {
	Score int
	Rect  *sdl.Rect
	font  *ttf.Font
}

func (t *Label) Draw(renderer *sdl.Renderer) {

	surface, err := t.font.RenderUTF8_Solid(strconv.Itoa(t.Score), sdl.Color{255, 255, 255, 255})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create surface: %s\n", err)
		os.Exit(1)
	}

	texture, err := renderer.CreateTextureFromSurface(surface)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create texture: %s\n", err)
		os.Exit(1)
	}

	src := &sdl.Rect{0, 0, 100, 100}

	renderer.Copy(texture, src, t.Rect)
}
