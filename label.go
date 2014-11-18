package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_ttf"
	"strconv"
)

type Label struct {
	Score int
	Rect  *sdl.Rect
	font  *ttf.Font
}

func (t *Label) Draw(renderer *sdl.Renderer) {

	surface := t.font.RenderText_Solid(strconv.Itoa(t.Score), sdl.Color{255, 255, 255, 255})
	texture := renderer.CreateTextureFromSurface(surface)

	src := &sdl.Rect{0, 0, 100, 100}

	renderer.Copy(texture, src, t.Rect)
}
