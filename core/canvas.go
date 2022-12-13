package core

import (
	"fmt"
	"strings"
)

const height, width int = 12, 12

type gallows struct {
	canvas [][]string
}

func newGallows() *gallows {
	gallow := &gallows{}

	for i := 0; i < height; i++ {
		line := []string{}
		for j := 0; j < width; j++ {
			line = append(line, " ")
		}
		gallow.canvas = append(gallow.canvas, line)
	}

	for i := 0; i < width; i++ {
		gallow.canvas[i][0] = "X"
	}

	for i := 0; i < height/2+1; i++ {
		gallow.canvas[0][i] = "X"
	}

	gallow.canvas[1][height/2] = "X"

	return gallow
}

func (g *gallows) render() {
	for _, line := range g.canvas {
		fmt.Println(strings.Join(line, ""))
	}
}

func drawHead(g *gallows) {
	g.canvas[2][5] = "-"
	g.canvas[2][6] = "-"
	g.canvas[2][7] = "-"
	g.canvas[3][4] = "("
	g.canvas[3][5] = "."
	g.canvas[3][7] = "."
	g.canvas[3][8] = ")"
	g.canvas[4][5] = "-"
	g.canvas[4][6] = "-"
	g.canvas[4][7] = "-"
}

func drawBody(g *gallows) {
	for i := 5; i < 9; i++ {
		g.canvas[i][height/2] = "X"
	}
}

func drawRightArm(g *gallows) {
	for i := 3; i < 7; i++ {
		g.canvas[i][i-1] = "\\"
	}
}

func drawLeftArm(g *gallows) {
	for i := 3; i < 7; i++ {
		g.canvas[i][height-i+1] = "/"
	}
}

func drawRightLeg(g *gallows) {
	g.canvas[9][5] = "/"
	g.canvas[10][4] = "/"
}

func drawLeftLeg(g *gallows) {
	g.canvas[9][7] = "\\"
	g.canvas[10][8] = "\\"
}

func drawLeftHand(g *gallows) {
	g.canvas[2][10] = "\\"
}

func drawRightHand(g *gallows) {
	g.canvas[2][2] = "/"
}

func drawLeftFoot(g *gallows) {
	g.canvas[11][9] = "\\"
	g.canvas[11][10] = "-"
}

func drawRightFoot(g *gallows) {
	g.canvas[11][2] = "-"
	g.canvas[11][3] = "/"
}
