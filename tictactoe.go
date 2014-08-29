package main

import "math/rand"

var board Board

// Players
type Player int

const (
	CROSS Player = iota
	CIRCLE
	EMPTY
)

// Outcomes
type Outcome int

const (
	CROSS_WIN Outcome = iota
	CIRCLE_WIN
	TIE
	NONE
)

func OutcomeToString(o Outcome) string {
	switch {
	case CROSS_WIN == o:
		return PlayerToString(CROSS)
	case CIRCLE_WIN == o:
		return PlayerToString(CIRCLE)
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
	fields map[Coord]Player
	turn   Player
}

func PlayerToString(i Player) string {
	switch {
	case EMPTY == i:
		return ""
	case CROSS == i:
		return "X"
	case CIRCLE == i:
		return "O"
	default:
		panic("Invalid sign, what you up to bro?")
	}
}

func playerToRepr(i Player) string {
	switch {
	case EMPTY == i:
		return "-"
	default:
		return PlayerToString(i)
	}
}

// Factory function to create a new board
func NewBoard() Board {
	board := Board{}
	board.fields = make(map[Coord]Player)
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			board.fields[Coord{x, y}] = EMPTY
		}
	}
	players := []Player{CROSS, CIRCLE}
	board.turn = players[rand.Intn(len(players))]
	return board
}

func (b *Board) repr() string {
	repr := ""
	for y := 0; y < 3; y++ {
		repr += "\n"
		for x := 0; x < 3; x++ {
			repr += playerToRepr(b.fields[Coord{x, y}]) + " "
		}
	}
	return repr
}

type Field struct {
	Player string
	X      int
	Y      int
}

func (b *Board) FieldsList() [][]Field {
	myFields := [][]Field{}
	for y := 0; y < 3; y++ {
		rowFields := []Field{}
		for x := 0; x < 3; x++ {
			field := Field{X: x, Y: y, Player: PlayerToString(b.fields[Coord{x, y}])}
			rowFields = append(rowFields, field)
		}
		myFields = append(myFields, rowFields)
	}
	return myFields
}

func (b *Board) Play(x int, y int) {
	b.fields[Coord{x, y}] = b.turn
	b.nextTurn()
}

func (b *Board) nextTurn() {
	if b.turn == CIRCLE {
		b.turn = CROSS
	} else {
		b.turn = CIRCLE
	}
}

func (b *Board) emptyCoords() []Coord {
	empty := []Coord{}
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			coord := Coord{x, y}
			if b.fields[coord] == EMPTY {
				empty = append(empty, coord)
			}
		}
	}
	return empty
}

func playerToOutcome(p Player) Outcome {
	switch {
	case CROSS == p:
		return CROSS_WIN
	case CIRCLE == p:
		return CIRCLE_WIN
	default:
		return NONE
	}
}

func (b *Board) allEqual(coords []Coord) bool {
	// See if all fields in the provided coordinates have the same player
	potential := b.fields[coords[0]]
	found := false
	if potential != EMPTY {
		found = true
		for i := 1; i < 3; i++ {
			if potential != b.fields[coords[i]] {
				found = false
				break
			}
		}
	}
	return found
}

func (b *Board) Winner() Outcome {
	// Check if any row contains only the same player
	for y := 0; y < 3; y++ {
		row := []Coord{}
		for x := 0; x < 3; x++ {
			row = append(row, Coord{x, y})
		}
		if b.allEqual(row) {
			return playerToOutcome(b.fields[row[0]])
		}
	}
	// Check if any column contains only the same player
	for x := 0; x < 3; x++ {
		col := []Coord{}
		for y := 0; y < 3; y++ {
			col = append(col, Coord{x, y})
		}
		if b.allEqual(col) {
			return playerToOutcome(b.fields[col[0]])
		}
	}
	// Check if any diagonal contains only the same player
	diagonal := []Coord{}
	for i := 0; i < 3; i++ {
		diagonal = append(diagonal, Coord{i, i})
	}
	if b.allEqual(diagonal) {
		return playerToOutcome(b.fields[diagonal[0]])
	}

	diagonal = []Coord{}
	for i := 0; i < 3; i++ {
		diagonal = append(diagonal, Coord{2 - i, i})
	}
	if b.allEqual(diagonal) {
		return playerToOutcome(b.fields[diagonal[0]])
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