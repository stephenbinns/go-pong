package main

func main() {
	game := NewGame()

	for game.running {
		game.EventLoop()
		game.Update()
		game.Draw()
	}

	game.Destroy()
}
