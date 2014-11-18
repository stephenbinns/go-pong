package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"math"
)

const InitialVelocity int32 = 6

type Ball struct {
	*Actor
}

func NewBall() *Ball {
	rect := &sdl.Rect{ScreenCenterX, ScreenCenterY, 10, 10}
	return &Ball{Actor: &Actor{rect, InitialVelocity, InitialVelocity}}
}

func (b *Ball) Update(g *Game) {
	b.Actor.Update(g)

	b.Collides(g.bat1.Actor)
	b.Collides(g.bat2.Actor)

	if b.Rect.Y >= winHeight || b.Rect.Y <= 0 {
		b.YVel *= -1
	}

	if b.Rect.X >= winWidth {
		b.Rect.X, b.Rect.Y = ScreenCenterX, ScreenCenterY
		b.XVel, b.YVel = InitialVelocity, InitialVelocity
		g.bat1_score.Score++
	}

	if b.Rect.X <= 0 {
		b.Rect.X, b.Rect.Y = ScreenCenterX, ScreenCenterY
		b.XVel, b.YVel = InitialVelocity, InitialVelocity
		g.bat2_score.Score++
	}
}

func (b *Ball) Collides(a *Actor) {
	if b.Rect.HasIntersection(a.Rect) {
		b.XVel *= -1
		b.YVel *= -(a.YVel / 2)

		paddleHeight := a.Rect.H

		intersect := float64((a.Rect.Y + (paddleHeight / 2.0)) - b.Rect.Y)
		normalized := float64(intersect / float64(paddleHeight/2.0))
		bounceAngle := normalized * 65.0 // 75 deg max

		// if the bat is the opposite side rotate by 180
		if a.Rect.X > 500 {
			bounceAngle += 180
		}

		// convert to radians - like a boss.
		bounceAngle *= math.Pi / 180

		b.XVel = int32(8.0 * math.Cos(float64(bounceAngle)))
		b.YVel = int32(8.0 * -math.Sin(float64(bounceAngle)))
	}
}
