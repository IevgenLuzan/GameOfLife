package main

import "time"

const StateAlive = 1
const StateDead = 0
const BoardSideSize = 25

type GameBoard struct {
	State         [][]int
	Rows, Columns int
}

type Cell struct {
	X, Y int
}

type printable interface {
	GetCells() []*Cell
}

type Figure struct {
	Cells []*Cell
}

func (f *Figure) GetCells() []*Cell {
	return f.Cells
}

//  0 1 0
//  0 0 1
//  1 1 1
func getGlider () printable {
	figure := new(Figure)
	figure.Cells = append(figure.Cells, &Cell{X: 0,Y: 2})
	figure.Cells = append(figure.Cells, &Cell{X: 1,Y: 0})
	figure.Cells = append(figure.Cells, &Cell{X: 2,Y: 2})
	figure.Cells = append(figure.Cells, &Cell{X: 2,Y: 1})
	figure.Cells = append(figure.Cells, &Cell{X: 1,Y: 2})

	return figure
}

func getStartFigureCells(figureName string) printable {
	switch figureName {
	case "glider":
		return getGlider()
	default:
		// Just for test
		return getGlider()
	}
}

func main() {

	gliderBoard := InitBoard(BoardSideSize, "glider")
	gliderBoard.Print()
	for {
		gliderBoard.NewGeneration()
		gliderBoard.Print()
		time.Sleep(1 * time.Second)
	}

}
func InitBoard(size int, startFigure string) GameBoard {

	state := make([][]int, size)
	for i := range state {
		state[i] = make([]int, size)
	}

	// Print Glider at the middle
	figure := getStartFigureCells(startFigure)
	for _, cell := range figure.GetCells() {
		state[11+cell.Y][11+cell.X] = 1
	}

	return GameBoard{State: state, Rows: size, Columns: size}
}

func (b *GameBoard) NewGeneration() {

	board := make([][]int, b.Rows)

	for row := range board {
		board[row] = make([]int, b.Columns)
		for col := range board[row] {
			board[row][col] = getStateForCell(b, row, col)
		}
	}

	b.State = board
}

func getStateForCell(b *GameBoard, i, j int) int {

	neighborsAlive := 0
	cellState := b.State[i][j]

	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if i+x < 0 || i+x > (b.Rows-1) || y+j < 0 || y+j > (b.Columns-1) {
				continue
			}
			neighborsAlive += b.State[i+x][y+j]
		}
	}
	neighborsAlive -= cellState

	if cellState == StateDead && neighborsAlive == 3 {
		return StateAlive
	} else if cellState == StateAlive && (neighborsAlive < 2 || neighborsAlive > 3) {
		return StateDead
	}

	return cellState
}

func (b *GameBoard) Print() {

	for i := range b.State {
		for j := range b.State[i] {
			print(b.State[i][j])
		}
		println()
	}
	println()
}
