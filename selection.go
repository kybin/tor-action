package action

import "github.com/kybin/tor/cell"

type Selection struct {
	rng cell.Range
}

func (s *Selection) Range() cell.Range {
	return s.rng
}

func (s *Selection) Empty() bool {
	return s.rng.Start == s.rng.End
}

func (s *Selection) Min() cell.Pt {
	return s.rng.Min()
}

func (s *Selection) Max() cell.Pt {
	return s.rng.Max()
}
