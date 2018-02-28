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

func (s *Selection) Left() cell.Pt {
	if s.rng.Start.Compare(s.rng.End) <= 0 {
		return s.rng.Start
	}
	return s.rng.End
}

func (s *Selection) Right() cell.Pt {
	if s.rng.Start.Compare(s.rng.End) <= 0 {
		return s.rng.End
	}
	return s.rng.Start
}
