package soft

import (
	"fmt"
	"github.com/treeforest/go-patterns/structural/bridge"
)

type game struct {
}

func NewGame() bridge.ISoft {
	return new(game)
}

func (game *game) Run() error {
	fmt.Println("Run the handset game.")
	return nil
}
