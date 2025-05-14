package engine

type Game interface {
	Next() // move mext time
	next(gameMap *GameMap)
	TimeMachine(count Time) *GameMap // make current time to count
	MultiBus(count Time)             // get count time game map

	GetMapHash(gameMap *GameMap) string

	AddPattern(pattern *Pattern) Time // add pattern and get added time
}

func (e *Engine) Next() {
	e.gameMap = e.next(e.gameMap)
}

func (e *Engine) next(gameMap *GameMap) *GameMap {
	var newMap GameMap

	for rowIdx, row := range gameMap {
		for colIdx, isLive := range row {
			xStart := max(colIdx-1, 0)
			xEnd := min(colIdx+2, WIDTH+1)
			yStart := max(rowIdx-1, 0)
			yEnd := min(rowIdx+2, HEIGHT+1)

			nearCount := 0
			for i := xStart; i < xEnd; i++ {
				for j := yStart; j < yEnd; j++ {
					if gameMap[i][j] {
						nearCount++
					}
				}
			}

			if nearCount == 3 || (nearCount == 2 && isLive) {
				newMap[rowIdx][colIdx] = true
			}
		}
	}

	return &newMap
}

func (e *Engine) TimeMachine(count Time) *GameMap {

}

func (e *Engine) MultiBus(count Time) {

}

func (e *Engine) GetMapHash(gameMap *GameMap) string {

}

func (e *Engine) AddPattern(pattern *Pattern) Time {

}
