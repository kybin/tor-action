package action

import "github.com/kybin/tor/cell"

type Action interface {
	Do(*Text, *Cursor, *Selection)
	Undo(*Text, *Cursor, *Selection)
	Merge(Action) bool
}

type ActionGenerator func(*Text, *Cursor, *Selection) Action

type GroupAction struct {
	actions []Action
}

func (g GroupAction) Do(t *Text, c *Cursor, sel *Selection) {
	for _, a := range g.actions {
		a.Do(t, c, sel)
	}
}

func (g GroupAction) Undo(t *Text, c *Cursor, sel *Selection) {
	for i := range g.actions {
		j := len(g.actions) - 1 - i
		a := g.actions[j]
		a.Undo(t, c, sel)
	}
}

func (g GroupAction) Merge(a Action) bool {
	return false
}

type MoveAction struct {
	old cell.Pt
	new cell.Pt
}

func (a MoveAction) Do(t *Text, c *Cursor, sel *Selection) {
	c.SetBPos(a.new)
}

func (a MoveAction) Undo(t *Text, c *Cursor, sel *Selection) {
	c.SetBPos(a.old)
}

func (a MoveAction) Merge(b Action) bool {
	m, ok := b.(MoveAction)
	if !ok {
		return false
	}
	a.new = m.new
	return true
}

type InsertAction struct {
	p        cell.Pt
	inserted string
}

type DeleteAction struct {
	p         cell.Pt
	backspace bool
}

type SelectAction struct {
	old cell.Range
	new cell.Range
}
