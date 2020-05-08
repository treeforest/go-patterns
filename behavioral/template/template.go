package main

import "fmt"

type Player interface {
	Play()
}

// 抽象基类
type Game struct {
	Initialize func()
	StartPlay func()
	EndPlay func()
	Player
}

// 模板
func (game *Game) Play() {
	// 初始化游戏
	game.Initialize()
	// 开始游戏
	game.StartPlay()
	// 结束游戏
	game.EndPlay()
}

type Cricket struct {
	Game
}

func NewCricket() Player {
	c := new(Cricket)
	c.Game.Initialize = c.Initialize
	c.Game.StartPlay = c.StartPlay
	c.Game.EndPlay = c.EndPlay
	return c
}

func (c *Cricket) Initialize() {
	fmt.Println("Cricket Game Initialized! Start playing.")
}

func (c *Cricket) StartPlay() {
	fmt.Println("Cricket Game Started. Enjoy the game!")
}

func (c *Cricket) EndPlay() {
	fmt.Println("Cricket Game Finished!")
}

type Football struct {
	Game
}

func NewFootball() Player {
	f := new(Football)
	f.Game.Initialize = f.Initialize
	f.Game.StartPlay = f.StartPlay
	f.Game.EndPlay = f.EndPlay
	return f
}

func (f *Football) Initialize() {
	fmt.Println("Football Game Initialized! Start playing.")
}

func (f *Football) StartPlay() {
	fmt.Println("Football Game Started. Enjoy the game!")
}

func (f *Football) EndPlay() {
	fmt.Println("Football Game Finished!")
}

func main() {
	var play Player

	play = NewCricket()
	play.Play()

	play = NewFootball()
	play.Play()
}