package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_ttf"
	"os"
)

const winWidth, winHeight int32 = 800, 600
const ScreenCenterX, ScreenCenterY int32 = winWidth / 2, winHeight / 2

type Game struct {
	window                 *sdl.Window
	renderer               *sdl.Renderer
	running                bool
	bat1_score, bat2_score *Label
	divider                *sdl.Rect
	bat1, bat2             *Bat
	ball                   *Ball
	font                   *ttf.Font
}

func NewGame() *Game {
	window := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int(winWidth), int(winHeight), sdl.WINDOW_SHOWN)

	if window == nil {
		fmt.Fprintf(os.Stderr, "Failed to create window", sdl.GetError())
		os.Exit(1)
	}

	renderer := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if renderer == nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", sdl.GetError())
		os.Exit(2)
	}

	ttf.Init()

	font, err := ttf.OpenFont("Courier New.ttf", 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load font: %s\n", sdl.GetError())
		os.Exit(3)
	}

	bat1 := NewBat(10)
	bat2 := NewBat(winWidth - 20)

	ball := NewBall()

	divider := &sdl.Rect{(ScreenCenterX) - 5, 0, 10, winHeight}

	bat1_score := &Label{0, &sdl.Rect{100, 50, 50, 75}, font}
	bat2_score := &Label{0, &sdl.Rect{600, 50, 50, 75}, font}

	return &Game{window, renderer, true, bat1_score, bat2_score, divider, bat1, bat2, ball, font}
}

func (g *Game) EventLoop() {

	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			g.running = false

		case *sdl.KeyDownEvent:
			switch t.Keysym.Sym {
			case sdl.K_a:
				g.bat1.MoveUp()
			case sdl.K_q:
				g.bat1.MoveDown()
			case sdl.K_l:
				g.bat2.MoveUp()
			case sdl.K_o:
				g.bat2.MoveDown()
			}
		}
	}
}

func (g *Game) Update() {
	g.bat1.Update(g)
	g.bat2.Update(g)
	g.ball.Update(g)
}

func (g *Game) Draw() {
	g.renderer.SetDrawColor(0, 0, 0, 255)
	g.renderer.Clear()

	g.renderer.SetDrawColor(255, 255, 255, 255)
	g.renderer.FillRect(g.divider)
	g.renderer.DrawRect(g.divider)

	g.bat1.Draw(g.renderer)
	g.bat2.Draw(g.renderer)
	g.ball.Draw(g.renderer)

	g.bat1_score.Draw(g.renderer)
	g.bat2_score.Draw(g.renderer)

	g.renderer.Present()
}

func (g *Game) Destroy() {
	g.renderer.Destroy()
	g.window.Destroy()
}
