package engine

type Lexicon interface {
	SavePattern(pattern *Pattern)
	GetPattern() []*Pattern
}

func (e *Engine) SavePattern(pattern *Pattern) {

}

func (e *Engine) GetPattern() []*Pattern {

}
