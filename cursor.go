package action

import (
	"unicode/utf8"

	"github.com/kybin/tor/cell"
)

type Cursor struct {
	l int
	b int
	v int
}

func (c *Cursor) BPos() cell.Pt {
	return cell.Pt{c.l, c.b}
}

func (c *Cursor) SetBPos(b cell.Pt) {
	c.l = b.L
	c.b = b.O
}

func (c *Cursor) VPos() cell.Pt {
	return cell.Pt{c.l, c.v}
}

func (c *Cursor) Left(t *Text) cell.Pt {
	l := c.l
	b := c.b
	if b == 0 {
		if l == 0 {
			return cell.Pt{0, 0}
		}
		return cell.Pt{l - 1, len(t.Line(l)) - 1}
	}
	remain := t.Line(l)[:b]
	_, size := utf8.DecodeLastRuneInString(remain)
	b -= size
	return cell.Pt{l, b}

}

func (c *Cursor) Right(t *Text) cell.Pt {
	l := c.l
	b := c.b
	if b == len(t.Line(l))-1 {
		if l == len(t.Lines())-1 {
			return cell.Pt{l, b}
		}
		return cell.Pt{l + 1, 0}
	}
	remain := t.Line(l)[b:]
	_, size := utf8.DecodeRuneInString(remain)
	b += size
	return cell.Pt{l, b}
}

func MoveLeft(t *Text, c *Cursor, sel *Selection) Action {
	if !sel.Empty() {
		return MoveAction{
			c.BPos(),
			sel.Min(),
		}
	}
	return MoveAction{
		c.BPos(),
		c.Left(t),
	}
}

func MoveRight(t *Text, c *Cursor, sel *Selection) Action {
	if !sel.Empty() {
		return MoveAction{
			c.BPos(),
			sel.Max(),
		}
	}
	return MoveAction{
		c.BPos(),
		c.Right(t),
	}
}
