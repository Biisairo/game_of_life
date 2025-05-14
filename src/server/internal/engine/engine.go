package engine

const WIDTH = 256
const HEIGHT = 256

type GameMap [HEIGHT][WIDTH]bool
type Pattern [][]*Cell
type Time int

type Cell struct {
	X      int
	Y      int
	IsLive bool
}

type Record struct {
	Pattern *Pattern
}

type Engine struct {
	gameMap *GameMap
	count   Time

	lexicon []*Record
}
