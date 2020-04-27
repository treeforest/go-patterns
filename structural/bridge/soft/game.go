package soft

import "fmt"

type Game struct {
	ISoft
}

func (game *Game)Run() error {
	fmt.Println("Run the handset game.")
	return nil
}