# Template Pattern 

模板模式：定义一个操作中的算法的骨架，而将一些步骤

# Implementation

```
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
```

# Usage

```
var play Player

play = NewCricket()
play.Play()

play = NewFootball()
play.Play()
```

```
Cricket Game Initialized! Start playing.
Cricket Game Started. Enjoy the game!
Cricket Game Finished!
Football Game Initialized! Start playing.
Football Game Started. Enjoy the game!
Football Game Finished!
```

# Rule of Thumb

模板模式提供了一个很好的代码服用平台。

当不变的和可变的行为在方法的子类实现中混合在一起的时候，不变的行为就会在子类中重复出现。我们通过模板模式把这些行为搬移到单一的地方，这样就帮助子类摆脱重复的不变的行为的纠缠。