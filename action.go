package action

type History struct {
	cur     int
	actions []Action
}

func (h *History) Append(a Action) {
	h.actions = append(h.actions, a)
}

func (h *History) Undo() {
	if h.cur == 0 {
		return
	}
	h.actions[cur].Undo()
	cur--
}

func (h *History) Redo() {
	if h.cur == len(h.actions)-1 {
		return
	}
	h.actions[h.cur].Do()
	h.cur++
}

type Action interface {
	Do()
	Undo()
	Merge(Action) bool
}

type ActionGenerator func(t *Text, c *Cursor, s *Selection) Action

type GroupAction struct {
	actions []Action
}

func (g GroupAction) Do() {
	for _, a := range g.actions {
		a.Do()
	}
}

func (g GroupAction) Undo() {
	for i := range g.actions {
		j = len(g.actions) - 1 - i
		a := g.actions[j]
		a.Undo()
	}
}

func (g GroupAction) Merge(a Action) bool {
	return false
}

type MoveAction struct {
	old Pos
	new Pos
}

type InsertAction struct {
	p        Pos
	inserted string
}

type DeleteAction struct {
	p         Pos
	backspace bool
}

type SelectAction struct {
	old Range
	new Range
}
