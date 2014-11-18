package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Actor struct {
	Rect *sdl.Rect
	XVel int32
	YVel int32
}

func (b *Actor) Update(g *Game) {
	b.Rect.X += b.XVel
	b.Rect.Y += b.YVel
}

func (b *Actor) Draw(renderer *sdl.Renderer) {
	renderer.SetDrawColor(255, 255, 255, 255)
	renderer.FillRect(b.Rect)
	renderer.DrawRect(b.Rect)
}
