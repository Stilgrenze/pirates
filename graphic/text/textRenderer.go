package text

import (
	"Pirates/models/ocean"
	"errors"
	"fmt"
	"os"
	"os/exec"
)

type TextRenderer struct {

}

func (tr *TextRenderer) clear() {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func (tr *TextRenderer) Draw(ocean *ocean.Ocean) error {
	tr.clear()

	fmt.Println("--------- OCEAN ---------")
	tiles := ocean.Tiles

	width := len(tiles)
	if width == 0 {
		return errors.New("Ocean has no width")
	}
	height := len(tiles[0])
	if height == 0 {
		return errors.New("Ocean has no height")
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			fmt.Print(tiles[x][y].Type)
		}
		fmt.Print("\n")
	}

	return nil
}
