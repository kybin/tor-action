package action

type History struct {
	cur     int
	actions []Action
}

func (h *History) Append(a Action) {
	h.actions = append(h.actions, a)
}

func (h *History) Undo(t *Text, c *Cursor, sel *Selection) {
	if h.cur == 0 {
		return
	}
	h.actions[h.cur].Undo(t, c, sel)
	h.cur--
}

func (h *History) Redo(t *Text, c *Cursor, sel *Selection) {
	if h.cur == len(h.actions)-1 {
		return
	}
	h.actions[h.cur].Do(t, c, sel)
	h.cur++
}
