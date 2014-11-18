package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const Velocity int32 = 5

type Bat struct {
	*Actor
}

func NewBat(x int32) *Bat {
	rect := &sdl.Rect{x, ScreenCenterY - 50, 10, 100}
	return &Bat{Actor: &Actor{rect, 0, 0}}
}

func (b *Bat) MoveUp() {
	b.YVel = Velocity
}

func (b *Bat) MoveDown() {
	b.YVel = -Velocity
}

func (b *Bat) Update(g *Game) {
	b.Actor.Update(g)

	if b.Rect.Y+b.Rect.H >= winHeight || b.Rect.Y <= 0 {
		b.YVel = 0
	}
}
