package tictactoe

import (
	"fmt"
	"math/rand"
)

// Size of the board
const SIZE = 3

// Players
type Player int

const (
	EMPTY Player = iota
	CROSS
	CIRCLE
)

var Players = []Player{CROSS, CIRCLE}

// Outcomes
type Outcome int

const (
	NONE Outcome = iota
	CROSS_WIN
	CIRCLE_WIN
	TIE
)

func (p Player) String() string {
	switch {
	case EMPTY == p:
		return ""
	case CROSS == p:
		return "X"
	case CIRCLE == p:
		return "O"
	default:
		panic("Invalid sign, what you up to bro?")
	}
}

func (p Player) toRepr() string {
	switch {
	case EMPTY == p:
		return "-"
	default:
		return fmt.Sprint(p)
	}
}

func (p Player) toOutcome() Outcome {
	switch {
	case CROSS == p:
		return CROSS_WIN
	case CIRCLE == p:
		return CIRCLE_WIN
	default:
		return NONE
	}
}

func (o Outcome) String() string {
	switch {
	case CROSS_WIN == o:
		return fmt.Sprint(CROSS)
	case CIRCLE_WIN == o:
		return fmt.Sprint(CIRCLE)
	case TIE == o:
		return "tie"
	case NONE == o:
		return "none"
	default:
		panic("Invalid outcome, what you up to bro?")
	}
}

type Coord struct {
	X, Y int
}

type Board struct {
	Fields map[Coord]Player
	Turn   Player
}

// Factory function to create a new board
func NewBoard() *Board {
	board := Board{}
	board.Fields = make(map[Coord]Player)
	for x := 0; x < SIZE; x++ {
		for y := 0; y < SIZE; y++ {
			board.Fields[Coord{x, y}] = EMPTY
		}
	}
	players := []Player{CROSS, CIRCLE}
	board.Turn = players[rand.Intn(len(players))]
	return &board
}

func (b *Board) String() string {
	repr := ""
	for y := 0; y < SIZE; y++ {
		repr += fmt.Sprintln()
		for x := 0; x < SIZE; x++ {
			repr += b.Fields[Coord{x, y}].toRepr() + " "
		}
	}
	return repr
}

type Field struct {
	Player string
	Coord  Coord
}

func (b *Board) FieldsList() [][]Field {
	myFields := [][]Field{}
	for y := 0; y < SIZE; y++ {
		rowFields := []Field{}
		for x := 0; x < SIZE; x++ {
			field := Field{Coord: Coord{X: x, Y: y}, Player: fmt.Sprint(b.Fields[Coord{x, y}])}
			rowFields = append(rowFields, field)
		}
		myFields = append(myFields, rowFields)
	}
	return myFields
}

func (b *Board) Play(x int, y int) {
	b.Fields[Coord{x, y}] = b.Turn
	b.nextTurn()
}

func (b *Board) nextTurn() {
	if b.Turn == CIRCLE {
		b.Turn = CROSS
	} else {
		b.Turn = CIRCLE
	}
}

func (b *Board) emptyCoords() []Coord {
	empty := []Coord{}
	for x := 0; x < SIZE; x++ {
		for y := 0; y < SIZE; y++ {
			coord := Coord{x, y}
			if b.Fields[coord] == EMPTY {
				empty = append(empty, coord)
			}
		}
	}
	return empty
}

func (b *Board) allEqual(coords []Coord) bool {
	// See if all fields in the provided coordinates have the same player
	potential := b.Fields[coords[0]]
	found := false
	if potential != EMPTY {
		found = true
		for i := 1; i < len(coords); i++ {
			if potential != b.Fields[coords[i]] {
				found = false
				break
			}
		}
	}
	return found
}

func (b *Board) Winner() Outcome {
	// Check if any row contains only the same player
	for y := 0; y < SIZE; y++ {
		row := []Coord{}
		for x := 0; x < SIZE; x++ {
			row = append(row, Coord{x, y})
		}
		if b.allEqual(row) {
			return b.Fields[row[0]].toOutcome()
		}
	}
	// Check if any column contains only the same player
	for x := 0; x < SIZE; x++ {
		col := []Coord{}
		for y := 0; y < SIZE; y++ {
			col = append(col, Coord{x, y})
		}
		if b.allEqual(col) {
			return b.Fields[col[0]].toOutcome()
		}
	}
	// Check if any diagonal contains only the same player
	diagonal := []Coord{}
	for i := 0; i < SIZE; i++ {
		diagonal = append(diagonal, Coord{i, i})
	}
	if b.allEqual(diagonal) {
		return b.Fields[diagonal[0]].toOutcome()
	}

	diagonal = []Coord{}
	for i := 0; i < SIZE; i++ {
		diagonal = append(diagonal, Coord{2 - i, i})
	}
	if b.allEqual(diagonal) {
		return b.Fields[diagonal[0]].toOutcome()
	}

	if len(b.emptyCoords()) == 0 {
		return TIE
	}
	return NONE
}

func (b *Board) RandomMove() Coord {
	empty := b.emptyCoords()
	return empty[rand.Intn(len(empty))]
}
