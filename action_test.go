package action

import (
	"testing"

	"github.com/kybin/tor/cell"
)

var text string = `This is a test string.
You can edit it.
Undo it.
Or redo it.
With tor.
`

type State struct {
	tStr     string
	cPos     cell.Pt
	selRange cell.Range
}

func TestUsage(t *testing.T) {
	cases := []struct {
		title  string
		action ActionGenerator
		want   State
	}{
		// These are series of actions,
		// prior state will be a start point of the next action.
		{
			"first move right",
			MoveRight,
			State{text, cell.Pt{0, 1}, cell.Range{cell.Pt{0, 0}, cell.Pt{0, 0}}},
		},
		{
			"return to origin",
			MoveLeft,
			State{text, cell.Pt{0, 0}, cell.Range{cell.Pt{0, 0}, cell.Pt{0, 0}}},
		},
		{
			"invalid move left",
			MoveLeft,
			State{text, cell.Pt{0, 0}, cell.Range{cell.Pt{0, 0}, cell.Pt{0, 0}}},
		},
	}
	txt := NewText(text)
	cur := &Cursor{}
	sel := &Selection{}
	for _, c := range cases {
		a := c.action(txt, cur, sel)
		a.Do(txt, cur, sel)
		got := State{txt.String(), cur.BPos(), sel.Range()}
		if got != c.want {
			t.Fatalf("%v: got %v, want %v", c.title, got, c.want)
		}
	}
}
