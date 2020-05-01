package soft

import (
	"fmt"
	"github.com/treeforest/go-patterns/structural/bridge"
)

// 游戏
type game struct {
}

func NewGame() bridge.Soft {
	return new(game)
}

func (game *game) Run() error {
	fmt.Println("Run the handset game.")
	return nil
}
