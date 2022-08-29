package game

type scoring struct {
	score int
}

func (s *scoring) Score() int {
	return s.score
}

func (s *scoring) RecordNewSquaresCompleted(count int) {
	s.score += count
}
